package comm_test

import (
	"errors"
	"io"
	"testing"

	"github.com/vanti-dev/assessment-syseng-go/comm"
)

func TestMockTransport_Connect_Success(t *testing.T) {
	tt := &comm.MockTransport{}

	err := tt.Connect()

	if err != nil {
		t.Errorf("Connect() returned an error: %s", err)
	}

	if !tt.Connected {
		t.Error("Connect() did not establish the connection")
	}
}

func TestMockTransport_Connect_AlreadyConnected(t *testing.T) {
	tt := &comm.MockTransport{}
	tt.Connect()

	err := tt.Connect()

	expectedErr := errors.New("already connected")
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("Connect() returned unexpected error: %s, expected: %s", err, expectedErr)
	}
}

func TestMockTransport_Read_Success(t *testing.T) {
	tt := &comm.MockTransport{
		Data: map[string]string{
			"Data1": "Data1",
		},
	}
	tt.Connect()

	readData := make([]byte, 5)
	_, err := tt.Write([]byte("Data1"))
	if err != nil {
		t.Errorf("Write() returned an error: %s", err)
	}
	n, err := tt.Read(readData)

	if err != nil {
		t.Errorf("Read() returned an error: %s", err)
	}

	expectedData := "Data1"
	if string(readData[:n]) != expectedData {
		t.Errorf("Read() returned unexpected data: %s, expected: %s", string(readData[:n]), expectedData)
	}
}

func TestMockTransport_Read_NotConnected(t *testing.T) {
	tt := &comm.MockTransport{}

	readData := make([]byte, 5)

	_, err := tt.Read(readData)

	expectedErr := errors.New("not connected")
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("Read() returned unexpected error: %s, expected: %s", err, expectedErr)
	}
}

func TestMockTransport_Read_EOF(t *testing.T) {
	tt := &comm.MockTransport{}
	tt.Connect()

	readData := make([]byte, 5)

	_, err := tt.Read(readData)

	if err != io.EOF {
		t.Errorf("Read() returned unexpected error: %s, expected: %s", err, io.EOF)
	}
}

func TestMockTransport_Write_Success(t *testing.T) {
	tt := &comm.MockTransport{}
	tt.Connect()

	writeData := []byte("Data")

	n, err := tt.Write(writeData)

	if err != nil {
		t.Errorf("Write() returned an error: %s", err)
	}

	if n != len(writeData) {
		t.Errorf("Write() returned unexpected number of bytes written: %d, expected: %d", n, len(writeData))
	}

}

func TestMockTransport_Write_NotConnected(t *testing.T) {
	tt := &comm.MockTransport{}

	writeData := []byte("Data")

	_, err := tt.Write(writeData)

	expectedErr := errors.New("not connected")
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("Write() returned unexpected error: %s, expected: %s", err, expectedErr)
	}
}

func TestMockTransport_Close_Success(t *testing.T) {
	tt := &comm.MockTransport{}
	tt.Connect()

	err := tt.Close()

	if err != nil {
		t.Errorf("Close() returned an error: %s", err)
	}

	if tt.Connected {
		t.Error("Close() did not close the connection")
	}
}

func TestMockTransport_Close_NotConnected(t *testing.T) {
	tt := &comm.MockTransport{}

	err := tt.Close()

	expectedErr := errors.New("not connected")
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("Close() returned unexpected error: %s, expected: %s", err, expectedErr)
	}
}
