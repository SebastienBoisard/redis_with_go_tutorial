package tutorial

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

func PlayTutorial03() {

	// Establish a connection to the Redis server listening on port 6379 of the docker instance 'redis-tutorial'.
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}

	// Use `defer` to ensure the connection is always properly closed before exiting the main() function.
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Add a key to hold a int value (Cf. https://redis.io/commands/set)
	fmt.Println(`Send 'SET integer_key 1' to Redis`)
	result, err := conn.Do("SET", "integer_key", 4)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Redis replied with '%s'\n\n", result)

	// Retrieve the integer value of the previously defined key (Cf. https://redis.io/commands/get)
	fmt.Println(`Send 'GET integer_key' to Redis`)
	result, err = redis.Int(conn.Do("GET", "integer_key"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Redis replied with '%d'\n\n", result)

	// Increment the integer stored at `integer_key` (Cf. https://redis.io/commands/incr)
	fmt.Println(`Send 'INCR integer_key' to Redis`)
	result, err = redis.Int(conn.Do("INCR", "integer_key"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Redis replied with '%d'\n\n", result)

	// Increment the integer stored at `integer_key` by 10 (Cf. https://redis.io/commands/incrby)
	fmt.Println(`Send 'INCRBY integer_key 10' to Redis`)
	result, err = redis.Int(conn.Do("INCRBY", "integer_key", 10))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Redis replied with '%d'\n\n", result)

	// Increment the integer stored at `integer_key` (Cf. https://redis.io/commands/incr)
	fmt.Println(`Send 'DECR integer_key' to Redis`)
	result, err = redis.Int(conn.Do("DECR", "integer_key"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Redis replied with '%d'\n\n", result)

	// Decrement the integer stored at `integer_key` by 5 (Cf. https://redis.io/commands/incrby)
	fmt.Println(`Send 'DECRBY integer_key 5' to Redis`)
	result, err = redis.Int(conn.Do("DECRBY", "integer_key", 5))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Redis replied with '%d'\n\n", result)
}
