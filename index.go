// Main is a special name declaring an executable rather than a library.
package main

// Import declaration declares library packages referenced in this file.
import (
	"fmt" // A package in the Go standard library.
	"os"  // OS functions like working with the file system
	// "io/ioutil" // Implements some I/O utility functions.
	// "net/http"  // Yes, a web server!
	// "strconv"   // String conversions.
)
func main() {
	var current int;
	var inputBytes,err = os.ReadFile("./index.html");

	if err != nil {
		fmt.Println(err)
		return
	}
	
	for current < len(inputBytes) {
		fmt.Print(string(inputBytes[current]));
		current++
	}
	fmt.Println(properties.padding);
}


type property struct {
	shortName, self, related string
	value []string
}

type CSSProperties struct {
	padding property
}

var padding property = property{
	shortName: "p",
	self: "padding",
	related: "boxModel",
	value: []string{},
};

var properties CSSProperties = CSSProperties{
	padding: padding,
}

func (p *CSSProperties) add(value... string) {
	p.padding.value = append(p.padding.value,value...)
}

func (p CSSProperties) pop() string {
	var removedValue string;
	if len(p.padding.value) > 0 {
		removedValue = p.padding.value[len(p.padding.value)-1]
    p.padding.value = p.padding.value[:len(p.padding.value)-1]
	}
	return removedValue;
}

func give(arg string) string {
	switch arg {
		case "p":
			return "padding"
		case "m":
			return "margin"
		case "w":
			return "width"
		case "h":
			return "height"
		case "bg":
			return "background"
		default:
			return ""
	}
}