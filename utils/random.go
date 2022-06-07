package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijkmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // return value 0 -> max-min because Int63n return a random integer between 0 and n-1 -> so if you plus min -> [min,max]
}

//RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// --- RANDOM ACCOUNTS ---

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// --- RANDOM TRANSFERS ---

func RandomAmout() int64 {
	return RandomInt(0, 20)
}

func RandomFromID() int64 {
	return RandomInt(0, 40)
}
func RandomToID() int64 {
	return RandomInt(40, 80)
}
