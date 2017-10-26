//Declare the server url
var serverUrl="http://localhost:8080/ajax/";

//Add the click event to the button
window.onload = function() {
    document.getElementById("sendButton").addEventListener('click',sendButtonClick);
    //Start conversation
    sendButtonClick();
}

function sendRequest(){
    var request = new XMLHttpRequest();

    var textarea=document.getElementById("newwrittenmessage");
    //start animation
    sendandreceivedAnimation();

    //Open connection
    request.open("POST", serverUrl);
    //Indicate this is a json
    request.setRequestHeader("Content-Type", "application/json");
    
    //Receive the response
    request.onload = function() {
        console.log(request);
        if (request.status >= 200 && request.status < 400) {
            //Parse data
            var data = JSON.parse(request.responseText);
            if(data.Error==""){
                //Add the user bubble
                if(textarea.value!="intitial_greetings"){
                    addBubble(textarea.value,"user");
                }
                addBubble(data.Answer,"bot");
                textarea.value="";
            }
        } else {
            // We reached our target server, but it returned an error

        }
        //start animation
        sendandreceivedAnimation();
    };

    request.onerror = function() {
        console.log(request);
        // There was a connection error of some sort
        //start animation
            sendandreceivedAnimation();
    };

    //Send the data
    request.send(JSON.stringify({"Question":textarea.value}));
}

function addBubble(text,owner){
    var side="left";
    //Change the side if the owner is the user
    if(owner=="user") side="right";
    //Get the parent container for the message
    var parentContainer=document.getElementById("messagesContainer");
    //Create parent element
    var messageParentDiv=document.createElement('div');
    //Create the message container
    var newmessagebox=document.createElement('div');

    //Add the text message to the box
    newmessagebox.innerHTML=text.trim();
    //Add the styleing
    newmessagebox.classList.add("messagebubble");
    newmessagebox.classList.add("message"+side);
    messageParentDiv.classList.add("indivualdmessage");
    //Add to the parent container
    messageParentDiv.appendChild(newmessagebox);
    parentContainer.appendChild(messageParentDiv);
}


function sendButtonClick(){
    //Get the textarea
    var textarea=document.getElementById("newwrittenmessage");
    //Check if there is anything to send
    if(textarea.value.trim()!=""){
        //Send the request
        sendRequest();
    }
}


function sendandreceivedAnimation(){
    var text=document.getElementById("sendtext");
    var loading=document.getElementById("loading");
    var textarea=document.getElementById("newwrittenmessage");

    if(text.classList.contains("hidden")){
        text.classList.remove("hidden");
        loading.classList.add("hidden");
        textarea.addAttribute("disabled");
    }else{
        textarea.removeAttribute("disabled");
        text.classList.add("hidden");
        loading.classList.remove("hidden");
        
    }
}