package main

import (
	"fmt"
	"net/http"
	v1 "online-library/api/v1"
	"online-library/config"
	"os"
)

func main() {
	if err := config.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n Couldn't start the my-app service, err: ", err)
		os.Exit(1)
	}

	fmt.Println("starting  my-app service on ", config.Conf.HTTPPort)
	config.CreateDir()
	v1.Routes()
	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.Conf.HTTPPort), nil); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n Couldn't start the server, err: ", err)
		os.Exit(1)
	}
}
