
// go parallel :: sample
// r.20241021.2358
// (c) 2024 unix-world.org

package main

import (
	"fmt"
	"time"

	parallel "github.com/unix-world/go-parallel"
)


func main() {

	input := []any{1, 2, "a", "b", nil, []byte{}, true}

	parallel.ForEach(input, func(elem any) {
		fmt.Printf("Processing %v\n", elem)
	})

	fmt.Println("=============================")
	output := parallel.Map(input, func(elem any) any {
		fmt.Printf("Processing %v\n", elem)
		return elem
	})
	fmt.Printf("The final result is %v\n", output)

	fmt.Println("=============================")
	maxConcurrency := 2
	parallel.ForEachLimit(input, maxConcurrency, func(elem any) {
		executionTime := time.Now().UTC().Format("15:04:05.999")
		fmt.Printf("%s - Processing %v\n", executionTime, elem)
		time.Sleep(1 * time.Second)
	})

	fmt.Println("=============================")
	output2 := parallel.MapLimit(input, maxConcurrency, func(elem any) any {
		executionTime := time.Now().UTC().Format("15:04:05.999")
		fmt.Printf("%s - Processing %v\n", executionTime, elem)
		time.Sleep(1 * time.Second)
		return elem
	})
	fmt.Printf("The final result is %v\n", output2)

}

// #end
