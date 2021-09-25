# Chatroom

### Prerequisite
* Must have Docker and make installed on system

### Usage
1. Clone this repo 
2. Run make help to perform suitable operation 
3. Run Either by Binary or Docker 

Remove all auth handlers to bypass authentication and set username as clientid in gin context


### Features

* Any User Can Create/Join room
* Admin Can Delete Room 
* Admin Can List User in Room 
* Admin Can Remove User from room
* Any User Can join Rooms
* Any User can List Rooms
* Chat history is stored in postgresql in demo.chathistory
* User is identified by username and socket connection
* Username and password is stored in postgresql along with admin flag
* To access Url get JWT access token from /token endpoint passing username and password
* All endpoints are authenticated by oauth
* Admin is identified by token and username as admin
* Rooms and user joined in rooms are stored in memory

#### Note:- Need Browser that supports passing Bearer token to use authentication and use html for chats

#### Reference 
1) https://github.com/gorilla/websocket/tree/master/examples/chat

#### Postan collection added for reference

