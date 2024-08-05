package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/miletic94/fm-language/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hei %s! This is fm programming language \n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
