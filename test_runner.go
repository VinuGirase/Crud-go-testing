package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url" // ✅ Add this
	"os"
	"os/exec"
	"runtime"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)


type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
// from vinu - now test - jj - qwertyuiop
func TestRunner() {
	rate := vegeta.Rate{Freq: 100, Per: time.Second}
	duration := 10 * time.Second
	attacker := vegeta.NewAttacker()

	rand.Seed(time.Now().UnixNano())

	// ------------------------------
	// POST /users - dynamic users
	// ----------------------------
	postTargeter := func(t *vegeta.Target) error {
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

	var postMetrics vegeta.Metrics
	for res := range attacker.Attack(postTargeter, rate, duration, "POST /users Load Test") {
		postMetrics.Add(res)
		if res.Code != http.StatusOK {
			fmt.Printf("POST Failed: Status=%d, Body=%s\n", res.Code, string(res.Body))
		}
	}
	postMetrics.Close()

	// ----------------------------
	// GET /users
	// ----------------------------
	getTargeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: http.MethodGet,
		URL:    "http://localhost:8080/users",
	})

	var getMetrics vegeta.Metrics
	for res := range attacker.Attack(getTargeter, rate, duration, "GET /users Load Test") {
		getMetrics.Add(res)
		if res.Code != http.StatusOK {
			fmt.Printf("GET Failed: Status=%d, Body=%s\n", res.Code, string(res.Body))
		}
	}
	getMetrics.Close()

	// ----------------------------
	// Cleanup test users (NEW)
	// ----------------------------
	// fmt.Println("\n🧹 Cleaning up test users...")
	// resp, err := http.DefaultClient.Do(&http.Request{
	// 	Method: http.MethodDelete,
	// 	URL:    mustParseURL("http://localhost:8080/users/cleanup"),
	// })
	// if err != nil {
	// 	fmt.Println("❌ Cleanup request failed:", err)
	// } else {
	// 	defer resp.Body.Close()
	// 	fmt.Printf("✅ Cleanup status: %s\n", resp.Status)
	// }

	// ----------------------------
	// Report Results
	// ----------------------------
	fmt.Println("\n--- POST /users Load Test Report ---")
	vegeta.NewTextReporter(&postMetrics).Report(os.Stdout)

	fmt.Println("\n--- GET /users Load Test Report ---")
	vegeta.NewTextReporter(&getMetrics).Report(os.Stdout)

	// ----------------------------
	// Test Coverage Report
	// ----------------------------
	coverProfile := "coverage.out"
	fmt.Println("📦 Running tests with coverage...")

	cmd := exec.Command("go", "test", "./controllers", "-v", "-coverprofile="+coverProfile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Tests failed:", err)
		return
	}

	fmt.Println("✅ Tests passed. Generating HTML coverage report...")

	htmlFile := "coverage.html"
	cmd = exec.Command("go", "tool", "cover", "-html="+coverProfile, "-o", htmlFile)
	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Failed to generate HTML:", err)
		return
	}

	fmt.Println("🌐 Opening HTML coverage report...")
	openBrowser(htmlFile)
}

func mustParseURL(raw string) *url.URL {
	parsed, err := url.ParseRequestURI(raw)
	if err != nil {
		panic(fmt.Sprintf("invalid URL: %s", raw))
	}
	return parsed
}


func openBrowser(file string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", file)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", file)
	default:
		cmd = exec.Command("xdg-open", file)
	}
	_ = cmd.Start()
}
