package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logwriter struct{}

func main() {
	resp, err := http.Get("http://www.google.com")

	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}

	lw := logwriter{}
	// io.Copy(os.Stdout, resp.Body)
	io.Copy(lw, resp.Body)

}

//by specifying logwter in the below gunction the logwriter is implementing Write inteface

func (logwriter) Write(bs []byte) (i int, err error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote these many bytes :", len(bs))
	return len(bs), nil
}
