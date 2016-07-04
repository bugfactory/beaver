package elastsear

import (
	

	elastigo "github.com/mattbaird/elastigo/lib"
)

func New() {
	client, err := elastigo.NewConn()
	fmt.Println(err, client)
}