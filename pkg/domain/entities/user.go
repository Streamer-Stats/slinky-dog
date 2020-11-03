package entities

type User struct {
	tableName struct{} `pg:"Auth.Users"`
	// Id
	ID int64 `pg:"Id"`

	// email
	Email string `pg:"Email"`

	// password
	Password string `pg:"Password"`

	// token
	Token string `pg:"-"`

	// username
	Username string `pg:"Username"`
}
