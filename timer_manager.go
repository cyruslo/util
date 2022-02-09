package util

import (
	"github.com/cyruslo/util/gtimer"
	"sync"
	"time"
)

const (
	//保留10000以下的timer id作为全局id
	INIT_TIMER_ID = 10000
)

var (
	genID    int32
	timerMap sync.Map
	mutex    sync.Mutex
)

func init() {
	genID = INIT_TIMER_ID
	timerMap = sync.Map{}
}

//Once run once after delay millsecond
func Once(delay int64, callback func()) int32 {
	mutex.Lock()
	defer mutex.Unlock()
	timerID := genTimerID()
	ch := gtimer.Once(time.Duration(delay*1000*1000), func() {
		delTimer(timerID)
		callback()
	})
	addTimer(timerID, ch)
	return timerID
}

// Loop start timer in millsecond
func Loop(timerID int32, duration int64, callback func()) int32 {
	mutex.Lock()
	defer mutex.Unlock()
	if timerID <= 0 {
		timerID = genTimerID()
	} else {
		if timerHandler := getTimer(timerID); timerHandler != nil {
			timerHandler <- struct{}{}
			delTimer(timerID)
		}
	}
	ch := gtimer.Forever(time.Duration(duration*1000*1000), callback)
	addTimer(timerID, ch)
	return timerID
}

// Loop with lock start timer in millsecond
func LoopLock(timerID int32, duration int64, callback func(), lock *sync.Mutex) int32 {
	mutex.Lock()
	defer mutex.Unlock()
	if timerID <= 0 {
		timerID = genTimerID()
	} else {
		if timerHandler := getTimer(timerID); timerHandler != nil {
			timerHandler <- struct{}{}
			delTimer(timerID)
		}
	}
	ch := gtimer.ForeverLocked(time.Duration(duration*1000*1000), callback, lock)
	addTimer(timerID, ch)
	return timerID
}

//Cancel stop timer by timer id
func Cancel(timerID int32) bool {
	mutex.Lock()
	defer mutex.Unlock()
	if timerHandler := getTimer(timerID); timerHandler == nil {
		return false
	} else {
		timerHandler <- struct{}{}
		delTimer(timerID)
		return true
	}
}

func genTimerID() int32 {
	return AtomicAdd(&genID)
}

func addTimer(timerID int32, ch chan struct{}) {
	timerMap.Store(timerID, ch)
}

func getTimer(timerID int32) chan struct{} {
	timerHandler, ok := timerMap.Load(timerID)
	if !ok {
		return nil
	}
	return timerHandler.(chan struct{})
}

func delTimer(timerID int32) {
	timerMap.Delete(timerID)
}
