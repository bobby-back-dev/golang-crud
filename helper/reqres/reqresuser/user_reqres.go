package user

type UserRequestLogin struct {
	ID       int    `json:"id"`
	Name     string `json:"username"`
	Password string `json:"password"`
}

type UserRequestRegistOrUpdate struct {
	ID       int    `json:"id"`
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponseLogin struct {
	ID    int    `json:"id"`
	Name  string `json:"username"`
	Email string `json:"email"`
}
