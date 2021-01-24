package questionanswer

import (
	"context"
	"strings"
	"time"

	"github.com/gyanesh-mishra/slackbot-winston/config"
	"github.com/gyanesh-mishra/slackbot-winston/internal/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// QuestionAnswer base database model
type QuestionAnswer struct {
	Question      string    `bson:"question" json:"question"`
	Keywords      []string  `bson:"keywords" json:"keywords"`
	Answer        string    `bson:"answer" json:"answer"`
	TeamID        string    `bson:"teamID" json:"teamID"`
	LastUpdated   time.Time `bson:"lastUpdated" json:"lastUpdated"`
	LastUpdatedBy string    `bson:"lastUpdatedBy" json:"lastUpdatedBy"`
}

// QuestionAnswers is a list of QuestionAnswer
type QuestionAnswers []*QuestionAnswer

// CreateIndexes creates the indexes and escalates any errors
func CreateIndexes() error {
	collection := getCollection()
	indexes := []mongo.IndexModel{
		{
			Keys: bson.M{"teamID": 1, "question": 1, "keyword": 1},
		},
	}

	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	_, err := collection.Indexes().CreateMany(context.TODO(), indexes, opts)

	return err
}

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

// AddOrUpdate inserts or updates a record and return it
func AddOrUpdate(question string, answer string, updatedBy string, teamID string) (interface{}, error) {

	// Store all questions as lowercase for saving time on case sensitivity search
	question = strings.ToLower(question)

	// Break the question into keywords
	keywords := helpers.GetNLPKeywords(question)

	// Get the database collection
	collection := getCollection()

	// Construct the object
	data := QuestionAnswer{
		Question:      question,
		Keywords:      keywords,
		Answer:        answer,
		TeamID:        teamID,
		LastUpdatedBy: updatedBy,
		LastUpdated:   time.Now().UTC(),
	}

	// Construct filters for upsert
	filter := bson.M{
		"teamID": teamID,
		"$or": []interface{}{
			bson.M{"question": question},
			bson.M{"keywords": keywords},
		},
	}
	update := bson.M{"$set": data}

	// Upsert the answer
	res, err := collection.UpdateOne(context.TODO(), filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return nil, err
	}
	return res, nil

}

// GetByQuestion returns an answer from the database matching the question passed
func GetByQuestion(question string, teamID string) (QuestionAnswer, error) {

	// Convert incoming question to lowercase for search
	question = strings.ToLower(question)

	// Break the question into keywords
	keywords := helpers.GetNLPKeywords(question)

	// Get ALL the database objects under the collection
	var result QuestionAnswer
	collection := getCollection()

	// Match either question or keywords
	filters := bson.M{
		"teamID": teamID,
		"$or": []interface{}{
			bson.M{"question": question},
			bson.M{"keywords": keywords},
		},
	}
	err := collection.FindOne(context.TODO(), filters).Decode(&result)
	if err != nil {
		return QuestionAnswer{}, err
	}

	return result, nil
}
