package schemas

import "time"

type SchemaAuth struct {
	ID        string    `json:"id" validate:"uuid"`
	Fullname  string    `json:"fullname" validate:"required,lowercase"`
	Email     string    `json:"email" validate:"required,email"`
	Token     string    `json:"token" validate:"required"`
	Password  string    `json:"password" validate:"required,gte=8"`
	Cpassword string    `json:"cpassword" validate:"required,gte=8"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
