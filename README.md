# Redis with Go tutorial
A simple Redis tutorial with Go




## Run Redis

For this tutorial, we'll use Redis from a Docker image.


Pull the latest redis image from the docker hub repository (Cf. https://hub.docker.com/_/redis)
```
> docker pull redis
```

Run a detached instance of the Redis image ( named `redis-tutorial`) on port 6379
```
> docker run -d --name redis-tutorial -p 6379:6379 redis
```

Test the Redis server directly on the docker instance with the redis-cli
```
> docker container exec -ti redis-tutorial redis-cli PING
PONG
```

You can also install the redis-cli... 
```
> sudo apt-get install redis-tools
```

and play with the Redis server on Docker from your OS
```
> redis-cli PING
PONG
```


## Redigo

There are several Go clients for Redis (Cf. https://redis.io/clients#go) and 2 highly maintained projects have more stars and contributors than the others:
  - Redigo (7.7k stars and 57 contributors) - https://github.com/gomodule/redigo
  - Go Redis (9.4k stars and 137 contributors) - https://github.com/go-redis/redis
Strangely, only Redigo is approved by Redis. Maybe because it uses a print-like API, whereas Go-Redis encapsulates all the Redis commands in Go functions.

Nevertheless, this tutorial is based on Redigo.


## Connect to Redis

The connection is established through the Dial() function 

```
	// Establish a connection to the Redis server listening on port 6379 of the docker instance 'redis-tutorial'.
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}

	// Use `defer` to ensure the connection is always properly closed before exiting the main() function.
	defer conn.Close()

	fmt.Println("Sent 'PING' to Redis")
	// Send a Redis command across the connection.
	// The first parameter to Do() is always the name of the Redis command (here, the `PING` command,
	// Cf. https://redis.io/commands/ping) optionally followed by any necessary arguments.
	result, err := conn.Do("PING")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Redis replied with '%s'\n\n", result)
```

Run the first tutorial to see the result
```
go run tutorial-01.go
```
