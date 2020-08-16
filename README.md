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

```go
// Establish a connection to the Redis server listening on port 6379 of the docker instance 'redis-tutorial'.
conn, err := redis.Dial("tcp", "localhost:6379")
if err != nil {
	log.Fatal(err)
}

// Use `defer` to ensure the connection is always properly closed before exiting the main() function.
defer conn.Close()

fmt.Println("Send 'PING' to Redis")
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
go run main.go 1
```

## Set and Get a value

Redis can add and retrieve string value from keys with the commands ["SET"](https://redis.io/commands/set) 
and ["GET"](https://redis.io/commands/get)

```go
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
```

The ["EXISTS"](https://redis.io/commands/exists) command can check if a key exists or not.

```go
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
```

Run the second tutorial to see the result
```
go run main.go 2
```

----

Note for setting a password to access the Redis Server

You can define the password while running the docker instance:

```bash
> docker container run --name some-redis -d -e REDIS_PASSWORD="$REDIS_PASSWORD" redis sh -c 'exec redis-server --requirepass "$REDIS_PASSWORD"'      


> docker container run --name some-redis -d -e REDIS_PASSWORD="$REDIS_PASSWORD" redis redis-server
```

@TODO: test these 2 lines.

Or you can set the password in a Dockerfile:
```dockerfile
FROM redis

ENV REDIS_PASSWORD default-password

CMD ["sh", "-c", "exec redis-server --requirepass \"$REDIS_PASSWORD\""]
```


Or you can define the password in the `docker-compose.yml` file:
```yaml
  redis:
    image: 'redis:latest'
    command: redis-server --requirepass yourpassword
    ports:
      - '6379:6379'
```


