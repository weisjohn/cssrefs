package cssrefs

import (
	"strings"
	"testing"
)

func TestAll(t *testing.T) {

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

        body {
            background: url('../img/light_honeycomb_@2X.png');
            background-image: url('../img/light_honeycomb.png');
            background-size: 270px 289px;
            background-repeat: repeat; 
        }
	`)

}
