package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

    "io"
) // "os"

// réponse à la requête GET 
func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

// lance le serveur local
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}


type event struct {
    ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`

	// id          string `json:"id"`
	// // spot       string `json:"Destination"`
	// createdTime string `json:"createdTime"`
}

type allEvents []event
var events = allEvents{
	{
        ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
        // id :"rec1VNeSJWYdZgoTo",
        // createdTime :"2018-05-31T00:16:16.000Z",
        // "fields":{"Surf Break":["Reef Break"],
        // "Destination":"Pipeline",
        // "Geocode":"eyJpIjoiUGlwZWxpbmUsIE9haHUsIEhhd2FpaSIsIm8iOnsic3RhdHVzIjoiT0siLCJmb3JtYXR0ZWRBZGRyZXNzIjoiRWh1a2FpIEJlYWNoIFBhcmssIEhhbGVpd2EsIEhJIDk2NzEyLCBVbml0ZWQgU3RhdGVzIiwibGF0IjoyMS42NjUwNTYyLCJsbmciOi0xNTguMDUxMjA0Njk5OTk5OTd9LCJlIjoxNTM1MzA3MDE5OTE1fQ==",
        // "Influencers":["reczMHu0P0DOPSyCG","recXaBgQfHWuPviL1"],"Magic Seaweed Link":"https://magicseaweed.com/Pipeline-Backdoor-Surf-Report/616/",
        // "Photos":[{"id":"attf6yu03NAtCuv5L","width":2233,"height":1536,"url":"https://v5.airtableusercontent.com/v3/u/24/24/1705413600000/SaUSFzjZJAJUK-B10NRoXQ/ekSr8mAa4w0e92OhyIepzn6gnt_lPJd8cPCTBb0J-SvCFUvxnHSr78IgqJSKElMVyfJ4-38Et9hYn1W-ii_BaeJr_Eip7ffiJ6r8sl4itm_D2424yzX-mfljdf0cnnao/-IfGTrQTmTwxMkB61M7qGE0Zod3GJhOQOhQKo-Atls0","filename":"thomas-ashlock-64485-unsplash.jpg","size":688397,"type":"image/jpeg","thumbnails":{"small":{"url":"https://v5.airtableusercontent.com/v3/u/24/24/1705413600000/krBBqyQ1RC42GuWYstpNqg/XrtBaBPlWNzvuheo3zaiV0MU7CjULPWnWkQinWk5tod9Oq0fUp_l74qmyj-ChaCr722MzENivgpuO9xURD9U9l7Z-4_wCfAjJAqHgkFAsQPR0CxAp9u0WzpPojGw-zGXaiol4HFdSSoxJZ5xFzt4Vw/q8J_AbSYun3K0TFOAbsWMC-u2KPOW1kJsfHgbFTiODY","width":52,"height":36},"large":{"url":"https://v5.airtableusercontent.com/v3/u/24/24/1705413600000/Tmz6f56ilB8qG5xK6t0dLA/kpeAYMipCB9Z66Sx1J2_Cpw-imbmU9HDFDvK_ywYQDYzFjSKKRdXTD-tjdT9oBgo4guugmZyrySej8ke3OEmcPVIMllgTYWyF_1ZC8guqy-zVSHXWKvxGfreRGvCaH13iZDxRhYKboZhgmdIqjZfcQ/CbHu5AnZEIgu6YHDyWM8SuibCFwNmzOGsu3lIJ-z8Zc","width":744,"height":512},"full":{"url":"https://v5.airtableusercontent.com/v3/u/24/24/1705413600000/9UIYWr3isEpYf-tSbNwJKA/j1Zad2nhT7KTzsTuedWgKvx_u5Xb8L7gD2nk2TeCVtnY1t3gE9-y_-6ysQV19zuTkB1fJIEhMGmoP5rJ2_MBqWRC14a-IDXIi2Bk0I14HYFNd19gbOeJd4imQfOu2BOb/HdGJFNKyzvOCyGMa9QVBRSfiZ__1Hn6TvXUgb4W5HZg","width":2233,"height":1536}}}],
        // "Peak Surf Season Begins":"2024-07-22",
        // "Destination State/Country":"Oahu, Hawaii",
        // "Peak Surf Season Ends":"2024-08-31",
        // "Difficulty Level":4,
        // "Address":"Pipeline, Oahu, Hawaii"}
    },
}

// créée un event en récupérant les datas 
func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event

    // en go, := est disponible uniquement dans une fonction
    // c'est une version raccourcie pour déclarer & initialiser une variable
	reqBody, err := io.ReadAll(r.Body)
	if err != nil { // nil = network interface layer (gère les protocoles?)
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	
    // Unmarshal permet de transformer de la data en byte en son format d'origine
    // CAD ici pour décoder de la data json
	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}