package main

import (
	"bytes"
	"github.com/gorilla/mux"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"net/http"
	"strconv"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/download/{image}", func(w http.ResponseWriter, r *http.Request) {
		//img, err := ioutil.ReadFile("Observation.png")
		//if err!=nil{
		//	log.Print(err.Error())
		//	return
		//}
		v := mux.Vars(r)
		if v["image"] == "blue" {
			m := image.NewRGBA(image.Rect(0, 0, 240, 240))
			blue := color.RGBA{0, 0, 255, 255}
			draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)
			var img image.Image = m
			buffer := new(bytes.Buffer)
			if err := jpeg.Encode(buffer, img, nil); err != nil {
				log.Println("unable to encode image.")
			}
			w.Header().Set("Content-Type", "image/jpeg")
			w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
			if _, err := w.Write(buffer.Bytes()); err != nil {
				log.Println("unable to write image.")
			}
		} else if v["image"] == "red" {
			m := image.NewRGBA(image.Rect(0, 0, 240, 240))
			red := color.RGBA{255, 0, 0, 255}
			draw.Draw(m, m.Bounds(), &image.Uniform{red}, image.ZP, draw.Src)
			var img image.Image = m
			buffer := new(bytes.Buffer)
			if err := jpeg.Encode(buffer, img, nil); err != nil {
				log.Println("unable to encode image.")
			}
			w.Header().Set("Content-Type", "image/jpeg")
			w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
			if _, err := w.Write(buffer.Bytes()); err != nil {
				log.Println("unable to write image.")
			}
		}
	}).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9080", r))
}
