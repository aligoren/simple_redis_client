package main

import (
	"fmt"
	"log"
)

func main() {

	redisConfig := &RedisConfig{
		Address:  "localhost",
		Port:     49153,
		Password: "redispw",
	}

	connect, err := redisConfig.Connect()
	if err != nil {
		log.Fatalf("Error while connecting: %v", err)
	}

	_, err = connect.Auth()
	if err != nil {
		log.Fatalf("Error while auth: %v", err)
	}

	response, err := connect.Set("DENEME", "DEGERI")
	if err != nil {
		log.Fatalf("Error while set value: %v", err)
	}

	fmt.Printf("Response: %s", response.Message)

}
