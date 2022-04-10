package main

import (
	"fmt"
	"go/go-lang-api/routes"
	"go/go-lang-api/database"
)

func main() {
	database.Conn()
	fmt.Println("Hello world com Go")
	routes.HandleRequest()
}
