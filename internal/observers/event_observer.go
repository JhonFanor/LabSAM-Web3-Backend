package observers

type EventObserverType string

const (
	Created  EventObserverType = "created"
	Updated  EventObserverType = "updated"
	Approved EventObserverType = "approved"
	Rejected EventObserverType = "rejected"
)

type EventObserver struct {
	Type         EventObserverType
	SenderID     uint
	ReceiverID   uint
	ResourceID   int
	ResourceType string
	Message      string
	Action       string
}
