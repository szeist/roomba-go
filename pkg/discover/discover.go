package discover

import (
	"encoding/json"
	"net"
	"time"
)

const roombaDiscoverPort string = "5678"

func Discover(timeout time.Duration) ([]*DiscoveryResult, error) {
	var results = []*DiscoveryResult{}

	conn, err := createBroadcastConnection(timeout)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	err = sendDiscoveryMessage(conn)
	if err != nil {
		return nil, err
	}

	for {
		res, err := readDiscoveryResponse(conn)
		if err != nil {
			break
		}
		results = append(results, res)
	}

	return results, nil
}

func createBroadcastConnection(timeout time.Duration) (*net.UDPConn, error) {
	localAddr, err := net.ResolveUDPAddr("udp4", ":0")
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP("udp4", localAddr)
	if err != nil {
		return nil, err
	}

	conn.SetReadDeadline(time.Now().Add(timeout))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func sendDiscoveryMessage(conn *net.UDPConn) error {
	broadcastAddr, err := net.ResolveUDPAddr("udp4", "255.255.255.255:"+roombaDiscoverPort)
	if err != nil {
		return err
	}

	_, err = conn.WriteToUDP([]byte("irobotmcs"), broadcastAddr)
	return err
}

func readDiscoveryResponse(conn *net.UDPConn) (*DiscoveryResult, error) {
	buf := make([]byte, 1024)
	readLen, addr, err := conn.ReadFrom(buf)
	if err != nil {
		return nil, err
	}

	result := &DiscoveredRoomba{}
	err = json.Unmarshal(buf[0:readLen], &result)
	if err != nil {
		return nil, err
	}

	return &DiscoveryResult{
		Address: addr,
		Roomba:  result,
	}, nil
}
