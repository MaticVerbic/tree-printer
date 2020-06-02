# Tree-printer
![test](https://github.com/MaticVerbic/tree-printer/workflows/test/badge.svg?branch=master)
Simple functionality for printing trees. 

## Example: 
```go
tr := tree.New("root", "root")
tr.Insert("two", "two", "root")
tr.Insert("three", "three", "root")
tr.Insert("four", "four", "two")
tr.Insert("five", "five", "two")
tr.Insert("six", "six", "four")
tr.Insert("seven", "seven", "three")

buf := bytes.NewBuffer([]byte{})

p := printer.New(tr, buf)

err := p.Parse()
if err != nil {
    panic(err)
}
fmt.Println(buf) 
```
output: 
```  
root
├─ two
|   ├─ four
|   |   └─ six
|   └─ five
└─ three
    └─ seven
```  

### Documentation

```go
// Parse the three into a provided buffer.
func (p *Printer) Parse() error

// Print the provided tree. If buffer other than os.Stdout
// is provided to printer, then replace buf with os.Stdout
// while printing.
func (p *Printer) Print() error
```