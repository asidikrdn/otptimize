# OTPtimize

OTPtimize adalah sebuah package Go yang membantu dalam menghasilkan, mengirim, dan memvalidasi OTP (One-Time Password). Package ini dirancang untuk memudahkan penggunaan OTP dalam aplikasi Anda dengan fokus pada pengoptimalan penggunaan sumber daya.

## Instalasi

Anda dapat menginstal package OTPtimize dengan perintah:

```bash
go get -u github.com/asidikrdn/otptimize
```

## Penggunaan

Berikut adalah contoh penggunaan dari package OTPtimize:

```go
package main

import (
	"fmt"
	"github.com/asidikrdn/otptimize"
	"time"
)

func main() {
	// Inisialisasi koneksi
	mailConfig := otptimize.MailConfig{ /* Isi dengan konfigurasi email */ }
	redisConfig := otptimize.RedisConfig{ /* Isi dengan konfigurasi Redis */ }
	otptimize.ConnectionInit(mailConfig, redisConfig)

	// Generate dan kirim OTP
	err := otptimize.GenerateAndSendOTP(6, 5, "MyApp", "John Doe", "johndoe@example.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Validasi OTP
	valid, err := otptimize.ValidateOTP("johndoe@example.com", "123456")
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

Pastikan untuk mengganti nilai-nilai konfigurasi email dan Redis sesuai dengan kebutuhan Anda.

## Dokumentasi

### ConnectionInit(mailConfig MailConfig, redisConfig RedisConfig)

Inisialisasi koneksi dengan konfigurasi email dan Redis yang diberikan.

### GenerateAndSendOTP(otpLength int, appName string, targetName string, targetEmail string) error

Generate dan kirim OTP ke alamat email yang ditentukan. OTP akan disimpan dalam Redis.

### ValidateOTP(email string, otpToken string) (bool, error)

Memvalidasi OTP yang diberikan untuk alamat email tertentu.

## Kontribusi

Kontribusi dipersilakan! Anda dapat membantu dengan mengajukan _pull request_ pada repositori ini.

## Lisensi

Proyek ini dilisensikan di bawah [MIT License](LICENSE).
