package wsp

// Get migration models
func DBModels() []any {
	return []interface{}{
		&Message{},
	}
}
