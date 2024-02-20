package entity

type Users struct {
	ID      string
	Name    string
	Address string
}

// create struct UserRequest for request body and all colomn is required
type UserRequest struct {
	ID      string `json:"id"`
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
}
type Student struct {
	ID      string
	Name    string
	Address string
	Grade   int
}
