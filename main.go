package main

import (
    "fmt"
    "runtime/pprof"
	"os"
	"time"
)

func myFunction(id int) {
    // Get the goroutine profile
    profile := pprof.Lookup("goroutine")

    // Print the information of each goroutine
    fmt.Printf("Goroutine profile for myFunction %d:\n", id)
    profile.WriteTo(os.Stdout, 1)

    // ... Your function code here ...
}

func main() {
    // Start multiple goroutines that call myFunction
    for i := 1; i <= 3; i++ {
        go myFunction(i)
    }

    // Wait for a few seconds to allow the goroutines to run
    time.Sleep(5 * time.Second)

    // Get the goroutine profile again
    profile := pprof.Lookup("goroutine")

    // Print the information of each goroutine
    fmt.Println("Goroutine profile for all running goroutines:")
    profile.WriteTo(os.Stdout, 1)
}
