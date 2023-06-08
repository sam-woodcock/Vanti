package und6iobt

import (
	"errors"
	"fmt"
	"github.com/vanti-dev/assessment-syseng-go/bluetooth"
	_ "github.com/vanti-dev/assessment-syseng-go/bluetooth"
	"github.com/vanti-dev/assessment-syseng-go/comm"
	"regexp"
	"strconv"
	"strings"
)

// Driver implements interrogation and control of a Q-SYS unD6IO-BT device.
type Driver struct {
	// Comm abstracts the underlying transport communication with the device.
	// We can assume that the underlying implementation of Comm is compatible with our device model.
	Comm comm.Transport
}

func (dc *Driver) Announce() error {

	err := dc.Comm.Connect()
	if err != nil {
		return err
	}

	n, err := dc.Write([]byte(ActivatePairing))
	if err != nil {
		return err
	}

	fmt.Printf("Bytes written: %d\n", n)

	//read response
	command := dc.Read(32)
	if string(command) == ActivatePairingResponseOk {
		return nil
	}

	if string(command) == ActivatePairingResponseNo {
		return errors.New(PairingFailed)
	}

	return errors.New(command) //push message back up
}
func (dc *Driver) Name() (string, error) {

	err := dc.Comm.Connect()
	if err != nil {
		return "", err
	}

	n, err := dc.Write([]byte(bluetooth.GetName))
	if err != nil {
		return "", err
	}

	fmt.Printf("Bytes written: %d\n", n)

	//read response
	command := dc.Read(16)
	if string(command) != Error {
		return strings.TrimSuffix(command, "<CR"), nil // Output: e.g unD6IO-BT-010203 from ACK BTN unD6IO-BT-010203<CR
	}
	return "", errors.New(bluetooth.NameFailed)

}
func (dc *Driver) ConnectionChanged(last bluetooth.Connection) (bluetooth.Connection, error) {

	var current, _ = dc.CheckConnection()
	if last != current {
		return current, nil
	}
	return bluetooth.ConnectionUnknown, errors.New(Error)
}

func (dc *Driver) CheckConnection() (bluetooth.Connection, error) {
	err := dc.Comm.Connect()
	if err != nil {
		return bluetooth.ConnectionUnknown, err
	}

	n, err := dc.Write([]byte(BTStatus))
	if err != nil {
		return bluetooth.ConnectionUnknown, err
	}

	fmt.Printf("Bytes written: %d\n", n)

	//read response
	command := dc.Read(16)

	re := regexp.MustCompile(`\d+`)
	stringNumber := re.FindString(command)
	match, err := regexp.MatchString(`\d+`, command)
	if err != nil {
		fmt.Println("Error:", err)
		return bluetooth.ConnectionUnknown, errors.New(Error)
	}
	if match {
		num, err := strconv.Atoi(stringNumber)
		if err != nil {
			fmt.Println("Error:", err)
			return bluetooth.ConnectionUnknown, errors.New(Error)
		}

		//extract value back
		//0 = Idle
		//1 = Discoverable
		//2 = Connected – Unknown AVRCP support
		//3 = Connected – AVRCP Not Supported
		//4 = Connected – AVRCP Supported
		//5 = Connected – AVRCP & PDU Supported

		if num > 2 {
			return bluetooth.ConnectionConnected, nil // Output: e.g unD6IO-BT-010203 from ACK BTN unD6IO-BT-010203<CR
		}
	}
	return bluetooth.ConnectionUnknown, errors.New(Error)

}

func (dc *Driver) Write(p []byte) (int, error) {
	n, err := dc.Comm.Write(p)
	if err != nil {
		return 0, err
	}
	return n, nil
}
func (dc *Driver) Read(count int) string {
	readData := make([]byte, count)
	command, err := dc.Comm.Read(readData)
	if err != nil {
		return Error
	}
	return string(command)
}
