package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	m := http.NewServeMux()

	m.HandleFunc("/", handlePage)
	
	port := os.Getenv("PORT")

	// this blocks forever, until the server
	// has an unrecoverable error
	fmt.Println("server started on ", port)
	err := http.ListenAndServe(":"+port, m)
	log.Fatal(err)
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	const page = `<html>
<head></head>
<body>
	<p> Hello from Docker! I'm a Go server. </p>
</body>
</html>
`
	w.Write([]byte(page))
}
