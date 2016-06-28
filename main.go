package main

import (
    "fmt"
    "github.com/bugfactory/beaver/config"
    // elastigo "github.com/mattbaird/elastigo/lib"
)

const configFile = "config/config.yml"

func main() {
	conf, err := configuration.New(configFile)
	fmt.Print(conf, err)
	
	// client, err := elastigo.NewConn()
	// fmt.Println(err, client)
	
}

