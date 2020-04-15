package events_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/events"
	"github.com/SevereCloud/vksdk/object"
	"github.com/stretchr/testify/assert"
)

func TestMessageNewObject_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	f := func(b []byte, want events.MessageNewObject, wantErr bool) {
		var obj events.MessageNewObject

		err := obj.UnmarshalJSON(b)
		if (err != nil) != wantErr {
			t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, wantErr)
		}

		assert.Equal(t, obj, want)
	}

	f(
		[]byte(""),
		events.MessageNewObject{},
		true,
	)
	f(
		[]byte(`{"id":1}`),
		events.MessageNewObject{Message: object.MessagesMessage{ID: 1}},
		false,
	)
	f(
		[]byte(`{"message":{"id":1},"client_info":{}}`),
		events.MessageNewObject{Message: object.MessagesMessage{ID: 1}},
		false,
	)
}
