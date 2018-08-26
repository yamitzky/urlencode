package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	fmt.Printf("%s", url.PathEscape(os.Args[1]))
}
