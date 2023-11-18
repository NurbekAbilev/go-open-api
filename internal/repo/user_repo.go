package repo

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(pgx *pgxpool.Pool) UserRepo {
	return UserRepo{
		db: pgx,
	}
}

func (r UserRepo) CreateUser(ctx context.Context, user User) (string, error) {
	query := `
		insert into users(login, password, created_at) 
			values ($1, $2, $3) 
		returning id
	`

	var lastInsertId string
	err := r.db.QueryRow(
		ctx, query,
		user.Login, user.Password, time.Now(),
	).Scan(&lastInsertId)

	if err != nil {
		return "", err
	}

	return lastInsertId, nil
}

func (r UserRepo) GetUserByLogin(ctx context.Context, login string) (*User, error) {
	query := `
		select id, login, password, created_at from users
			where login = $1
	`

	u := User{}
	err := r.db.QueryRow(ctx, query, login).Scan(
		&u.ID, &u.Login, &u.Password, &u.CreatedAt,
	)
	if err != nil {
		log.Println("Error during find employee by login: %w", err)
		return nil, err
	}

	return &u, nil
}
