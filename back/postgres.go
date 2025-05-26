package main

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pg *pgxpool.Pool

func init() {
	var err error
	pg, err = pgxpool.New(context.TODO(), "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func PGUserCreate(username, password, email string) error {
	_, err := pg.Exec(context.TODO(), `INSERT INTO users (username,password,email) VALUES ($1,$2,$3)`, username, password, email)
	return err
}

func PGFindPassword(username string) (string, error) {
	row := pg.QueryRow(
		context.TODO(),
		`SELECT password FROM users WHERE username=$1`,
		username)
	var hashedPassword string
	return hashedPassword, row.Scan(&hashedPassword)
}
