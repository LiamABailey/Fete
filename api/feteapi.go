package api

import (
  "net/http"

  "github.com/gorilla/mux"
  "go.mongodb.org/mongo-driver/bson/primitive"
)


func BuildRouter() mux.Router {
  r := mux.NewRouter()
  r.HandleFunc("/featuredefinition", postFeatureDefinition).Methods("POST")
  r.HandleFunc("/featuredefinition/{featureDefinitionId}", getFeatureDefinition).Methods("GET")
  r.HandleFunc("/featuredefinition", getMatchingFeatureDefinitions).Methods("GET")
  r.HandleFunc("/featuredefinition/{featureDefinitionId}", patchFeatureDefinition).Methods("PATCH")
  return r
}

// post a feature definition
func postFeatureDefinition(w http.ResponseWriter, r *http.Request) {

}

// get a feature definiton
func getFeatureDefinition(w http.ResponseWriter, r *http.Request) {
  searchFdid = mux.Vars(r)["featureDefinitionId"]
}

// get all feature definitions matching the query params
func getMatchingFeatureDefinitions(w http.ResponseWriter, r *http.Request) {

}

// support for partial or full updates of an existing feature definition.
func patchFeatureDefintion(w http.ResponseWriter, r *http.Request) {
  updateFdid = mux.Vars(r)["featureDefinitionId"]
}
