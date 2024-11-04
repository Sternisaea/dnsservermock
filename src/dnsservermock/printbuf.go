package dnsservermock

import "fmt"

func printbuffer(buf []byte) {
	fmt.Println("Buffer:")
	chrs := ""
	for i, b := range buf {
		if i != 0 && i%16 == 0 {
			if len(chrs) != 0 {
				fmt.Printf("%s", chrs)
			}
			fmt.Println()
			chrs = ""
		}

		// if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9') || b == '-' {
		if b >= ' ' && b <= '~' {
			chrs += string(b)
		} else {
			chrs += "."
		}

		fmt.Printf("%02x ", b)
	}
	fmt.Println()
}
