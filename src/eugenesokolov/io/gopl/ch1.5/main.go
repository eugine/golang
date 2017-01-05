// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(prepareURL(url))
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// buffer, err := ioutil.ReadAll(resp.Body)
		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		fmt.Printf("status code %v\n", resp.StatusCode)
	}
}

func prepareURL(url string) string {
    if !strings.HasPrefix(url, "http") {
        return "http://" + url
    }
    return url
}
//!-
