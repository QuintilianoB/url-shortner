package util

import (
	"database/sql"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
	"unsafe"
)

// Constants used in the random string generator.
const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func GetEnv(key string, defaultVal string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultVal
	}
	return value
}

func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func DbClose(db *sql.DB) {
	err := db.Close()

	if err != nil {
		log.Println("Error while closing a DB connection.")
	}
}

func HttpClose(resp *http.Response) {

	err := resp.Body.Close()
	if err != nil {
		log.Println("Error while closing a Http connection.")
	}
}

/*
	RandString generates a string of fixed size n. Used to generate the shortened URL.
	https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
*/
func RandString(n int) string {
	src := rand.NewSource(time.Now().UnixNano())

	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
