package main

import (
	"fmt"
	"log"
	"main/db"
	"main/handler"
	"main/interfaces"
	"main/service"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {

	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "root:@tcp(127.0.0.1:3306)/adserver"
	}

	if err := db.InitDB(dsn); err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	db.DB.SetMaxOpenConns(100)
	db.DB.SetMaxIdleConns(20)
	db.DB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB := db.DB
	repo := interfaces.NewSQLCampaignPersistence(sqlDB)
	service := service.NewCampaignService(repo)
	http.HandleFunc("/v1/delivery", handler.DeliveryHandler(*service))

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
