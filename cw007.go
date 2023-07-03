package goschool

import (
	"fmt"
	"strconv"
	"strings"
)

/* func Int32ToIp(n uint32) string {} */
func Codewars007(n uint32) string {
	binary := fmt.Sprintf("%032b", n)
	octets := []string{binary[0:8], binary[8:16], binary[16:24], binary[24:32]}

	var decimalOctets []string
	for _, octet := range octets {
		decimal, _ := strconv.ParseInt(octet, 2, 64)
		decimalOctets = append(decimalOctets, strconv.Itoa(int(decimal)))
	}

	ipAddress := strings.Join(decimalOctets, ".")
	return ipAddress
}

/*
   Take the following IPv4 address: 128.32.10.1
   This address has 4 octets where each octet is a single byte (or 8 bits).
   1st octet 128 has the binary representation: 10000000
   2nd octet 32 has the binary representation: 00100000
   3rd octet 10 has the binary representation: 00001010
   4th octet 1 has the binary representation: 00000001
   So 128.32.10.1 == 10000000.00100000.00001010.00000001
   Because the above IP address has 32 bits, we can represent it as the unsigned 32 bit number: 2149583361

   Complete the function that takes an unsigned 32 bit number and returns a string representation of its IPv4 address.
   2149583361 ==> "128.32.10.1"
   32         ==> "0.0.0.32"
   0          ==> "0.0.0.0"
*/
