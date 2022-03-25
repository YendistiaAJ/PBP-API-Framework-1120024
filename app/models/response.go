package models

type Response struct {
	Status      int
	ContentType string
}

type UserResponse struct {
	Status      int
	ContentType string
	Data        User
}

type UsersResponse struct {
	Status      int
	ContentType string
	Data        []User
}

type ProductResponse struct {
	Status      int
	ContentType string
	Data        Product
}

type ProductsResponse struct {
	Status      int
	ContentType string
	Data        []Product
}

type TransactionResponse struct {
	Status      int
	ContentType string
	Data        Transaction
}

type TransactionsResponse struct {
	Status      int
	ContentType string
	Data        []Transaction
}
