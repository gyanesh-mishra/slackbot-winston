package helpers

import (
	"regexp"
	"strings"

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
	// Make all strings lowercase
	question := strings.ToLower(message)

	// Filter out @user mentions
	re := regexp.MustCompile(`<@[^>]*>`)
	question = re.ReplaceAllString(question, "")

	// Replace any greetings
	greetings := []string{"hi", "hello", "hey", "good morning", "morning", "good day", "good afternoon",
		"good evening", "greetings", "how's it going", "what's up", "howdy", " i "}
	for _, greeting := range greetings {
		question = strings.ReplaceAll(question, greeting, "")
	}

	// Trim any additional whitespace
	question = strings.TrimSpace(question)

	return question
}
