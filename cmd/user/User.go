package user

type User struct {
	Username string
	Password string
}

type createUserRequest struct {
	Username string
	Password string
}
