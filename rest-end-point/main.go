package main

import (
	"os"

	"github.com/Kinfe123/rest-end-point/shared"
	"github.com/bmdavis419/rest-end-point/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)


func main(){
  app := fiber.New()
  app.Get("/" , func(c *fiber.Ctx) error{
    return c.SendString("Hello world")
  })
  fmt.Printf("The server is started at: " , 3000)
  app.Listen(":3000")



}
func runnable() error {
  err := common.LoadEnv()
  if err != nil {
    return err
  }
  // start the mongoDb
  err := common.InitDB()
  if err != nil {
    return err
  }
  defer common.CloseMongo()
  
  app := fiber.New()
  app.Use(logger.New())
  app.Use(recover.New())
  app.Use(cors.New())


  // routes
   var port string = os.Getenv("PORT") 
   if port == ""{
    port = 8080
  }
  fmt.Printf("The port is running at " , port)
  app.Listen(":" + port)

  
}


