# Tree-printer

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

p.Parse()
fmt.Println(buf) 
```
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

```
// Parse the three into a provided buffer.
func (p *Printer) Parse() 

// Print the provided tree. If buffer other than os.Stdout
// is provided to printer, then replace buf with os.Stdout
// while printing.
func (p *Printer) Print() 
```