package observers

type Observer interface {
	Handle(event EventObserver)
}

type Dispatcher struct {
	observers map[EventObserverType][]Observer
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		observers: make(map[EventObserverType][]Observer),
	}
}

func (d *Dispatcher) Register(eventType EventObserverType, observer Observer) {
	d.observers[eventType] = append(d.observers[eventType], observer)
}

func (d *Dispatcher) Dispatch(event EventObserver) {
	if observers, found := d.observers[event.Type]; found {
		for _, observer := range observers {
			go observer.Handle(event)
		}
	}
}
