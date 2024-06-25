# Remote Notes

## Supported OS
currently the only supported OS is linux, however there is hope for windows and macos users in the future. 
 


### What is it?
> Remote Notes is a powerful, blazingly fast CLI file server. Users have unqiue keys to increase secruity, no password required!

### How Does It work?
Remote Notes uses advanced Encryption algorithms to encrypt the data inside a file, and upload it to a self hosted server, which it is then compressed. encrpytion and decryption happen on device, making the info inside of the filesquite secure. 

### Why?
This project solves a personal problem of mine, and thats easily accessing my notes from anywhere. I use a desktop computer and two laptops, a personal and work/school one. Paying for google cloud is expensive and many other options are just complex, or require a GUI, something that isnt possible for headless usage. This tool aims to solve that problem, simply run the startscript* (there is a bit more work but not much), and its ready to go

## Setup
You heard me sing the praises of how easy it is to setup, now let me show you.
### Client

for the client, it is simply just one command, all you need to know is the ipaddr/domain for your server. It should look something like this

`bash
$ ClientStartup http://ipaddr:Port
`

And just like that your client is ready to go! 
> yes the http is required 

### Server 
For the server, its a tad bit more complex, but not really. All you need to do is create the directory you want the notes to be stored in, as there is some weird behavior on immutable distros if the program creates the directory.

You may create the dir by typing 

`bash
$ mkdir /Path/To/Dir
`

After that, simply use the startup script and pass in the directory

`bash
$ ServerStartup /Path/To/Dir
`

And boom! your Remote Notes server is ready to go!

> Sidenote for setup, you will also need to move the executable to your bin folder, to be globally accessible. 

## Usage

### Client
The Client has three main commands, these being to upload, download and list

#### list 
List simply lists all possible files to download. You can pass in an optional directory to list the given directory, if none is given it simply lists the root Notes_Directory

#### Download
Downloads the given files, it can take multiple arguments at once

#### Upload
Uploads the given files, it can also take multiple arugments

#### Help
you can also type help to display the message

#### Examples 
`bash
$ Remote_Notes help
`

`bash
$ Remote_Notes list
`

### Server
Simply execute the server executable and it will run and startup
