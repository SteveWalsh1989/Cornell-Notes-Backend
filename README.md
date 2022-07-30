# Cornell Notes backend repo
This program serves as the backend for my final year project: 'Cornell Notes: A student Focused Note Taking Appliation' written in Go

It is a note taking application based around the cornel note taking method including:
- create new user profile
- create folders for notes
- create notes
- Notes contain cues, answers and summary sections. Can also be tagged
- When adding new Cues/answers user can earn points
- Notes contain option to review which opens a quiz interface where user can test knowledge and earn points
- Earned Points are reflected in a badge area as incentive to use the applications features and reflect learning levels 

Grade Recieved was 82% for the project research phase and 92% for the implementation phase


![Cornell Notes demo](cn-demo.gif)


## Features
It uses Gorilla/mux for its server
It uses sql to connect to a server running on my local machine
A Vue.Js frontend is currently used to interact with the server

## Additional Information
Cornell_Notes.pdf contains all information about projects background research, planning and implementation
