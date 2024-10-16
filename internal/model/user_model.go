package model

import (
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"` // Unique
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Bio       string    `json:"bio"`
	Role 			string 		`json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateUserTable() error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(50) NOT NULL UNIQUE,
		email VARCHAR(50) NOT NULL,
		password VARCHAR(255) NOT NULL,
		bio VARCHAR(255) DEFAULT '',
		role ENUM('student', 'instructor', 'admin') NOT NULL DEFAULT 'student',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	)
	`
	_, err := DB.Exec(createTableQuery)
	return err
}

func InsertUser(user *User) (int64, error) {
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"

	result, err := DB.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func FindUserByID(id int64) (*User, error) {
	query := "SELECT id, username, email, password, bio, role, created_at, updated_at FROM users WHERE id = ?"

	var user User
	err := DB.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Bio, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func FindUserByUserName(username string) (*User, error) {
	query := "SELECT id, username, email, password, bio, role, created_at, updated_at FROM users WHERE username = ?"

	var user User
	err := DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Bio, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func FindUserByEmail(email string) (*User, error) {
	query := "SELECT id, username, email, password, bio, role, created_at, updated_at FROM users WHERE email = ?"

	var user User
	err := DB.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Bio, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(user *User) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	query := "UPDATE users SET username = ?, email = ?, bio = ? WHERE id = ?"
	_, err = DB.Exec(query, user.Username, user.Email, user.Bio, user.ID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

