package sse

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
	writeEvent(w, event.Event)
	err := writeData(w, event.Data)
	writeId(w, event.Id)
	writeRetry(w, event.Retry)

	return err
}

func WriteHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Content-Type", "text/event-stream")
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

func writeData[T any](w io.Writer, data T) error {
	w.Write([]byte("data: "))
	switch kindOfData(data) {
	case reflect.Struct, reflect.Slice, reflect.Map:
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			return err
		}
		w.Write([]byte("\n"))
	default:
		fmt.Printf("Writing data %s to event\n", fmt.Sprint(data))

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
