package main

import (
	"fmt"
	"github.com/jancimertel/envs"
	"os"
)

type EnvsBundle struct {
	VAR1 bool
	VAR2 string
	VAR3 int
	VAR4 float32
}

func main() {
	os.Setenv("VAR1", "1")
	os.Setenv("VAR2", "test")
	os.Setenv("VAR3", "123")
	os.Setenv("VAR4", "0.2")

	bundle := EnvsBundle{}
	if err := envs.MustHave(&bundle); err == nil {
		fmt.Println(fmt.Sprintf("%v, %v, %v, %v", bundle.VAR1, bundle.VAR2, bundle.VAR3, bundle.VAR4))
	}
}
