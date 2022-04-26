package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/Msksgm/go-itddd-12-aggregate/application"
	"github.com/Msksgm/go-itddd-12-aggregate/domain/model/circle"
	"github.com/Msksgm/go-itddd-12-aggregate/infrastructure/persistence"
)

func main() {
	uri := fmt.Sprintf("postgres://%s/%s?sslmode=disable&user=%s&password=%s&port=%s&timezone=Asia/Tokyo",
		os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"))
	db, err := sql.Open("postgres", uri)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	log.Println("successfully connected to database")

	circleRepository, err := persistence.NewCircleRepository(db)
	if err != nil {
		panic(err)
	}
	circleService, err := circle.NewCircleService(circleRepository)
	if err != nil {
		panic(err)
	}
	circleApplicationService, err := application.NewCircleApplicationService(circleRepository, *circleService)
	if err != nil {
		panic(err)
	}

	if err := circleApplicationService.Register("test-circle-name"); err != nil {
		log.Println(err)
	}
}
