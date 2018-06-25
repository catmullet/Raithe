package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"strconv"
	"time"
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
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
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

// SetAgents Sets Agents on Redis
func SetAgents(key string, message []byte) error {
	r := redisClient.Set(fmt.Sprintf("%v", key), message, 120*time.Hour)
	return r.Err()
}

// Get Retrieves a message from file
func Get(queue string) ([]byte, error) {
	return readFile(queue)
}

// GetAgents Retrieves registered agents from Redis
func GetAgents(key string) ([]byte, error) {

	msg := []byte{}
	msg = []byte(redisClient.Get(key).Val())

	return msg, nil
}

// Invalidates Agents
func InvalidateAgents(key string) error {
	r := redisClient.Set(fmt.Sprintf("%v", key), []byte{}, 120*time.Hour)
	return r.Err()
}

func writeFile(key string, message []byte) error {
	r := redisClient.Set(fmt.Sprintf("%v_%v", key, makeTimestamp()), message, 48*time.Hour)
	return r.Err()
}

func readFile(key string) ([]byte, error) {
	iter := redisClient.Scan(0, key+"_*", 1).Iterator()
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
