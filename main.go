package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

    "io"
    "encoding/json"
) // "os"

// réponse à la requête GET 
func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home Team!")
}

// lance le serveur local
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)  // ici quand url de la requête fini par :8080/ lance la fonction homelink
    router.HandleFunc("/event", createEvent).Methods("POST")
    router.HandleFunc("/events", getAllEvents).Methods("GET")
    router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
    router.HandleFunc("/events/{id}", updateEvent).Methods("PATCH")
    router.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}


type event struct {
    ID          string `json:"ID"`
	Name  		string  `json:"Name"`
    City  		string `json:"City"`
    Geocode 	string  `json:"Geocode"`
    Address 	string  `json:"Address"`
    AdditionnalInformations string `json:"Informations"`
    WavesTypes 	string  `json:"WavesTypes"`
    DifficultyLevel string `json:"DifficultyLevel"`
    PeakSurfSeasonStart string `json:"PeakSurfSeasonStart"`
    PeakSurfSeasonEnd string `json:"PeakSurfSeasonEnd"`
    ImageURL string `json:"ImageURL"`
    Liked bool `json:"Liked"`
}

type allEvents []event
var events = allEvents{
        {
		ID:                      "1",
		Name:                    "Festival de Surf",
		City:                    "Oceanville",
		Geocode:                 "123.456,789.012",
		Address:                 "Plage principale",
		AdditionnalInformations:  "Événement annuel",
		WavesTypes:              "Grandes vagues",
		DifficultyLevel:         "Difficile",
		PeakSurfSeasonStart:     "2024-06-01",
		PeakSurfSeasonEnd:       "2024-06-10",
		ImageURL:                "https://images.mnstatic.com/ab/e2/abe2a17c724d7348e0831f4437ba5ba0.jpg?quality=75&format=pjpg&fit=crop&width=980&height=880&aspect_ratio=980%3A880",
		Liked:                   false,
    },
	{
		ID:                      "2",
		Name:                    "Compétition de Surf",
		City:                    "Côte Azur",
		Geocode:                 "45.678,23.901",
		Address:                 "Plage de la compétition",
		AdditionnalInformations:  "Ouvert à tous les surfeurs",
		WavesTypes:              "Vagues moyennes",
		DifficultyLevel:         "Intermédiaire",
		PeakSurfSeasonStart:     "2024-07-15",
		PeakSurfSeasonEnd:       "2024-07-25",
		ImageURL:                "https://images.pexels.com/photos/7166574/pexels-photo-7166574.jpeg?cs=srgb&dl=pexels-nathan-shingleton-7166574.jpg&fm=jpg",
		//Liked:                   true,
	},
}


// créée un event en récupérant les datas 
func createEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "On entre dans createEvent")
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

    events = append(events, newEvent) // ajoute le nouvel évènement à la liste d'évènements
	w.WriteHeader(http.StatusCreated) // indique que la création de l'événement a été réussie en définissant le code de statut de la réponse à "201 Created"

    // Encode l'événement nouvellement créé en JSON et le renvoie dans le corps de la réponse HTTP (w)
	json.NewEncoder(w).Encode(newEvent)
}

// get an event based on the number at the end of the url
func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"] // récupère et stock l'id (numéro) dans l'url de la requête après /events/

	for _, singleEvent := range events {    // _, 
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

// get all the events
func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}

// modify elements of an event
func updateEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var updatedEvent event

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedEvent)

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			singleEvent.Name = updatedEvent.Name
			singleEvent.City = updatedEvent.City
			events = append(events[:i], singleEvent)
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

// delete an event
func deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			events = append(events[:i], events[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", eventID)
		}
	}
}