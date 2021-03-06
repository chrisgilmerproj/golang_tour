/*
Packages
Every Go program is made up of packages.

Programs start running in package main.

This program is using the packages with import paths "fmt" and "math/rand".

By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.

Note: the environment in which these programs are executed is deterministic, so rand.Intn will always return the same number. (To see a different number, seed the number generator; see rand.Seed.)
*/

package main

import (
    "fmt"
    "math/rand"
)

func main() {
    fmt.Println("My favorite number is", rand.Intn(10))
}
