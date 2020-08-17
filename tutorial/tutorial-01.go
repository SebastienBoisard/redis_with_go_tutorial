package tutorial

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

func PlayTutorial01() {

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

	fmt.Println("Send 'PING' to Redis")
	// Send a Redis command across the connection.
	// The first parameter to Do() is always the name of the Redis command (here, the `PING` command,
	// Cf. https://redis.io/commands/ping) optionally followed by any necessary arguments.
	result, err := conn.Do("PING")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Redis replied with '%s'\n\n", result)

	fmt.Println("Send 'PING \"hello\"' to Redis")
	// Send another `PING`command but with a "hello" message, so Redis can reply with this message.
	result, err = conn.Do("PING", "hello")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Redis replied with '%s'\n\n", result)
}
