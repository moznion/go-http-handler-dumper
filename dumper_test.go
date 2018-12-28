package hhdumper

import (
	"net/http"
	"strings"
	"testing"
)

func foo(w http.ResponseWriter, r *http.Request) {
}

func TestDumpBy(t *testing.T) {
	var mux http.ServeMux
	mux.Handle("/foo", http.HandlerFunc(foo))
	mux.HandleFunc("/bar/buz", func(w http.ResponseWriter, r *http.Request) {
	})

	routes := DumpBy(&mux)
	if len(routes) != 2 {
		t.Errorf("invalid length: expected=2, got=%d", len(routes))
	}

	fooHandler := routes["/foo"]
	if fooHandler == nil {
		t.Errorf("handler is unexpected nil: route=/foo")
	}

	expectedFuncName := "foo"
	if fooHandler.FuncName != expectedFuncName {
		t.Errorf("invalid func name: expected=%s, got=%s", expectedFuncName, fooHandler.FuncName)
	}

	fileNameLeaves := strings.Split(fooHandler.FileName, "/")
	fileName := fileNameLeaves[len(fileNameLeaves)-1]
	expectedFileName := "dumper_test.go"
	if fileName != expectedFileName {
		t.Errorf("invalid file name: expected=%s, got=%s", expectedFileName, fileName)
	}

	// XXX fragile!
	expectedLineNum := 10
	if fooHandler.LineNumber != expectedLineNum {
		t.Errorf("invalid line number: expected=%d, got=%d", expectedLineNum, fooHandler.LineNumber)
	}

	barBuzHandler := routes["/bar/buz"]
	if barBuzHandler == nil {
		t.Errorf("handler is unexpected nil: route=/bar/buz")
	}

	expectedFuncName = "func1"
	if barBuzHandler.FuncName != expectedFuncName {
		t.Errorf("invalid func name: expected=%s, got=%s", expectedFuncName, barBuzHandler.FuncName)
	}

	if fileName != expectedFileName {
		t.Errorf("invalid file name: expected=%s, got=%s", expectedFileName, fileName)
	}

	// XXX fragile!
	expectedLineNum = 16
	if barBuzHandler.LineNumber != expectedLineNum {
		t.Errorf("invalid line number: expected=%d, got=%d", expectedLineNum, barBuzHandler.LineNumber)
	}
}

func TestDump(t *testing.T) {
	http.Handle("/foo", http.HandlerFunc(foo))
	http.HandleFunc("/bar/buz", func(w http.ResponseWriter, r *http.Request) {
	})

	routes := Dump()
	if len(routes) != 2 {
		t.Errorf("invalid length: expected=2, got=%d", len(routes))
	}

	fooHandler := routes["/foo"]
	if fooHandler == nil {
		t.Errorf("handler is unexpected nil: route=/foo")
	}

	expectedFuncName := "foo"
	if fooHandler.FuncName != expectedFuncName {
		t.Errorf("invalid func name: expected=%s, got=%s", expectedFuncName, fooHandler.FuncName)
	}

	fileNameLeaves := strings.Split(fooHandler.FileName, "/")
	fileName := fileNameLeaves[len(fileNameLeaves)-1]
	expectedFileName := "dumper_test.go"
	if fileName != expectedFileName {
		t.Errorf("invalid file name: expected=%s, got=%s", expectedFileName, fileName)
	}

	// XXX fragile!
	expectedLineNum := 10
	if fooHandler.LineNumber != expectedLineNum {
		t.Errorf("invalid line number: expected=%d, got=%d", expectedLineNum, fooHandler.LineNumber)
	}

	barBuzHandler := routes["/bar/buz"]
	if barBuzHandler == nil {
		t.Errorf("handler is unexpected nil: route=/bar/buz")
	}

	expectedFuncName = "func1"
	if barBuzHandler.FuncName != expectedFuncName {
		t.Errorf("invalid func name: expected=%s, got=%s", expectedFuncName, barBuzHandler.FuncName)
	}

	if fileName != expectedFileName {
		t.Errorf("invalid file name: expected=%s, got=%s", expectedFileName, fileName)
	}

	// XXX fragile!
	expectedLineNum = 70
	if barBuzHandler.LineNumber != expectedLineNum {
		t.Errorf("invalid line number: expected=%d, got=%d", expectedLineNum, barBuzHandler.LineNumber)
	}
}
