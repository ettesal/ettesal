package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(url string) string {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}
	defer res.Body.Close()

	// Read all body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read error:", err)
		return ""
	}

	// fmt.Println("Status:", res.Status)
	// fmt.Println("Body:", string(body))

	return string(body)
}
