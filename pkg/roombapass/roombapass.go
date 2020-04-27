package roombapass

import (
	"crypto/tls"
	"errors"
)

func GetPassword(address string) (string, error) {
	conn, err := connectToRoomba(address)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	err = sendMagicPacket(conn)
	if err != nil {
		return "", err
	}

	err = readInitialResponse(conn)
	if err != nil {
		return "", err
	}

	password, err := readPasswordResponse(conn)
	if err != nil {
		return "", err
	}

	return password, nil
}

func connectToRoomba(address string) (*tls.Conn, error) {
	conn, err := tls.Dial("tcp", address+":8883", &tls.Config{
		InsecureSkipVerify: true,
	})
	return conn, err
}

func sendMagicPacket(conn *tls.Conn) error {
	msg := []byte{0xf0, 0x05, 0xef, 0xcc, 0x3b, 0x29, 0x00}
	_, err := conn.Write(msg)
	return err
}

func readInitialResponse(conn *tls.Conn) error {
	buf := make([]byte, 128)
	readLen, err := conn.Read(buf)
	if err != nil {
		return err
	}
	if readLen != 2 {
		return errors.New("Unknown initial response")
	}
	return nil
}

func readPasswordResponse(conn *tls.Conn) (string, error) {
	buf := make([]byte, 128)
	readLen, err := conn.Read(buf)
	if err != nil {
		return "", err
	}
	if readLen < 8 {
		return "", errors.New("Cannot parse password response")
	}

	return string(buf[5:readLen]), nil
}
