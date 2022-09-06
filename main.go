package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//cacher := redis.NewCacher(redis.Config{
	//	Addr: "localhost:9096",
	//})
	r := mux.NewRouter()
	r.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {

		v := r.URL.Query().Get("file")
		b, err := ioutil.ReadFile("uploads/" + v)
		//b, err := cacher.Get(v)
		if err != nil {
			log.Print(err.Error())
			return
		}
		w.Header().Set("Content-Disposition", "attachment; filename="+v)
		w.Header().Set("Content-Type", "application/octet-stream")

		if _, err := w.Write(b); err != nil {
			log.Println("unable to write image.")
			return
		}
	}).Methods("GET")
	r.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("File Upload Endpoint Hit")
		err := r.ParseMultipartForm(10240)
		if err != nil {
			log.Print(err.Error())
			return
		}
		file, handler, err := r.FormFile("file")
		if err != nil {
			log.Println("Error Retrieving the File")
			log.Println(err)
			return
		}
		defer file.Close()
		log.Printf("Uploaded File: %+v\n", handler.Filename)
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		//cacher.Set(handler.Filename, fileBytes, 0)
		err = ioutil.WriteFile("uploads/"+handler.Filename, fileBytes, 0644)
		if err != nil {
			log.Printf("failed writing to file: %s", err)
			return
		}
		log.Println("Successfully Uploaded File\n")
	}).Methods("POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9080", r))
}
