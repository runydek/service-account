package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRekeningNumber() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("REK-%06d", rand.Intn(1000000))
}
