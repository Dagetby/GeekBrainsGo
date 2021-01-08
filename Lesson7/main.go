package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	Version  string `yamL:"version"`
	Services services
}
type services struct {
	Testdb_p testdb_p
}
type testdb_p struct {
	Image       string `yaml:"image"`
	Restart     string `yaml:"restart"`
	Ports       string `yaml:"ports"`
	Environment environment
}
type environment struct {
	PostgreSPassword string `yaml:"POSTGRES_PASSWORD"`
}

func main() {
	c := Config{}

	file, err := os.Open("/Users/t.mirzakhanyan/goMod/GeekBraisGo/Lesson7/docker-compose.yaml")
	if err != nil {
		log.Fatal("Не удалось открыть файл ", err)
	}

	defer func() {
		err := file.Close()

		if err != nil {
			log.Fatal("Не удалось закрыть файл")
		}
	}()

	err = yaml.NewDecoder(file).Decode(&c)

	if err != nil {
		log.Fatal("Не удалось распаковать файл ", err)
	}

	fmt.Println(c)
}
