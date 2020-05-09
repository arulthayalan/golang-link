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
  <h1>Hello!</h1>
  <a href="/page-one">A link to 
  to  first page
  </a>
  <a href="/page-two">A link to second page</a>
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
