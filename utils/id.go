package utils

import (
	"github.com/hashicorp/go-uuid"
	"log"
	"sync"
	"time"
)

const (
	workerIdBits     int64 = 5
	datacenterIdBits int64 = 5
	sequenceBits     int64 = 12

	maxWorkerId     int64 = -1 ^ (-1 << uint64(workerIdBits))
	maxDatacenterId int64 = -1 ^ (-1 << uint64(datacenterIdBits))
	maxSequence     int64 = -1 ^ (-1 << uint64(sequenceBits))

	timeLeft uint8 = 22
	dataLeft uint8 = 17
	workLeft uint8 = 12

	twepoch int64 = 1525705533000
)

func UUID() (string, error) {
	return uuid.GenerateUUID()
}

type SnowFlake struct {
	mu           sync.Mutex
	laststamp    int64
	workerid     int64
	datacenterid int64
	sequence     int64
}

func (w *SnowFlake) getCurrentTime() int64 {
	return time.Now().UnixNano() / 1e6
}

//var i int = 1
func (w *SnowFlake) nextId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	timestamp := w.getCurrentTime()
	if timestamp < w.laststamp {
		log.Fatal("can not generate id")
	}
	if w.laststamp == timestamp {
		// 这其实和 <==>
		// w.sequence++
		// if w.sequence++ > maxSequence  等价
		w.sequence = (w.sequence + 1) & maxSequence
		if w.sequence == 0 {
			// 之前使用 if, 只是没想到 GO 可以在一毫秒以内能生成到最大的 Sequence, 那样就会导致很多重复的
			// 这个地方使用 for 来等待下一毫秒
			for timestamp <= w.laststamp {
				//i++
				//fmt.Println(i)
				timestamp = w.getCurrentTime()
			}
		}
	} else {
		w.sequence = 0
	}
	w.laststamp = timestamp

	return ((timestamp - twepoch) << timeLeft) | (w.datacenterid << dataLeft) | (w.workerid << workLeft) | w.sequence
}

func (w *SnowFlake) tilNextMillis() int64 {
	timestamp := w.getCurrentTime()
	if timestamp <= w.laststamp {
		timestamp = w.getCurrentTime()
	}
	return timestamp
}
