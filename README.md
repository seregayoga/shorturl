# SHORTURL

## Simple but powerful url shortener in Go and Redis

Dependcies:
Docker and docker compose

Run:
`make up`

Stop:
`make down`

Tests:
`make test`

Usage:

Open http://localhost:8080 or run
`curl -s -X POST http://localhost:8080 -d '{"long_url":"http://www.example.com"}'`
