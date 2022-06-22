package main

import "net"

type RedisClient interface {
	Connect() (RedisConnection, error)
	Auth() (RedisConnection, error)
	Set(key string, value string) RedisResponse
}

type RedisConfig struct {
	Address    string
	Port       int
	Password   string
	Connection RedisConnection
}

type RedisConnection struct {
	Stream net.Conn
}

type RedisResponse struct {
	Message string
	Success bool
}
