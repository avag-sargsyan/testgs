package main

import (
	"fmt"
	"github.com/avag-sargsyan/testgs/internal/adapter/handler"
	"github.com/avag-sargsyan/testgs/internal/conf"
	"github.com/avag-sargsyan/testgs/internal/err"
	"github.com/avag-sargsyan/testgs/internal/logger"
	"github.com/caarlos0/env/v9"
	"net/http"
)

func main() {
	logger.Setup()

	confApp := conf.App{}
	e := env.Parse(&confApp)
	err.FatalIfError(e)

	initHTTPServer(&confApp)
}

func initHTTPServer(conf *conf.App) {
	ph := handler.NewPackHandler(conf)
	http.HandleFunc("/api/packs", ph.Pack)
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	err.FatalIfError(http.ListenAndServe(conf.HTTPAddress, nil))
	fmt.Println("Server is running on port ", conf.HTTPAddress)
}
