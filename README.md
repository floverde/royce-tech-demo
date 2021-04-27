# Royce Tecnology Demo - User API
This is a demo project commissioned by Royce Technology.

It is a stand-alone web service that exposes an external
a REST interface that implements CRUD operations _(create/read/update/delete)_ on a **User** entity.

### User Model
This is the data model used for users.
```
{
  "id": "xxx",                  // user ID (must be unique)
  "name": "backend test",       // user name
  "dob": "",                    // date of birth
  "address": "",                // user address
  "description": "",            // user description
  "createdAt": ""               // user created date
  "updatedAt": ""               // user updated date
}
```

### REST endpoints
The REST interface exposes the following endpoints:
- FIND all users
 - `HTTP GET /users`
- FIND a user by ID
 - `HTTP GET /users/:id`
- CREATE a new user
 - `HTTP POST /users`
- UPDATE an existing user
 - `HTTP PATCH /users/:id`
- DELETE an existing user
 - `HTTP DELETE /users/:id`
- FIND all places associated with a user's address.
 - `HTTP GET /users/:id/places`
 
The last endpoint goes beyond classic CRUD operations, but provides an example of integration with the [Mapbox](https://www.mapbox.com/api-documentation/) APIs.

### How to build
A `Makefile` is included in the project through which you can perform certain operations.

To build the project you can use this command:
```
make win_build
```
If you are using a Window operation system
```
make unix_build
```

### How to test
To run unit tests use this command:
```
make test
```

### How to run
After compiling the project, there is a command to run it:
```
make run
```
Or simply run the executable `sample-rest-api` inside the `target` folder.

It will start a stand-alone server that will respond to http://localhost:8080/.

### View code documentation
To view the project code documentation, run this command:
```
make godoc
```
Another standalone server will be started, this time on port 6060.

Clicking on this [link](http://localhost:6060/pkg/roycetechnology.com/floverde/sample-rest-api/) will open a browser window with the documentation of the project code.

## Usage examples

Once compiled and executed, the application will start a small stand-alone web server listening on port 8080, which will respond to the endpoints listed in the previous section.

For the sake of convenience, we list a few examples of calls that can be made to test the correct functioning of the application.

#### 1st use case: *Creation of a new user*

To create a new user, we can make the following request through any HTTP client:

Request:
```
POST http://localhost:8080/users

{
   "name": "John Smith",
   "dob": "1990-01-01T00:00:00+00:00",
   "address":"2 Lincoln Memorial Circle NW",
   "description": "A nice guy"
}
```

Response:
```
HTTP 201 Created

{
  "id": 1
  "name": "John Smith",
  "dob": "1990-01-01T00:00:00+00:00",
  "address":"2 Lincoln Memorial Circle NW",
  "description": "A nice guy",
  "createdAt" "2021-04-28T10:30:00+00:00",
  "updatedAt" "2021-04-28T10:30:00+00:00"
}
```

#### 2st use case: *Get list of all users*

If we imagine that we have created another user, making this request will give us a list of all the users in the database.

Request:
```
GET http://localhost:8080/users
```

Response:
```
HTTP 200 OK

[
  {
    "id": 1,
    "name": "John Smith",
    "dob": "1990-01-01T00:00:00+00:00",
    "address":"2 Lincoln Memorial Circle NW",
    "description": "A nice guy",
    "createdAt" "2021-04-28T10:30:00+00:00",
    "updatedAt" "2021-04-28T10:30:00+00:00"
  },
  {
    "id": 2,
	"name": "Ashley Bennett",
	"dob": "1995-10-24T00:00:00+00:00",
    "address":"111 S Grand Ave Los Angeles CA",
    "description": "A lovely girl",
    "createdAt" "2021-04-28T11:40:15+00:00",
    "updatedAt" "2021-04-28T11:40:15+00:00"
  }
]
```

#### 3st use case: *Get a single user by knowing his ID*

We can also request the details of an single user if we know their ID.

Request:
```
GET http://localhost:8080/users/1
```

Response:
```
HTTP 200 OK

{
  "id": 1
  "name": "John Smith",
  "dob": "1990-01-01T00:00:00+00:00",
  "address":"2 Lincoln Memorial Circle NW",
  "description": "A nice guy",
  "createdAt" "2021-04-28T10:30:00+00:00",
  "updatedAt" "2021-04-28T10:30:00+00:00"
}
```

But for example, if we request a user that doesn't exist (e.g. requesting user ID 999) we get this:

Request:
```
GET http://localhost:8080/users/999
```

Response:
```
HTTP 404 Not Found

{
   "error": "User not found!"
}
```

#### 4st use case: *Update user details*

Suppose we want to update some details about one of our users. For example, we have changed our opinion about John.

Request:
```
PATCH http://localhost:8080/users/1

{
   "description": "Not as nice as I thought"
}
```

Response:
```
HTTP 200 OK

{
  "id": 1
  "name": "John Smith",
  "dob": "1990-01-01T00:00:00+00:00",
  "address":"2 Lincoln Memorial Circle NW",
  "description": "Not as nice as I thought",
  "createdAt" "2021-04-28T10:30:00+00:00",
  "updatedAt" "2021-04-28T10:30:00+00:00"
}
```

#### 5st use case: *Deleting a user*

If we wanted to remove Ashley from the system instead, we could make this other HTTP request:

Request:
```
DELETE http://localhost:8080/users/2
```

Response:
```
HTTP 200 OK
```

#### 6st use case: *Knowing the places associated with the user's address*

As an example of integration with Mapbox APIs, we can know the places associated with the user's address:

Request:
```
GET http://localhost:8080/users/1/places
```

Response:
```
HTTP 200 OK

{
  "address": "2 Lincoln Memorial Circle NW",
  "places": {
    // information retrieved from Mapbox APIs
  }
}
```

## Class Diagram

As a brief introduction to the structure of the application, we provide below a class diagram showing the relationship between the various classes of the application.

![](https://github.com/floverde/royce-tech-demo/blob/main/doc/class-diagram.svg)

_Author: Fabrizio Lo Verde_