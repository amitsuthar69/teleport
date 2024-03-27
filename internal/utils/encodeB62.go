package encode

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

func Base62(id int64) string {
	if id == 0 {
		return "0"
	}

	base62Chars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	var hash string
	salt, err := strconv.Atoi(os.Getenv("SALT"))
	if err != nil {
		return ""
	}

	number := id * int64(salt)
	for number > 0 {
		rem := number % 62
		hash = string(base62Chars[rem]) + hash
		number /= 62
	}

	return hash
}
