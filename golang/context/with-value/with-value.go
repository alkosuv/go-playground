package main

import (
	"context"
	"log"
)

// Key импользуется для ихбежания коллизий (не обязательно)
type Key string

func main() {
	ctx := context.WithValue(context.Background(), Key("name"), "gen95mis")

	log.Printf("name (type key) = %v", ctx.Value(Key("name")))
	log.Printf("name (type string)= %v", ctx.Value("name"))
	log.Printf("login (type key) = %v", ctx.Value(Key("login")))
}
