package account

type SignUp struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmpassword"`
	Email           string `json:"email"`
}

type User struct {
	Name    string  `json:"name"`
	Account Account `json:"account"`
}
type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
