/*
Exercise: Maps
Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.

May need to:

    $ go get code.google.com/p/go-tour/wc
*/

package main

import (
    "strings"
    "code.google.com/p/go-tour/wc"
)

var wmap map[string]int

func WordCount(s string) map[string]int {
    wmap = make(map[string]int)
    for _, word := range strings.Fields(s) {
      wmap[word] += 1
    }
    return wmap
}

func main() {
    wc.Test(WordCount)
}

