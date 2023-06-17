package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kangata/p100-api-go/services"
)

func GetDevice(c *fiber.Ctx) error {
	device, err := services.GetDeviceInstance()

	if err != nil {
		c.Status(fiber.StatusServiceUnavailable)

		return c.JSON(fiber.Map{
			"code":    fiber.ErrServiceUnavailable.Code,
			"message": fiber.ErrServiceUnavailable.Message,
		})
	}

	status, err := device.GetDeviceInfo()

	if err != nil {
		c.Status(fiber.StatusServiceUnavailable)

		return c.JSON(fiber.Map{
			"code":    fiber.ErrServiceUnavailable.Code,
			"message": fiber.ErrServiceUnavailable.Message,
		})
	}

	return c.JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Ok",
		"data":    status.Result,
	})
}

func SwitchDeviceStatus(c *fiber.Ctx) error {
	device, err := services.GetDeviceInstance()

	if err != nil {
		c.Status(fiber.StatusServiceUnavailable)

		return c.JSON(fiber.Map{
			"code":    fiber.ErrServiceUnavailable.Code,
			"message": fiber.ErrServiceUnavailable.Message,
		})
	}

	type Body struct {
		Status bool `json:"status"`
	}

	body := new(Body)

	if err := c.BodyParser(body); err != nil {
		c.Status(fiber.StatusBadRequest)

		return c.JSON(fiber.Map{
			"code":    fiber.ErrBadRequest.Code,
			"message": fiber.ErrBadRequest.Message,
		})
	}

	if err := device.Switch(body.Status); err != nil {
		c.Status(fiber.StatusServiceUnavailable)

		return c.JSON(fiber.Map{
			"code":    fiber.ErrServiceUnavailable.Code,
			"message": fiber.ErrServiceUnavailable.Message,
		})
	}

	return c.JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "OK",
	})
}
