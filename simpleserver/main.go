package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// func main() {
// 	http.HandleFunc("/", rootHandler)
// 	log.Fatal(http.ListenAndServe(":8888", nil))
// }

// func rootHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hi %s", "back")
// }

func setupRoutes(app *fiber.App) {
	// Define and map a route
	app.Get("/", homePage)
}

func homePage(ctx *fiber.Ctx) error {
	// Create response
	msg := "Welcome to Web Server Using Fiber"
	// Send
	return ctx.SendString(msg)
}

func main() {
	// Create new fiber instance
	app := fiber.New()

	// Setup routes
	setupRoutes(app)

	// Start fiber app
	fmt.Println("Server is running on port :3000")
	app.Listen(":3000")
}
