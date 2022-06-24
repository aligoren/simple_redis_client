package main

import "net"

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
