# cssrefs

Package cssrefs returns a slice of `Reference{URI, Token string}`s from an `io.Reader`.

### usage

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/weisjohn/cssrefs"
)

func main() {
    resp, _ := http.Get("http://johnweis.com/css/main.css")
    refs := cssrefs.All(resp.Body)

    for _, ref := range refs {
        fmt.Println(ref.Token, ":", ref.URI)
    }
}
```

### output

```golang
$ go run example-cssrefs.go
css : fineprint.css
font : ../fonts/bootstrap/glyphicons-halflings-regular.eot
font : ../fonts/bootstrap/glyphicons-halflings-regular.eot?#iefix
font : ../fonts/bootstrap/glyphicons-halflings-regular.woff
font : ../fonts/bootstrap/glyphicons-halflings-regular.ttf
font : ../fonts/bootstrap/glyphicons-halflings-regular.svg#glyphicons_halflingsregular
img : ../img/light_honeycomb_@2X.png
img : ../img/light_honeycomb.png
```

### struct

`cssrefs` returns a slice of `Reference`s

```
type Reference struct {
    URI, Token string
}
```

### credits

Inspired by https://github.com/weisjohn/htmlrefs