package mgo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	IDFieldName = "_id"
)

// defines the object id
type IDField struct {
	ID primitive.ObjectID `bson:"_id"`
}

type UpdateAtField struct {
	UpdateAt int64 `bson:"update_at"`
}

// NewObjID generates a new object id
var NewObjID = primitive.NewObjectID

// UpdatedAt returns the current unix nanosecond timestamp
var UpdatedAt = func() int64 {
	return time.Now().UnixNano()
}

// Set returns a mongo $set document.
func Set(v interface{}) bson.M {
	return bson.M{
		"$set": v,
	}
}

func SetOnInsert(v interface{}) bson.M {
	return bson.M{
		"$setOnInsert": v,
	}
}
