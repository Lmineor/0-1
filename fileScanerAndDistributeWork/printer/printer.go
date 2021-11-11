package printer

import "fmt"

// type Print interface {
// 	PrintInfo(string)
// 	PrintWarning(string)
// 	PrintAdd(string)
// 	PrintDelete(string)
// }

type Printer struct{}

func NewPrinter() *Printer {
	return &Printer{}
}

func (p *Printer) PrintInfo(ln string) {
	str := fmt.Sprintf("[info]-----:%s", ln)
	fmt.Println(str)
}

func (p *Printer) PrintWarning(ln string) {
	str := fmt.Sprintf("[warnning]-----:%s", ln)
	fmt.Println(str)
}

func (p *Printer) PrintAdd(ln string) {
	str := fmt.Sprintf("[add]-----:%s", ln)
	fmt.Println(str)
}

func (p *Printer) PrintDelete(ln string) {
	str := fmt.Sprintf("[delete]-----:%s", ln)
	fmt.Println(str)
}
