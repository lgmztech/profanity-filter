package middleware

import (
	"strings"
)

func CheckMessage(message string) (string, bool) {
	// check the users input for some profanity

	// List of banned words
	bannedWords := []string{"arsehole", "asshat", "asshole", "bitch", "bloody", "blowjob", "bugger", "bullshit", "clusterfuk", "cock", "cocksucker", "coonass", "cornhole", "cunt", "damn", "dick", "enshittification", "faggot", "feck", "fuck", "fuckery", "healslut", "kike", "motherfucker", "nigga", "nigger", "paki", "poof", "poofter", "prick", "pussy", "ratfucking", "retard", "shit", "shithouse", "shitter", "slut", "spic", "twat", "wanker", "whore", "spastic"}
	// convert the message to lowercase for case-insensitive comparison
	lowerMessage := strings.ToLower(message)

	// Initialize a flag to track if the message has been modified
	modified := false

	// Replace each occurrence of filtered words with "****"
	for _, word := range bannedWords {
		lowerWord := strings.ToLower(word)
		for {
			index := strings.Index(lowerMessage, lowerWord)
			if index == -1 {
				break
			}
			message = message[:index] + strings.Repeat("*", len(word)) + message[index+len(word):]
			lowerMessage = lowerMessage[:index] + strings.Repeat("*", len(word)) + lowerMessage[index+len(word):]
			modified = true
		}
	}

	return message, modified
}
