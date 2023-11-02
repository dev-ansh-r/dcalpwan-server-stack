//go:build !linux && !darwin
// +build !linux,!darwin

package store

import "net"

func checkConn(conn net.Conn) error {
	return nil
}
