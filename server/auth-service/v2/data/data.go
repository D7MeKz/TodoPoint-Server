package data

// Credential is a struct that contains user credential httpdata.
type Credential struct {
	Email    string
	Password string
}

type RedisParams struct {
	Key     string
	Value   string
	Expires int64
}
