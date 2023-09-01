package router


import (
	"github.com/Kinfe123/rest-end-point/shared"
	"github.com/Kinfe123/rest-end-point/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func AddBookGroup(app *fiber.App){
  // getting the group of book routes
  bookGroup := app.Group("/books")
  // all the below routes will be prefixed by /books

  bookGroup.Get("/" , getBooks)
  bookGroup.Get("/:id" , getBook)
  bookGroup.Post("/" , createBook)
  bookGroup.Put("/:id" , updateBook)
  bookGroup.Delete("/:id" , deleteBook)
}


// for getting a lists of books
// we use pointer since we need to take the 
// the reference rather copying the data which really costs

func getBooks(c *fiber.Ctx) error {
   collection = common.GetCollection("books")
   

}

