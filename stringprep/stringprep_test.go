package stringprep

import (
	"github.com/domainr/mustang/test"

	"fmt"
	"testing"
)

var nameprepStrings = map[string]string{
	"example.com": "example.com",
	"example.org": "example.org",
	"example.net": "example.net",
	"examplé.com": "examplé.com",
	"xn--ninja.com": "xn--ninja.com",
	"中文网":         "中文网",
}

func TestNameprepNoUnassigned(t *testing.T) {
	for input, expectedOutput := range nameprepStrings {
		output := NameprepNoUnassigned(input)
		fmt.Printf("NameprepNoUnassigned %s -> %s (%s)\n", input, output, expectedOutput)
		// test.Expect(t, err == nil, true)
		test.Expect(t, output, expectedOutput)
	}
}
