package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/eduardonunesp/goaiml"
)

//Function used to initialize eliza
func intializeEliza() (*goaiml.AIML, error) {
	//Load the aim parser
	aiml := goaiml.NewAIML()
	//Load the aiml file aka Eliza in.
	if err := aiml.Learn("eliza.aiml"); err != nil {
		//If elize could not be loaded in for some reason
		return nil, errors.New("Eliza could not be loaded")
	}
	//Return nil on success
	return aiml, nil
}

//Function used to process the request
func processRequestWithBot(r *http.Request, bot *goaiml.AIML) Response {
	var request Request
	var response Response
	//Check if there is a body
	if r.Body != nil {
		//Try to decode the request body
		err := json.NewDecoder(r.Body).Decode(&request)

		//Check if the encode vas succesful
		if err == nil {
			if request.Question == "intitial_greetings" {
				response.Answer = "What is bothering you again?"
			} else if request.Question == "" {
				response.Answer = "You should say/ask something"
			} else {
				//Get the answer from the bot and check if there was any errors
				if resp, rerr := bot.Respond(request.Question); rerr == nil {
					//If there was no error then set the response answer
					response.Answer = resp
				} else {
					//If there was an error then set the error
					response.Error = rerr.Error()
					fmt.Println("Response error:", rerr)
				}
			}

			fmt.Println("Request: ", request.Question)
			fmt.Println("Answer: ", response.Answer)

		} else {
			response.Error = err.Error()
			fmt.Println(err)
		}

	}

	if response.Answer == "" && response.Error == "" {

		response.Error = "Invalid Request"
		fmt.Println("Invalid request")
	}

	return response
}
