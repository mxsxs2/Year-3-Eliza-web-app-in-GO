package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)
//Structure for the json request
type Request struct{
    Question string
}

//Strustire for the json response
type Response struct{
    Answer string //The answer for the question
    Error string  //Error if the question is actually not a question
}



//The handler for http request-response
func handler(w http.ResponseWriter, r *http.Request) {
    //process the request and write back
    writeResponse(w,processRequest(r));
    //fmt.Fprintf(w, "Hi there, I love %s!", r)
}

//Function used to process the request
func processRequest(r *http.Request) Response{
    var request Request
    var response Response
    //Check if there is a body
    if(r.Body!=nil){
        //Try to decode the request body
        err:=json.NewDecoder(r.Body).Decode(&request)

        //Check if the encode vas succesful
        if(err==nil){
            if(request.Question=="intitial_greetings"){
                response.Answer="What is bothering you again?"    
            }else if(request.Question==""){
                response.Answer="You should say/ask something"
            }else{
                response.Answer="Well, tough luck. bye"
            }
            
            fmt.Println("Request: ",request.Question)
            fmt.Println("Answer: ",response.Answer)
            
        }else{
            response.Error=err.Error();
            fmt.Println(err)
        }
        
    }

    if(response.Answer=="" && response.Error==""){
        
        response.Error="Invalid Request"
        fmt.Println("Invalid request")
    }

    return response
}

//Function used to write the response back
func writeResponse(w http.ResponseWriter, response Response){
    // allow cross domain AJAX requests
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers")
    //Write response
    json.NewEncoder(w).Encode(response);
}

//Check if the request is correct
func isValidRequest(r *http.Request) bool{
    return r.Body == nil
}


func main() {
    //Set the handler for http
    http.HandleFunc("/", handler)
    //Listen on port 8080
    http.ListenAndServe(":8080", nil)
    fmt.Println("Listen started on port 8080");
}





func parserequest(s string){

}