package roomba

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type stateMessage struct {
	State reportedMessage `json:"state"`
}

type reportedMessage struct {
	Reported map[string]interface{} `json:"reported"`
}

func (r *Roomba) stateMessageHandler(client mqtt.Client, msg mqtt.Message) {
	reportedState := stateMessage{}
	err := json.Unmarshal(msg.Payload(), &reportedState)
	if err != nil {
		fmt.Printf("Roomba state unmarshal error: %s", err.Error())
	}

	r.statusMutex.Lock()
	statusValue := reflect.ValueOf(r.status).Elem()
	for field, value := range reportedState.State.Reported {
		statusValue.FieldByName(strings.Title(field)).Set(reflect.ValueOf(value))
	}
	r.statusMutex.Unlock()
}
