package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pg *pgxpool.Pool

func init() {
	var err error
	pg, err = pgxpool.New(context.TODO(), "postgres://aysa:pass@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}
	var kaka string
	pg.QueryRow(context.TODO(), "SELECT name FROM kids_zones").Scan(&kaka)
	fmt.Printf("%+v\n", kaka)
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
