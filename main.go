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
	val string
)

const usage = `'getOpt43' <url>
This program encodes/decodes an ACS URL string to/from HEX.
The format is able to be received by Innbox CPE and applied to their ACS client.
Encode input must begin with http and end with the port number:
\texample http://10.5.100.205:7547 or https://acs.lab.local:10302
Decode input must begin with a HEX identifier 0x01:
\texample 0x0118687474703a2f2f31302e352e3130302e3230353a37353437 
\tor 0x011b68747470733a2f2f6163732e6c61622e6c6f63616c3a3130333032`

func main() {
	flag.Parse()
	if *helpFlag {
		flag.PrintDefaults()
		fmt.Println(usage)
		return
	}
	if flag.NArg() < 1 {
		fmt.Print("Supply ACS URL in format protocol://hostname[ip]:port\nOr encoded value to reverse\n>> ")
		val = readFromStdin()
	} else {
		// drop all trailing args
		val = flag.Args()[0]
	}
	if strings.HasPrefix(val, "http") {
		// encode
		tmp := strings.Split(val, ":")
		if len(tmp) != 3 {
			fmt.Println(usage)
			return
		}
		src := []byte(val)
		enc := hex.EncodeToString(src)
		num := (len(enc)) / 2
		// format:
		// Hex Identifier and Option Code: 0x01
		// Hex value of Data Length without Hex Identifier
		// Hex value of ACS server URL
		out := fmt.Sprintf("0x01%x%s", num, enc)
		fmt.Println(out)
	} else if strings.HasPrefix(val, "0x01") {
		// decode
		// num := url[4:5] // data length in hex
		dec, err := hex.DecodeString(val[6:])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(dec))
		}
	} else {
		fmt.Println(usage)
		return
	}
  
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
