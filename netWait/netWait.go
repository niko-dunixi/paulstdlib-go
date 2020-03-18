package netWait

import (
	"fmt"
	"net"
	"time"
)

type TimeoutError struct {
	Host string
	Port int
}

func (timeoutError TimeoutError) Error() string {
	return fmt.Sprintf("timeout occurred for host: %s port: %d", timeoutError.Host, timeoutError.Port)
}

func WaitPort(host string, port int, timeout time.Duration) error {
	timer := time.NewTimer(0)
	defer func() {
		if !timer.Stop() {
			_ = <-timer.C
		}
	}()
	timeoutChannel := time.After(timeout)
	for {
		select {
		case _ = <-timer.C:
			portString := fmt.Sprintf("%d", port)
			connection, _ := net.DialTimeout("tcp", net.JoinHostPort(host, portString), time.Second*2)
			if connection != nil {
				connection.Close()
				return nil
			}
			timer.Reset(time.Millisecond * 250)
		case _ = <-timeoutChannel:
			return TimeoutError{
				Host: host,
				Port: port,
			}
		}
	}
}
