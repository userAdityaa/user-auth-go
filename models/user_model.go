package models

type User struct {
	ID       uint   `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username" validate:"required,min=4,max=10"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
