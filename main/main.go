package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	. "jug/algorithm"
	"log"
	"net/http"
)

func main() {
	startServer()
}

func solveAndPrint(in SolveInput, w http.ResponseWriter) {
	fmt.Printf("start %v\n", in)
	solution := ShortestSolutionIfExists(Solve(in))
	fmt.Printf("complete %v\n", in)
	if solution == nil {
		fmt.Fprintf(w, "No Solution")
		return
	}
	data, err := json.Marshal(solution)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(data))
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		in := SolveInput{}
		err = json.Unmarshal(body, &in)
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		solveAndPrint(in, w)
	default:
		fmt.Fprintf(w, "use POST method with ")
	}
}

func startServer() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
