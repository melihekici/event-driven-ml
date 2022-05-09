package repository

import (
	"authapp/internal/dto"
	"log"
)

type UserQuery interface {
	CreateUser(user dto.User) (*int64, error)
	GetUser(id int64) (*dto.User, error)
	DeleteUser(id int64) error
	UpdateUser(user dto.User) (*dto.User, error)
}

type userQuery struct{}

func (u *userQuery) CreateUser(user dto.User) (*int64, error) {
	db, err := NewDB()
	if err != nil {
		log.Println("Unable to get database connection. ", err)
		return nil, err
	}
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO user (username, password, email) VALUES($1, $2, $3)")
	if err != nil {
		log.Println("Error while preparing statemtnt for CreateUser ", err)
		return nil, err
	}

	result, err := statement.Exec(user.Username, user.Password, user.Email)
	if err != nil {
		log.Println("Error while creating user. ", err)
		return nil, err
	}

	var insertedID int64
	insertedID, err = result.LastInsertId()
	if err != nil {
		log.Println("Error getting last inserted id. ", err)
		return nil, err
	}

	return &insertedID, nil
}

func (u *userQuery) GetUser(id int64) (*dto.User, error) {
	db, err := NewDB()
	if err != nil {
		log.Println("Unable to get database connection. ", err)
		return nil, err
	}
	defer db.Close()

	row := db.QueryRow("SELECT * FROM user WHERE id=$1", id)

	var user dto.User
	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		log.Println("Error executing sql", err)
		return nil, err
	}

	return &user, nil
}

func (u *userQuery) DeleteUser(id int64) error {
	db, err := NewDB()
	if err != nil {
		log.Println("Unable to get database connection. ", err)
		return err
	}
	db.Close()

	stmt := "DELETE FROM user WHERE id=$1"

	_, err = db.Exec(stmt, id)
	if err != nil {
		log.Fatalf("Error while deleting user. %v", err)
		return err
	}

	return nil
}

func (u *userQuery) UpdateUser(user dto.User) (*dto.User, error) {
	db, err := NewDB()
	if err != nil {
		log.Println("Unable to get database connection. ", err)
		return nil, err
	}
	defer db.Close()

	stmt := "UPDATE user SET username = $2, password = $3, email = $4 WHERE id = $1"
	_, err = db.Exec(stmt, user.ID, user.Username, user.Password, user.Email)
	if err != nil {
		log.Printf("Error updating user. %v", err)
		return nil, err
	}

	var updatedUser *dto.User
	updatedUser, err = u.GetUser(user.ID)
	if err != nil {
		log.Printf("Error while reading updated user from database. %v", err)
		return nil, err
	}

	return updatedUser, nil
}
