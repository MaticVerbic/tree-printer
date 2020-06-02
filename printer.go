package printer

import (
	"fmt"
	"io"
	"os"
)

// Node defines behavior of nodes.
type Node interface {
	Children() []Node
	Data() interface{}
}

// Tree defines the behavior of trees.
type Tree interface {
	// Naming chosen not to conflict with package specific required names.
	RootNode() Node
}

// Printer ...
type Printer struct {
	t   Tree
	buf io.Writer
}

// New returns a New printer. If no buffer (buf) is provided,
// os.Stdout is chosen as the buffer.
func New(t Tree, buf io.Writer) *Printer {
	if buf == nil {
		buf = os.Stdout
	}

	return &Printer{
		t:   t,
		buf: buf,
	}
}

// Parse the three into a provided buffer.
func (p *Printer) Parse() {
	p.print(p.t.RootNode(), "", true, 0)
}

// Print the provided tree. If buffer other than os.Stdout
// is provided to printer, then replace buf with os.Stdout
// while printing.
func (p *Printer) Print() {
	if p.buf == os.Stdout {
		p.print(p.t.RootNode(), "", true, 0)
		return
	}

	temp := p.buf
	p.buf = os.Stdout

	p.print(p.t.RootNode(), "", true, 0)

	p.buf = temp
}

// Unexported worker method.
func (p *Printer) print(n Node, prefix string, isLast bool, height int) {
	if height > 0 {
		p.buf.Write([]byte(prefix))
		if isLast {
			p.buf.Write([]byte("└─ "))
			prefix += "   "
		} else {
			p.buf.Write([]byte("├─ "))
			prefix += "|   "
		}
	}
	p.buf.Write([]byte(fmt.Sprintf("%v\n", n.Data())))

	for i, child := range n.Children() {
		p.print(child, prefix, i == len(n.Children())-1, height+1)
	}
}
