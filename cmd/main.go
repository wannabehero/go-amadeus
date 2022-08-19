package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/wannabehero/go-amadeus/api"
	"github.com/wannabehero/go-amadeus/requests"
	"github.com/wannabehero/go-amadeus/types"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type AvailabilityAPIRequest struct {
	Hotels     []string         `json:"hotels"`
	CheckIn    string           `json:"checkIn"`
	CheckOut   string           `json:"checkOut"`
	Currency   string           `json:"currency"`
	Adults     int              `json:"adults"`
	InfoSource types.InfoSource `json:"infoSource"`
}

func checkAvailabilityHandler(config types.AmadeusConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request AvailabilityAPIRequest
		if err := c.BodyParser(&request); err != nil {
			return err
		}

		envelope, action := requests.NewAvailabilityRequest(
			request.InfoSource,
			request.CheckIn,
			request.CheckOut,
			request.Currency,
			"US",
			request.Adults,
			request.Hotels,
			nil,
			config,
		)

		data, err := amadeusClient.SendRequest(action, envelope)
		if err != nil {
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		c.Set("content-type", "text/xml")
		return c.Send(data)
	}
}

type GetDescriptionAPIRequest struct {
	Hotels []string `json:"hotels"`
}

func getDescriptionHandler(config types.AmadeusConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request GetDescriptionAPIRequest
		if err := c.BodyParser(&request); err != nil {
			return err
		}

		envelope, action := requests.NewDescriptiveInfoRequest(request.Hotels, nil, config)
		data, err := amadeusClient.SendRequest(action, envelope)
		if err != nil {
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		c.Set("content-type", "text/xml")
		return c.Send(data)
	}
}

var amadeusClient *api.AmadeusAPIClient

func main() {
	config := types.LoadConfigFromFile(os.Getenv("CONFIG_FILE"))
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Use(logger.New())

	amadeusClient = api.NewClient(config)

	gracefully(func() {
		app.Shutdown()
	})

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "3000"
	}

	app.Post("/api/availability", checkAvailabilityHandler(config))
	app.Post("/api/description", getDescriptionHandler(config))

	log.Printf("Starting server at %v", port)
	app.Listen(fmt.Sprintf("0.0.0.0:%v", port))
}

func gracefully(block func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		block()
	}()
}
