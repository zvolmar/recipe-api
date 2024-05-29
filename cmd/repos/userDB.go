package repos

import (
	
	"nomcom-api/cmd/database"
	"nomcom-api/cmd/models"
	"time"
)

func CreateUser(user models.User) (models.User,error) {
	db := database.GetDB()
	sqlStatement := `INSERT INTO users (username, email, password)
	VALUES ($1, $2, $3) RETURNING id`

	err := db.QueryRow(sqlStatement, user.Username, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return user, err
	}
	return user, nil
}

func UpdateUser(user models.User, id int) (models.User,error) {
	db := database.GetDB()
	sqlStatement := `UPDATE users SET username = $2, email = $3, password = $4, updated_at = $5
	WHERE id = $1
	RETURNING id, username, email, updated_at`

	err := db.QueryRow(sqlStatement, id, user.Username, user.Email, user.Password, time.Now()).Scan(&user.ID, &user.Username, &user.Email, &user.UpdatedAt)
	if err != nil {
		return models.User{}, err
	}

	user.ID = id
	return user, nil
}

func GetUser(user models.User, id int) (models.User, error) {
	db := database.GetDB()
	sqlStatement := `SELECT id, username, email, created_at, updated_at FROM users WHERE id = $1`
	
	err := db.QueryRow(sqlStatement, id).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return models.User{}, err
	}
	return user, err
}
