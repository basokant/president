package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

const ContentType = "text/event-stream;charset=utf-8"
const NoCache = "no-cache"
const KeepAlive = "keep-alive"

var fieldReplacer = strings.NewReplacer(
	"\n", "\\n",
	"\r", "\\r")

var dataReplacer = strings.NewReplacer(
	"\n", "\ndata:",
	"\r", "\\r")

type Event[T any] struct {
	Data  T
	Event string
	Id    string
	Retry uint
}

func Encode[T any](w io.Writer, event Event[T]) error {
	writeId(w, event.Id)
	writeEvent(w, event.Event)
	writeRetry(w, event.Retry)
	return writeData(w, event.Data)
}

func WriteHeaders(w http.ResponseWriter) {
	header := w.Header()
	header.Set("Content-Type", ContentType)

	if _, exist := header["Cache-Control"]; !exist {
		header.Set("Cache-Control", NoCache)
	}

	if _, exist := header["Connection"]; !exist {
		header.Set("Connection", KeepAlive)
	}
}

func writeId(w io.Writer, id string) {
	if len(id) > 0 {
		w.Write([]byte("id:"))
		fieldReplacer.WriteString(w, id)
		w.Write([]byte("\n"))
	}
}

func writeEvent(w io.Writer, event string) {
	if len(event) > 0 {
		w.Write([]byte("event:"))
		fieldReplacer.WriteString(w, event)
		w.Write([]byte("\n"))
	}
}

func writeRetry(w io.Writer, retry uint) {
	if retry > 0 {
		w.Write([]byte("retry:"))
		w.Write([]byte(strconv.FormatUint(uint64(retry), 10)))
		w.Write([]byte("\n"))
	}
}

func writeData(w io.Writer, data interface{}) error {
	w.Write([]byte("data:"))
	switch kindOfData(data) {
	case reflect.Struct, reflect.Slice, reflect.Map:
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			return err
		}
		w.Write([]byte("\n"))
	default:
		dataReplacer.WriteString(w, fmt.Sprint(data))
		w.Write([]byte("\n\n"))
	}
	return nil
}

func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
