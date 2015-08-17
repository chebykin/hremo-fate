package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

var port int

func main() {
	action := "s"

	flag.IntVar(&port, "port", 8080, "port for http interface")
	flag.Parse()

	fmt.Printf("Serving current dir on %d port\n", port)

	if action == "s" {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			panic(err)
		}

		err = http.ListenAndServe(
			fmt.Sprintf(":%d", port),
			http.FileServer(http.Dir(dir)))
		if err != nil {
			panic(err)
		}
	}
}
