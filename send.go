package proto

import "github.com/golang/protobuf/proto"

func sendWithCode(t EventCode, a string, b, x, y int64) {
	var c proto.Message
	if t == EventCode_C {
		c = &EventWithCode{
			Type: EventCode_C,
			EventC: &EventWithCodeC{
				A: a,
				B: b,
				C: &EventWithCodeC_Context{
					X: x,
					Y: y,
				},
			},
		}
	}

	// some manual validate with EventC:
	// type is corresponds to EventX

	send(c)
}

func sendOneOf(t string, a string, b, x, y int64) {
	c := &Event{}
	if t == "C" {
		c.Event = &Event_EventC{
			EventC: &EventC{
				A: a,
				B: b,
				C: &EventC_Context{
					X: x,
					Y: y,
				},
			},
		}
	}

	if err := c.Validate(); err != nil {
		panic(err.Error())
	}

	send(c)
}

func send(msg proto.Message) {}
