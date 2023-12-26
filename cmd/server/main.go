package main

// simple auth server for emqx mqttsn-gateway using echo
import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ryodeushii/mqttsn-bombardier/utils"
)

func main() {
	log := utils.NewLogger()
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	e.POST("*", func(c echo.Context) error {
		jsonmap := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&jsonmap)
		if err != nil {
			return c.JSON(400, map[string]interface{}{
				"error": "invalid json",
			})
		}

		log.Info("Req body: ", jsonmap)

		return c.JSON(200, map[string]interface{}{
			"result":       "allow", // allow, deny, ignore
			"is_superuser": false,
		})
	})
	err := e.Start("0.0.0.0:42069")
	if err != nil {
		log.Panic("Failed to start server", err)
	}

}
