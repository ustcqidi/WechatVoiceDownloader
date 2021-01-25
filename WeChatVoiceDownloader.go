package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	args := os.Args

	url := args[1]
	filename := args[2]

	fmt.Print("\nurl ", url)
	fmt.Print("\nname ", filename)

	response,_:=http.Get(url)
	defer response.Body.Close()
	body,_:=ioutil.ReadAll(response.Body)
	s := string(body)
	index := strings.Index(s, "voice_encode_fileid=")
	encode_fileid := s[index+21:index+21+28]

	download_link := "https://res.wx.qq.com/voice/getvoice?mediaid=" + encode_fileid

	fmt.Print("download_link is ", download_link)

	// Get the data
	fmt.Print("\ndownload file...")
	resp, err := http.Get(download_link)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Print("\nread stream...")

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Print("\nwrite file...")

	ioutil.WriteFile(filename, data, 0644)
}
