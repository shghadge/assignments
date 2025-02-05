#### Simple Client Server Application

**Steps:**
If you don't want to download and run Go on your system you can download the precompiled binaries and run the application.

Navigate to the following repository and donwload the particular binary according to your processor architecture:
https://github.com/shghadge/assignments/tree/main/CSC258/assignment1/bin

**OR**

#### Following is a proper way to build and run the Go code

**Step 1:**
Download and set up Go according to the instructions from the following website 
https://go.dev/doc/install

Check whether the installation was successful by typing `go` in your terminal

**Step 2:**
Change directory to *server* and run the following commands to build and run the server

```go
go build
```
```go
./server
```

**Step 3:**
Change directory to *client* and run the following commands to build and run the client
```go
go build
```
```go
./client
```