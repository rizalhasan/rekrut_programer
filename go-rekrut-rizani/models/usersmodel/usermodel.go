package usersmodel

import (
	"go-rekrut-rizani/config"
	"go-rekrut-rizani/entities"
)

func GetAll() []entities.Users {
	rows, err := config.DB.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var user []entities.Users

	for rows.Next() {
		var users entities.Users
		if err := rows.Scan(&users.Id, &users.Username, &users.Password, &users.Role, &users.CreatedAt); err != nil {
			panic(err)
		}

		user = append(user, users)

	}

	return user

}
