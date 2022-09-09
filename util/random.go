package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Generate random number
func RandomInt(min, max int64) int64 {
	return rand.Int63n(max - min + 1)
}

// Generate random string where letters are picked up randomly from the aplhabet
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Generate random name for owner
func RandomOwner() string {
	return RandomString(6)
}

// Generate random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// Select currency randomly
func RandomCurrency() string {
	currencies := []string{EUR, USD, CAD}
	n := len(currencies)

	return currencies[rand.Intn(n)]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
