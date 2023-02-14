package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/mjpancheri/imersao-full-stack-2023-02/simulator/infra/kafka"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	producer := kafka.NewKafkaProducer()
	kafka.Publish("Olá", "readtest", producer)
	for {
		_ = 1
	}
	// route := route2.Route{
	// 	ID:       "1",
	// 	ClientID: "1",
	// }
	// route.LoadPositions()
	// stringJson, _ := route.ExportJsonPositions()
	// fmt.Println(stringJson[0])
}
