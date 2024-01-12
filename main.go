package main

import (
	"fmt"
	"net/http"
)

type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

type Logger interface {
	Log(message string)
}

type Logic interface {
	SayHello(userID string) (string, error)
}

func LogOutput(message string) {
	fmt.Println(message)
}

func main() {
	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	http.HandleFunc("/hello", c.SayHello)
	http.ListenAndServe(":8080", nil)
}

//http://localhost:8080/hello?user_id=1
//http://localhost:8080/hello?user_id=2
//http://localhost:8080/hello?user_id=3
