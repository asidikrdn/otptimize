package otptimize

// =============================================== CUSTOM ERROR =============================================== //
type customError struct {
	Message string
}

func (e *customError) Error() string {
	return e.Message
}

// =============================================== MAIL CONFIG =============================================== //
type MailConfig struct {
	Host     string
	Port     int
	Email    string
	Password string
}

// =============================================== REDIS CONFIG =============================================== //
type RedisConfig struct {
	Host     string
	Port     string
	Password string
}
