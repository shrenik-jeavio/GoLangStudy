package main

import (
	"fmt"
	//"os"
	"io/ioutil"
)

func main(){
	bs, err := ioutil.ReadFile("test_file.txt")
	if err != nil{ return }
	str := string(bs)
	fmt.Println(str)
	/*
	file, err := os.Open("test_file.txt")
	if err != nil { return }
	defer file.Close()

	// get file size
	stat, err := file.Stat()
	if err != nil{ return }
	//fmt.Println(stat)

	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil { return }

	str := string(bs)
	fmt.Println(str)
	*/

}