package schemas

import "time"

type SchemaStudent struct {
	ID        string    `json:"id" validate:"uuid"`
	Name      string    `json:"name" validate:"required,lowercase"`
	Npm       int       `json:"npm" validate:"required,numeric"`
	Fak       string    `json:"fak" validate:"required,lowercase"`
	Bid       string    `json:"bid" validate:"required,lowercase"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
