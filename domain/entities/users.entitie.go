package entitie

import (
	"time"

	"github.com/guregu/null/v6/zero"
)

type UsersEntitie struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Status    string    `db:"status"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt zero.Time `db:"updated_at"`
	DeletedAt zero.Time `db:"deleted_at"`
}
