package main

import (
	"log"
	"places/config"
	"places/db"
	"places/internal/repository"
)

func main() {
	log.Printf("Started app\n")

	es_store, err := db.ConnectWithElasticSearch()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected with Elastic Search")
	log.Println(es_store.ClassicClient, es_store.TypedClient)

	data, err := repository.ParsePlacesFromCsv(config.PlacesFilePath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Parsed places from CSV")

	_, err = es_store.InsertPlaces(data)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("success upload files")

	log.Println(es_store.ClassicClient.Get("places", "id"))
	log.Println(es_store.ClassicClient.Get("places", "1"))

	config.ConfigServer()

}
