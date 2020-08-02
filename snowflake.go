package snowflake

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	Epoch           int64 = 1420041600000
	SequenceBits    uint8 = 12
	SequenceMask    int64 = -1 ^ (-1 << SequenceBits)
	InstanceIdBits  uint8 = 10
	MaxInstanceId   int64 = -1 ^ (-1 << InstanceIdBits)
	InstanceIdShift uint8 = SequenceBits
	TimestampShift  uint8 = InstanceIdBits + SequenceBits
)

type Generator struct {
	mutex     sync.Mutex
	epoch     int64
	instance  int64
	timestamp int64
	sequence  int64
}

func (g *Generator) InstanceId() int64 {
	return g.instance
}

func (g *Generator) NextId() (int64, error) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	now := currentTimestamp()
	if now < g.timestamp {
		return 0, errors.New("Clock moved backwards")
	}

	if now > g.timestamp {
		g.timestamp = now
		g.sequence = 0
	} else { // now == g.timestamp
		g.sequence = (g.sequence + 1) & SequenceMask
		if g.sequence == 0 {
			for now <= g.timestamp {
				now = currentTimestamp()
			}
			g.timestamp = now
		}
	}

	id := int64((now-g.epoch)<<TimestampShift |
		(g.instance << InstanceIdShift) |
		(g.sequence))
	return id, nil
}

func currentTimestamp() int64 {
	return time.Now().UnixNano() / 1000000
}

func NewGeneratorWithEpoch(epoch int64, instanceId int64) (*Generator, error) {
	if instanceId < 0 || instanceId > MaxInstanceId {
		return nil, fmt.Errorf("Instance id can't be greater than %d or less than 0", MaxInstanceId)
	}
	return &Generator{
		epoch:     epoch,
		instance:  instanceId,
		timestamp: 0,
		sequence:  0,
	}, nil
}

func NewGenerator(instanceId int64) (*Generator, error) {
	return NewGeneratorWithEpoch(Epoch, instanceId)
}
