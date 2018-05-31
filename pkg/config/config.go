package config

// Config application config
type Config struct {
	RedisAddr     string `default:"redis:6379"`
	RedisPassword string `default:""`
	RedisDB       int    `default:"0"`
	ListenAddr    string `default:":8080"`
	RedirectHost  string `default:"http://localhost:8080/"`
	PathSize      int    `default:"5"`
}
