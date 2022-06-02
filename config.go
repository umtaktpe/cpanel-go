package cpanel

type Config struct {
	Username string
	Password string
}

func NewConfig(username, password string) *Config {
	return &Config{
		Username: username,
		Password: password,
	}
}
