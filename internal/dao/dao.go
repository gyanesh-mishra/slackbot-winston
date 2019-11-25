package dao

import (
	questionAnswerDAO "github.com/gyanesh-mishra/slackbot-winston/internal/dao/questionanswer"
)

// CreateAllIndexes creates the indexes for all tables and returns any errors
func CreateAllIndexes() error {
	err := questionAnswerDAO.CreateIndexes()

	return err
}
