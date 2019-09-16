# Go API 

<<<<<<< HEAD
A simple REST API with CRUD operations. The API will use an internal array, built and run on go version go1.12 but will probably work fine on earlier versions as well.

## How to build it

To build the api from source - you need to install go (use 1.12), pull the dependancies and compile.
```
user@local:~$ go get -d -v .
```
That will pull the dependancies based on the dependancie list in the repo.

```
user@local:~$ go build
```
Will compile the source code and create a binary for your target platform.
=======
A simple REST API with CRUD operations. The API will use an internal array, built and run on go version go1.10.4 (also tested and working fine on 1.12) but will probably work fine on earlier versions as well.
>>>>>>> 2b8c799181d27ba2e0ed45c06c78e174ceb41485


# Endpoints and usage:

## Health:
```
user@local:~$ curl 127.0.0.1:8080
"ok"
```

## List all users:
```
user@local:~$ curl 127.0.0.1:8080/users
[{"id":"1","firstname":"James","lastname":"Hetfield","age":56},{"id":"2","firstname":"Lars","lastname":"Ulrich","age":55},{"id":"3","firstname":"Kirk","lastname":"Hammett","age":56},{"id":"4","firstname":"Robert","lastname":"Trujillo","age":55}]
```

## Show single user by ID:
```
user@local:~$ curl 127.0.0.1:8080/users/1
{"id":"1","firstname":"James","lastname":"Hetfield","age":56}
```

```
user@local:~$ curl 127.0.0.1:8080/users/3
{"id":"3","firstname":"Kirk","lastname":"Hammett","age":56}
```

## Add user:
```
user@local:~$ curl -d '{"id":"5", "firstname":"Tom", "lastname": "Jones", "age":75}' -H "Content-Type: application/json" -X POST http://localhost:8080/users/5
[{"id":"1","firstname":"James","lastname":"Hetfield","age":56},{"id":"2","firstname":"Lars","lastname":"Ulrich","age":55},{"id":"3","firstname":"Kirk","lastname":"Hammett","age":56},{"id":"4","firstname":"Robert","lastname":"Trujillo","age":55},{"id":"5","firstname":"Tom","lastname":"Jones","age":75}]
```

## Delete user:
```
user@local:~$ curl -X DELETE http://localhost:8080/users/5
[{"id":"1","firstname":"James","lastname":"Hetfield","age":56},{"id":"2","firstname":"Lars","lastname":"Ulrich","age":55},{"id":"3","firstname":"Kirk","lastname":"Hammett","age":56},{"id":"4","firstname":"Robert","lastname":"Trujillo","age":55}]
```
