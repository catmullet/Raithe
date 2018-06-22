package cache

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"github.com/go-redis/redis"
)

var redis_client *redis.Client

func InitializeRedisClient() {
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))

	if err != nil {
		fmt.Println("Failed to Get Redis DB Environment Variable.  Defaulting to 0")
		db = 0
	}

	redis_client = redis.NewClient(&redis.Options{
		Addr:os.Getenv("REDIS_URL"),
		Password:os.Getenv("REDIS_PASSWORD"),
		DB:db,
	})

	ping, err := redis_client.Ping().Result()

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
	r := redis_client.Set(fmt.Sprintf("%v_%v", key, makeTimestamp()), message, 48 * time.Hour)
	return r.Err()
}

func readFile(key string) ([]byte, error) {
	iter := redis_client.Scan(0, key + "_*",1).Iterator()
	iter.Next()

	msg := []byte{}

	if iter.Err() != nil {
		return msg, iter.Err()
	}

	msg = []byte(redis_client.Get(iter.Val()).Val())

	redis_client.Del(iter.Val())

	return msg, nil
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
