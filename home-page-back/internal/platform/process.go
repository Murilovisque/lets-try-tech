package platform

import "sync/atomic"

type ProcessCounter struct {
	currentProcess int64
}

func (p *ProcessCounter) IncrementProcess() {
	atomic.AddInt64(&p.currentProcess, 1)
}

func (p *ProcessCounter) DecrementProcess() {
	atomic.AddInt64(&p.currentProcess, -1)
}

func (p *ProcessCounter) WaitForAllProcessesToComplete() {
	for {
		if atomic.LoadInt64(&p.currentProcess) == 0 {
			break
		}
	}
}
