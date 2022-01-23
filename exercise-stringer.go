package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

/* NOTE: Stringers
One of the most ubiquitous interfaces is Stringer defined by the fmt package.

type Stringer interface {
    String() string
}
A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.
*/

func (ip IPAddr) String() string {
	var strArray [len(ip)]string
	for i, v := range ip {
		strArray[i] = fmt.Sprintf("%v", v)
	}
	return strings.Join(strArray[:], ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v %T\n", name, ip, ip)
	}
}
