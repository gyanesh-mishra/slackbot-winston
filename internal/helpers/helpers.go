package helpers

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"gopkg.in/jdkato/prose.v2"
)

// GetNLPKeywords extracts the keywords from a string
func GetNLPKeywords(message string) []string {
	keywords := []string{}
	keywordTags := []string{"NN", "VBG", "NNS", "JJ"}
	doc, _ := prose.NewDocument(message, prose.WithExtraction(false), prose.WithSegmentation(false))

	for _, tok := range doc.Tokens() {

		if IsStringInSlice(tok.Tag, keywordTags) {
			keywords = append(keywords, tok.Text)
		}
	}
	return keywords
}

// IsStringInSlice checks if a given input string is in a slice
func IsStringInSlice(input string, slice []string) bool {
	for _, currentWord := range slice {
		if input == currentWord {
			return true
		}
	}
	return false
}

// ExtractQuestionFromMessage sanitizes input message and extracts query from it
func ExtractQuestionFromMessage(message string) string {

	question := message

	// Filter out any greetings
	greetings := []string{"hi", "hello", "hey", "good morning", "morning", "good day", "good afternoon",
		"good evening", "greetings", "how's it going", "what's up", "howdy", "sup"}
	for _, greeting := range greetings {
		// Match whole words only using word boundaries
		currentGreetingRegex := regexp.MustCompile(fmt.Sprintf(`\b(?i)(%s)\b`, greeting))
		question = currentGreetingRegex.ReplaceAllString(question, "")
	}

	// Trim any un-necessary whitespace
	question = strings.TrimSpace(question)

	return question
}

// RemoveUserMention removes any user mentions from the string
func RemoveUserMention(message string) string {

	userRegex := regexp.MustCompile(`<@[^>]*>`)
	message = userRegex.ReplaceAllString(message, "")

	// Remove any trailing whitespaces
	message = strings.ReplaceAll(message, "  ", " ")
	message = strings.TrimSpace(message)

	return message
}

// GetRandomStringFromSlice returns a random item from a slice
func GetRandomStringFromSlice(slice []string) string {
	rand.Seed(time.Now().UnixNano())

	return slice[rand.Intn(len(slice))]
}
