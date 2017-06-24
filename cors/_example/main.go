package main

// $ go get github.com/rs/cors
// $ go run main.go

import (
	"github.com/go-siris/siris"
	"github.com/go-siris/siris/context"

	"github.com/go-siris/middleware/cors"
)

func main() {

	app := siris.New()
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
	})

	v1 := app.Party("/api/v1")
	v1.Use(crs)
	{
		v1.Get("/home", func(ctx context.Context) {
			ctx.WriteString("Hello from /home")
		})
		v1.Get("/about", func(ctx context.Context) {
			ctx.WriteString("Hello from /about")
		})
		v1.Post("/send", func(ctx context.Context) {
			ctx.WriteString("sent")
		})
	}

	// or use that to wrap the entire router
	// even before the path and method matching
	// this should work better and with all cors' features.
	// Use that instead, if suits you.
	// app.WrapRouter(cors.WrapNext(cors.Options{
	// 	AllowedOrigins:   []string{"*"},
	// 	AllowCredentials: true,
	// }))
	app.Run(siris.Addr("localhost:8080"))
}
