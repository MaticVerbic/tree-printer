package printer_test

import (
	"bytes"
	"fmt"
	"printer"
	"printer/tree"
	"testing"

	"github.com/kylelemons/godebug/diff"
)

func TestParse(t *testing.T) {
	tr := tree.New("root", "root")
	tr.Insert("two", "two", "root")
	tr.Insert("three", "three", "root")
	tr.Insert("four", "four", "two")
	tr.Insert("five", "five", "two")
	tr.Insert("six", "six", "four")
	tr.Insert("seven", "seven", "three")

	s := `root
├─ two
|   ├─ four
|   |   └─ six
|   └─ five
└─ three
   └─ seven
`

	buf := bytes.NewBuffer([]byte{})

	p := printer.New(tr, buf)

	p.Parse()

	if diff := diff.Diff(s, buf.String()); diff != "" {
		fmt.Println(diff)
		t.Fail()
	}

}
