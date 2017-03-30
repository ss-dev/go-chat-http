package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type RawData struct {
	Name string
	Text string
}

type Message struct {
	Id        int
	Timestamp time.Time
	Name      string
	Text      string
}

var bufCounter int = -1
var messageList []Message = []Message{}
var mesBuffer chan RawData = make(chan RawData, 100)

func bufWriter() {
	for {
		var item RawData = <-mesBuffer
		bufCounter++
		messageList = append(messageList, Message{
			Id:        bufCounter,
			Timestamp: time.Now(),
			Name:      item.Name,
			Text:      item.Text,
		})
	}
}

func chatServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if req.Method == "GET" {
		rawId := req.FormValue("id")
		if rawId == "" {
			rawId = "-1"
		}
		id, err := strconv.Atoi(rawId)
		if err != nil {
			fmt.Println("error:", err)
		} else if id > bufCounter {
			w.WriteHeader(http.StatusBadRequest)
		} else if id < bufCounter {
			newData, err := json.Marshal(messageList[id+1 : bufCounter+1])
			if err != nil {
				fmt.Println("error:", err)
			} else {
				io.WriteString(w, string(newData))
			}
		}
	} else {
		req.ParseForm()
		mesBuffer <- RawData{
			string(req.Form["name"][0]),
			string(req.Form["text"][0]),
		}
	}
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	data, err := ioutil.ReadFile("static/index.html")
	if err == nil {
		fmt.Fprintf(w, string(data))
	} else {
		log.Println(err)
		fmt.Fprintf(w, "<html><head></head><body><h1>Error on reading index.html</h1></body></html>")
	}
}

func main() {
	go bufWriter()

	fs := http.FileServer(http.Dir("static/js"))
	http.Handle("/js/", http.StripPrefix("/js/", fs))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/chat/", chatServer)

	var appPort = flag.String("port", "8080", "HTTP service port")
	flag.Parse()

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":"+*appPort, nil))
}
