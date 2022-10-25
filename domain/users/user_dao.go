package users

import (
	"fmt"
	"github.com/dayroMeli/bookstore_userss-api/logger"

	"github.com/dayroMeli/bookstore_userss-api/datasources/mysql/users_db"
	"github.com/dayroMeli/bookstore_userss-api/utils/errors"
)

const (
	queryInsertUser        = "INSERT INTO users(first_name,last_name,email,date_created, password,status)VALUES(?,?,?,?,?,?);"
	queryGetUser           = "SELECT id, first_name, last_name, email, date_created,status FROM users WHERE id=?; "
	queryUpdateUser        = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?"
	queryDeleteUser        = "DELETE FROM users WHERE id=?;"
	queryFindUsersByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
)

func (user *User) Get() *errors.RestErr {

	stmt, err := users_db.UsersDB.Prepare(queryGetUser)
	if err != nil {
		logger.Error("errro when trying to prepare get user statemnt ", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		fmt.Println(err)
		logger.Error("errro when trying to  get user by id ", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.UsersDB.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password, user.Status)
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save: %s", err.Error()))
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.UsersDB.Prepare(queryUpdateUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)

	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to update : %s", err.Error()))
	}

	return nil

}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.UsersDB.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id)

	if _, err = stmt.Exec(user.Id); err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to delete : %s", err.Error()))
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.UsersDB.Prepare(queryFindUsersByStatus)

	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	rows, err := stmt.Query(status)
	defer rows.Close()
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}
