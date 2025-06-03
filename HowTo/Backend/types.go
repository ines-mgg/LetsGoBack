package main

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

type User struct {
	ID        int
	Username  string
	ProfilPic string
	Email     string
	Password  string
	Role      string
}
type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type ProfileRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

var Users = []User{
	{ID: 1, Username: "user1", Email: "user1@example.com", Password: "password1", Role: RoleAdmin},
	{ID: 2, Username: "user2", Email: "user2@example.com", Password: "password2", Role: RoleUser},
	{ID: 3, Username: "user3", Email: "user3@example.com", Password: "password3", Role: RoleUser},
	{ID: 4, Username: "user4", Email: "user4@example.com", Password: "password4", Role: RoleUser},
}

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Stock       int
	Pictures    []string
}

type ProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

var Products = []Product{
	{ID: 1, Name: "Product 1", Description: "Description for product 1", Price: 10.99, Stock: 100},
	{ID: 2, Name: "Product 2", Description: "Description for product 2", Price: 20.99, Stock: 50},
	{ID: 3, Name: "Product 3", Description: "Description for product 3", Price: 30.99, Stock: 75},
	{ID: 4, Name: "Product 4", Description: "Description for product 4", Price: 40.99, Stock: 25},
	{ID: 5, Name: "Product 5", Description: "Description for product 5", Price: 50.99, Stock: 10},
	{ID: 6, Name: "Product 6", Description: "Description for product 6", Price: 60.99, Stock: 5},
}
