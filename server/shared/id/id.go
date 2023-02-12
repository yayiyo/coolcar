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

// IdentityID defines the identity id type
type IdentityID string

func (i IdentityID) String() string {
	return string(i)
}

type CarID string

func (c CarID) String() string {
	return string(c)
}
