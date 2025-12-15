package database

import (
	"database/sql"

	"github.com/fletiszi/goteste/schemas"
)

func GetUsers(db *sql.DB) ([]schemas.User, error) {
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []schemas.User

	for rows.Next() {
		var u schemas.User
		err := rows.Scan(&u.ID, &u.Name, &u.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
