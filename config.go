package otptimize

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

// =============================================== BCRYPT =============================================== //
// hashing token
func hashingToken(token string) (string, error) {
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(token), 10)
	if err != nil {
		return "", err
	}
	return string(hashedByte), nil
}

// compare token
func checkToken(token, hashedToken string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedToken), []byte(token))
	return err == nil
}

// =============================================== GOMAIL =============================================== //
type MailConfig struct {
	Host     string
	Port     int
	Email    string
	Password string
}

// email connection
var MailConnection *gomail.Dialer

// initialize connection
func mailConnectionInit(config MailConfig) {
	MailConnection = gomail.NewDialer(
		config.Host,
		config.Port,
		config.Email,
		config.Password,
	)
}

// =============================================== REDIS =============================================== //
type RedisConfig struct {
	Host     string
	Port     string
	Password string
}

// redis connection
var RedisClient *redis.Client

func redisInit(config RedisConfig) {
	// Client is setting connection with redis
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: config.Password,
		DB:       0, // use default DB
	})
}

// setValue sets the key value pair
func setRedisValue(key string, value string, expiry time.Duration) error {
	errr := RedisClient.Set(key, value, expiry).Err()
	if errr != nil {
		return errr
	}
	return nil
}

// get value from redis
func getRedisValue(key string) (string, error) {
	value, err := RedisClient.Get(key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}
