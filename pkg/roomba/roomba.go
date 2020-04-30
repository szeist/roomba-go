package roomba

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"os"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/szeist/roomba-go/pkg/config"
	"github.com/szeist/roomba-go/pkg/status"
)

const statusPollIntervalMs time.Duration = 500

type Roomba struct {
	client       mqtt.Client
	status       *status.Status
	statusMutex  *sync.Mutex
	isConnected  bool
	debug        bool
	statusWriter io.Writer
}

func New(cfg *config.Config) *Roomba {
	r := &Roomba{
		status:       &status.Status{},
		statusMutex:  &sync.Mutex{},
		debug:        cfg.Debug,
		statusWriter: *cfg.StatusWriter,
	}

	if cfg.Debug {
		mqtt.DEBUG = log.New(os.Stderr, cfg.LogPrefix, 0)
	}
	mqtt.ERROR = log.New(os.Stderr, cfg.LogPrefix, 0)
	opts := mqtt.NewClientOptions().AddBroker("ssl://" + cfg.Address + ":8883").SetClientID(cfg.User)
	opts.Username = cfg.User
	opts.Password = cfg.Password
	opts.ProtocolVersion = 4
	opts.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	opts.SetDefaultPublishHandler(r.stateMessageHandler)

	r.client = mqtt.NewClient(opts)

	return r
}

func (r *Roomba) Connect() error {
	token := r.client.Connect()
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (r *Roomba) Disconnect() {
	r.client.Disconnect(250)
}

func (r *Roomba) SendCommand(cmd string) error {
	roombaCommand, err := createCommand(cmd)
	if err != nil {
		return err
	}

	jsonCmd, err := json.Marshal(roombaCommand)
	if err != nil {
		return err
	}

	token := r.client.Publish("cmd", 0, false, jsonCmd)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func (r *Roomba) GetStatus(timeout time.Duration) status.Status {
	for i := 0; i < int(timeout/statusPollIntervalMs); i++ {
		if r.status.IsAllValuesPresent() {
			break
		}
		time.Sleep(statusPollIntervalMs * time.Millisecond)
	}
	return *r.status
}

func (r *Roomba) IsConnected() bool {
	return r.isConnected
}

func GetSupportedCommands() []string {
	return commands
}
