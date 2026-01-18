package maps

import "fmt"

type dnsMap map[string]string

func Driver() {
	//using type alias
	dns := make(dnsMap, 5)

	dns = map[string]string{
		"google":   "https://www.google.com",
		"facebook": "https://www.facebook.com",
	}

	dns["youtube"] = "https://www.youtube.com"

	for i, j := range dns {
		fmt.Printf("%s: %s\n", i, j)
	}
}
