package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"redis_client_example/builder"
	"redis_client_example/commands"
	"strings"
)

func (c *RedisConfig) Connect() (*RedisConfig, error) {

	stream, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.Address, c.Port))

	c.Connection = RedisConnection{
		Stream: stream,
	}

	if c.Password != "" {
		_, err := c.Auth()
		if err != nil {
			log.Fatalf("Error while auth: %v", err)
		}
	}

	return c, err
}

func (c *RedisConfig) Auth() (*RedisConfig, error) {

	buf := make([]byte, 0, 4096)
	tmp := make([]byte, 256)

	authCommand := builder.BuildCommandtring(commands.AUTH, c.Password)

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

func (c *RedisConfig) Info() (RedisResponse, error) {

	buf := make([]byte, 0, 4096)
	tmp := make([]byte, 256)

	infoCommand := builder.BuildCommandtring(commands.INFO)

	command := []byte(infoCommand)

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
		}, errors.New(fmt.Sprintf("Command failed"))
	}

	return RedisResponse{
		Message: result,
		Success: true,
	}, nil
}

func (c *RedisConfig) Set(key string, value string) (RedisResponse, error) {

	buf := make([]byte, 0, 4096)
	tmp := make([]byte, 256)

	setCommand := builder.BuildCommandtring(commands.SET, key, value)

	command := []byte(setCommand)

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

func (c *RedisConfig) Get(key string) (*RedisResponse, error) {

	buf := make([]byte, 0, 4096)
	tmp := make([]byte, 256)

	setCommand := builder.BuildCommandtring(commands.GET, key)

	command := []byte(setCommand)

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

	if string(result[0:3]) == "$-1" {
		return nil, errors.New(fmt.Sprintf("Key is not found: %s", key))
	}

	resultStringArr := strings.Split(result, "\r\n")

	return &RedisResponse{
		Message: strings.Join(resultStringArr[1:len(resultStringArr)-1], ""),
		Success: true,
	}, nil
}

func (c *RedisConfig) InsertArray(listName string, values ...interface{}) (RedisResponse, error) {

	buf := make([]byte, 0, 4096)
	tmp := make([]byte, 256)

	setCommand := builder.BuildArrayString(listName, values)

	command := []byte(setCommand)

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
		}, errors.New(fmt.Sprintf("Couldn't set array: %s with values %s", listName, values))
	}

	return RedisResponse{
		Message: result,
		Success: true,
	}, nil
}
