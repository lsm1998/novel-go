package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	start := time.Now().Unix()
	token, err := GenerateToken(1, []int64{1, 2})
	end := time.Now().Unix()
	fmt.Println(token, err)
	time.Sleep(1 * time.Second)
	fmt.Println(ValidToken(token, 1))
	fmt.Println(end - start)
}
