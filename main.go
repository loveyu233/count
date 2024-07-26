package main

import (
	"count/handler"
	"count/model"
	"count/utils/pgsql"
	"github.com/gofiber/fiber/v2"
)

func main() {
	dsn := "host=pgsqlcount user=pgsql password=pgsql dbname=pgsql port=5432 sslmode=disable"
	pgsql.InitPgsql(dsn)
	model.InitModel()
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")                   // 允许所有来源访问
		c.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // 允许的请求方法
		c.Set("Access-Control-Allow-Headers", "Content-Type")       // 允许的请求头

		// 检查是否是预检请求（OPTIONS），如果是直接返回
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusNoContent)
		}

		// 继续处理请求
		return c.Next()
	})
	app.Get("/all", handler.GetAll())
	app.Post("/add", handler.Add())
	app.Listen(":9999")
}
