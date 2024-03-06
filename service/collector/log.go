package collector

import (
	"fmt"
	"strconv"

	logApi "github.com/I-m-Surrounded-by-IoT/backend/api/log"
	"github.com/sirupsen/logrus"
	"github.com/zijiren233/stream"
	"google.golang.org/protobuf/proto"
)

var _ logrus.Formatter = (*kafkaDeviceLogKeyFormatter)(nil)

type kafkaDeviceLogKeyFormatter struct{}

func getDeviceID(entry *logrus.Entry) (uint64, error) {
	deviceIDI, ok := entry.Data["device_id"]
	if !ok {
		return 0, fmt.Errorf("missing device_id field")
	}
	deviceID, ok := deviceIDI.(uint64)
	if !ok {
		return 0, fmt.Errorf("invalid device_id type")
	}
	return deviceID, nil
}

func getTopic(entry *logrus.Entry) (string, error) {
	topicI, ok := entry.Data["topic"]
	if !ok {
		return "", fmt.Errorf("missing topic field")
	}
	topic, ok := topicI.(string)
	if !ok {
		return "", fmt.Errorf("invalid topic type")
	}
	return topic, nil
}

func (k *kafkaDeviceLogKeyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	deviceID, err := getDeviceID(entry)
	if err != nil {
		return nil, err
	}
	return stream.StringToBytes(strconv.FormatUint(deviceID, 10)), nil
}

var _ logrus.Formatter = (*kafkaDeviceLogValueFormatter)(nil)

type kafkaDeviceLogValueFormatter struct{}

func (k *kafkaDeviceLogValueFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	deviceID, err := getDeviceID(entry)
	if err != nil {
		return nil, err
	}
	topic, err := getTopic(entry)
	if err != nil {
		return nil, err
	}
	dl := logApi.DeviceLogData{
		DeviceId:  deviceID,
		Topic:     topic,
		Timestamp: entry.Time.UnixMilli(),
		Message:   entry.Message,
		Level:     uint32(entry.Level),
	}
	return proto.Marshal(&dl)
}

// type DeviceStreamLog struct {
// 	Time    time.Time
// 	Level   uint32
// 	Message string
// }

// type DeviceStreamLogChan struct {
// 	ch          chan *DeviceStreamLog
// 	levelFilter string
// }

// type DeviceStreamLogChans struct {
// 	chans  map[uint64]*DeviceStreamLogChan
// 	lock   sync.RWMutex
// 	closed bool
// }

// func (c *DeviceStreamLogChans) Watch(levelFilter string) (<-chan *DeviceStreamLog, func(), error) {
// 	c.lock.Lock()
// 	defer c.lock.Unlock()
// 	if c.closed {
// 		return nil, nil, fmt.Errorf("device log chans already closed")
// 	}
// 	id := rand.Uint64()
// 	if _, ok := c.chans[id]; ok {
// 		return nil, nil, fmt.Errorf("device log chan already exists")
// 	}
// 	ch := &DeviceStreamLogChan{
// 		levelFilter: levelFilter,
// 		ch:          make(chan *DeviceStreamLog),
// 	}
// 	c.chans[id] = ch
// 	return ch.ch, func() {
// 		c.lock.Lock()
// 		defer c.lock.Unlock()
// 		delete(c.chans, id)
// 		if !c.closed {
// 			close(ch.ch)
// 		}
// 	}, nil
// }

// func (c *DeviceStreamLogChans) WriteLog(log *DeviceStreamLog) {
// 	c.lock.RLock()
// 	defer c.lock.RUnlock()
// 	for _, ch := range c.chans {
// 		if strings.Contains(ch.levelFilter, strconv.FormatUint(uint64(log.Level), 10)) {
// 			continue
// 		}
// 		select {
// 		case ch.ch <- log:
// 		default:
// 		}
// 	}
// }

// func (c *DeviceStreamLogChans) Close() {
// 	c.lock.Lock()
// 	defer c.lock.Unlock()
// 	if c.closed {
// 		return
// 	}
// 	c.closed = true
// 	for id, ch := range c.chans {
// 		delete(c.chans, id)
// 		close(ch.ch)
// 	}
// }

// func NewDeviceStreamLogChans() *DeviceStreamLogChans {
// 	return &DeviceStreamLogChans{
// 		chans: make(map[uint64]*DeviceStreamLogChan),
// 	}
// }

// type DeviceStreamLogRegistor struct {
// 	m *rwmap.RWMap[uint64, *DeviceStreamLogChans]
// }

// func NewDeviceStreamLogRegistor() *DeviceStreamLogRegistor {
// 	return &DeviceStreamLogRegistor{
// 		m: &rwmap.RWMap[uint64, *DeviceStreamLogChans]{},
// 	}
// }

// func (r *DeviceStreamLogRegistor) RegisterDevice(id uint64) (*DeviceStreamLogChans, error) {
// 	dlc, loaded := r.m.LoadOrStore(id, NewDeviceStreamLogChans())
// 	if loaded {
// 		return nil, fmt.Errorf("device log chans already exists")
// 	}
// 	return dlc, nil
// }

// func (r *DeviceStreamLogRegistor) UnregisterDevice(id uint64, dlc *DeviceStreamLogChans) bool {
// 	return r.m.CompareAndDelete(id, dlc)
// }

// func (r *DeviceStreamLogRegistor) GetDeviceLogChans(id uint64) (*DeviceStreamLogChans, bool) {
// 	return r.m.Load(id)
// }

// func (r *DeviceStreamLogRegistor) Levels() []logrus.Level {
// 	return logrus.AllLevels
// }

// func marshalLog(entry *logrus.Entry) *DeviceStreamLog {
// 	return &DeviceStreamLog{
// 		Time:    entry.Time,
// 		Level:   uint32(entry.Level),
// 		Message: entry.Message,
// 	}
// }

// func (r *DeviceStreamLogRegistor) Fire(entry *logrus.Entry) error {
// 	deviceIDI, ok := entry.Data["device_id"]
// 	if !ok {
// 		return nil
// 	}
// 	deviceID, ok := deviceIDI.(uint64)
// 	if !ok {
// 		return fmt.Errorf("invalid device_id type")
// 	}
// 	dlc, ok := r.GetDeviceLogChans(deviceID)
// 	if !ok {
// 		return fmt.Errorf("device log chans not found")
// 	}
// 	dlc.WriteLog(marshalLog(entry))
// 	return nil
// }
