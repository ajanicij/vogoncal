package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type config struct {
	RootDir        string
	CalFilePattern string
}

func main() {
	fmt.Println("hello world")
	buf, err := os.ReadFile("./myconfig.cfg")
	if err != nil {
		log.Fatal(err)
	}
	var cfg config
	err = toml.Unmarshal(buf, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("configuration:")
	fmt.Printf("rootdir: %s\n", cfg.RootDir)
	fmt.Printf("calfilepattern: %s\n", cfg.CalFilePattern)
}
