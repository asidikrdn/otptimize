# OTPtimize

OTPtimize is a Go package that helps in generating, sending, and validating OTP (One-Time Password). This package is designed to simplify the usage of OTP in your application with a focus on optimizing resource utilization.

## Installation

You can install the OTPtimize package using the following command:

```bash
go get -u github.com/asidikrdn/otptimize
```

## Usage

Here's an example of how to use the OTPtimize package:

### **Generate OTP**

```go
package main

import (
  "fmt"
  "github.com/asidikrdn/otptimize"
)

func main() {
  // Initialize connections
  mailConfig := otptimize.MailConfig{
    Host:     "<mail_server_host>",      // e.g : "smtp.gmail.com"
    Port:     "<mail_server_port>",      // e.g : 587
    Email:    "<your_email>",            // e.g : "your_email@mail.com"
    Password: "<your_email_password>",   // e.g : "asszsdweaqw2e"
  }
  redisConfig := otptimize.RedisConfig{
    Host:     "<redis_server_host>",     // e.g : "172.17.0.1"
    Port:     "<redis_server_port>",     // e.g : "6379"
    Password: "<redis_server_password>", // e.g : "1j2oda982jskxzi"
  }
  otptimize.ConnectionInit(mailConfig, redisConfig)

  // Generate and send OTP
  err := otptimize.GenerateAndSendOTP(6, 25, "MyApp", "John Doe", "sidikrudini16@gmail.com")
  if err != nil {
    fmt.Println("Error:", err)
    return
  }
}
```

### **Validate OTP**

```go
package main

import (
  "fmt"
  "github.com/asidikrdn/otptimize"
)

func main() {
  // Initialize connections
  mailConfig := otptimize.MailConfig{
    Host:     "<mail_server_host>",      // e.g : "smtp.gmail.com"
    Port:     "<mail_server_port>",      // e.g : 587
    Email:    "<your_email>",            // e.g : "your_email@mail.com"
    Password: "<your_email_password>",   // e.g : "asszsdweaqw2e"
  }
  redisConfig := otptimize.RedisConfig{
    Host:     "<redis_server_host>",     // e.g : "172.17.0.1"
    Port:     "<redis_server_port>",     // e.g : "6379"
    Password: "<redis_server_password>", // e.g : "1j2oda982jskxzi"
  }
  otptimize.ConnectionInit(mailConfig, redisConfig)

  // Validate OTP
  valid, err := otptimize.ValidateOTP("your_email@mail.com", "085481")
  if err != nil {
    fmt.Println("Error:", err)
    return
  }

  if valid {
    fmt.Println("OTP Valid!")
  } else {
    fmt.Println("Invalid OTP!")
  }
}
```

Make sure to replace email and Redis configuration values according to your needs.

## Documentation

### ConnectionInit

The `ConnectionInit` function is used to initialize the connections required by the OTPtimize package. This function takes two parameters: `mailConfig` of type `MailConfig` and `redisConfig` of type `RedisConfig`. The `mailConfig` parameter contains configurations for connecting to the email server, such as host, port, sender's email, and password. The `redisConfig` parameter contains configurations for connecting to the Redis server, such as host, port, and password. This function needs to be called before using other functions within this package.

### GenerateAndSendOTP

The `GenerateAndSendOTP` function is responsible for generating an OTP (One-Time Password), sending it to the specified email address, and storing it in the Redis server. This function accepts several parameters:

- `otpLength` is the length of the OTP to be generated.
- `tokenExpirationMinutes` is the duration in minutes for which the OTP will expire.
- `appName` is the name of the application to be used in the OTP message.
- `targetName` is the name of the OTP recipient.
- `targetEmail` is the email address of the OTP recipient.

This function will generate an OTP, store it in the Redis server along with an expiration time calculated based on `tokenExpirationMinutes`, and then send the OTP to the specified email address.

### ValidateOTP

The `ValidateOTP` function is responsible for validating the given OTP for a specific email address. This function accepts two parameters:

- `email` is the email address to be validated.
- `otpToken` is the OTP to be validated.

This function will check whether the provided OTP matches the OTP stored in the Redis server for the given email address. If they match, the function will return `true`. If they do not match or an error occurs, the function will return `false` along with the error that occurred.

## Contribution

Contributions are welcome! You can contribute by submitting pull requests to this repository.

## License

This project is licensed under the [MIT License](LICENSE).
