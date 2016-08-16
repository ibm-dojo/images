package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<html>"))
	name, _ := os.Hostname()
	w.Write([]byte(fmt.Sprintf("<b>Hostname: %s<br>\n", name)))
	w.Write([]byte("<br>Interfaces:</b><br>\n"))
	w.Write([]byte("<pre>"))
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Can't get addrs: %s\n", err)))
		return
	}
	for _, addr := range addrs {
		str := addr.String()
		ch := str[0]
		if ch < '0' || ch > '9' {
			continue
		}
		str = strings.Split(str, "/")[0]
		w.Write([]byte(fmt.Sprintf("%s\n", str)))
	}
	w.Write([]byte("</pre>"))
	w.Write([]byte(fmt.Sprintf("<br><b>Date: %s</b>\n", time.Now().String())))
}

func main() {
	addr := "0.0.0.0"
	port := "80"

	if len(os.Args) > 2 {
		addr = os.Args[1]
		port = os.Args[2]
	} else if len(os.Args) > 1 {
		port = os.Args[1]
	}

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(addr+":"+port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
