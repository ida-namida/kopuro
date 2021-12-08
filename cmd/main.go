package main

import (
    "kopuro/controller/httpserver"
    "kopuro/service"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    baseFilePath := os.Getenv("BASE_FILE_PATH")
    jsonFileService := service.NewJsonFileService(baseFilePath)
    server := httpserver.NewServer(port, jsonFileService)
    server.Start()
}