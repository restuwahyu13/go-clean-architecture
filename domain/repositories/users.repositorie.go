package repo

import (
	"context"

	"github.com/jmoiron/sqlx"
	entitie "github.com/restuwahyu13/go-clean-architecture/domain/entities"
	inf "github.com/restuwahyu13/go-clean-architecture/shared/interfaces"
)

type usersRepositorie struct {
	ctx     context.Context
	db      *sqlx.DB
	entitie *entitie.UsersEntitie
}

func NewUsersRepositorie(ctx context.Context, db *sqlx.DB) inf.IUsersRepositorie {
	return &usersRepositorie{ctx: ctx, db: db, entitie: new(entitie.UsersEntitie)}
}

func (r usersRepositorie) Find(query string, args ...any) ([]entitie.UsersEntitie, error) {
	users := []entitie.UsersEntitie{}

	if args == nil {
		if err := r.db.SelectContext(r.ctx, &users, sqlx.Rebind(sqlx.DOLLAR, query)); err != nil {
			return nil, err
		}
	} else {
		if err := r.db.SelectContext(r.ctx, &users, sqlx.Rebind(sqlx.DOLLAR, query), args...); err != nil {
			return nil, err
		}
	}

	return users, nil
}

func (r usersRepositorie) FindOne(query string, args ...any) (*entitie.UsersEntitie, error) {
	if err := r.db.GetContext(r.ctx, r.entitie, sqlx.Rebind(sqlx.DOLLAR, query), args...); err != nil {
		return nil, err
	}

	return r.entitie, nil
}

func (r usersRepositorie) Create(dest any, query string, args ...any) error {
	if err := r.db.GetContext(r.ctx, dest, query, args...); err != nil {
		return err
	}

	return nil
}

func (r usersRepositorie) Update(dest any, query string, args ...any) error {
	if err := r.db.GetContext(r.ctx, dest, query, args...); err != nil {
		return err
	}

	return nil
}

func (r usersRepositorie) Delete(dest any, query string, args ...any) error {
	if err := r.db.GetContext(r.ctx, dest, query, args...); err != nil {
		return err
	}

	return nil
}
