/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком.
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func main() {
	err := Wget("https://go.dev/blog/strings")
	if err != nil {
		fmt.Println(err)
	}
}

func Wget(URL string) error {
	resp, err := http.Get(URL)
	if err != nil {
		return fmt.Errorf("error fetching url: %v", err)
	}
	defer resp.Body.Close()

	u, err := url.Parse(URL)
	if err != nil {
		return fmt.Errorf("error parsing url: %v", err)
	}

	fileDir := "downloaded" + u.Path
	if u.Path == "/" {
		fileDir = "downloaded/index.html"
	}

	err = os.MkdirAll(fileDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create dir: %v", err)
	}

	file, err := os.Create(filepath.Join(fileDir, "index.html"))
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	fmt.Println("Downloaded to:", fileDir)

	return nil
}
