# go-process-killer
a program that hunts for all go processes and kills them.

# How To Run
- Clone repository
- Build `go run main.go` and run `./main`
- Or you can just run it directly using `go run main.go`

# Note
This program will actually terminate all processes the kernel creates for golang binaries. so things like your go programming language server and what not will stop functioning.
So unless you want to code without intelisense, then by all means, be my guest. 

# Proof
So to validate my program, I created a tiny go program that prints the value returned from a ticker every second. <br/> 
Once it starts running, I run the process killer and you notice that it shuts down with a code of 9. <br/>
See image below.  <img width="1440" alt="Screen Shot 2021-03-15 at 10 47 21 PM" src="https://user-images.githubusercontent.com/24861123/111231311-f5ec3a80-85e0-11eb-83de-1e523694731f.png">

# Notes
- Ideally, you should always use SIGTERM, instead of SIGKILL, as the former allows the process to cleanup.
- You want to make sure you're not killing yourself in the process (no pun intended)
- I took the code for checking if a process was created by a go binary from here -> https://github.com/google/gops/blob/55729d43c2835298a7e05e576a77da2989ff5b7b/goprocess/gp.go#L113
