package roomba

import (
	"encoding/json"
	"log"
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
	if r.debug {
		log.Printf("RECV: %s", string(msg.Payload()))
	}

	if r.stateWriter != nil {
		r.stateWriter.Write(append(msg.Payload(), byte('\n')))
	}

	reportedState := stateMessage{}
	err := json.Unmarshal(msg.Payload(), &reportedState)
	if err != nil {
		log.Printf("Roomba state unmarshal error: %s", err.Error())
	}

	r.statusMutex.Lock()
	statusValue := reflect.ValueOf(r.status).Elem()
	for field, value := range reportedState.State.Reported {
		statusValue.FieldByName(strings.Title(field)).Set(reflect.ValueOf(value))
	}
	r.statusMutex.Unlock()
}
