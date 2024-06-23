package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var (
	DbClient *pgx.Conn
)

func init() {
	if err := Initialize(); err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	log.Println("successfully connected")
}

func Initialize() error {

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	fmt.Println("DBURL:",dbURL)
	
	//create connection
	DbClient, err = pgx.Connect(context.Background(), dbURL)
	if err != nil {
		return err
	}
	

	return nil
}
