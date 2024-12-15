package protobq

import (
	"time"
)

type MaterializedViewOptions struct {
	EnableRefresh   bool
	RefreshInterval time.Duration

	// TODO: クラスタ化列、パーティション列を定義する。
}

type MaterializedView interface {
	Name() string
	Options() MaterializedViewOptions
	InsertDTO() InsertDTO
}

type InsertDTO interface {
	TableName() string
	Value() map[string]any
}
