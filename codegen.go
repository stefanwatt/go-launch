package main

import (
	"os"
	"path"
	"strings"
)

func generateIframeGoJs() {
	basePath := "~/.config/go-launch/frontend/src/lib/wailsjs/go/main"
	jsPath := path.Join(basePath, "App.js")
	// dtsPath := path.Join(basePath, "App.d.ts")
	// read js file line into string
	bytes, err := os.ReadFile(jsPath)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(bytes), "\n")
	// init empty string array
	chunk := []string{}
	functions := [][]string{}

	for _, line := range lines {
		if strings.Contains(line, "//") || len(strings.TrimSpace(line)) == 0 {
			continue
		}
		if strings.Contains(line, "export function") {
			functions = append(functions, chunk)
			chunk = []string{}
		}
		chunk = append(chunk, line)
	}
}

func getFunctionName(function []string) string {
	return strings.Split(function[0], " ")[2]
}
