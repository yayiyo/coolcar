package dao

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/mongotesting"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoURI string
)

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunMongoInDocker(m, &mongoURI))
}

func TestCreateTrip(t *testing.T) {
	mongoURI = "mongodb://localhost:27017"
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	m := NewMongo(client.Database("coolcar"))
	trip, err := m.CreateTrip(context.Background(), &rentalpb.Trip{
		AccountId: "abc",
		CarId:     "123",
		Start: &rentalpb.LocationStatus{
			Location: &rentalpb.Location{
				Latitude:  34,
				Longitude: 120,
			},
			FeeCent:  7539,
			KmDriven: 3975903,
			PoiName:  "北京",
		},
		End: &rentalpb.LocationStatus{
			Location: &rentalpb.Location{
				Latitude:  34,
				Longitude: 110,
			},
			FeeCent:  54,
			KmDriven: 309797,
			PoiName:  "北京",
		},
		Status: rentalpb.TripStatus_TS_NOT_SPECIFIED,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", trip)
}
