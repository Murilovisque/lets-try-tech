package platform

type ProcessCounter struct {
	amountOfProcessChannel   chan struct{}
	shutdownMode             bool
	errLimitReached          error
	errShutdownModeActivated error
}

func (p *ProcessCounter) Setup(limit int, errorLimitReached, errorShutdownModeActivated error) {
	p.amountOfProcessChannel = make(chan struct{}, limit)
	p.errLimitReached = errorLimitReached
	p.errShutdownModeActivated = errorShutdownModeActivated
	p.shutdownMode = false
}

func (p *ProcessCounter) IncrementProcess() error {
	if p.shutdownMode {
		return p.errShutdownModeActivated
	}
	select {
	case p.amountOfProcessChannel <- struct{}{}:
		return nil
	default:
		return p.errLimitReached
	}
}

func (p *ProcessCounter) DecrementProcess() {
	<-p.amountOfProcessChannel
}

func (p *ProcessCounter) Shutdown() {
	p.shutdownMode = true
	close(p.amountOfProcessChannel)
	p.waitForAllProcessesToComplete()
}

func (p *ProcessCounter) IsUp() bool {
	return !p.shutdownMode
}

func (p *ProcessCounter) waitForAllProcessesToComplete() {
	for {
		if len(p.amountOfProcessChannel) == 0 {
			break
		}
	}
}
