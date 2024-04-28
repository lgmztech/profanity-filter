package database

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lgmztech/profanity-filter/models"
)

func CreateNewToken(token string) error {
	// Create new token in database
	newToken := models.RateLimitedModel{
		Token:     token,
		CallsUsed: 0,
		RateMax:   10,
	}
	result := DB.Create(&newToken)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CheckForValidToken(accessToken string) (*models.RateLimitedModel, error) {
	findToken, err := findTokenDetails(accessToken)
	if err != nil {
		return nil, err
	}

	hasBeenLimited := checkIfLimitReached(findToken)
	if hasBeenLimited {
		return findToken, fiber.ErrLocked
	}

	return findToken, nil
}

func AppendTokenUse(accessToken string) error {
	findToken, err := findTokenDetails(accessToken)
	if err != nil {
		return err
	}
	if err := DB.Model(findToken).Update("calls_used", findToken.CallsUsed+1).Error; err != nil {
		return err
	}
	return nil
}

func findTokenDetails(accessToken string) (*models.RateLimitedModel, error) {
	var token models.RateLimitedModel
	if err := DB.Where("token = ?", accessToken).First(&token).Error; err != nil {
		return nil, err
	}
	return &token, nil
}

func checkIfLimitReached(token *models.RateLimitedModel) bool {
	return token.CallsUsed >= token.RateMax
}
