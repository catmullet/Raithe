package cache

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"github.com/go-redis/redis"
)

var redisClient *redis.Client

// InitializeRedisClient Initializes the Redis Client from Environment variables.
func InitializeRedisClient() {
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))

	if err != nil {
		fmt.Println("Failed to Get Redis DB Environment Variable.  Defaulting to 0")
		db = 0
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:os.Getenv("REDIS_URL"),
		Password:os.Getenv("REDIS_PASSWORD"),
		DB:db,
	})

	ping, err := redisClient.Ping().Result()

	if err != nil {
		fmt.Println("Failed to Initialize Redis, " + ping)
	}
}

// Set Writes a message to file
func Set(key string, message []byte) error {
	return writeFile(key, message)
}

// Get Retrieves a message from file
func Get(queue string) ([]byte, error) {
	return readFile(queue)
}

func writeFile(key string, message []byte) error {
	r := redisClient.Set(fmt.Sprintf("%v_%v", key, makeTimestamp()), message, 48 * time.Hour)
	return r.Err()
}

func readFile(key string) ([]byte, error) {
	iter := redisClient.Scan(0, key + "_*",1).Iterator()
	iter.Next()

	msg := []byte{}

	if iter.Err() != nil {
		return msg, iter.Err()
	}

	msg = []byte(redisClient.Get(iter.Val()).Val())

	redisClient.Del(iter.Val())

	return msg, nil
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
