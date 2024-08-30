package main

import (
	"crider/technical/request"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		const (
			MOTIVE_URL   = "https://api.gomotive.com/oauth/authorize"
			CLIENT_ID    = "98a670ed21a9b27a7e104160d61d51396577283d942b630202e12557a39a76f4"
			REDIRECT_URI = "https://eovvvgjxrp54hso.m.pipedream.net"
		)

		motiveURL := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&response_type=code&scope=assets.read users.read vehicles.read", MOTIVE_URL, CLIENT_ID, REDIRECT_URI)
		return c.Redirect(motiveURL)
	})

	app.Get("/callback", func(c *fiber.Ctx) error {
		access_token := c.Query("access_token")
		if access_token == "" {
			return c.Status(fiber.StatusBadRequest).SendString("No code provided")
		}

		request.HandleRequest(access_token)

		return c.Redirect("https://pipedream.com/@axle-interview/projects/proj_ELsRy5R/crider-data-responses-p_yKCLpDg/inspect")
	})

	err = app.Listen(":" + os.Getenv("PORT"))

	log.Fatal("Error starting server: ", err)
}
