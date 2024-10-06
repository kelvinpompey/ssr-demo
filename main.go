package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"example.com/gin/models"
	gossr "github.com/natewong1313/go-react-ssr"
)

var APP_ENV string = "production"

func main() {
	// Serve static files like favicon.ico and assets
	http.Handle("/favicon.ico", http.FileServer(http.Dir("./frontend/public")))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./frontend/public"))))

	// Initialize the SSR engine
	engine, err := gossr.New(gossr.Config{
		AppEnv:             APP_ENV,
		AssetRoute:         "/assets",
		FrontendDir:        "./frontend/src",
		GeneratedTypesPath: "./frontend/src/generated.d.ts",
		TailwindConfigPath: "./frontend/tailwind.config.js",
		LayoutCSSFilePath:  "globals.css",
		PropsStructsPath:   "./models/props.go",
	})
	if err != nil {
		log.Fatal("Failed to init go-react-ssr")
	}

	// Route for Home page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		content := engine.RenderRoute(gossr.RenderConfig{
			File:  "Home.tsx",
			Title: "Go built-in server example",
			MetaTags: map[string]string{
				"og:title":    "Go example app",
				"description": "Hello world!",
			},
			Props: &models.IndexRouteProps{
				InitialCount: rand.Intn(100),
			},
		})
		w.Header().Set("Content-Type", "text/html")
		w.Write(content)
	})

	// Route for About page
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		content := engine.RenderRoute(gossr.RenderConfig{
			File:  "About.tsx",
			Title: "About page",
			MetaTags: map[string]string{
				"og:title":    "Go example app",
				"description": "Hello world!",
			},
			Props: &models.AboutRouteProps{
				Content: "This is my about content",
			},
		})
		w.Header().Set("Content-Type", "text/html")
		w.Write(content)
	})

	// Start the server
	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}