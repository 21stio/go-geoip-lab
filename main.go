package main

import (
	geoip2 "github.com/oschwald/geoip2-golang"
	"log"
	"net"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	db, err := geoip2.Open("geoip_database.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ip := net.ParseIP("195.82.154.225")
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(record)
}
