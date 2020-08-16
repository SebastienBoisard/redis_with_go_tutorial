package tutorial

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

func PlayTutorial02() {

	// Establish a connection to the Redis server listening on port 6379 of the docker instance 'redis-tutorial'.
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}

	// Use `defer` to ensure the connection is always properly closed before exiting the main() function.
	defer conn.Close()


	// Add a key to hold a string value (Cf. https://redis.io/commands/set)
	fmt.Println(`Send 'SET my_first_key "my_first_value"' to Redis`)
	result, err := conn.Do("SET", "my_first_key", "my_first_value")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Redis replied with '%s'\n\n", result)

	// Retrieve the value of the previously defined key (Cf. https://redis.io/commands/get)
	fmt.Println(`Send 'GET my_first_key' to Redis`)
	result, err = conn.Do("GET", "my_first_key")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Redis replied with '%s'\n\n", result)

	// Test how Redis responds when we try to retrieve an unknown key.
	fmt.Println(`Send 'GET another_key' to Redis`)
	result, err = conn.Do("GET", "another_key")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Redis replied with '%#v' (because 'another_key' doesn't exist in Redis)\n\n", result)

	// Send a command to check if a key exists (Cf. https://redis.io/commands/exists)
	fmt.Println(`Send 'EXISTS my_first_key' to Redis`)
	result, err = conn.Do("EXISTS", "my_first_key")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Redis replied with '%d' (which means this key exists)\n\n", result)

	// Send a command to check if a key exists (Cf. https://redis.io/commands/exists)
	fmt.Println(`Send 'EXISTS another_key' to Redis`)
	result, err = conn.Do("EXISTS", "another_key")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Redis replied with '%d' (which means this key doesn't exist)\n\n", result)
}
