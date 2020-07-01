package db

import (
	"context"
	"fmt"
	"time"

	"github.com/ddrugeon/go-rest-api/internal/model"
	"github.com/gomodule/redigo/redis"
)

type redisRepository struct {
	// Declare a pool variable to hold the pool of Redis connections.
	pool *redis.Pool
}

func NewRedisRepository(addr string) Repository {
	pool := &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
	}
	db := redisRepository{pool: pool}

	return &db
}

func (r *redisRepository) Ping() error {
	conn := r.pool.Get()
	defer conn.Close()

	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		return fmt.Errorf("cannot 'PING' db: %v", err)
	}

	return nil
}

func (r *redisRepository) Get(ctx context.Context) []model.Droid {
	droids := make([]model.Droid, 0)

	keys, err := r.getKeys()
	if err != nil {
		return droids
	}

	for _, key := range keys {
		current, err := r.GetByID(ctx, key)
		if err == nil {
			droids = append(droids, current)
		}
	}

	return droids
}

func (r *redisRepository) GetByID(ctx context.Context, id string) (model.Droid, error) {
	conn := r.pool.Get()
	defer conn.Close()

	value, err := conn.Do("HGETALL", id)
	if err != nil {
		return model.Droid{}, err
	}

	droid, err := redis.StringMap(value, err)
	if err != nil {
		return model.Droid{}, err
	}
	return populateDroid(droid)
}

func (r *redisRepository) Put(d model.Droid) error {
	conn := r.pool.Get()
	defer conn.Close()

	hash := d.ID

	// Begin transaction to add new object
	err := conn.Send("MULTI")
	if err != nil {
		return err
	}

	err = conn.Send("HSET", hash, "id", d.ID)
	if err != nil {
		return err
	}

	err = conn.Send("HSET", hash, "class", d.Class)
	if err != nil {
		return err
	}

	err = conn.Send("HSET", hash, "company", d.Company)
	if err != nil {
		return err
	}

	err = conn.Send("HSET", hash, "height", d.Height)
	if err != nil {
		return err
	}

	err = conn.Send("HSET", hash, "model", d.Model)
	if err != nil {
		return err
	}

	err = conn.Send("HSET", hash, "type", d.Type)
	if err != nil {
		return err
	}

	err = conn.Send("HSET", hash, "vehicles", d.Vehicles)
	if err != nil {
		return err
	}

	err = conn.Send("HSET", hash, "name", d.Name)
	if err != nil {
		return err
	}

	err = conn.Send("SADD", "droid:keys", hash)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXEC")
	if err != nil {
		return err
	}

	return nil
}

func (r *redisRepository) Version() string {
	return "Redis database"
}

func (r *redisRepository) getKeys() ([]string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	iter := 0
	keys := []string{}
	for {
		arr, err := redis.Values(conn.Do("SSCAN", "droid:keys", iter))
		if err != nil {
			return keys, fmt.Errorf("error retrieving keys")
		}

		iter, _ = redis.Int(arr[0], nil)
		k, _ := redis.Strings(arr[1], nil)
		keys = append(keys, k...)

		if iter == 0 {
			break
		}
	}

	return keys, nil
}

func populateDroid(reply map[string]string) (model.Droid, error) {
	droid := model.Droid{}

	if len(reply) == 0 {
		return droid, model.ErrNoRecord
	}

	droid = model.Droid{
		ID:       reply["id"],
		Class:    reply["class"],
		Company:  reply["company"],
		Height:   reply["height"],
		Model:    reply["model"],
		Type:     reply["type"],
		Vehicles: reply["vehicules"],
		Name:     reply["name"],
	}
	return droid, nil
}
