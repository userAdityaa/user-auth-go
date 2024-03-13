package models

type User struct {
	ID       uint   `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
