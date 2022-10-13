package db

import (
	"errors"
	"log"
	"vasya_project/pkg/models"
)

func GetUserById(dbConnection string, userId string) (user models.User, error error) {
	query := "SELECT * FROM public.users WHERE id=$1"
	rows := Query(dbConnection, query, userId)
	var users []models.User = []models.User{}

	for rows.Next() {
		var (
			id       uint
			email    string
			name     string
			role     int
			password string
		)

		if err := rows.Scan(&id, &name, &email, &role, &password); err != nil {
			log.Fatal(err)
		}

		users = append(users, models.User{Id: id, Email: email, Name: name, Role: role, Password: password})
	}

	if len(users) == 0 {
		return models.User{}, errors.New("user not found")
	}

	return users[0], nil
}

func GetUserByEmail(dbConnection string, email string) (user models.User, error error) {
	query := "SELECT * FROM public.users WHERE email=$1"
	rows := Query(dbConnection, query, email)
	var users []models.User = []models.User{}

	for rows.Next() {
		var (
			id       uint
			email    string
			name     string
			role     int
			password string
		)

		if err := rows.Scan(&id, &name, &email, &role, &password); err != nil {
			log.Fatal(err)
		}

		users = append(users, models.User{Id: id, Email: email, Name: name, Role: role, Password: password})
	}

	if len(users) == 0 {
		return models.User{}, errors.New("user not found")
	}

	return users[0], nil
}
