# Eliza Web App
Eliza chat bot written in GO language as a web app, with a simple HTML/JS chat.

### How to install and run GO

To install, simply go to GO's website and download the installer and run it: https://golang.org/

When the installation is done, the "go" command is going to be avaialable in terminal.(You have to restart an opened terminal)

### How to use this repository

To run any of the applications, navigate to the desired folder and type this command: 
```
go build 
```

The previous command will compile the go file into a runnable.

Once the runnable is created then it can be run in terminal, for example: 
```
./Year-3-Eliza-web-app-in-GO.exe
```

## How it works
### Webserver
The webserver listens on port 8080. 
The landing page is the index.html page from the www folder. 
This page is the main chat window, which sends the messages to the server with AJAX.
The AJAX request are sent to ```http://domain:8080/ajax/```which is not a real directory, whoever the program recognizes the path and begins the initialization of eliza. 
### File structure
```
--www
    --css
        --index.css
    --js
        --index.js
    --index.html
```
### AIML Parser
The aiml parser is originally from https://github.com/eduardonunesp/goaiml
I forked and tweaked this parser to work with my project.
The modifications includes
* Fix ```<star/>``` tag replaces
* Fix recognition of ```<srai>``` tags

### Eliza
Once Eliza is initialized. It reads in the ```eliza.aiml``` which contains the the categories, patterns and templates for parsing the request and chosing a right answer. 
The original file for Eliza is from https://github.com/massyn/ELIZA which I altered for better communication and also fixed issues with the ```<star/>``` tags.

### Simple chat
The chat is a simple "window" in the broswer where the user can send a message to Eliza and she answers.
It is written in HTML and JavaScript. The communication with Eliza is happening throgh AJAX calls which delivers JSON responses/request.



# References
* What is eliza - https://en.wikipedia.org/wiki/ELIZA
* AIML - http://www.alicebot.org/aiml.html
* AIML Parser - https://github.com/eduardonunesp/goaiml
* Simple webserver - https://github.com/mxsxs2/Go-Problem-set-Web-applications
* JSON - https://blog.golang.org/json-and-go
* Original Eliza AIML - https://github.com/massyn/ELIZA