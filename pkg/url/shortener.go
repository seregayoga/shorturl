package url

import (
	"math/rand"
	"time"

	"github.com/go-redis/redis"
	"github.com/seregayoga/shorturl/pkg/config"
)

var alphabet = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-")

const (
	url2shortPrefix = "url2short:"
	short2urlPrefix = "short2url:"
)

// Shortener creates short urls and stores it in Redis
type Shortener struct {
	redisClient *redis.Client
	cfg         *config.Config
}

// NewShortener creates new Shortener
func NewShortener(redisClient *redis.Client, cfg *config.Config) *Shortener {
	return &Shortener{
		redisClient: redisClient,
		cfg:         cfg,
	}
}

// CreateShortURL creates short url from long
func (s *Shortener) CreateShortURL(longURL string) (string, error) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	shortURLRaw := make([]rune, s.cfg.PathSize)
	for i := range shortURLRaw {
		shortURLRaw[i] = alphabet[rnd.Intn(len(alphabet))]
	}

	shortURL := string(shortURLRaw)

	url2short := url2shortPrefix + longURL
	wasSet, err := s.redisClient.SetNX(url2short, shortURL, 0).Result()
	if err != nil {
		return "", err
	}

	if !wasSet {
		shortURL, err := s.redisClient.Get(url2short).Result()
		if err != nil {
			return "", err
		}

		return shortURL, nil
	}

	_, err = s.redisClient.Set(short2urlPrefix+shortURL, longURL, 0).Result()
	if err != nil {
		return "", err
	}

	return shortURL, nil
}

// GetLongURL returns long url from short
func (s *Shortener) GetLongURL(shortURL string) (string, error) {
	longURL, err := s.redisClient.Get(short2urlPrefix + shortURL).Result()
	if err == redis.Nil {
		return "", nil
	}
	if err != nil {
		return "", err
	}

	return longURL, nil
}
