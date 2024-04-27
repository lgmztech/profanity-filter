package models

type MessageResponse struct {
	OriginalMessage string
	NewMessage      string
	HasProfanity    bool
}
