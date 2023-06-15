package comm

import (
	"errors"
	"io"
	"sync"
)

type MockTransport struct {
	Connected bool
	Written   []string
	Data      map[string]string
	mutex     sync.RWMutex
}

func (t *MockTransport) Connect() error {
	if t.Connected {
		return errors.New("already connected")
	}
	t.Connected = true
	return nil
}

func (t *MockTransport) Read(p []byte) (n int, err error) {
	if !t.Connected {
		return 0, errors.New("not connected")
	}

	t.mutex.Lock()
	defer t.mutex.Unlock()

	if len(t.Written) == 0 {
		return 0, io.EOF
	}

	readData := t.Data[t.Written[0]]
	copy(p, []byte(readData))
	n = len(readData)

	t.Written = t.Written[1:] // Remove the read data from the written list

	return n, nil
}

func (t *MockTransport) Write(p []byte) (n int, err error) {
	if !t.Connected {
		return 0, errors.New("not connected")
	}

	writtenData := string(p)

	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.Written = append(t.Written, writtenData)

	return len(p), nil
}

func (t *MockTransport) Close() error {
	if !t.Connected {
		return errors.New("not connected")
	}
	t.Connected = false
	return nil
}

// Ensure that MockTransport implements the io.ReadWriteCloser interface
var _ io.ReadWriteCloser = &MockTransport{}
