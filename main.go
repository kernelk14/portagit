package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"net"
	"strings"
)

func help() {
	fmt.Println("Help:")
	fmt.Println("    init    Creates a new repo [arg]")
	fmt.Println("    help    Prints this help")
	fmt.Println("    rem     Deletes the repo [arg]")
}

func main() {
	ip_addr := ""
	uname := ""

	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err.Error())
	}

	uname = currentUser.Username

	addrs, err := net.InterfaceAddrs()
    if err != nil {
        panic(err)
    }
    for _, addr := range addrs {
        ipNet, ok := addr.(*net.IPNet)
        if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
            if !strings.HasPrefix(ipNet.IP.String(), "172") {
				ip_addr = ipNet.IP.String()
			}

        }
    }

	args := os.Args

	if len(args) < 2 {
		help()
	} else {
		if args[1] == "init" {
			if len(args) > 2 {
				PG_WORKDIR := fmt.Sprintf("/home/%s/Git/", uname)
				for _, a := range args[2:] {
					bare_dir := fmt.Sprintf("%s%s", PG_WORKDIR, a)
					bare_repo := fmt.Sprintf("%s/%s.git", bare_dir, a)
					err := os.MkdirAll(bare_dir, os.ModePerm)
					if err != nil {
						log.Fatal(err)
					}

					cmd := exec.Command("git", "init", "--bare", bare_repo)
					err1 := cmd.Run()
					if err1 != nil {
						log.Fatal(err)
					} else {
						fmt.Printf("Created new repository on %s\n", bare_repo)
						fmt.Printf("GIT SSH MIRROR: %s@%s:%s\n", uname, ip_addr, bare_repo)
					}
				}
			} else {
				// TODO:  <10-05-25, kernelk14> //
				fmt.Println("Creating repo in the same directory is not yet done.")
			}
		}
	}

}
