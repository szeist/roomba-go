package status

import (
	"reflect"
)

type Status struct {
	Audio              interface{} `json:"audio"`
	BatPct             interface{} `json:"batPct"`
	BatteryType        interface{} `json:"batteryType"`
	Bbchg              interface{} `json:"bbchg"`
	Bbchg3             interface{} `json:"bbchg3"`
	Bbmssn             interface{} `json:"bbmssn"`
	Bbnav              interface{} `json:"bbnav"`
	Bbpanic            interface{} `json:"bbpanic"`
	Bbpause            interface{} `json:"bbpause"`
	Bbrstinfo          interface{} `json:"bbrstinfo"`
	Bbrun              interface{} `json:"bbrun"`
	Bbswitch           interface{} `json:"bbswitch"`
	Bbsys              interface{} `json:"bbsys"`
	Bin                interface{} `json:"bin"`
	BinPause           interface{} `json:"binPause"`
	BootloaderVer      interface{} `json:"bootloaderVer"`
	Cap                interface{} `json:"cap"`
	CarpetBoost        interface{} `json:"carpetBoost"`
	CleanMissionStatus interface{} `json:"cleanMissionStatus"`
	CleanSchedule      interface{} `json:"cleanSchedule"`
	CloudEnv           interface{} `json:"cloudEnv"`
	Country            interface{} `json:"country"`
	Dock               interface{} `json:"dock"`
	EcoCharge          interface{} `json:"ecoCharge"`
	HardwareRev        interface{} `json:"hardwareRev"`
	Langs              interface{} `json:"langs"`
	Language           interface{} `json:"language"`
	LastCommand        interface{} `json:"lastCommand"`
	Localtimeoffset    interface{} `json:"localtimeoffset"`
	Mac                interface{} `json:"mac"`
	MapUploadAllowed   interface{} `json:"mapUploadAllowed"`
	MobilityVer        interface{} `json:"mobilityVer"`
	Name               interface{} `json:"name"`
	NavSwVer           interface{} `json:"navSwVer"`
	Netinfo            interface{} `json:"netinfo"`
	NoAutoPasses       interface{} `json:"noAutoPasses"`
	NoPP               interface{} `json:"noPP"`
	OpenOnly           interface{} `json:"openOnly"`
	Pose               interface{} `json:"pose"`
	SchedHold          interface{} `json:"schedHold"`
	Signal             interface{} `json:"signal"`
	Sku                interface{} `json:"sku"`
	SoftwareVer        interface{} `json:"softwareVer"`
	SoundVer           interface{} `json:"soundVer"`
	SvcEndpoints       interface{} `json:"svcEndpoints"`
	Timezone           interface{} `json:"timezone"`
	TwoPass            interface{} `json:"twoPass"`
	Tz                 interface{} `json:"tz"`
	UiSwVer            interface{} `json:"uiSwVer"`
	UmiVer             interface{} `json:"umiVer"`
	Utctime            interface{} `json:"utctime"`
	VacHigh            interface{} `json:"vacHigh"`
	Wifistat           interface{} `json:"wifistat"`
	WifiSwVer          interface{} `json:"wifiSwVer"`
	Wlcfg              interface{} `json:"wlcfg"`
}

func (s *Status) IsAllValuesPresent() bool {
	v := reflect.ValueOf(s).Elem()

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).IsNil() {
			return false
		}
	}

	return true
}
