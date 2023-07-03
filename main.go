package main

import (
	"pin"
)

func main() {

	r := pin.New()
	r.Get("/", func(ctx *pin.Context) {
		name := ctx.Query("name")
		ctx.JSON(200, name)
	})

	r.Run("localhost:8080")
}
