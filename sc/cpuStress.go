/*
*
* This project is destinated to stress CPU cores maintining them 100% busy
* Its very helpful if you want to test your current cooling system
* Cross Platform and Working for any Intel/AMD proccesor available in the market 
*
* Contact : Rodrigo Higareda <rhs_21@hotmail.com>
*/


package main

import (
	"fmt"
	"runtime"
	"math/rand"
	"sort"
)

// Workload to be ran on each CPU core
func Sort(exit chan bool) {
	// create a random list to be ordered
	var list []int = rand.Perm(300000)
	for  {
		 sort.Ints(list)
	}
	// notify when this thread ends
	exit <- true
}

// Prints CPU cores available to be tested
func proc_info(cores *uint8) {
	fmt.Println("**********\n")
	fmt.Println("CPU: ", *cores)
	fmt.Println("by Rodrigo Higareda <rhs_21@hotmail.com>\n")
	fmt.Println("**********")
}

// Main Function to be run
func main() {
	var	cores uint8 = uint8(runtime.NumCPU())
	var count uint8
	var start string
	exit := make(chan bool)
	runtime.GOMAXPROCS(int(cores))

	proc_info(&cores)
    fmt.Println("Press Enter to begin the test and CTRL + C to stop it....")
    fmt.Scanln(&start)
	fmt.Println("Stressing CPU....")
	// Start a Thread for each core found
	for count = 0; count < cores; count++ {
		go Sort(exit)
	}
	<- exit
}