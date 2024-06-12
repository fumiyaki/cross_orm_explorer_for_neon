package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fumiyaki/cross_orm_explorer_for_neon/src/gorm/generated"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connStr, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		fmt.Println("DATABASE_URL is not set")
	}

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	query := generated.Use(db)

	ctx := context.Background()
	userData, err := query.User.WithContext(ctx).Where(query.User.ID.Eq(1)).First()
	if err != nil {
		log.Fatal("failed to get user data:", err)
	}

	fmt.Printf("User Data: %+v\n", userData)
}
