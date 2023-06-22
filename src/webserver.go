package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	
	"strings"
	"strconv"
	"time"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
		//Get timestamp as string
		timestamp := strconv.FormatInt(time.Now().UTC().Unix(), 10)

		//Process remote address to get IP (TODO: Fix this for IPv6)
		addrArray := strings.Split(r.RemoteAddr, ":")
		ip := addrArray[0]

		//Log to console
		fmt.Print("[", timestamp)
		fmt.Print("] Connection from ", r.RemoteAddr)
		fmt.Print("\n")

		//Output to HTTP writer (TODO: Make this JSON)
		fmt.Fprintf(w, "Hello %q", html.EscapeString(ip))
		fmt.Fprintf(w, ", unix time is %q", html.EscapeString(timestamp))
		fmt.Fprintf(w, "\n")
	})

	//Listen for HTTP and log fatal errors
	fmt.Println("Listening for connections on *:8081\n---")
	log.Fatal(http.ListenAndServe(":8081", nil))

}