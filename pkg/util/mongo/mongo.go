package mongoutil

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func ToBSONDocument(data any) (bson.D, error) {
	bsonBytes, err := bson.Marshal(data)
	if err != nil {
		return nil, err
	}
	var document bson.D
	err = bson.Unmarshal(bsonBytes, &document)
	if err != nil {
		return nil, err
	}
	return document, nil
}

func GetMongoId(result *mongo.InsertOneResult) string {
	return result.InsertedID.(primitive.ObjectID).Hex()
}
