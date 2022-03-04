package main

import "kedamaonsen/react-recipe-app-api/internal/infra/router"

func main() {

	r := router.CreateRouter()
	r.StartApp()
}
