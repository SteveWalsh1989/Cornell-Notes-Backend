# FYP_Backend
This program serves as the backend for my final year project: 'Cornell Notes: A student Focused Note Taking Appliation'

## Features
It uses Gorilla/mux for its server
It uses sql to connect to a server running on my local machine
A Vue.Js frontend can be used to interact with the server


# API EndPoints 
## GET
[/folders](#GET-folders-) <br/>

## POST

___




### GET /folders/
Gets basic folder information for user 

**Parameters**

|          Name | Required |  Type   | Description                                                                                                                                                           |
| -------------:|:--------:|:-------:| -------------------------------------------------------------------------------------------------------------------------------------------------------|
|     `user_id` | required | string  | the user id for the user     |
