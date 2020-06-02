package printer

import (
	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"
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
func (p *Printer) Parse() error {
	return p.print(p.t.RootNode(), "", true, 0)
}

// Print the provided tree. If buffer other than os.Stdout
// is provided to printer, then replace buf with os.Stdout
// while printing.
func (p *Printer) Print() error {
	if p.buf == os.Stdout {
		return p.print(p.t.RootNode(), "", true, 0)
	}

	temp := p.buf
	p.buf = os.Stdout

	err := p.print(p.t.RootNode(), "", true, 0)

	p.buf = temp

	return err
}

// Unexported worker method.
func (p *Printer) print(n Node, prefix string, isLast bool, height int) error {
	if height > 0 {
		if _, err := p.buf.Write([]byte(prefix)); err != nil {
			return errors.Wrap(err, "failed to parse prefix")
		}
		if isLast {
			if _, err := p.buf.Write([]byte("└─ ")); err != nil {
				return errors.Wrap(err, "failed to parse last child")
			}
			prefix += "   "
		} else {
			if _, err := p.buf.Write([]byte("├─ ")); err != nil {
				return errors.Wrap(err, "failed to parse child")
			}
			prefix += "|   "
		}
	}
	if _, err := p.buf.Write([]byte(fmt.Sprintf("%v\n", n.Data()))); err != nil {
		return errors.Wrap(err, "failed to parse data")
	}

	for i, child := range n.Children() {
		if err := p.print(child, prefix, i == len(n.Children())-1, height+1); err != nil {
			return errors.Wrap(err, "failed to parse child")
		}

	}

	return nil
}
