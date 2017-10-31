//Declare the server url
var serverUrl="http://localhost:8080/ajax/";

//Add the click event to the button
window.onload = function() {
    //Register click event listener onto the send button
    document.getElementById("sendButton").addEventListener('click',sendButtonClick);
    //Register the enter pressed event onto the text box
    document.getElementById("newwrittenmessage").addEventListener('keypress',sendOnEnterPress);
    //Set the focus to the textarea
    document.getElementById("newwrittenmessage").focus();
    //Start conversation
    //sendButtonClick();
}

//Function used to send the ajax request
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
        if (request.status >= 200 && request.status < 400) {
            //Parse data
            var data = JSON.parse(request.responseText);
            if(data.Error==""){
                //Add the user bubble
                if(textarea.value!="intitial_greetings"){
                    addBubble(textarea.value,"user");
                    //Empty the text area
                    textarea.value="";
                }

                //Wait 500ms to imitate typing
                setTimeout(function(){
                    //Add the chat bubble from the bot
                    addBubble(data.Answer,"bot");
                    //Empty the text area
                    textarea.value="";
                    //Scroll to bottom of the chat window
                    scrollToBottom()
                },500)

                //Scroll to bottom of the chat window
                scrollToBottom()
            }
        } else {
            // We reached our target server, but it returned an error

        }
        //start animation
        sendandreceivedAnimation();
    };

    request.onerror = function() {
        //console.log(request);
        // There was a connection error of some sort
        //start animation
            sendandreceivedAnimation();
    };

    //Send the data
    request.send(JSON.stringify({"Question":textarea.value}));
}

//Function used to roll to the bottom of the chat window
function scrollToBottom(){
    //Get the window
    var chatWindow = document.getElementById("messagesContainer");
    //Scroll to bottom
    chatWindow.scrollTop = chatWindow.scrollHeight;
}
//Function used to add a chat bubble
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

//Function used to caputure the button click
function sendButtonClick(){
    //Get the textarea
    var textarea=document.getElementById("newwrittenmessage");
    //Check if there is anything to send
    if(textarea.value.trim()!=""){
        //Send the request
        sendRequest();
        //Set the focus back onto the textarea
        textarea.focus();
    }
}
//Function used to send the text on enter press
function sendOnEnterPress(event){
    //If enter was pressed
    if(event.charCode==13){
        //Check if there is anything to send
        if(event.target.value.trim()!=""){
            //Send the request
            sendRequest();
        }
    }
}

//Function used to change the send button into a loading animation, and back
function sendandreceivedAnimation(){
    var text=document.getElementById("sendtext");
    var loading=document.getElementById("loading");
    var textarea=document.getElementById("newwrittenmessage");

    //If the text us hidden
    if(text.classList.contains("hidden")){
        //Make it visible
        text.classList.remove("hidden");
        loading.classList.add("hidden");
        //textarea.disabled=true;
    }else{
        //Make it hidden
        textarea.disabled=false;
        text.classList.add("hidden");
        loading.classList.remove("hidden");
        
    }
}