package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"syscall"

	"github.com/vishvananda/netns"
)

func main() {
	// Lock the OS Thread so we don't accidentally switch namespaces
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// Save the current network namespace
	origns, _ := netns.Get()
	defer origns.Close()

	// Do something with the network namespace
	ifs, _ := net.Interfaces()
	fmt.Printf("Interfaces: %v\n", ifs)

	// Create a new network namespace
	newns, _ := netns.New()
	defer newns.Close()

	nsPath := filepath.Join("/var/run/netns", "ns8")
	f, err := os.Create(nsPath)
	if err != nil {
		log.Println("create dir fail", err)
	}

	if err = f.Close(); err != nil {
		log.Fatal("close file fail")
	}
	err = syscall.Mount("/proc/self/ns/net", nsPath, "", syscall.MS_BIND, "")
	if err != nil {
		log.Println("mount faild")
	}
	// Do something with the network namespace
	fmt.Println("ip interface inside namespace")
	ifaces, _ := net.Interfaces()
	fmt.Printf("Interfaces: %v\n", ifaces)

	// Switch back to the original namespace
	netns.Set(origns)
}
