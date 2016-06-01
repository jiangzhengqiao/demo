package main

import (
	"io"
	"log"
	"net/http"
	"os"
	// "time"
)

// version 1
// func main() {
// 	http.HandleFunc("/", sayHello)

// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func sayHello(rw http.ResponseWriter, req *http.Request) {
// 	io.WriteString(rw, "Hello Word!")
// }

// // version 2
// func main() {
// 	mux := http.NewServeMux()
// 	mux.Handle("/", &myHandler{})

// 	err := http.ListenAndServe(":8080", mux)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// type myHandler struct{}

// func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "URL:"+r.URL.String())
// }

// version 3
func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/hello", sayHello)

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	mux.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir(wd))))

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "URL:"+r.URL.String())
}

func sayHello(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "Hello Word!")
}

// version 4
// var mux map[string]func(http.ResponseWriter, *http.Request)

// func main() {
// 	server := http.Server{
// 		Addr:        ":8080",
// 		Handler:     &myHandler{},
// 		ReadTimeout: 5 * time.Second,
// 	}

// 	mux = make(map[string]func(http.ResponseWriter, *http.Request))
// 	mux["/hello"] = sayHello
// 	mux["/bye"] = sayBye

// 	err := server.ListenAndServe()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// type myHandler struct{}

// func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	log.Println(r.URL.String())
// 	h, ok := mux[r.URL.String()]
// 	log.Println(mux)
// 	if ok {
// 		h(w, r)
// 		return
// 	}

// 	io.WriteString(w, "URL:"+r.URL.String())
// }

// func sayHello(rw http.ResponseWriter, req *http.Request) {
// 	io.WriteString(rw, "Hello Word!")
// }

// func sayBye(rw http.ResponseWriter, req *http.Request) {
// 	io.WriteString(rw, "bye bye!")
// }
