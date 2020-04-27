package discover

import (
	"fmt"
	"net"
	"strings"
)

type DiscoveryResult struct {
	Address net.Addr
	Roomba  *DiscoveredRoomba
}

type DiscoveredRoomba struct {
	Ver       string `json:"ver"`
	Hostname  string `json:"hostname"`
	Robotname string `json:"robotname"`
	Ip        string `json:"ip"`
	Mac       string `json:"mac"`
	Sw        string `json:"sw"`
	Sku       string `json:"sku"`
	Nc        int    `json:"nc"`
	Proto     string `json:"proto"`
	Cap       Cap    `json:"cap"`
}

type Cap struct {
	Pose          int `json:"pose"`
	Ota           int `json:"ota"`
	MultiPass     int `json:"multiPass"`
	Pp            int `json:"pp"`
	BinFullDetect int `json:"binFullDetect"`
	LangOta       int `json:"langOta"`
	Maps          int `json:"maps"`
	Edge          int `json:"edge"`
	Eco           int `json:"eco"`
	SvcConf       int `json:"svcConf"`
}

func (d *DiscoveryResult) String() string {
	return fmt.Sprintf("%s\tRobotname: %s\tMAC: %s\tSW: %s",
		strings.Split(d.Address.String(), ":")[0],
		d.Roomba.Robotname,
		d.Roomba.Mac,
		d.Roomba.Sw,
	)
}
