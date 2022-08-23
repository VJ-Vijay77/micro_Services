package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type RequestPayload struct {
	Action string `json:"action"`
	Auth AuthPayload `json:"auth,omitempty"`
}

type AuthPayload struct {
	Email string `json:"email"`
	Password string `json:"password"`
}


func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {

	payload := jsonResponse{
		Error:   false,
		Message: "Hit on Broker Service ",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)

}


func(app *Config) HandleSubmission(w http.ResponseWriter,r *http.Request) {

	var requestPayload RequestPayload

	err := app.readJSON(w,r,&requestPayload)
	if err != nil {
		app.errorJSON(w,errors.New("error in read json"))
		return
	}

	switch requestPayload.Action {
	case "auth":
		app.authenticate(w,requestPayload.Auth)
	default:
		app.errorJSON(w,errors.New("unknown action"))
	}


}

func (app *Config) authenticate(w http.ResponseWriter,a AuthPayload) {
	// create some json and sending it to auth mcroservice
	jsonData,_ := json.MarshalIndent(a,"","\t")


	//!call the service
	request,err := http.NewRequest("POST","http://authentication-service/authenticate",bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w,err)
		return
	}

	client := &http.Client{}
	response,err := client.Do(request)
	if err != nil {
		app.errorJSON(w,err)
		return
	}
	defer response.Body.Close()

	//? getting the correct status code in return 
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w,errors.New("invalid credentials"))
		return
	}else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w,errors.New("error calling auth service"))
		return
	} 

	//!create a varialble to read response.Body coming from authentication-service
		var jsonFromService jsonResponse

	//decode the json from the auth service
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w,err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w,err,http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Authenticated"
	payload.Data = jsonFromService.Data

	app.writeJSON(w,http.StatusAccepted,payload)

}