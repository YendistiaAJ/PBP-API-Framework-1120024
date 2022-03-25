package controllers

import (
	"net/http"
	d "test_revel/app/db"
	m "test_revel/app/models"

	"github.com/revel/revel"
)

type User struct {
	*revel.Controller
}

func (c User) CreateUser() revel.Result {
	var response m.Response
	db := d.Connect()
	defer db.Close()

	name := c.Params.Form.Get("name")
	age := c.Params.Form.Get("age")
	address := c.Params.Form.Get("address")
	email := c.Params.Form.Get("email")
	password := c.Params.Form.Get("password")
	userType := c.Params.Form.Get("user_type")

	if name == "" || age == "" || address == "" || email == "" || password == "" {
		c.Log.Errorf("Missing insert field")
		response = m.Response{
			Status:      http.StatusBadRequest,
			ContentType: "Missing insert field",
		}
		return c.RenderJSON(response)
	}

	rows, err := db.Exec(m.InsertUserSQL, name, age, address, email, password, userType)
	if err != nil {
		c.Log.Errorf("Query Error: %v", err)
		response = m.Response{
			Status:      http.StatusBadRequest,
			ContentType: "Failed to process query",
		}
		return c.RenderJSON(response)
	}

	affRows, _ := rows.RowsAffected()

	if affRows == 0 {
		c.Log.Errorf("Insert Failed")
		response = m.Response{
			Status:      http.StatusBadRequest,
			ContentType: "No rows were inserted",
		}
		return c.RenderJSON(response)
	}

	response = m.Response{
		Status:      http.StatusOK,
		ContentType: "Success",
	}
	return c.RenderJSON(response)
}

func (c User) GetUserById() revel.Result {
	id := c.Params.Route.Get("id")
	var response m.UserResponse
	db := d.Connect()
	defer db.Close()

	rows, err := db.Query(m.GetUserByIdSQL, id)

	if err != nil {
		c.Log.Errorf("Query Error: %v", err)
		response = m.UserResponse{
			Status:      http.StatusBadRequest,
			ContentType: "Failed to process query",
		}
		return c.RenderJSON(response)
	}

	var user m.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.Email, &user.Password, &user.UserType); err != nil {
			c.Log.Errorf("Rows Scan Error: %v", err)
			response = m.UserResponse{
				Status:      http.StatusBadRequest,
				ContentType: "Failed to retrieve user data",
			}
			return c.RenderJSON(response)
		}
	}

	response = m.UserResponse{
		Status:      http.StatusOK,
		ContentType: "Success",
		Data:        user,
	}
	return c.RenderJSON(response)
}

func (c User) GetUsers() revel.Result {
	var response m.UsersResponse
	db := d.Connect()
	defer db.Close()

	rows, err := db.Query(m.GetUserSQL)
	if err != nil {
		c.Log.Errorf("Query Error: %v", err)
		response = m.UsersResponse{
			Status:      http.StatusBadRequest,
			ContentType: "Failed to process query",
		}
		return c.RenderJSON(response)
	}

	var users []m.User
	for rows.Next() {
		var user m.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.Email, &user.Password, &user.UserType); err != nil {
			c.Log.Errorf("Rows Scan Error: %v", err)
			response = m.UsersResponse{
				Status:      http.StatusBadRequest,
				ContentType: "Failed to retrieve users data",
			}
			return c.RenderJSON(response)
		} else {
			users = append(users, user)
		}
	}

	response = m.UsersResponse{
		Status:      http.StatusOK,
		ContentType: "Success",
		Data:        users,
	}
	return c.RenderJSON(response)
}

func (c User) UpdateUser() revel.Result {
	id := c.Params.Route.Get("id")
	name := c.Params.Form.Get("name")
	age := c.Params.Form.Get("age")
	address := c.Params.Form.Get("address")

	var response m.Response
	db := d.Connect()
	defer db.Close()

	if name == "" || age == "" || address == "" {
		c.Log.Errorf("Missing update field")
		response = m.Response{
			Status:      http.StatusBadRequest,
			ContentType: "Missing update field",
		}
		return c.RenderJSON(response)
	}

	rows, err := db.Exec(m.UpdateUserSQL, name, age, address, id)
	if err != nil {
		c.Log.Errorf("Query Error: %v", err)
		response = m.Response{
			Status:      http.StatusBadRequest,
			ContentType: "Failed to process query",
		}
		return c.RenderJSON(response)
	}

	affRows, _ := rows.RowsAffected()

	if affRows == 0 {
		c.Log.Errorf("Update Failed")
		response = m.Response{
			Status:      http.StatusBadRequest,
			ContentType: "No rows were updated",
		}
		return c.RenderJSON(response)
	}

	response = m.Response{
		Status:      http.StatusOK,
		ContentType: "Success",
	}
	return c.RenderJSON(response)
}

func (c User) DeleteUser() revel.Result {
	id := c.Params.Route.Get("id")
	var response m.Response
	db := d.Connect()
	defer db.Close()

	success := m.DeleteTransactionByUserId(id)
	if success == 0 {
		c.Log.Errorf("Delete Failed")
		response = m.Response{
			Status:      http.StatusBadRequest,
			ContentType: "No rows were deleted",
		}
		return c.RenderJSON(response)
	}

	rows, err := db.Exec(m.DeleteUserSQL, id)
	if err != nil {
		c.Log.Errorf("Query Error: %v", err)
		response = m.Response{
			Status:      http.StatusBadRequest,
			ContentType: "Failed to process query",
		}
		return c.RenderJSON(response)
	}

	affRows, _ := rows.RowsAffected()

	if affRows == 0 {
		c.Log.Errorf("Delete Failed")
		response = m.Response{
			Status:      http.StatusBadRequest,
			ContentType: "No rows were deleted",
		}
		return c.RenderJSON(response)
	}

	response = m.Response{
		Status:      http.StatusOK,
		ContentType: "Success",
	}
	return c.RenderJSON(response)
}
