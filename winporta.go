package main

import (
    "fmt"
    "log"
    "os"
	"os/user"
    "strings"
)


func main() {
    uname := ""
    currentUser, err := user.Current()
    if err != nil {
        log.Fatal(err.Error())
    } else {
        uname = currentUser.Username
    }

    name := strings.Split(uname, "\\")
    
    desktop_name := name[1] // The name of the user
    
    path_name := fmt.Sprintf("C:\\Users\\%s\\Documents\\Git\\", desktop_name)
    if _, err := os.Stat(path_name); os.IsNotExist(err) {
        fmt.Println("Git directory not found!")
        fmt.Printf("Create this directory first: '%s', then run this program again.", path_name)
    } else {
        fmt.Printf("Found folder: %s", path_name)
    }

}
