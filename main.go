package main

import (
	"net/http"

	"github.com/apaarshrm39/learningGO/controllers"
)

func main() {
	controllers.RegisterContollers()
	http.ListenAndServe(":3000", nil)
}
