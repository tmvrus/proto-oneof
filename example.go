package proto

func init() {
	_ = oneOf
	_ = withCode
}

func oneOf() {
	_ = Event{
		Event: &Event_EventA{
			EventA: &EventA{
				A: "This is A",
			},
		},
	}
}

func withCode() {
	_ = EventWithCode{
		Type: EventCode_A,
		EventA: &EventWithCodeA{
			A: "This is A",
		},
	}
}
