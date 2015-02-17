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

// TODO:

### struct

`cssrefs` returns a slice of `Reference`s

```
type Reference struct {
    URI, Token string
}
```

### credits

Inspired by https://github.com/weisjohn/htmlrefs