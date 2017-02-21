# API-Golang
this is simple project for creating API using Go. 
This is based on my opinion for the folder structure and the code structure.
If you find any thing that can be improved I'm open for pull request.

# Structure Folder
--> model : this is where I store database transaction like read update delete etc.<br><br>
--> controller : this is where business logic or handler to handle the incoming request.<br><br>
--> database : all database connection define here.<br><br>
--> logs : this folder is used for storing the log files.<br><br>
--> util : all the function that will be used in all others package.<br><br>

# Code Structure
<b> Model </b>
    - the file it self to describt the object.
    - instance : this package is for set the struct get it on differenct case such as mock and as the real struct.
    - interface : define the interface of the struct.
    - mock : define the mock struct for unit test purposes.
And there is `global` package for define the global response on API


<b> Util </b>
This package is used to declare the global function that used for all packages.


<b> Database</b>
This package is to define all the database connection.

<b> Controller </b>
this package is used for declare the business logic and the hanlder for handling the incoming request.

<b>Flow code </b>
In app.go : call all the connection database (database) -> pass it to model (model) -> call hanlder for hanlde the incoming request (controller)
Basically liket this : main->database-> model -> handler->controller

I hope this repository could help you to learn some basic developing API using Go.

