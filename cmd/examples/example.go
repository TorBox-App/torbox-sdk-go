package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"torbox-sdk-go/pkg/torboxapi"
	"torbox-sdk-go/pkg/torboxapiconfig"
	"torbox-sdk-go/pkg/torrents"
	"torbox-sdk-go/pkg/util"
)

func main() {
	loadEnv()

	config := torboxapiconfig.NewConfig()
	client := torboxapi.NewTorboxApi(config)

	request := torrents.CreateTorrentRequest{
		AllowZip: util.ToPointer("AllowZip"),
		AsQueued: util.ToPointer("AsQueued"),
		File:     []byte{},
		Magnet:   util.ToPointer("Magnet"),
		Name:     util.ToPointer("Name"),
		Seed:     util.ToPointer("Seed"),
	}

	response, err := client.Torrents.CreateTorrent(context.Background(), "apiVersion", request)
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
