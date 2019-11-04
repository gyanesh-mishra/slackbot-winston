package helpers

import (
	"gopkg.in/jdkato/prose.v2"
)

// GetNLPKeywords extracts the keywords from a string
func GetNLPKeywords(message string) []string {
	keywords := []string{}
	doc, _ := prose.NewDocument(message, prose.WithExtraction(false))
	for _, tok := range doc.Tokens() {
		if tok.Tag == "NN" || tok.Tag == "VBG" || tok.Tag == "NNS" {
			keywords = append(keywords, tok.Text)
		}
	}
	return keywords
}
