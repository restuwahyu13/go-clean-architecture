package repositorys

import (
	"gorm.io/gorm"
)

type RepositoryPing interface {
	PingRepository() string
}

type repositoryPing struct {
	db *gorm.DB
}

func NewRepositoryPing(db *gorm.DB) *repositoryPing {
	return &repositoryPing{db: db}
}

func (r *repositoryPing) PingRepository() string {
	return "Ping Test"
}
