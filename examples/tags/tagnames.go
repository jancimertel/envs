package main

import (
	"fmt"
	"github.com/jancimertel/envs"
	"os"
)

type EnvsTagBundle struct {
	VAR1 bool    `envs:"TAG1"`
	VAR2 string  `envs:"TAG2"`
	VAR3 int     `envs:"TAG3"`
	VAR4 float32 `envs:"TAG4"`
}

func main() {
	os.Setenv("TAG1", "1")
	os.Setenv("TAG2", "test")
	os.Setenv("TAG3", "123")
	os.Setenv("TAG4", "0.2")

	bundle := EnvsTagBundle{}
	if err := envs.MustHave(&bundle); err == nil {
		fmt.Println(fmt.Sprintf("%v, %v, %v, %v", bundle.VAR1, bundle.VAR2, bundle.VAR3, bundle.VAR4))
	}
}
