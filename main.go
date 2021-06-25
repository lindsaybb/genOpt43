package main

import (
	"fmt"
	"flag"
	"io"
	"bufio"
	"os"
	"strings"
	"encoding/hex"
)

var (
	helpFlag = flag.Bool("h", false, "Show this help")
	url string
)

const usage = `'getOpt43' <url>
This program encodes an ACS URL string to a HEX format able to be received by Innbox CPE and applied to their ACS client.
Input format must begin with http and end with the port number, for example http://10.5.5.32:7547 or https://172.17.1.7:10302`

func main() {
	flag.Parse()
	if *helpFlag {
		flag.PrintDefaults()
		fmt.Println(usage)
		return
	}
	if flag.NArg() < 1 {
		fmt.Print("Supply ACS URL in format protocol://hostname[ip]:port\n>> ")
		url = readFromStdin()
	} else {
		// drop all trailing args
		url = flag.Args()[0]
	}
	if !strings.HasPrefix(url, "http") {
		fmt.Println(usage)
		return
	}
	tmp := strings.Split(url, ":")
	if len(tmp) != 3 {
		fmt.Println(usage)
		return
	}
	src := []byte(url)
	enc := hex.EncodeToString(src)
	num := (len(enc)) / 2
  // format:
  //  Hex Identifier and Option Code: 0x01
  //  Hex value of Data Length without Hex Identifier
  //  Hex value of ACS server URL
	out := fmt.Sprintf("0x01%x%s", num, enc)
	fmt.Println(out)
}

func readFromStdin() string {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	a, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	} else if err != nil {
		panic(err)
	}

	return strings.TrimRight(string(a), "\r\n")
}
