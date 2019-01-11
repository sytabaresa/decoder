package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sin := bufio.NewScanner(os.Stdin)
	serr := bufio.NewWriter(os.Stderr)
	defer serr.Flush()
	sout := bufio.NewWriter(os.Stdout)
	defer sout.Flush()

	sin.Split(ReadToken)
	for sin.Scan() {
		token := sin.Bytes()
		//fmt.Printf("%v \n", token)
		e, err := ParseToken(token)
		if err != nil {
			serr.WriteString(fmt.Sprintf("dato no válido: %s\n", err))
			serr.Flush()
		} else {
			sout.WriteString(fmt.Sprintf("%v \n", e))
			//fmt.Printf("%s \n", e)
			//sout.Write(e)
			sout.Flush()
		}
		//sout.WriteString(fmt.Sprintf("%x\n", token))
		// sout.Write(token)
		// sout.Flush()
	}
}
