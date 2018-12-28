package hhdumper

import (
	"net/http"
	"reflect"
	"runtime"
	"strings"
)

// HandlerInfo represents handler detail information.
type HandlerInfo struct {
	// FuncName is a function name of a handler.
	FuncName string
	// FileName indicates the name of the file that contains the handler function.
	FileName string
	// LineNumber indicates the line number of the handler function.
	LineNumber int
}

// Dump dumps handler information with routes of `http.DefaultServeMux`.
// It returns a map that points `path-info` to `handler-information`.
func Dump() map[string]*HandlerInfo {
	return DumpBy(http.DefaultServeMux)
}

// DumpBy dumps handler information with routes of given mux.
// It returns a map that points `path-info` to `handler-information`.
func DumpBy(mux *http.ServeMux) map[string]*HandlerInfo {
	routeMap := reflect.ValueOf(mux).Elem().FieldByName("m")
	keys := routeMap.MapKeys()

	result := make(map[string]*HandlerInfo, len(keys))
	for _, key := range keys {
		handler := routeMap.MapIndex(key).FieldByName("h")
		funk := runtime.FuncForPC(handler.Elem().Pointer())
		funcNameLeaves := strings.Split(funk.Name(), ".")
		funcName := funcNameLeaves[len(funcNameLeaves)-1]
		fileName, lineNumber := funk.FileLine(handler.Elem().Pointer())

		result[key.String()] = &HandlerInfo{
			FuncName:   funcName,
			FileName:   fileName,
			LineNumber: lineNumber,
		}
	}

	return result
}
