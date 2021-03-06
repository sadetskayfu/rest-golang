// Package main point of entry
package main

import (
	"context"
	// "github.com/jackc/pgx/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sadetskayfu/rest-golang/internal/config"
	"github.com/sadetskayfu/rest-golang/internal/handler"
	"github.com/sadetskayfu/rest-golang/internal/model"
	"github.com/sadetskayfu/rest-golang/internal/repository"
	"github.com/sadetskayfu/rest-golang/internal/service"
	"github.com/sadetskayfu/rest-golang/pkg/logging"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// Initial logger
	log :=logging.NewLogger()
	log.SetLevel()

	// CONFIG POSTGRES DB
	// connStr := config.GetPostgresConnectionString()
	// CONNECT POSTGRES DB
	// conn, err := pgx.Connect(context.Background(), connStr)
	// if err != nil {
	// 	panic(err)
	// }

	// CONFIG MONGO DB
	connStrMongo := config.GetMongoConnectionString()
	clientOptions := options.Client().ApplyURI(connStrMongo)
	// CONNECT MONGO DB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
  
	// repo := repository.NewPostgresRepository(conn) //
	repo := repository.NewMongoRepository(client)
	srv := service.NewService(repo,log)
	hand := handler.NewHandler(srv)

	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ROUTE CUSTOMER
	e.POST("/customer", hand.CreateCustomer)
	e.GET("/customer", hand.GetAllCustomer)
	e.GET("/customer/:id", hand.GetByIdCustomer)
	e.DELETE("/customer/:id", hand.DeleteByIDCustomer)
	e.PUT("/customer/:id", hand.UpdateByIDCustomer)

	// ROUT Authentication
	e.POST("/registration", hand.Registration)
	e.POST("/auth", hand.LogIn)

	// Configure middleware with the custom claims type
	r := e.Group("/cat")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}))
	// PRIVATE ROUTE CAT
	r.POST("", hand.CreateCat)
	r.GET("/:id", hand.GetByIDCat)
	r.DELETE("/:id", hand.DeleteByIDCat)
	r.PUT("/:id", hand.UpdateByIDCat)
	r.GET("", hand.GetAllCat)
	// Start server
	e.Logger.Fatal(e.Start(":3001"))
}
