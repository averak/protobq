package internal

import (
	"reflect"
	"testing"
)

func TestInsertDTOImpl_Value(t *testing.T) {
	type fields struct {
		tableName string
		fields    []BQField
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]any
	}{
		{
			fields: fields{
				fields: []BQField{
					{
						Path:  []string{"a", "b", "c"},
						Value: 1,
					},
					{
						Path:  []string{"a", "d"},
						Value: 2,
					},
					{
						Path:  []string{"a", "e", "f"},
						Value: 3,
					},
				},
			},
			want: map[string]any{
				"a": map[string]any{
					"b": map[string]any{
						"c": 1,
					},
					"d": 2,
					"e": map[string]any{
						"f": 3,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := InsertDTOImpl{
				tableName: tt.fields.tableName,
				fields:    tt.fields.fields,
			}
			if got := r.Value(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
