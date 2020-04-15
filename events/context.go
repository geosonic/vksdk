package events // import "github.com/SevereCloud/vksdk/events"

import (
	"context"

	"github.com/SevereCloud/vksdk/internal"
)

// GroupIDFromContext returns the GroupID from context
func GroupIDFromContext(ctx context.Context) int {
	return ctx.Value(internal.GroupIDKey).(int)
}

// EventIDFromContext returns the EventID from context
func EventIDFromContext(ctx context.Context) string {
	return ctx.Value(internal.EventIDKey).(string)
}
