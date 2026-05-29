package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Person struct {
	Name      string    `json:"name"`
	Height    string    `json:"height"`
	Mass      string    `json:"mass"`
	HairColor string    `json:"hair_color"`
	SkinColor string    `json:"skin_color"`
	EyeColor  string    `json:"eye_color"`
	BirthYear string    `json:"birth_year"`
	Gender    string    `json:"gender"`
	Homeworld string    `json:"homeworld"`
	Films     []string  `json:"films"`
	Species   []string  `json:"species"`
	Vehicles  []string  `json:"vehicles"`
	Starships []string  `json:"starships"`
	Created   time.Time `json:"created"`
	Edited    time.Time `json:"edited"`
	URL       string    `json:"url"`
}

func main() {
	// Меня timeout  можно контролировать время запроса
	timeout := time.Second * 15
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://swapi.dev/api/people/1/", nil)
	if err != nil {
		panic(fmt.Sprintf("failed to create request with ctx: %v", err))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(fmt.Sprintf("failed to perform http request: %v", err))
	}

	var r Person
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		panic(fmt.Sprintf("failed read body: %v", err))
	}

	fmt.Printf("%+v", r)
}
