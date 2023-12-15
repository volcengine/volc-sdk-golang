package hi_sse

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"regexp"
)

var (
	fieldData = []byte("data")

	pat2Eols = regexp.MustCompile(`[\r\n]{2}`)
)

type (
	Event struct {
		Data []byte
	}

	EventStream struct {
		scanner *bufio.Scanner
	}
)

// NewEventStreamFromReader creates an instance of EventStream.
func NewEventStreamFromReader(stream io.Reader, maxBufferSize int) *EventStream {
	scanner := bufio.NewScanner(stream)
	initBufferSize := maxBufferSize
	if initBufferSize > 4096 {
		initBufferSize = 4096
	}
	scanner.Buffer(make([]byte, initBufferSize), maxBufferSize)

	scanner.Split(func(data []byte, atEOF bool) (int, []byte, error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if loc := pat2Eols.FindIndex(data); len(loc) > 1 && loc[0] >= 0 {
			return loc[1], data[:loc[0]], nil
		}

		if atEOF {
			return len(data), data, nil
		}

		// Request more data.
		return 0, nil, nil
	})

	return &EventStream{
		scanner: scanner,
	}
}

func (e *EventStream) Next() (*Event, error) {
	if e.scanner.Scan() {
		raw := e.scanner.Bytes()
		s := string(raw)

		fmt.Sprintf(s)
		colonIndex := bytes.Index(raw, []byte("data:"))
		if colonIndex != -1 {
			colonIndex += len("data:") - 1
		}
		if colonIndex == 0 {
			// skip comment
			return nil, nil
		}

		var value []byte
		if colonIndex < 0 {
			//field = raw
			value = raw
		} else {
			//field = raw[:colonIndex]
			value = raw[colonIndex+1:]
		}

		//if bytes.Equal(field, fieldData) {
		return &Event{Data: value}, nil
		//}
		// invalid event, just ignore it
		//return nil, nil
	}
	if err := e.scanner.Err(); err != nil {
		if errors.Is(err, context.Canceled) {
			return nil, io.EOF
		}
		return nil, err
	}
	return nil, io.EOF
}
