package id

// AccountID defines the account id type
type AccountID string

func (a AccountID) String() string {
	return string(a)
}

// TripID defines the trip id type
type TripID string

func (t TripID) String() string {
	return string(t)
}
