package env

import (
	"fmt"
	"testing"
)

func TestGenerateEnv(t *testing.T) {
	list := GenerateEnv("env_fabric_peer","./")
	fmt.Println(list)
}