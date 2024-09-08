package utilities

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz"

func GenerateMoniker() (string, error) {
	var length, err = strconv.Atoi(os.Getenv("CAMPAIGN_MONIKER_LENGTH"))
	if err != nil {
		// If there is an error in conversion, log the error and handle it appropriately
		log.Fatalf("Error converting CAMPAIGN_MONIKER_LENGTH to integer: %v", err)
	}
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b), err
}
