package bridge

import (
	"strings"
	"testing"
)

func TestPrintAPI1(t *testing.T) {
	var api1 PrinterImpl1 = PrinterImpl1{}

	var err error = api1.PrintMessage("Hello")
	if err != nil {
		t.Errorf("Error trying to use the API1 implementation: Message: %s\n", err.Error())
	}
}

func TestPrintAPI2(t *testing.T) {
	var api2 PrinterImpl2 = PrinterImpl2{}

	var err error = api2.PrintMessage("Hello")
	if err != nil {
		var expectedErrorMessage string = "You need to pass an io.Writer to PrinterImpl2"
		if !strings.Contains(err.Error(), expectedErrorMessage) {
			t.Errorf("\nError message was not correct.\nActual: %s\nExpected:%s\n", err.Error(), expectedErrorMessage)
		}
	}

	var testWriter TestWriter = TestWriter{}

	api2 = PrinterImpl2{
		Writer: &testWriter,
	}

	var expectedMessage string = "Hello"
	err = api2.PrintMessage(expectedMessage)
	if err != nil {
		t.Errorf("Error trying to use the API2 implementation: %s\n", err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Fatalf("\nAPI2 did not write correctly on the io.Writer. \nActual: %s\nExpected: %s\n", testWriter.Msg, expectedMessage)
	}
}

func TestNormalPrinter_Print(t *testing.T) {
	var expectedMessage string = "Hello io.Writer"

	normal := NormalPrinter{
		Msg:     expectedMessage,
		Printer: &PrinterImpl1{},
	}

	err := normal.Print()
	if err != nil {
		t.Errorf(err.Error())
	}

	testWriter := TestWriter{}
	normal = NormalPrinter{
		Msg: expectedMessage,
		Printer: &PrinterImpl2{
			Writer: &testWriter,
		},
	}

	err = normal.Print()
	if err != nil {
		t.Error(err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Errorf("The expected message on the io.Writer doesn't match actual.\nActual: %s\nExpected: %s\n", testWriter.Msg, expectedMessage)
	}
}

func TestPactPrinter_Print(t *testing.T) {
	var passedMessage string = "Hello io.Writer"
	var expectedMessage string = "Message from Packt: Hello io.Writer"

	packt := PacktPrinter{
		Msg:     passedMessage,
		Printer: &PrinterImpl1{},
	}

	err := packt.Print()
	if err != nil {
		t.Errorf(err.Error())
	}

	testWriter := TestWriter{}
	packt = PacktPrinter{
		Msg: passedMessage,
		Printer: &PrinterImpl2{
			Writer: &testWriter,
		},
	}

	err = packt.Print()
	if err != nil {
		t.Error(err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Errorf("The expected message on the io.Writer doesn't match actual.\nActual: %s\nExpected: %s\n", testWriter.Msg, expectedMessage)
	}
}
