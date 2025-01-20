package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"baserouter"
)

func main() {
	htmlFiles := html.New("./static/html", ".html")
	htmlFiles.Reload(true)

	// aşşağıda yeni bir fiber serveri hasıl ediyoruz.
	server := fiber.New(fiber.Config{
		Views:       htmlFiles,
		ViewsLayout: "layouts/main",
	})

	server.Static("/css", "./static/css")
	server.Static("/js", "./static/js")
	server.Static("/files", "./static/files")

	fmt.Printf("Merhaba! \n")

	baserouter.BaseRouter(server)

	server.Listen(":2000")
}
