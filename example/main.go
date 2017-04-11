package main

import (
	"huffman"
	"os"
	"log"
	"fmt"
)

func main() {
	fp, err := os.Open("example/data.txt")
	if err != nil {
		log.Println(err)
		return
	}
	//defer fp.Close()

	ht := huffman.New(fp)
	//fmt.Println("Traverse: ", ht.Traverse())
	if err := ht.Encode(); err != nil {
		log.Println(err)
	}

	de := ht.Decode("0101010101010101010101010101010101010101010101010101")
	fmt.Println(string(de))
	fmt.Println(string(ht.Decode("0000101100000110011100010101101101001111101011111100011001111110100100101")))
	ht.Traverse()
}
