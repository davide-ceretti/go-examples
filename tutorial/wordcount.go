package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    word_map := make(map[string]int)
    words := strings.Fields(s)
    for _, word := range words {
        if _, exist := word_map[word]; exist {
            word_map[word] += 1
        } else {
            word_map[word] = 1
        }
    }
    return word_map
}

func main() {
    wc.Test(WordCount)
}
