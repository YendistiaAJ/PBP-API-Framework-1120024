package models

import d "test_revel/app/db"

const (
	InsertUserSQL  = `INSERT INTO users (name, age, address, email, password, usertype) VALUES (?, ?, ?, ?, ?, ?)`
	GetUserSQL     = `SELECT id, name, age, address, email, password, usertype FROM users`
	GetUserByIdSQL = GetUserSQL + ` WHERE id=?`
	UpdateUserSQL  = `UPDATE users SET name=?, age=?, address=? WHERE id=?`
	DeleteUserSQL  = `DELETE FROM users WHERE id=?`
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType int    `json:"user_type"`
}

func GetUserById(id string) (User, int) {
	db := d.Connect()
	defer db.Close()

	var user User
	rows, err := db.Query(GetUserByIdSQL, id)
	if err != nil {
		return user, 0
	}

	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.Email, &user.Password, &user.UserType); err != nil {
			return user, 0
		}
	}

	return user, 1
}

func (u *User) InsertUser(name string, age string, address string, email string, password string, userType string) int {
	db := d.Connect()
	defer db.Close()

	_, err := db.Query(InsertUserSQL, name, age, address, email, password, userType)
	if err != nil {
		return 0
	}
	return 1
}
