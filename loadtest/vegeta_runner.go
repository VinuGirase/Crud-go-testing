package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

// Match your models.User struct
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	rate := vegeta.Rate{Freq: 100, Per: time.Second}
	duration := 10 * time.Second
	attacker := vegeta.NewAttacker()

	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Dynamic targeter that creates new users on each request
	targeter := func(t *vegeta.Target) error {
		user := User{
			Name:  fmt.Sprintf("VegetaUser%d", rand.Intn(1_000_000)),
			Email: fmt.Sprintf("vegeta%d@example.com", rand.Intn(1_000_000)),
		}

		body, err := json.Marshal(user)
		if err != nil {
			return err
		}

		t.Method = http.MethodPost
		t.URL = "http://localhost:8080/users"
		t.Header = http.Header{
			"Content-Type": []string{"application/json"},
		}
		t.Body = body

		return nil
	}

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Dynamic User Load Test") {
		metrics.Add(res)

		if res.Code != http.StatusOK {
			fmt.Printf("Failed Request: Status=%d, Body=%s\n", res.Code, string(res.Body))
		}
	}
	metrics.Close()

	// Use Vegeta's built-in text reporter
	reporter := vegeta.NewTextReporter(&metrics)
	fmt.Println("\n--- Load Test Report ---")
	reporter.Report(os.Stdout)
}
