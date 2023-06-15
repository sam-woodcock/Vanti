package und6iobt

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/vanti-dev/assessment-syseng-go/bluetooth"
	"github.com/vanti-dev/assessment-syseng-go/comm"
)

const (
	CR    = "<CR>"
	Error = "Error"
)

// Driver implements interrogation and control of a Q-SYS unD6IO-BT device.
type Driver struct {
	// Comm abstracts the underlying transport communication with the device.
	// We can assume that the underlying implementation of Comm is compatible with our device model.
	Comm      comm.Transport
	IPAddress string
	Port      int
}

func New(transport comm.Transport, ipAddress string, port int) *Driver {
	return &Driver{
		Comm:      transport,
		IPAddress: ipAddress,
		Port:      port,
	}
}

func (dc *Driver) Announce() error {
	if err := dc.connectAndClose(func() error {
		n, err := dc.write([]byte(ActivatePairing))
		if err != nil {
			return err
		}

		log.Printf("Bytes written: %d\n", n)

		// Read response
		command := dc.read(32)
		if string(command) == ActivatePairingResponseOk {
			return nil
		}

		if string(command) == ActivatePairingResponseNo {
			return errors.New(PairingFailed)
		}

		return errors.New(command) // Push message back up
	}); err != nil {
		return err
	}

	return nil
}

func (dc *Driver) Name() (string, error) {
	var name string
	if err := dc.connectAndClose(func() error {
		n, err := dc.write([]byte(bluetooth.GetName))
		if err != nil {
			return err
		}

		log.Printf("Bytes written: %d\n", n)

		//Read response
		command := dc.read(32)
		if string(command) != Error {
			name = strings.TrimSuffix(command, CR)
			name = strings.TrimPrefix(name, "ACK BTN ")
			name = strings.TrimSpace(name)
			return nil
		}
		return errors.New(bluetooth.NameFailed)
	}); err != nil {
		return "", err
	}

	return name, nil
}

func (dc *Driver) ConnectionChanged(last bluetooth.Connection) (bluetooth.Connection, error) {
	var current bluetooth.Connection
	err := dc.connectAndClose(func() error {
		var err error
		current, err = dc.checkConnection()
		return err
	})

	if err != nil {
		return bluetooth.ConnectionUnknown, err
	}

	if last != current {
		return current, nil
	}

	return bluetooth.ConnectionUnknown, errors.New(ErrorMessage)
}

func (dc *Driver) checkConnection() (bluetooth.Connection, error) {
	var connection bluetooth.Connection

	n, err := dc.write([]byte(BTStatus))
	if err != nil {
		return bluetooth.ConnectionUnknown, err
	}

	log.Printf("Bytes written: %d\n", n)

	// Read response
	command := dc.read(16)

	re := regexp.MustCompile(`\d+`)
	stringNumber := re.FindString(command)
	match, err := regexp.MatchString(`\d+`, command)
	if err != nil {
		log.Println("Error:", err)
		return bluetooth.ConnectionUnknown, errors.New(Error)
	}
	if match {
		num, err := strconv.Atoi(stringNumber)
		if err != nil {
			log.Println("Error:", err)
			return bluetooth.ConnectionUnknown, errors.New(Error)
		}

		status := bluetooth.GetConnectionStatus(num)
		connection = bluetooth.Connection(status)
	}

	return connection, nil
}

func (dc *Driver) write(p []byte) (int, error) {
	n, err := dc.Comm.Write(p)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (dc *Driver) read(count int) string {
	readData := make([]byte, count)

	n, err := dc.Comm.Read(readData)
	if err != nil {
		fmt.Println("Failed to read from connection:", err)
		return ErrorMessage
	}
	return string(readData[:n])

}

func (dc *Driver) connectAndClose(action func() error) error {
	if err := dc.Comm.Connect(); err != nil {
		return err
	}
	defer dc.Comm.Close()

	return action()
}
