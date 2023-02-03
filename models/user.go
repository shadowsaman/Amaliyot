package models

type UserPrimarKey struct {
	Id string `json:"id"`
}

type CreateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateUser struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

type UpdateUserSwagger struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

type GetListUserRequest struct {
	Limit  int32
	Offset int32
}

type GetListUserResponse struct {
	Count int32   `json:"count"`
	Users []*User `json:"users"`
}
