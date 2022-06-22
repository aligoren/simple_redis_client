package main

import (
	"errors"
	"fmt"
	"net"
)

func (c *RedisConfig) Connect() (*RedisConfig, error) {

	stream, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.Address, c.Port))

	c.Connection = RedisConnection{
		Stream: stream,
	}

	return c, err
}

func (c *RedisConfig) Auth() (*RedisConfig, error) {

	buf := make([]byte, 0, 4096)
	tmp := make([]byte, 256)

	authKey := "AUTH"
	authLength := len(authKey)
	password := c.Password
	passwordLength := len(password)

	authCommand := fmt.Sprintf("*2\r\n$%d\r\n%s\r\n$%d\r\n%s\r\n", authLength, authKey, passwordLength, password)

	command := []byte(authCommand)

	_, err := c.Connection.Stream.Write(command)
	if err != nil {
		return nil, err
	}

	_, err = c.Connection.Stream.Read(tmp)
	if err != nil {
		return nil, err
	}

	for _, d := range tmp {
		buf = append(buf, d)
	}

	result := string(tmp)

	if string(result[0]) == "-" {
		return nil, errors.New(fmt.Sprintf("connection failed: %s", result))
	}

	return c, nil
}

func (c *RedisConfig) Set(key string, value string) (RedisResponse, error) {

	buf := make([]byte, 0, 4096)
	tmp := make([]byte, 256)

	set := "SET"
	setLength := len(set)
	keyLength := len(key)
	valueLength := len(value)

	output := fmt.Sprintf("*3\r\n$%d\r\n%s\r\n$%d\r\n%s\r\n$%d\r\n%s\r\n", setLength, set, keyLength, key, valueLength, value)

	command := []byte(output)

	_, err := c.Connection.Stream.Write(command)
	if err != nil {
		return RedisResponse{}, err
	}

	_, err = c.Connection.Stream.Read(tmp)
	if err != nil {
		return RedisResponse{}, err
	}

	for _, d := range tmp {
		buf = append(buf, d)
	}

	result := string(tmp)

	if string(result[0]) == "-" {
		return RedisResponse{
			Message: result[0:],
			Success: false,
		}, errors.New(fmt.Sprintf("Couldn't set key: %s with value %s", key, value))
	}

	return RedisResponse{
		Message: result,
		Success: true,
	}, nil
}
