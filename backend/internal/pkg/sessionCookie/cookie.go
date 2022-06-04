package sessionCookie

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func Get(c *fiber.Ctx, key string) string {
	return c.Cookies(key)
}

func Set(c *fiber.Ctx, name, val string) {
	cookie := fiber.Cookie{Name: name, Value: val, Expires: time.Now().Add(30 * 24 * time.Hour)}
	c.Cookie(&cookie)
}

func Clear(c *fiber.Ctx, name string) {
	c.ClearCookie(name)
}
