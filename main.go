package main

import (
	"flover-market/db"
	"flover-market/env"
	"flover-market/web"
	"log"
	"os"
)

func main() {
	env.Urldb = os.Getenv("DATABASE_URL")
	env.AdminPhone = os.Getenv("ADMIN_PHONE")
	env.Port = os.Getenv("PORT")
	//urldb := "postgresql://gardenia:BPDTnxIcT7qmoRQrL4ILmhrsfPTZaSy0@dpg-cprddgdumphs73c3i4sg-a.oregon-postgres.render.com/gardenia"

	// Connect to database
	err := db.Connect(env.Urldb)
	if err != nil {
		log.Fatal(err)
		return
	}

	if err = db.Init(); err != nil {
		panic(err)
	}

	// Start web server
	web.StartGin(":" + env.Port)
}
