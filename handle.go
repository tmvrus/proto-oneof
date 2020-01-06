package proto

import "github.com/golang/protobuf/proto"

func handleWithCode(data []byte) {
	c := EventWithCode{}
	if err := proto.Unmarshal(data, &c); err != nil {
		panic(err.Error())
	}

	if c.Type == EventCode_C {
		// some manual validate with EventC:
		// type is corresponds to EventC and not nill
		if c.EventC == nil {
			panic("got nil :(")
		}

		_ = c.EventC.A
		_ = c.EventC.B
		_ = c.EventC.C.X
		_ = c.EventC.C.Y
	}
}

func handleOneOf(data []byte) {
	c := Event{}

	if err := proto.Unmarshal(data, &c); err != nil {
		panic(err.Error())
	}

	if err := c.Validate(); err != nil {
		panic(err.Error())
	}

	switch event := c.Event.(type) {
	case *Event_EventC:
		_ = event.EventC.A
		_ = event.EventC.B
		_ = event.EventC.C.X
		_ = event.EventC.C.Y

	}

}
