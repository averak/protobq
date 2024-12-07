package protobq

import "time"

type MaterializedViewOptions struct {
	EnableRefresh   bool
	RefreshInterval time.Duration
}

type MaterializedView interface {
	Options() MaterializedViewOptions
}
