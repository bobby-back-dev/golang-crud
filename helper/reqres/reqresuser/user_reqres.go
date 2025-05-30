package reqresuser

type UserRequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRequestRegistOrUpdate struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponseLogin struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
