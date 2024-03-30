package main

import (
	"go-mongo/src/db"
	"go-mongo/src/dbaccess"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

func main() {
	setupEnv()
	setupLogging()
	setupDB()
	//testing DB Query
	ordersCRUD()
	validatorCreate()
}

func setupEnv() {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
		panic(err)
	}

}

func setupLogging() {
	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("os.OpenFile error:", err)
	}

	log.SetOutput(file)
}

func setupDB() {
	db := db.Init()
	dbaccess.SetupService(db.Client, db.Ctx, db.CtxCancel, db.Database)
}

func ordersCRUD() {
	// create
	order := dbaccess.Order{
		CurrentNode:     1,
		OrderHeight:     1,
		TimeSinceSwitch: time.Now(),
	}
	output, err := dbaccess.CreateOrder(&order)
	if err != nil {
		log.Error("err", err)
	}
	log.Debug("Order Created", output)

	// find all orders
	orders := dbaccess.FindOrders(&dbaccess.Order{CurrentNode: 1})
	log.Debug("orders", orders)
}

func validatorCreate() {
	validator := dbaccess.Validator{
		NodeOrder:    1,
		Moniker:      "xy1s",
		PubKey:       "asda1sdas",
		LastPingTime: 122,
		IP:           "127.0.0.11",
	}
	output, err := dbaccess.CreateValidator(&validator)
	if err != nil {
		log.Error("err", err)
	}
	log.Debug("Validator Created", output)
}
