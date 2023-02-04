package objid

import (
	"fmt"

	"coolcar/shared/id"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FromID(id fmt.Stringer) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id.String())
}

// MustFromID the id must be object id, it will panic if id is not valid primitive.ObjectID.
func MustFromID(id fmt.Stringer) primitive.ObjectID {
	i, err := primitive.ObjectIDFromHex(id.String())
	if err != nil {
		panic(err)
	}
	return i
}

func ToAccountID(oid primitive.ObjectID) id.AccountID {
	return id.AccountID(oid.Hex())
}

func ToTripID(oid primitive.ObjectID) id.TripID {
	return id.TripID(oid.Hex())
}
