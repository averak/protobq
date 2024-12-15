package internal

type InsertDTOImpl struct {
	tableName string
	fields    []BQField
}

func NewInsertDTOImpl(tableName string, fields []BQField) *InsertDTOImpl {
	return &InsertDTOImpl{
		tableName: tableName,
		fields:    fields,
	}
}

func (r InsertDTOImpl) TableName() string {
	return r.tableName
}

func (r InsertDTOImpl) Value() map[string]any {
	res := make(map[string]any)
	for _, f := range r.fields {
		currentMap := res
		for i, key := range f.Path {
			if i == len(f.Path)-1 {
				currentMap[key] = f.Value
			} else {
				if _, ok := currentMap[key]; !ok {
					currentMap[key] = make(map[string]any)
				}
				currentMap = currentMap[key].(map[string]any) //nolint:forcetypeassert
			}
		}
	}
	return res
}

func (r *InsertDTOImpl) AddField(f BQField) {
	r.fields = append(r.fields, f)
}

type BQField struct {
	Path  []string
	Value any
}
