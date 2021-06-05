package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	log.Printf("port: " + port)
	http.HandleFunc("/user-fullname", getFullName)
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/user-interests", getUserInterests)
	http.HandleFunc("/user-info-id", getUserInfoById)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getFullName(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Yersultan Assylbekov")
	w.WriteHeader(http.StatusOK)
}

func getUserInfoById(w http.ResponseWriter, r *http.Request) {

	UserId, ok1 := r.URL.Query()["user_id"]
	if !ok1 || len(UserId[0]) < 1 {
		json.NewEncoder(w).Encode(nil)
		w.WriteHeader(404)
		return
	}
	userId := UserId[0]
	var userInfo = ""
	if userId == "1" {
		userInfo = "{'name': 'First User', 'age': '22', 'weight': '55'}"
	} else {
		userInfo = "{'name': 'Other User', 'age': '23', 'weight': '56'}"
	}
	json.NewEncoder(w).Encode(userInfo)
	return

}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := []user_item{
		{
			full_name: "Yersultan Assylbekov",
			interests: []string{"football", "chess", "guitar", "travel"},
			age:       22,
			weight:    55.4,
		},
		{
			full_name: "Aibyn Rassulov",
			interests: []string{"swimming", "piano", "travel"},
			age:       21,
			weight:    57.4,
		},
	}
	json.NewEncoder(w).Encode(users)
	return

}

func getUserInterests(w http.ResponseWriter, r *http.Request) {
	hobbies := []string{"football", "chess", "guitar", "travel"}
	json.NewEncoder(w).Encode(hobbies)
	return
}

type user_item struct {
	full_name string
	interests []string
	age       int
	weight    float64
}
