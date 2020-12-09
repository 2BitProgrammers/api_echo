# 2bitprogrammers/api_echo
An HTTP Echo API which returns JSON info about the requested endpoints. 
This is meant to be used for testing and debugging purposes only.

Essentially, this app will output whatever is sent to the server. 
This includes:
* web client info - ip, port, user agent
* method - GET, POST, etc.
* headers - Content-Type, Cookie, etc.
* payload - if one is sent, the string of the request body data

The API listens on port:  1234


## Run as Standalone GoLang App
This will run the application with the go application.  It assumes that you have installed your golang environment correctly.

```bash
$ cd src
$ go run main.go

2bitprogrammers/api_echo v2018.11a
www.2BitProgrammers.com
Copyright (C) 2020. All Rights Reserved.

Listening on:  1234

CTRL+C
```


## Run within Docker 
This will run the components on your local system without using minikube or kubernetes.

### Building the Docker Image
```bash
$ docker build . -t 2bitprogrammers/api_echo

Sending build context to Docker daemon  114.7kB
Step 1/11 : FROM golang:alpine AS builder
 ---> b3bc898ad092
Step 2/11 : ENV GO111MODULE=on     CGO_ENABLED=0     GOOS=linux     GOARCH=amd64
 ---> Using cache
 ---> 8462443c0070
Step 3/11 : WORKDIR /build
 ---> Using cache
 ---> 99600623930c
Step 4/11 : COPY $PWD/src/go.mod .
 ---> Using cache
 ---> 04466d71935c
Step 5/11 : COPY $PWD/src/main.go .
 ---> Using cache
 ---> 91a1e7c623ba
Step 6/11 : RUN go mod download
 ---> Using cache
 ---> ec172095ad7c
Step 7/11 : RUN go build -o api_echo .
 ---> Using cache
 ---> 61393a21a25b
Step 8/11 : FROM scratch
 --->
Step 9/11 : WORKDIR /
 ---> Using cache
 ---> a66c59ea194a
Step 10/11 : COPY --from=builder /build/api_echo .
 ---> Using cache
 ---> 22f5a780ab79
Step 11/11 : ENTRYPOINT [ "/api_echo" ]
 ---> Using cache
 ---> 3adb0272900e
Successfully built 3adb0272900e
Successfully tagged 2bitprogrammers/api_echo:latest
SECURITY WARNING: You are building a Docker image from Windows against a non-Windows Docker host. All files and directories added to build context will have '-rwxr-xr-x' permissions. It is recommended to double check and reset permissions for sensitive files and directories.
```

### Image Status
```bash
$ docker images

REPOSITORY                     TAG          IMAGE ID           CREATED              SIZE
2bitprogrammers/api_echo       latest       3adb0272900e       4 minutes ago       6.56MB
```

### Running the Container
```bash
$ docker run --rm --name "api_echo" -p 1234:1234 2bitprogrammers/api_echo 

2bitprogrammers/api_echo v2018.11a
www.2BitProgrammers.com
Copyright (C) 2020. All Rights Reserved.

Listening on:  1234

CTRL+C
```

### Check the Container Status (docker)
```bash
$ docker ps

CONTAINER ID    IMAGE                           COMMAND             CREATED              STATUS              PORTS                      NAMES
0e421af8fcfd    2bitprogrammers/api_echo        "/api_echo"         31 seconds ago       Up 27 seconds       0.0.0.0:1234->1234/tcp     api_echo
```

### Watch Container Logs
```bash
$ docker logs -f 2bitprogrammers/api_echo

2bitprogrammers/api_echo v2018.11a
www.2BitProgrammers.com
Copyright (C) 2020. All Rights Reserved.

Listening on:  1234

2020-12-09 19:48:07.7554606 +0000 UTC    GET             /abc    200     172.17.0.1      Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36
2020-12-09 19:48:13.1801922 +0000 UTC    GET             /test   200     172.17.0.1      Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36
2020-12-08 17:30:31.6822105 -0800 PST    POST            /test/my/uri   200     127.0.0.1       curl/7.55.1
2020-12-08 17:31:14.3062692 -0800 PST    POST            /test/my/uri   200     127.0.0.1       curl/7.55.1

CTRL+C
```


### Stopping the Container
```bash
$ docker stop api_echo
```


## Using the API
For the below examples, we will assume the following:
* Server:  locahost (127.0.0.1)
* Bind Port: 1234
* Method: GET, POST, etc.
* URI: Can be anything, but using:  /test/my/uri 
* Body Data:   Can be anything, but using: _{ "name": "bubba", "int_value": 987 }_

A simple GET request:
```bash
$ curl http://localhost:1234/test/my/uri 

{"statusCode":200,"headers":["Sec-Fetch-Site::[none]","Sec-Fetch-Mode::[navigate]","Accept-Language::[en-US,en;q=0.9]","Cookie::[Pycharm-42083dc3=45bb8e9b-de30-425f-acf2-5c413b44818c]","Connection::[keep-alive]","Upgrade-Insecure-Requests::[1]","User-Agent::[Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36]","Accept::[text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9]","Sec-Fetch-Dest::[document]","Accept-Encoding::[gzip, deflate, br]"],"method":"GET","uri":"/test/my/uri","payload":"","clientIp":"172.17.0.1","clientPort":"39340","userAgent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36","healthy":true}
```

POST Example - for Linux:
```bash
$ curl -X POST -H "Content-Type: application/json" -d '{ "name": "bubba", "int_value": 987 }' http://127.0.0.1:1234/test/my/uri 

{"statusCode":200,"headers":["User-Agent::[curl/7.55.1]","Accept::[*/*]","Content-Type::[application/json]","Content-Length::[37]"],"method":"POST","uri":"/test/my/uri","payload":"{ \"name\": \"bubba\", \"int_value\": 987 }","clientIp":"172.17.0.1","clientPort":"39346","userAgent":"curl/7.55.1","healthy":true}
```

POST Example - for Windows (cmd.exe):
```powershell
C:\>  curl -X POST -H "Content-Type: application/json" -d "{ """name""": """bubba""", """int_value""": 987 }" http://127.0.0.1:1234/test/my/uri 

{"statusCode":200,"headers":["User-Agent::[curl/7.55.1]","Accept::[*/*]","Content-Type::[application/json]","Content-Length::[37]"],"method":"POST","uri":"/test/my/uri","payload":"{ \"name\": \"bubba\", \"int_value\": 987 }","clientIp":"172.17.0.1","clientPort":"39346","userAgent":"curl/7.55.1","healthy":true}
```
