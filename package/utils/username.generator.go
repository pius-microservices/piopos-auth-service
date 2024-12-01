package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func GenerateUsername(email string) string {
	localPart := email[:len(email)-len("@"+email[strings.LastIndex(email, "@")+1:])]

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	randomDigits := r.Intn(10000000)
	username := localPart + strconv.Itoa(randomDigits)
	return username
}
