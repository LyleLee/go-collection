package main

import "fmt"
import "os/exec"

func main() {

	cmd := exec.Command("iptables", "--version")

	out, err := cmd.Output()

	fmt.Println(err)
	fmt.Println(string(out))

}
