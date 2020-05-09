package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/arulthayalan/link"
)

var exampleHtml = `
<html>
<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
</body>
</html>
`

func main() {

	r := strings.NewReader(exampleHtml)

	links, err := link.Parse(r)

	if err != nil {
		log.Fatal(fmt.Errorf("Unable to parse html document %v", err))
	}
	for _, link := range links {
		fmt.Printf("%+v\n", link)
	}
}
