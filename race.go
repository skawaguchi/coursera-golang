/*
In the program below, the two goroutines share `value`. If you run the program 10 times, you will not get the same result. There are multiple reasons:
- The processes execute in different order because the interleavings are non-deterministic, and so the operations are different each time.
- The variable that is shared between the two goroutines is vulnerable to modification in unexpected ways. The functions being applied are not thread-safe.
- The goroutines are non-blocking, so in some cases the main thread will complete and the goroutines will not compmlete. We can see this if the `Total Operations Executed` is not 20 - each operation should have executed twice.

Race conditions occur when the outcome of some code assumes a deterministic outcome but the code is written in such a way that outcomes are actually non-deterministic. Specifically, the code might assume that processes or threads need to execute in a certain order. This means that if the code assumes a specific order of commands and depends on the interleavings in multi-threaded being consistent. If those commands can somehow get out of order, then undesirable behaviours such as bugs can occur.

Race conditions are especially difficult to manage because they are hard to reproduce. So, we should take extra special care to code in such a way that they can’t occur. We call this type of code thread-safe.
*/

package main

import "fmt"

var value int

func main() {
	counter := 0
	value = 1
	for i := 0; i < 10; i++ {
		go add(&value, 2, &counter)
		go multiply(&value, 2, &counter)
	}
	fmt.Println("Final Value:", value)
	fmt.Println("Total Operations Executed:", counter)
}

func multiply(val *int, multiplier int, counter *int) {
	*val *= multiplier
	*counter += 1
}

func add(val *int, addend int, counter *int) {
	*val += addend
	*counter += 1
}
