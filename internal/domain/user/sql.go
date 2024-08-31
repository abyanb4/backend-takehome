package user

const (
	RegisterUserQuery string = "INSERT INTO users (id, username, email, password_hash) VALUES (?, ?, ?, ?)"
	LoginQuery        string = "SELECT id, password_hash FROM users WHERE username = ?"

	IsUsernameExistsQuery string = "SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)"
	IsEmailExistsQuery    string = "SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)"
)
