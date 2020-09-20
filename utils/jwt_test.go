package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken(1, []int64{1, 2})
	fmt.Println(token, err)

	time.Sleep(1 * time.Second)

	fmt.Println(ValidToken(token, 1))
}
