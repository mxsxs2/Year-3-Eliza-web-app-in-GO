package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Structure for the json request
type Request struct {
	Question string
}

//Strustire for the json response
type Response struct {
	Answer string //The answer for the question
	Error  string //Error if the question is actually not a question
}

//Function used to write the response back
func writeJSONResponse(w http.ResponseWriter, response Response) {
	// allow cross domain AJAX requests
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers")
	//Write response
	json.NewEncoder(w).Encode(response)
}

//Check if the request is correct
func isValidRequest(r *http.Request) bool {
	return r.Body == nil
}

//Handler for serving the css and js locally
func serveCSSandJS() {
	//Adapted from https://stackoverflow.com/questions/43601359/how-do-i-serve-css-and-js-in-go-lang
	//Serve the css files
	//http.HandleFunc("/css/", asd)
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./www/css"))))
	//Serve the JavaScript files
	http.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("./www/js"))))
}
func asd(w http.ResponseWriter, r *http.Request) {
	fmt.Println("css", r.URL)
}

//Handler for favicon
func serveFavicon(w http.ResponseWriter, r *http.Request) {
	//Icon from https://m.veryicon.com/icons/application/ios7-style-metro-ui/metroui-folder-os-game-center.html
	//Serve the favicon ico
	http.ServeFile(w, r, "favicon.ico")
}

//Handler for the landing page
func landingPageHandler(w http.ResponseWriter, r *http.Request) {
	//Set the html hconent type
	w.Header().Set("Content-Type", "text/html")
	//Serve the index file
	http.ServeFile(w, r, "www/index.html")
}

//Handler for the ajax request
func ajaxHandler(w http.ResponseWriter, r *http.Request) {
	//Try to initialize eliza
	if bot, err := intializeEliza(); err == nil {
		//process the request and write back
		writeJSONResponse(w, processRequestWithBot(r, bot))
	} else {
		fmt.Println(err)
		//Let the user know about the error
		writeJSONResponse(w, Response{"Sorry i can not talk right now.", ""})
	}

}

func main() {
	//Add the handler function for the landing page
	http.HandleFunc("/", landingPageHandler)
	//Add bootstrap file handling
	serveCSSandJS()
	//Add the handler for serving the favicon
	http.HandleFunc("/favicon.ico", serveFavicon)
	//Add the handler for the ajax
	http.HandleFunc("/ajax/", ajaxHandler)
	fmt.Println("Listen started on port 8080")
	//Listen on port 8080
	http.ListenAndServe(":8080", nil)
	
}
