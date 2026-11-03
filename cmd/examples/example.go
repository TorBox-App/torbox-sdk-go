package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"context"
	"example.com/main-api-sdk"
	"example.com/main-api-sdk/mainapisdk1"
)

func main() {
	loadEnv()

	config := mainapisdk.NewConfig()
	client := mainapisdk.NewMainAPISDK(config)

	request := mainapisdk1.CreateTorrentRequest{
		File:            mainapisdk.Nullable[[]byte]([]byte{}),
		Magnet:          mainapisdk.Nullable[string]("string"),
		Seed:            mainapisdk.Nullable[string]("int"),
		AllowZip:        mainapisdk.Nullable[string]("boolean"),
		Name:            mainapisdk.Nullable[string]("string"),
		AsQueued:        mainapisdk.Nullable[string]("boolean"),
		AddOnlyIfCached: mainapisdk.Nullable[string]("boolean"),
	}

	response, err := client.MainAPISDK1.CreateTorrent(context.Background(), "api_version", request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", response)
}

func loadEnv() error {
	file, err := os.Open(".env")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
