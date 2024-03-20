package sq3

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"gocrudapp/model"

	_ "github.com/mattn/go-sqlite3"
)

const (
	insertUser string = "INSERT INTO users (name, email, city, created_at) VALUES (?,?,?,?);"

	fetchAllUser string = "SELECT * FROM users;"

	fetchUserByID string = "SELECT id, name, email, city, created_at FROM users where id = ?;"

	updateUserByID string = "UPDATE users set city = ? where id = ?;"

	deleteUserByID string = "DELETE FROM users where id = ?;"

	deleteAllUser string = "DELETE FROM users;"
)

type UserRepository struct {
	DB *sql.DB
}

func (repo UserRepository) InsertUser(u *model.User) (int, error) {
	result, err := repo.DB.Exec(insertUser, u.Name,
		u.Email, u.City, time.Now().UTC())

	if err != nil {
		log.Println(err)
		return 0, err
	}

	insertID, err := result.LastInsertId()

	return int(insertID), err
}

func (repo UserRepository) FetchAllUser() ([]model.User, error) {
	rows, err := repo.DB.Query(fetchAllUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// An user slice to hold data from returned rows.
	var users []model.User

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email,
			&user.City, &user.CreatedAt); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return users, err
	}

	return users, nil
}

func (repo UserRepository) FetchUserByID(uid int) (model.User, error) {
	var user model.User
	// Query for a value based on a single row.
	err := repo.DB.QueryRow(fetchUserByID, uid).
		Scan(&user.ID, &user.Name, &user.Email,
			&user.City, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("id doesn't exist %d", uid)
		}
		return user, fmt.Errorf("fetch error %d: %v", uid, err)
	}

	return user, nil
}

func (repo UserRepository) UpdateUserByID(uid int, u *model.User) (int, error) {
	result, err := repo.DB.Exec(updateUserByID, u.City, uid)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	return int(rowsAffected), err
}

func (repo UserRepository) DeleteUserByID(uid int) (int, error) {
	result, err := repo.DB.Exec(deleteUserByID, uid)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	return int(rowsAffected), err
}

func (repo UserRepository) DeleteAllUser() (int, error) {
	result, err := repo.DB.Exec(deleteAllUser)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	return int(rowsAffected), err
}
