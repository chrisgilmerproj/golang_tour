/*
Exercise: Slices
Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y (to compute the latter function, use math.Pow).

(You need to use a loop to allocate each []uint8 inside the [][]uint8.)

(Use uint8(intValue) to convert between types.)

May need to:

    $ go get code.google.com/p/go-tour/pic
*/

package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {

    pixels := make([][]uint8, dy)

    for y, _ := range pixels {
        pixels[y] = make([]uint8, dx)

        for x, _ := range pixels[y] {
            pixels[y][x] = uint8( (x + y) / 2 )
        }
    }
    return pixels
}

func main() {
    pic.Show(Pic)
}
