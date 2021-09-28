package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"bufio"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}



func scanLine(scanner *bufio.Scanner) []string {
	s := make([]string, 2000000)
	counter := 0
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	for scanner.Scan() {
		log.Println(counter)
		scanner.Text()
		//text:= scanner.Text()
		//s[counter] = text
		counter++
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
	}
	fmt.Printf("Total:")
	log.Println(counter)
	return s
}

func main() {

	f, err := os.Open("./openfoodfacts-products.jsonl")
	check(err)
	scanner := bufio.NewScanner(f)
	results:= scanLine(scanner)

	handler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Result:" + results[42371])
	}
	http.HandleFunc("/hello", handler)
	log.Println("Listing for request at http://localhost:8080/hello")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
