package url

import (
	"testing"

	"github.com/go-redis/redis"
	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/assert"

	"github.com/seregayoga/shorturl/pkg/config"
)

func TestCreateShortURLAndGetLongURL(t *testing.T) {
	test := assert.New(t)

	cfg := &config.Config{}
	err := envconfig.Process("shorturl", cfg)
	test.NoError(err)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	shortener := NewShortener(redisClient, cfg)

	originalLongURL := "https://www.example.com/some/path"

	shortURL, err := shortener.CreateShortURL(originalLongURL)
	test.NoError(err)

	longURL, err := shortener.GetLongURL(shortURL)
	test.NoError(err)

	test.Equal(originalLongURL, longURL)
}

func TestCreateShortURLTwiceReturnsSameShortURL(t *testing.T) {
	test := assert.New(t)

	cfg := &config.Config{}
	err := envconfig.Process("shorturl", cfg)
	test.NoError(err)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	shortener := NewShortener(redisClient, cfg)

	longURL := "https://www.example2.com/some/path"

	firstShortURL, err := shortener.CreateShortURL(longURL)
	test.NoError(err)

	secondShortURL, err := shortener.CreateShortURL(longURL)
	test.NoError(err)

	test.Equal(firstShortURL, secondShortURL)
}
