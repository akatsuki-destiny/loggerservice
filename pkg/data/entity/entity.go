package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ErrLog struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Level     string             `bson:"level"`
	Source    string             `bson:"source"`
	Message   string             `bson:"message"`
	Error     string             `bson:"error"`
	Timestamp time.Time          `bson:"timestamp"`
}

type SuccessLog struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Source         string             `bson:"source"`
	Request        string             `bson:"request"`
	RequestHeader  string             `bson:"request_header"`
	Response       string             `bson:"response"`
	ResponseHeader string             `bson:"response_header"`
	Timestamp      time.Time          `bson:"timestamp"`
}
