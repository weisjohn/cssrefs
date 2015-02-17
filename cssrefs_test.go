package cssrefs

import (
	"strings"
	"testing"
)

func TestAll(t *testing.T) {

	// the required resources to be found in the example
	reqs := [...]Reference{
		{URI: "fineprint.css", Token: "css"},
		{URI: "../fonts/bootstrap/glyphicons-halflings-regular.eot", Token: "font"},
		{URI: "../fonts/bootstrap/glyphicons-halflings-regular.eot?#iefix", Token: "font"},
		{URI: "../fonts/bootstrap/glyphicons-halflings-regular.woff", Token: "font"},
		{URI: "../fonts/bootstrap/glyphicons-halflings-regular.ttf", Token: "font"},
		{URI: "../fonts/bootstrap/glyphicons-halflings-regular.svg#glyphicons_halflingsregular", Token: "font"},
		{URI: "../img/light_honeycomb_@2X.png", Token: "img"},
		{URI: "../img/light_honeycomb.png", Token: "img"},
	}

	// example HTML reader
	reader := strings.NewReader(`

        @import url("fineprint.css") print;

        @font-face {
          font-family: 'Glyphicons Halflings';
          src: url('../fonts/bootstrap/glyphicons-halflings-regular.eot');
          src: url('../fonts/bootstrap/glyphicons-halflings-regular.eot?#iefix') format('embedded-opentype'), 
            url('../fonts/bootstrap/glyphicons-halflings-regular.woff') format('woff'), 
            url('../fonts/bootstrap/glyphicons-halflings-regular.ttf') format('truetype'), 
            url('../fonts/bootstrap/glyphicons-halflings-regular.svg#glyphicons_halflingsregular') format('svg'); 
        }

        .foo {
            background: red;
        }

        body {
            background: url('../img/light_honeycomb_@2X.png');
            background-image: url('../img/light_honeycomb.png');
            background-size: 270px 289px;
            background-repeat: repeat; 
        }
	`)

	// get the refs from the implementation
	refs := All(reader)

	need, have := len(reqs), len(refs)
	if need != have {
		t.Errorf("Wrong number of refs returned. need: %d , have: %d", need, have)
	}

	// loop through and verify URI and Token names
	for i, req := range reqs {
		ref := refs[i]

		if req.URI != ref.URI {
			t.Errorf("Mismatch URI detected. need: %s , have: %s", req.URI, ref.URI)
		}

		if req.Token != ref.Token {
			t.Errorf("Mismatch Token detected. need: %s , have: %s", req.Token, ref.Token)
		}
	}
}
