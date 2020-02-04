package bridge

import (
	"errors"
	"fmt"
	"io"
)

// PrinterAPI ...
type PrinterAPI interface {
	PrintMessage(string) error
}

// PrinterImpl1 ...
type PrinterImpl1 struct{}

// PrintMessage ...
func (p *PrinterImpl1) PrintMessage(msg string) error {
	fmt.Printf("%s\n", msg)
	return nil
}

// PrinterImpl2 ...
type PrinterImpl2 struct {
	Writer io.Writer
}

// PrintMessage ...
func (d *PrinterImpl2) PrintMessage(msg string) error {
	if d.Writer == nil {
		return errors.New("You need to pass an io.Writer to PrinterImpl2")
	}
	fmt.Fprintf(d.Writer, "%s", msg)
	return nil
}

// TestWriter ...
type TestWriter struct {
	Msg string
}

// Write ...
func (t *TestWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 {
		t.Msg = string(p)
		return n, nil
	}
	err = errors.New("Content received on Writer was empty")
	return
}

// PrinterAbstraction ...
type PrinterAbstraction interface {
	Print() error
}

// NormalPrinter ...
type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

// Print ...
func (c *NormalPrinter) Print() error {
	c.Printer.PrintMessage(c.Msg)
	return nil
}

// PacktPrinter ...
type PacktPrinter struct {
	Msg     string
	Printer PrinterAPI
}

// Print ...
func (c *PacktPrinter) Print() error {
	c.Printer.PrintMessage(fmt.Sprintf("Message from Packt: %s", c.Msg))
	return nil
}
