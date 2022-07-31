package redis

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/robertlestak/procx/internal/flags"
	log "github.com/sirupsen/logrus"
)

type RedisList struct {
	Client   *redis.Client
	Host     string
	Port     string
	Password string
	Key      string
}

func (d *RedisList) LoadEnv(prefix string) error {
	if os.Getenv(prefix+"REDIS_HOST") != "" {
		d.Host = os.Getenv(prefix + "REDIS_HOST")
	}
	if os.Getenv(prefix+"REDIS_PORT") != "" {
		d.Port = os.Getenv(prefix + "REDIS_PORT")
	}
	if os.Getenv(prefix+"REDIS_PASSWORD") != "" {
		d.Password = os.Getenv(prefix + "REDIS_PASSWORD")
	}
	if os.Getenv(prefix+"REDIS_KEY") != "" {
		d.Key = os.Getenv(prefix + "REDIS_KEY")
	}
	return nil
}

func (d *RedisList) LoadFlags() error {
	d.Host = *flags.RedisHost
	d.Port = *flags.RedisPort
	d.Password = *flags.RedisPassword
	d.Key = *flags.RedisKey
	return nil
}

func (d *RedisList) Init() error {
	l := log.WithFields(log.Fields{
		"package": "cache",
		"method":  "Init",
	})
	l.Debug("Initializing redis list driver")
	d.Client = redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%s", d.Host, d.Port),
		Password:    d.Password, // no password set
		DB:          0,          // use default DB
		DialTimeout: 30 * time.Second,
		ReadTimeout: 30 * time.Second,
	})
	cmd := d.Client.Ping()
	if cmd.Err() != nil {
		l.Error("Failed to connect to redis")
		return cmd.Err()
	}
	l.Debug("Connected to redis")
	return nil
}

func (d *RedisList) GetWork() (*string, error) {
	l := log.WithFields(log.Fields{
		"package": "cache",
		"method":  "GetWork",
	})
	l.Debug("Getting work from redis list")
	msg, err := d.Client.LPop(d.Key).Result()
	if err != nil {
		// If the queue is empty, return nil
		if err == redis.Nil {
			l.Debug("Queue is empty")
			return nil, nil
		}
		l.WithError(err).Error("Failed to receive message")
		return nil, err
	}
	l.Debug("Received message")
	return &msg, nil
}

func (d *RedisList) ClearWork() error {
	l := log.WithFields(log.Fields{
		"package": "cache",
		"method":  "ClearWork",
	})
	l.Debug("Clearing work from redis list")
	return nil
}

func (d *RedisList) HandleFailure() error {
	l := log.WithFields(log.Fields{
		"package": "cache",
		"method":  "HandleFailure",
	})
	l.Debug("Handling failure")
	return nil
}