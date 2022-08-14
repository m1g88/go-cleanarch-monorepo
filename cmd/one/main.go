package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/m1g88/go-cleanarch-monorepo/pkg/server"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	log.Printf("db connected")

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		collection := client.Database("testing").Collection("numbers")

		ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		b := make([]byte, 10)
		_, err := rand.Read(b)
		if err != nil {
			c.JSON(400, err)
		}
		res, err := collection.InsertOne(ctx, bson.D{{"name", "random"}, {"value", hex.EncodeToString(b)}})
		if err != nil {
			c.JSON(400, err)
		}
		id := res.InsertedID

		c.JSON(200, gin.H{"id": id, "data": "hello world"})
	})

	srv := server.New("3000", r)
	srv.RunWithGracefulShutdown()
}
