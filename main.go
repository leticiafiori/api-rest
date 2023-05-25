package main

import (
	"fmt"

	"github.com/leticiafiori/routes"
)

func main() {
	fmt.Println("Iniciando o servidor rest com Go")
	routes.HandleRequest()
}
