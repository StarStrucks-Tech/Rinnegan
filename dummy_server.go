package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

func main2() {

	for i := 0; i < 5; i++ {
		fmt.Println("Hello World")
		time.Sleep(5 * time.Second)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	fmt.Println("Starting server on :8080")
	go openBrowser("http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		panic(err)
	}
}
