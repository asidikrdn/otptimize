package otptimize

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// generate random number with custom length
func otpGenerator(length int) string {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	token := make([]string, length)
	for i := range token {
		token[i] = strconv.Itoa(rng.Intn(10))
	}
	return strings.Join(token, "")
}

// =============================================== CONNECTION INIT =============================================== //
func ConnectionInit(mailConfig MailConfig, redisConfig RedisConfig) {
	mailConnectionInit(mailConfig)
	redisInit(redisConfig)
}

// =============================================== GENERATE OTP AND SEND IT BY EMAIL =============================================== //
func GenerateAndSendOTP(otpLength int, appName string, targetName string, targetEmail string) error {
	// generate otp code
	otpToken := otpGenerator(otpLength)

	// save hashed OTP Code on redis server
	hashedOTP, err := hashingToken(otpToken)
	if err != nil {
		fmt.Println("ERROR on hashing token : ", err)
		return err
	}

	err = setRedisValue(targetEmail, hashedOTP, time.Minute*5)
	if err != nil {
		fmt.Println("ERROR on save token to redis : ", err)
		return err
	}

	// send otp to target mail
	go sendVerificationEmail(appName, targetName, targetEmail, otpToken)

	return nil
}

// validating otp token
// =============================================== VALIDATE OTP =============================================== //
func ValidateOTP(email string, otpToken string) (bool, error) {

	// get token from redis
	token, err := getRedisValue(email)
	if err != nil {
		fmt.Println("ERROR on get token from redis : ", err)
		return false, err
	}

	// compare token
	if !checkToken(otpToken, token) {
		err := &customError{Message: "Token invalid"}
		fmt.Println("ERROR on compare token : ", err)
		return false, err
	}

	return true, nil
}
