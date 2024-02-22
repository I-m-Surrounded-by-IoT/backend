package collector

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/zijiren233/gencontainer/rwmap"
)

var _ logrus.Formatter = (*kafkaDeviceLogKeyFormatter)(nil)

type kafkaDeviceLogKeyFormatter struct{}

func (k *kafkaDeviceLogKeyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	deviceIDI, ok := entry.Data["device_id"]
	if !ok {
		return nil, fmt.Errorf("missing device_id field")
	}
	deviceID, ok := deviceIDI.(uint64)
	if !ok {
		return nil, fmt.Errorf("invalid device_id type")
	}
	return []byte(fmt.Sprintf("%v", deviceID)), nil
}

type DeviceLog struct {
	Time    time.Time
	Level   uint32
	Message string
}

type DeviceLogChan struct {
	ch       chan *DeviceLog
	minLevel uint32
}

type DeviceLogChans struct {
	chans  map[uint64]*DeviceLogChan
	lock   sync.RWMutex
	closed bool
}

func (c *DeviceLogChans) Watch(level uint32) (<-chan *DeviceLog, func(), error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.closed {
		return nil, nil, fmt.Errorf("device log chans already closed")
	}
	id := rand.Uint64()
	if _, ok := c.chans[id]; ok {
		return nil, nil, fmt.Errorf("device log chan already exists")
	}
	ch := &DeviceLogChan{
		minLevel: level,
		ch:       make(chan *DeviceLog),
	}
	c.chans[id] = ch
	return ch.ch, func() {
		c.lock.Lock()
		defer c.lock.Unlock()
		delete(c.chans, id)
		if !c.closed {
			close(ch.ch)
		}
	}, nil
}

func (c *DeviceLogChans) WriteLog(log *DeviceLog) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	for _, ch := range c.chans {
		if log.Level > ch.minLevel {
			continue
		}
		select {
		case ch.ch <- log:
		default:
		}
	}
}

func (c *DeviceLogChans) Close() {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.closed {
		return
	}
	c.closed = true
	for id, ch := range c.chans {
		delete(c.chans, id)
		close(ch.ch)
	}
}

func NewDeviceLogChans() *DeviceLogChans {
	return &DeviceLogChans{
		chans: make(map[uint64]*DeviceLogChan),
	}
}

type DeviceLogChanRegistor struct {
	m *rwmap.RWMap[uint64, *DeviceLogChans]
}

func NewDeviceLogChanRegistor() *DeviceLogChanRegistor {
	return &DeviceLogChanRegistor{
		m: &rwmap.RWMap[uint64, *DeviceLogChans]{},
	}
}

func (r *DeviceLogChanRegistor) RegisterDevice(id uint64) (*DeviceLogChans, error) {
	dlc, loaded := r.m.LoadOrStore(id, NewDeviceLogChans())
	if loaded {
		return nil, fmt.Errorf("device log chans already exists")
	}
	return dlc, nil
}

func (r *DeviceLogChanRegistor) UnregisterDevice(id uint64, dlc *DeviceLogChans) bool {
	return r.m.CompareAndDelete(id, dlc)
}

func (r *DeviceLogChanRegistor) GetDeviceLogChans(id uint64) (*DeviceLogChans, bool) {
	return r.m.Load(id)
}

func (r *DeviceLogChanRegistor) Levels() []logrus.Level {
	return logrus.AllLevels
}

func marshalLog(entry *logrus.Entry) *DeviceLog {
	return &DeviceLog{
		Time:    entry.Time,
		Level:   uint32(entry.Level),
		Message: entry.Message,
	}
}

func (r *DeviceLogChanRegistor) Fire(entry *logrus.Entry) error {
	deviceIDI, ok := entry.Data["device_id"]
	if !ok {
		return nil
	}
	deviceID, ok := deviceIDI.(uint64)
	if !ok {
		return fmt.Errorf("invalid device_id type")
	}
	dlc, ok := r.GetDeviceLogChans(deviceID)
	if !ok {
		return fmt.Errorf("device log chans not found")
	}
	dlc.WriteLog(marshalLog(entry))
	return nil
}
