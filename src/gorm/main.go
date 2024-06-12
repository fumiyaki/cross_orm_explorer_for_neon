package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
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
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	g := gen.NewGenerator(gen.Config{
		OutPath: "src/gorm/generated",
	})

	// データベースから構造体を生成
	g.UseDB(db)
	g.ApplyBasic(
		g.GenerateModel("Post"),
		g.GenerateModel("Profile"),
		g.GenerateModel("Category"),
		g.GenerateModel("User"),
		g.GenerateModel("Tag"),
	)

	// 構造体の生成
	g.Execute()
}
