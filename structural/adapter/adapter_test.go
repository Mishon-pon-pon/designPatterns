package adapter

import "testing"

func TestAdapter(t *testing.T) {
	var msg string = "Hello World!"
	var adapter PrinterAdapter = PrinterAdapter{OldPrinter: &MyLegacyPrinter{}, Msg: msg}
	var returnedMsg string = adapter.PrintStored()
	if returnedMsg != "Legacy Printer: Adapter: Hello World!\n" {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}

	adapter = PrinterAdapter{OldPrinter: nil, Msg: msg}
	returnedMsg = adapter.PrintStored()
	if returnedMsg != "Hello World!" {
		t.Errorf("Message didn't match: %s\n", returnedMsg)
	}
}
