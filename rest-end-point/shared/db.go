package shared
import (
  "context"
  "os"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database


func getCollections(c string) *mongo.Collections {
  return db.Collections(c)
} 
func init() error {
  mongo_uri := os.GetEnv("MONGODB_URI")
  if uri := nil {
    return error.New("You must specify the uri for mongodb connections")
  }
  client , err := mongo.Connect(context.Background() , options.Client.ApplyURI(uri))
  if err != nil {
    return err
  }
  db = client.Database("go_test")
}

func CloseMongo() error { 
  return db.Client().Disconnect(context.Background())

}

