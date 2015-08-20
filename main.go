package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/twanies/weavebox-docker-example/Godeps/_workspace/src/github.com/twanies/weavebox"
)

func main() {
	var listen = flag.Int("listen", 3000, "target listen address of the application")
	flag.Parse()

	app := weavebox.New()
	app.Get("/", helloHandler)
	log.Fatal(app.Serve(*listen))
}

func helloHandler(ctx *weavebox.Context) error {
	return ctx.Text(http.StatusOK, "Hello you")
}
