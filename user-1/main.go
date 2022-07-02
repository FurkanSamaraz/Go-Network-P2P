package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/perlin-network/noise"
)

var denee = make(chan string, 0)

func Gon(c *fiber.Ctx) error {
	adres := c.FormValue("adres")
	mesaj := c.FormValue("mesaj")

	furkan, err := noise.NewNode()
	check(err)

	c.Response().Header.Set("Content-Type", "application/json")

	defer furkan.Close()

	check(furkan.Listen())

	res, err := furkan.Request(context.TODO(), adres, []byte(mesaj))
	check(err)

	fmt.Printf("Gelen Mesaj(Eren) : '%s'\n", string(res))
	return c.JSON(furkan.Addr())
}
func Al(c *fiber.Ctx) error {
	furkan, err := noise.NewNode()
	mesaj := c.FormValue("mesaj")

	check(err)

	defer furkan.Close()

	furkan.Handle(func(ctx noise.HandlerContext) error {
		if !ctx.IsRequest() {
			return nil
		}

		fmt.Printf("Eren'den mesaj var: '%s'\n", string(ctx.Data()))

		return ctx.Send([]byte(mesaj))
	})

	check(furkan.Listen())

	check(err)

	fmt.Println("adress", furkan.Addr())
	return c.JSON(err)
}

func main() {
	app := fiber.New()
	app.Post("/Al", Al)
	app.Post("/Gonder", Gon)
	app.Listen(":8080")

}
func check(err error) {
	if err != nil {
		panic(err)
	}
}
