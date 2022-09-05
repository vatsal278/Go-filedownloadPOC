package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vatsal278/go-redis-cache"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	cacher := redis.NewCacher(redis.Config{
		Addr: "localhost:9096",
	})
	r := mux.NewRouter()
	r.HandleFunc("/download/{image}", func(w http.ResponseWriter, r *http.Request) {
		v := mux.Vars(r)
		b, err := cacher.Get(v["image"])
		if err != nil {
			log.Print(err.Error())
		}
		w.Header().Set("Content-Disposition", "attachment; filename=Observation.png")
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Length", fmt.Sprint(len(b)))
		if _, err := w.Write(b); err != nil {
			log.Println("unable to write image.")
		}
	}).Methods("GET")
	r.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("File Upload Endpoint Hit")
		file, handler, err := r.FormFile("file")
		if err != nil {
			log.Println("Error Retrieving the File")
			log.Println(err)
			return
		}
		defer file.Close()
		log.Printf("Uploaded File: %+v\n", handler.Filename)
		tempFile, err := ioutil.TempFile("", handler.Filename)
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		tempFile.Write(fileBytes)
		cacher.Set(handler.Filename, tempFile, 0)
		log.Println("Successfully Uploaded File\n")
	}).Methods("POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9080", r))
}
