package main

import (
	"fmt"
	"net/http"

	"github.com/honkytonktown/GoProject/controllers"
)

func main() {
	controllers.RegisterControllers()
	fmt.Println("Listening at http://localhost:3000/")
	http.ListenAndServe(":3000", nil)
}
