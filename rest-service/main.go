package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ottingbob/backend-developer-tests/rest-service/pkg/models"
	uuid "github.com/satori/go.uuid"
)

const (
	idParam          = "id"
	firstNameParam   = "firstName"
	lastNameParam    = "lastName"
	phoneNumberParam = "phoneNumber"
)

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func writeJsonResp(w http.ResponseWriter, people []*models.Person) {
	resp, err := json.Marshal(people)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}

func personGetByIdHandler(w http.ResponseWriter, r *http.Request) {
	personIdString := strings.Split(r.URL.Path, "/")[2]
	personId, err := uuid.FromString(personIdString)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	personById, err := models.FindPersonByID(personId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp, err := json.Marshal(personById)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}

func personGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received %s request on /people endpoint\n", r.Method)

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "The %s method isn't allowed on this endpoint\n", r.Method)
		return
	}

	// Check if there is an /:id on the route
	pathValues := strings.Split(r.URL.Path, "/")
	if len(pathValues) == 3 {
		personGetByIdHandler(w, r)
		return
	} else if len(pathValues) > 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Check if there are URL Query params:
	// In the case of a random query param we default to returning
	// no matched people records
	queryParams := r.URL.Query()
	if len(queryParams) > 0 {
		// Handle first and last name query params
		firstName, fn_ok := queryParams[firstNameParam]
		lastName, ln_ok := queryParams[lastNameParam]
		if fn_ok && ln_ok {
			nameParams := min(len(firstName), len(lastName))
			peopleByName := make([]*models.Person, 0)
			for i := 0; i < nameParams; i++ {
				morePeople := models.FindPeopleByName(firstName[i], lastName[i])
				peopleByName = append(peopleByName, morePeople...)
			}

			writeJsonResp(w, peopleByName)
			return
		}

		// Handle phone number query param
		phoneNumber, pn_ok := queryParams[phoneNumberParam]
		if pn_ok {
			peopleByPhone := make([]*models.Person, 0)
			for i := 0; i < len(phoneNumber); i++ {
				morePeople := models.FindPeopleByPhoneNumber(phoneNumber[i])
				peopleByPhone = append(peopleByPhone, morePeople...)
			}

			writeJsonResp(w, peopleByPhone)
			return
		}

		return
	}

	allPeople := models.AllPeople()
	writeJsonResp(w, allPeople)
}

func main() {
	fmt.Println("SP// Backend Developer Test - RESTful Service")
	fmt.Println()

	http.HandleFunc("/people/", personGetHandler)
	http.HandleFunc("/people", personGetHandler)
	port := 8080
	fmt.Printf("Starting REST server on port [%d]\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
