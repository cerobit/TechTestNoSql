package drivenadapters

import (
	"context"
	"strconv"
)
import "github.com/go-redis/redis/v8"



var rdb  *redis.Client
var ctx context.Context

type BirdVoteRedis struct {
	// redis conn
}

// Todo inject config properly
func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx = context.Background()
}

func (bv BirdVoteRedis) AddVote(i int) (int, error) {

	key := strconv.Itoa(i)

	val2, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		// Does not exits, so it does initialize values
		rdb.Set(ctx, key, 1, 0).Err()
		val2 = "0"
	} else if err != nil {
		panic(err)
	} else {
		rdb.Incr(ctx, key )
		n,err := strconv.Atoi(val2)
		// Invalid value
		if err != nil {
			return -1, err
		}
		return n+1, err
	}

	n,err := strconv.Atoi(val2)
	// Invalid value
	if err != nil {
		return -1, err
	}
	return n, err
}
