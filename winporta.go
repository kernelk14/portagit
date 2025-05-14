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
    path, err1 := os.Stat(path_name)
    if os.IsExist(err1) != true {
        fmt.Println(*path)
        fmt.Println("You don't have a Git folder!")
        fmt.Printf("Please create this folder: %s, then run this program again.", path_name)
    } else {
        fmt.Println(*path)
        fmt.Printf("Found folder: %s", path_name)
    }

}
