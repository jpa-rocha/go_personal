package main

import (
	// "html/template"
	"net/http"
	"os/exec"
	"fmt"
)


const dirname = "./"

func main() {
	http.Handle("/", http.FileServer(http.Dir(dirname)))

	exec.Command("xdg-open", "http://localhost:8000").Run()
	fmt.Println("Server started at http://localhost:8000")
	http.ListenAndServe(":8000", nil)

}
