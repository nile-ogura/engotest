package main

import (
	"fmt"

	"github.com/nile-ogura/engotest/internal/testapp/logic"
)

func main() {
	human := logic.NewHuman("hoge")
	fmt.Printf("Hello World %s\n", human.Name)
}
