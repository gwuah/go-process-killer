# go-process-killer
a program that hunts for all go processes and kills them.

# How To Run
- Clone repository
- Build `go run main.go` and run `./main`
- Or you can just run it directly using `go run main.go`

# Note
This program will actually terminate all processes the kernel creates for golang binaries. so things like your go programming language server and what not will stop functioning.
So unless you want to code without intelisense, then by all means, be my guest. 
