package dao

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
	mgo "coolcar/shared/mongo"
	"coolcar/shared/mongo/objid"
	"coolcar/shared/mongotesting"
	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunMongoInDocker(m))
}

func TestCreateTrip(t *testing.T) {
	ctx := context.Background()
	client, err := mongotesting.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("coolcar")
	err = mongotesting.SetupIndexes(ctx, db)
	if err != nil {
		log.Fatalf("failed to setup indexes: %+v", err)
	}

	m := NewMongo(db)
	cases := []struct {
		name       string
		tripID     string
		accountID  string
		tripStatus rentalpb.TripStatus
		wantErr    bool
	}{
		{
			name:       "finish-1",
			tripID:     "63de1910e8ff1a3d9eeab107",
			accountID:  "account-1",
			tripStatus: rentalpb.TripStatus_FINISHED,
		},
		{
			name:       "finish-2",
			tripID:     "63de1910e8ff1a3d9eeab108",
			accountID:  "account-1",
			tripStatus: rentalpb.TripStatus_FINISHED,
		},
		{
			name:       "in-progress-1",
			tripID:     "63de1910e8ff1a3d9eeab109",
			accountID:  "account-1",
			tripStatus: rentalpb.TripStatus_IN_PROGRESS,
			wantErr:    false,
		},
		{
			name:       "in-progress-2",
			tripID:     "63de1910e8ff1a3d9eeab110",
			accountID:  "account-1",
			tripStatus: rentalpb.TripStatus_IN_PROGRESS,
			wantErr:    true,
		},
		{
			name:       "in-progress-3",
			tripID:     "63de1910e8ff1a3d9eeab111",
			accountID:  "account-2",
			tripStatus: rentalpb.TripStatus_IN_PROGRESS,
		},
	}

	for _, c := range cases {
		mgo.NewObjID = func() primitive.ObjectID {
			return objid.MustFromID(id.TripID(c.tripID))
		}
		tr, err := m.CreateTrip(ctx, &rentalpb.Trip{
			AccountId: c.accountID,
			Status:    c.tripStatus,
		})

		if c.wantErr {
			if err == nil {
				t.Errorf("%s want error, but not got.", c.name)
			}
			t.Logf("%s tested successfully", c.name)
			continue
		}

		if err != nil {
			t.Errorf("%s want not error, but got.", c.name)
			continue
		}

		if c.tripID != tr.ID.Hex() {
			t.Errorf("%s trip want %q, but got %q", c.name, c.tripID, tr.ID.Hex())
		}

		t.Logf("%s tested successfully", c.name)
	}
}

func TestGetTrip(t *testing.T) {
	client, err := mongotesting.NewClient(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	m := NewMongo(client.Database("coolcar"))
	mgo.NewObjID = primitive.NewObjectID
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
		Status: rentalpb.TripStatus_FINISHED,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", trip)
	got, err := m.GetTrip(context.Background(), objid.ToTripID(trip.ID), id.AccountID(trip.Trip.AccountId))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", trip)

	if diff := cmp.Diff(trip, got, protocmp.Transform()); diff != "" {
		t.Errorf("results is different: %s", diff)
	}
}

func TestGetTrips(t *testing.T) {

}
