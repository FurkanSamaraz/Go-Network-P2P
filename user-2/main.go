package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/perlin-network/noise"
)

func Al(c *fiber.Ctx) error {
	Eren, err := noise.NewNode()
	check(err)
	mesaj := c.FormValue("mesaj")

	defer Eren.Close()
	Eren.Handle(func(ctx noise.HandlerContext) error {
		if !ctx.IsRequest() {
			return nil
		}

		fmt.Printf("Furkan'dan mesaj var: '%s'\n", string(ctx.Data()))

		return ctx.Send([]byte(mesaj))
	})

	check(Eren.Listen())

	check(err)
	fmt.Println("addres", Eren.Addr())

	return c.JSON(err)
}
func Gon(c *fiber.Ctx) error {
	adres := c.FormValue("adres")
	mesaj := c.FormValue("mesaj")
	Eren, err := noise.NewNode()
	check(err)

	defer Eren.Close()

	check(Eren.Listen())

	res, err := Eren.Request(context.TODO(), adres, []byte(mesaj))
	check(err)

	fmt.Printf("Gelen Mesaj(Furkan) : '%s'\n", string(res))
	return c.JSON(err)
}
func main() {
	app := fiber.New()
	app.Post("/Al", Al)
	app.Post("/Gonder", Gon)

	app.Listen(":8000")

}
func check(err error) {
	if err != nil {
		panic(err)
	}
}
