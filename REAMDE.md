# Eliza Web App
Eliza written in GO language as a web app, with a simple HTML/JS chat.

### How to install and run GO

To install, simply go to GO's website and download the installer and run it: https://golang.org/

When the installation is done, the "go" command is going to be avaialable in terminal.(You have to restart an opened terminal)

### How to use this repository

To run any of the applications, navigate to the desired folder and type this command: 
`go build <name_of_go_file>`

The previous command will compile the go file into a runnable.

Once the runnable is created then it can be run in terminal, for example: 
`./<name_of_runnable>` 

## How it works
### Eliza
Eliza "listens" at port 8080 and capable of communication through JSON with practically any other application throught the internet.
Once eliza receives the "intitial_greetings" request it begins communication with the user.
At the moment Eliza's main feature is not implemented, Therefore the only thing she says is "Well, tough luck, bye". This is going to be implemented in the near future.
### Simple chat
The chat is a simple "window" in the broswer where the user can send a message to Eliza and eliza answers.
It is written in HTML and JavaScript. The communication with Eliza is happening throgh AJAX calls which delivers JSON responses/request.