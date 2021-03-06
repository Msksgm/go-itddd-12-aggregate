package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/msksgm/go-itddd-12-aggregate/application"
	"github.com/msksgm/go-itddd-12-aggregate/domain/model/circle"
	"github.com/msksgm/go-itddd-12-aggregate/infrastructure/persistence"
)

var command = flag.String("usecase", "", "usercase of application")

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

	flag.Parse()
	log.Println(*command)
	switch *command {
	case "register":
		if err := circleApplicationService.Register("test-circle-name"); err != nil {
			log.Println(err)
		}
	case "get":
		circleData, err := circleApplicationService.Get("test-circle-name")
		if err != nil {
			log.Println(err)
		}
		log.Println(circleData)
	default:
		log.Printf("%s is not command. choose in ('register', 'get', 'update', 'delete')", *command)
	}
}
