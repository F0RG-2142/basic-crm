# basic-crm

A basic crm API as my first project woth Go Fiber instead of the standard library or gorilla/mux.
It is a customer relations management app that stores leads consisting of 
name, company, email, and phone number in a sqlite3 databse that is accessed with the gorm.io package.
The API has: get all, get one, update, and delete functions in order to manage entries in the sql database.
it listens on http://localhost:8080/api/lead/ for http requests.