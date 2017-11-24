# Eliza Web App
Eliza chat bot written in GO language as a web app, with a simple HTML/JS chat.
The application is written for Data Representation module in year 3 of Software Development(2017) course at Galway Mayo Institute of Technology, Galway Campus. 

### How to install and run GO

To install, simply go to GO's website and download the installer and run it: https://golang.org/

When the installation is done, the "go" command is going to be avaialable in terminal.(You have to restart an opened terminal)

### How to use this repository

#### Get dependencies
In order to be able to compile tha application the AIML parser package has to be downloaded.
To download the AIML parser, run the following command:
```
go get github.com/mxsxs2/goaiml
```
This will download the package from github into ```$GOPATH\src\github.com\mxsxs2\goaiml```
#### Build the application
To build the application, navigate to the desired folder and type this command: 
```
go build 
```

The previous command will compile the go project into a runnable.

Once the runnable is created then it can be run in terminal, for example: 
```
./Year-3-Eliza-web-app-in-GO.exe
```
##### When the port 8080 is taken
If the port 8080 is taken, the application opens and closes straight away with an error message. The port has to be freed before the application starts.

## How it works
### Webserver
The webserver listens on port 8080. 
The landing page is the index.html page from the www folder. 
This page is the main chat window, which sends the messages to the server with AJAX.
The AJAX request are sent to ```http://domain:8080/ajax/```which is not a real directory, however the program recognizes the path and begins the initialization of eliza. 
### File structure
```
--www
    --css
        --index.css
    --js
        --index.js
    --imga
        --blurry-background-1.jpg
    --index.html
```
### AIML Parser
The aiml parser is originally from https://github.com/eduardonunesp/goaiml
I forked and tweaked this parser to work with my project.
The modifications includes
* Added comment for every statement to support better understanding of the library
* ```processStar``` function now replaces every occurance of the ```<star/>``` tag in the templates before the random
* Added reflection support to the ```<star/>``` tag
* Added pre processor for better pattern matching
* Fix recognition of ```<srai>``` tags

#### How it works
1. The parser loads in the AIML file on every request, which is helpful since the server does not have to be restarted if the AIML file is changed.
2. When a request is sent to the parser it searches through the categories in the AIML file top to bottom.
3. Once match found the tags (srai,star,bot,set) are parsed.
4. On a random tag the parser choses one item and returns back to the fucntion which called the AIML parser.

### Eliza
Once Eliza is initialized. It reads in the ```eliza.aiml``` which contains the categories, patterns and templates for parsing the request and chosing a right answer. 
The original file for Eliza is from https://github.com/massyn/ELIZA which I altered for better communication and also fixed issues with the ```<star/>``` tags.

### Simple chat
The chat is a simple "window" in the browser where the user can send a message to Eliza and she answers.
It is written in HTML, CSS and JavaScript.
I decided not to use any thrid party library for the CSS or JavaScript as the plain version of them shows better understanding of the given solution.(Definetely the XMLHttpRequest part.)
The communication with Eliza is happening throgh AJAX calls which delivers JSON responses/request.



# References
* What is eliza - https://en.wikipedia.org/wiki/ELIZA
* AIML - http://www.alicebot.org/aiml.html
* AIML Parser - https://github.com/eduardonunesp/goaiml
* Simple webserver - https://github.com/mxsxs2/Go-Problem-set-Web-applications
* JSON - https://blog.golang.org/json-and-go
* Original Eliza AIML - https://github.com/massyn/ELIZA
* Background picture - http://wlpapers.com/blurry-background.html