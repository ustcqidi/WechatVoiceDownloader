package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func download(url string, file_idx int) {
	fmt.Print("\nurl is ", url)
	filename := strconv.Itoa(file_idx) + ".mp3"
	fmt.Print("\nfilename is ", filename)

	// Parse
	response,_:=http.Get(url)
	defer response.Body.Close()
	body,_:=ioutil.ReadAll(response.Body)
	s := string(body)
	index := strings.Index(s, "voice_encode_fileid=")
	encode_fileid := s[index+21:index+21+28]

	download_link := "https://res.wx.qq.com/voice/getvoice?mediaid=" + encode_fileid

	fmt.Print("download_link is ", download_link)

	// Download
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

func main() {

   filepath := "download.txt"
   file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
   if err != nil {
      fmt.Println("Open file error!", err)
      return
   }
   defer file.Close()

   fileIdx := 1

   buf := bufio.NewReader(file)
   for {
      line, err := buf.ReadString('\n')
	  line = strings.TrimSpace(line)
	  download(line, fileIdx)
	  fileIdx += 1
    //   fmt.Println(line)
      if err != nil {
         if err == io.EOF {
            fmt.Println("\nFinish!")
            break
         } else {
            fmt.Println("Read file error!", err)
            return
         }
      }
   }

	// args := os.Args

	// url := args[1]
	// filename := args[2]

	// fmt.Print("\nurl ", url)
	// fmt.Print("\nname ", filename)

}
