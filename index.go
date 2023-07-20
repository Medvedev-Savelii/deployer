package main

import (
    "fmt"
    "net/http"
    "os"
    "os/exec"
)

func main() {

    portStr := os.Getenv("PORT")
    http.Handle("/deploy", hwHandler{})
    http.ListenAndServe(":"+portStr, nil)

}

type hwHandler struct{}

func (hwHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
    // 1 Get User and Port From enviroments varss
    sshUser := os.Getenv("SSH_USER")
    sshHost := os.Getenv("SSH_HOST")
    // 2 Run as ssh using user and host

    cmd := exec.Command("./deploy", sshUser, sshHost)
    
    // 3 Process Error
    
    _, err := cmd.StdoutPipe()

    err = cmd.Start()

    if err != nil {
        fmt.Fprintln(os.Stderr, "Error starting cmd ne kurwa", err)
        return
    }

    err = cmd.Wait()

    if err != nil {
        fmt.Fprintln(os.Stderr, "Error waiting for cmd kurwa blat", err)
        return
    }
    // 4 Return 200 is success best


    writer.WriteHeader(200)
}
