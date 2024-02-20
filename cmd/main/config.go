package main

import (
	"errors"
	"log"
	"os"
	"strconv"
)

var ErrMissingEnv = errors.New("missing env variable")

type Configurations struct {
	Server ServerConfigurations `json:"server"`
}

type ServerConfigurations struct {
	Port int `json:"port"`
}

func (c *Configurations) ReadConfigs() error {
	portEnv := os.Getenv("PORT")
	if portEnv == "" {
		log.Fatal("env PORT empty")
		return ErrMissingEnv
	}

	port, err := strconv.Atoi(portEnv)
	if err != nil {
		log.Fatal("invalid PORT env")
		return ErrMissingEnv
	}

	c.Server.Port = port
	return nil
}
