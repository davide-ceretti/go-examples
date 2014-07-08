package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    i := 0
    slice := []int{1, 1}
    return func() int {
        k := len(slice)-1
        if i>k {
            slice = append(slice, slice[k]+slice[k-1])
        }
        value := slice[i]
        i += 1
        return value
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
