package adapter

import "fmt"

// LegacyPrinter ...
type LegacyPrinter interface {
	Print(s string) string
}

// MyLegacyPrinter ...
type MyLegacyPrinter struct{}

// Print ...
func (l *MyLegacyPrinter) Print(s string) (newMsg string) {
	newMsg = fmt.Sprintf("Legacy Printer: %s\n", s)
	fmt.Println(newMsg)
	return
}

// ModernPrinter ...
type ModernPrinter interface {
	PrintStored() string
}

// PrinterAdapter ...
type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

// PrintStored ...
func (p *PrinterAdapter) PrintStored() (newMsg string) {
	if p.OldPrinter != nil {
		newMsg = fmt.Sprintf("Adapter: %s", p.Msg)
		newMsg = p.OldPrinter.Print(newMsg)
	} else {
		newMsg = p.Msg
	}
	return
}
