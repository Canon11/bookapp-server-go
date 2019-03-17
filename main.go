package main

import (
	"bookapp/router"
)

func main() {
	r := router.GetRouter()
	r.Run()
}
