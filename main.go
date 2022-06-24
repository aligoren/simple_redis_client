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

	client, err := redisConfig.Connect()
	if err != nil {
		log.Fatalf("Error while connecting: %v", err)
	}

	_, err = client.Set("DENEME", "DEGERI")
	if err != nil {
		log.Fatalf("Error while set value: %v", err)
	}

	/*info, err := client.Info()

	fmt.Printf("Info: %s\n", info.Message)*/

	/*getKey, _ := client.Get("DENEME")

	fmt.Printf("Key: %s\n", getKey.Message)*/

	/*getKey, err := client.Get("DENEME")
	if err != nil {
		log.Fatalf("Err: %v", err)
	}*/

	arrResponse, _ := client.InsertArray("falanca", "x", "z", "c", "d")

	fmt.Printf("Array response: %v", arrResponse)

}
