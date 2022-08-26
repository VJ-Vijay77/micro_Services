package main

import (
	"log-service/data"
	"net/http"
)


type JOSNPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}



func(app *Config) WriteLog(w http.ResponseWriter,r *http.Request) {
	// read json into a variable
	var requstPayload JOSNPayload

	_ = app.readJSON(w,r,&requstPayload)


	//inserting data
	event := data.LogEntry{
		Name: requstPayload.Name,
		Data: requstPayload.Data,
	}

	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		app.errorJSON(w,err)
		return
	}

	resp := jsonResponse{
		Error: false,
		Message: "Logged",
	}
	app.writeJSON(w,http.StatusAccepted,resp)

}