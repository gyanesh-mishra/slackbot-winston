package questionanswer

import (
	"context"
	"fmt"
	"log"

	"github.com/gyanesh-mishra/slackbot-winston/config"
	"github.com/gyanesh-mishra/slackbot-winston/internal/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// QuestionAnswer base database model
type QuestionAnswer struct {
	Question string   `json:"question"`
	Keywords []string `json:"keywords"`
	Answer   string   `json:"answer"`
}

// QuestionAnswers is a list of QuestionAnswer
type QuestionAnswers []*QuestionAnswer

// getCollection returns the QuestionAnswer mongo collection object
func getCollection() *mongo.Collection {
	// Get the database client object
	configuration := config.GetConfig()
	client := configuration.Database.Client

	collection := client.Database("winston").Collection("questions")
	return collection
}

// GetAll returns a list of QuestionAnswer objects
func GetAll() (QuestionAnswers, error) {
	// Get the database client object
	collection := getCollection()

	// Declare list of questionAnswers to fetch
	var results QuestionAnswers

	// Get ALL the database objects under the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		return nil, err
	}

	// Iterate over the database cursor and add pointers to each questionAnswer object to the result set
	for cur.Next(context.TODO()) {

		var elem QuestionAnswer
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	cur.Close(context.TODO())
	return results, nil
}

// Add inserts a new record and returns the ID
func Add(question string, answer string) (interface{}, error) {

	// Break the question into keywords
	keywords := helpers.GetNLPKeywords(question)

	// Store the answer with keywords
	record := QuestionAnswer{Question: question, Keywords: keywords, Answer: answer}
	collection := getCollection()
	res, err := collection.InsertOne(context.TODO(), record)
	if err != nil {
		return nil, err
	}
	id := res.InsertedID

	return id, nil
}

// GetAnswerByQuestion returns an answer from the database matching the question passed
func GetAnswerByQuestion(question string) (string, error) {

	// Break the question into keywords
	keywords := helpers.GetNLPKeywords(question)

	// Get ALL the database objects under the collection
	var result QuestionAnswer
	collection := getCollection()

	// Match either question or keywords
	filters := bson.M{
		"$or": []interface{}{
			bson.M{"question": question},
			bson.M{"keywords": keywords},
		},
	}
	err := collection.FindOne(context.TODO(), filters).Decode(&result)
	if err != nil {
		log.Print(fmt.Sprintf("Error while searching for question: %s, with keywords: %s", question, keywords))
		return "", err
	}

	return result.Answer, nil
}
