package systemEvents

type EventListener interface {
	EventOccured(ev *EventData)
}

type EventGenerator interface {
	GetName() string
	Attach(ev EventListener)
}

type EventData struct {
	Name string
}

type EventWatcher interface {
	EventListener
	Register(eg *EventGenerator)
}

// Implements EventWatcher as well as EventListener
type  EventWatcherImpl struct{}

func (ewi *EventWatcherImpl) Register(eg *EventGenerator) {
	eg.Attach(ewi)
}

func (ewi *EventWatcherImpl) EventOccured(ev *EventData) {
	switch (ev.Name)
	// scheduler: 
	// call other handlers
}

type usbEvents struct{}

int main() {
	ew := NewEventWather()
	se1 := NewUsbEvents()
	se2 := NewNetworkEvents()
	ew.Register(se1)
	ew.Register(se2)
}