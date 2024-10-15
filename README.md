# Url Shortener

A URL shortening REST API built with Golang, Gin, MongoDB, and Redis. It uses JWT for authentication, with tokens securely stored in cookies. Key features include rate limiting, link expiration, and click tracking

https://url-shortener-go-u8dq.onrender.com/api

## Clone the repository

git clone [https://github.com/manlikehenryy/url-shortener-go.git](https://github.com/manlikehenryy/url-shortener-go.git)

## After cloning navigating to the directory

## Create a .env file, copy the content of example.env to the .env file

    cp example.env .env

## Install Dependencies

    go get ./...

## Run the app

    go run main.go

.

# REST API

## Signup

### Request

`POST /api/register`

     http://localhost:5000/api/register

     {
    "firstName": "",
    "lastName": "",
    "phone": "",
    "email": "",
    "password": ""
    }

### Response

    HTTP/1.1 201 CREATED
    Status: 201 CREATED
    Content-Type: application/json

    {
    "data": {
        "_id": "670eccd115ff67fa6d3fab2e",
        "firstName": "name",
        "lastName": "lastname",
        "email": "ss@gmail.com",
        "phone": "09022002000",
        "createdAt": "2024-10-15T21:13:04.3139111+01:00",
        "updatedAt": "2024-10-15T21:13:04.3139111+01:00"
    },
    "message": "Account created successfully"
    }

## Login

### Request

`POST /api/login`

    http://localhost:5000/api/login

    {
    "email": "",
    "password": ""
    }

### Response

    HTTP/1.1 200 OK
    Status: 200 OK
    Content-Type: application/json
   

    {
    "data": {
        "_id": "670eccd115ff67fa6d3fab2e",
        "firstName": "name",
        "lastName": "lastname",
        "email": "ss@gmail.com",
        "phone": "09022002000",
        "createdAt": "2024-10-15T20:13:04.313Z",
        "updatedAt": "2024-10-15T20:13:04.313Z"
    },
    "message": "Logged in successfully"
    }

## Shorten a url

### Request

`POST /api/url`

    http://localhost:5000/api/url

    token needs to be stored in cookies

    {
    "originalUrl": "https://longurl/jdjdjeuuuednffms/sjsjsjsjsjsjnnsnssssshh/msmsmsmssmsmsmsmmsmmsmsm",
    "expiration": 240  // 0 for no expiration
    }

### Response

    HTTP/1.1 201 Created
    Status: 201 Created
    Content-Type: application/json


    {
    "data": {
        "_id": "670ece9b15ff67fa6d3fab2f",
        "shortUrl": "localhost:5000/7761ea45",
        "originalUrl": "https://longurl/jdjdjeuuuednffms/sjsjsjsjsjsjnnsnssssshh/msmsmsmssmsmsmsmmsmmsmsm",
        "expiration": 240,
        "ClickCount": 0,
        "ClickDetails": [],
        "userId": "670eccd115ff67fa6d3fab2e",
        "createdAt": "2024-10-15T21:20:43.2041692+01:00",
        "updatedAt": "2024-10-15T21:20:43.2041692+01:00"
    },
    "message": "Url created successfully"
    }

## List all url

### Request

`GET /url`

    http://localhost:5000/url

    token needs to be stored in cookies

### Response

    HTTP/1.1 200 OK
    Status: 200 OK
    Content-Type: application/json


    {
    "data": [
        {
            "_id": "670ece9b15ff67fa6d3fab2f",
            "shortUrl": "7761ea45",
            "originalUrl": "https://longurl/jdjdjeuuuednffms/sjsjsjsjsjsjnnsnssssshh/msmsmsmssmsmsmsmmsmmsmsm",
            "expiration": 240,
            "ClickCount": 0,
            "ClickDetails": [],
            "userId": "670eccd115ff67fa6d3fab2e",
            "createdAt": "2024-10-15T20:20:43.204Z",
            "updatedAt": "2024-10-15T20:20:43.204Z"
        }
    ],
    "message": "Data fetched successfully",
    "meta": {
        "hasNextPage": false,
        "hasPrevPage": false,
        "nextPage": 0,
        "page": 1,
        "pageCount": 1,
        "perPage": 10,
        "total": 1
    }
    }

## Update a url

### Request

`POST /api/url/:urlId`

    http://localhost:5000/api/url/670ece9b15ff67fa6d3fab2f

    token needs to be stored in cookies

    {
    "originalUrl": "https://longurl/jdjdjeuuuednffms/sjsjsjsjsjsjnnsnssssshh/msmsmsmssmsmsmsmmsmmsmsm",
    "expiration": 300  // 0 for no expiration
    }

### Response

    HTTP/1.1 200 OK
    Status: 200 OK
    Content-Type: application/json


   {
    "message": "Url updated successfully"
   }

## Delete a url

### Request

`POST /api/url/:urlId`

    http://localhost:5000/api/url/670ece9b15ff67fa6d3fab2f

    token needs to be stored in cookies

### Response

    HTTP/1.1 200 OK
    Status: 200 OK
    Content-Type: application/json


   {
    "message": "Url deleted successfully"
   }