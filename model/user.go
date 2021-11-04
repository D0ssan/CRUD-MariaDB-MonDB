package model

// User is the data struct in the db.
type User struct {
	FirstName      string `json:"firstName" bson:"firstname" db:"first_name"`
	LastName       string `json:"lastName" bson:"lastname" db:"last_name"`
	Specialization string `json:"specialization" bson:"specialization" db:"specialization"`
	DOB            string `json:"dob" bson:"dob" db:"dob"`
	ID             int64  `json:"id,string" db:"id"`
}
