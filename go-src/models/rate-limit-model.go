package models

type RateLimitedModel struct {
	ID        uint
	Token     string
	CallsUsed uint8
	RateMax   uint8
}
