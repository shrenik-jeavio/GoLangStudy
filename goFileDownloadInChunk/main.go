//Download file in chunks from server.

package main

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"fmt"
	"os"
	"log"
	"bufio"
)

var wg sync.WaitGroup

var fileName string = "sample_word.txt"
var server_file_path = "http://localhost/" + fileName
var newFileAt = "./created"
var limit = 5


func printError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

// merge chunk files into one downloaded.txt
func mergeChunkDataFiles() {

	fileHandle, err := os.Create(newFileAt + "/" + "downloaded.txt")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < limit; i++ {
		// Open file for reading
		file, err := os.Open(newFileAt + "/" + strconv.Itoa(i) + ".txt")
		if err != nil {
			log.Fatal(err)
		}

		//Read data
		data, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}

		//write data
		writer := bufio.NewWriter(fileHandle)
		defer fileHandle.Close()

		writer.Write(data)
		writer.Flush()
	}

}

func main(){

	res, _ := http.Head(server_file_path); // 1 MB file of random numbers per line
	maps := res.Header

	length, _ := strconv.Atoi(maps["Content-Length"][0]) // Get the content length from the header request
	fmt.Println("length size : ", length)
	len_sub := length / limit // Bytes for each Go-routine
	diff := length % limit // Get the remaining for the last request

	body := make([]string, 11) // Make up a temporary array to hold the data to be written to the file

	for i := 0; i < limit; i++ {
		wg.Add(1)

		min := len_sub * i
		max := len_sub * ( i + 1 )

		if(i == limit - 1){
			max += diff // Add the remaining bytes in the last request
		}

		go func(min int, max int, i int) {

			client := &http.Client {}
			req, _ := http.NewRequest("GET", server_file_path, nil)

			range_header := "bytes=" + strconv.Itoa(min) + "-" + strconv.Itoa(max-1) // Add the data for the Range header of the form "bytes=0-100"
			fmt.Println("range_header : ", range_header)
			req.Header.Add("Range", range_header)

			resp, _ := client.Do(req)
			defer resp.Body.Close()

			reader, _ := ioutil.ReadAll(resp.Body)
			body[i] = string(reader)

			var path = newFileAt + "/" + strconv.Itoa(i) + ".txt"
			if _, err := os.Stat(path); os.IsNotExist(err) {
				os.Mkdir(newFileAt, os.FileMode(0777))
			}

			ioutil.WriteFile(path, []byte(string(body[i])), 0x777) // Write to the file i as a byte array
			wg.Done()
		}(min, max, i)
	}
	wg.Wait()

	mergeChunkDataFiles()
}