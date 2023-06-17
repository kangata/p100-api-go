package services

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/artemvang/p100-go"
)

var device *p100.P100Device

var lock = &sync.Mutex{}

func GetDeviceInstance() (*p100.P100Device, error) {
	lock.Lock()

	defer lock.Unlock()

	if device == nil {
		fmt.Printf("[%s] Initial device instance", time.Now().Format("2006-01-02 15:04:05"))

		device = p100.New(os.Getenv("DEVICE_IP"), os.Getenv("USER_EMAIL"), os.Getenv("USER_PASSWORD"))

		if err := device.Handshake(); err != nil {
			return nil, err
		}

		if err := device.Login(); err != nil {
			return nil, err
		}
	}

	return device, nil
}
