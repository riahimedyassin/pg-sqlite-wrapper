package main

import (
	"fmt"
	pglitewrapper "github.com/riahimedyassin/pg-sqlite-wrapper"
)

func main() {
	c := pglitewrapper.NewDatabaseConfig("dsa", "postgres")
	r := pglitewrapper.NewReconnectConfig(true, 3)
	_, err := pglitewrapper.NewDatabaseWrapper(c, r)
	if err != nil {
		fmt.Println(err)
	}
}
