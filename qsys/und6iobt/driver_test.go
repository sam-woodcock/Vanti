package und6iobt_test

import (
	"errors"
	"github.com/vanti-dev/assessment-syseng-go/bluetooth"
	mockT "github.com/vanti-dev/assessment-syseng-go/comm"
	"github.com/vanti-dev/assessment-syseng-go/qsys/und6iobt"
	"testing"
)

func TestDriver_Announce_Success(t *testing.T) {
	mock := &mockT.MockTransport{
		Data: map[string]string{

			"BTB<CR>": "ACK BTB<CR>",
		},
	}
	driver := und6iobt.New(mock, "127.0.0.1", 22)

	err := driver.Announce()

	if err != nil {
		t.Errorf("Announce() returned an error: %s", err)
	}
}

func TestDriver_Announce_PairingFailed(t *testing.T) {
	mock := &mockT.MockTransport{
		Data: map[string]string{
			"BTB<CR>": "NACK BTB<CR>",
		},
	}
	driver := und6iobt.New(mock, "127.0.0.1", 22)

	err := driver.Announce()

	expectedErr := errors.New(und6iobt.PairingFailed)
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("Announce() returned unexpected error: %s, expected: %s", err, expectedErr)
	}
}

func TestDriver_Name_Success(t *testing.T) {
	mock := &mockT.MockTransport{
		Data: map[string]string{
			"BTN<CR>": "ACK BTN unD6IO-BT <CR>",
		},
	}
	driver := und6iobt.New(mock, "127.0.0.1", 22)

	name, err := driver.Name()

	if err != nil {
		t.Errorf("Name() returned an error: %s", err)
	}

	if name != "unD6IO-BT" {
		t.Errorf("Name() returned unexpected name: %s, expected: unD6IO-BT", name)
	}
}

func TestDriver_Name_Failed(t *testing.T) {
	mock := &mockT.MockTransport{
		Data: map[string]string{
			"BTN<CR>": "Error",
		},
	}
	driver := und6iobt.New(mock, "127.0.0.1", 22)

	_, err := driver.Name()

	expectedErr := errors.New(bluetooth.NameFailed)
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("Announce() returned unexpected error: %s, expected: %s", err, expectedErr)
	}
}

func TestDriver_ConnectionChanged_ConnectionChanged(t *testing.T) {
	mock := &mockT.MockTransport{
		Data: map[string]string{
			"BTS<CR>": "ACK BTS 2<CR>",
		},
	}
	driver := und6iobt.New(mock, "127.0.0.1", 22)

	last := bluetooth.ConnectionUnknown
	expectedConn := bluetooth.ConnectionConnected

	conn, err := driver.ConnectionChanged(last)

	if err != nil {
		t.Errorf("ConnectionChanged() returned an error: %s", err)
	}

	if conn != expectedConn {
		t.Errorf("ConnectionChanged() returned unexpected result: %d, expected: %d", conn, expectedConn)
	}

	t.Logf("ConnectionChanged() returned result: %d, expected: %d", conn, expectedConn)
}
