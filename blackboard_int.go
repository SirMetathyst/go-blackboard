package blackboard

// KI ...
type KI struct {
	Key   string
	Value *int
}

// SetIntP ...
func (bb *BB) SetIntP(key string, value *int) {
	bb.SetValue(key, value)
}

// SetIntP ...
func SetIntP(key string, value *int) {
	blackboard.SetIntP(key, value)
}

// SetInt ...
func (bb *BB) SetInt(key string, value int) {
	bb.SetValue(key, &value)
}

// SetInt ...
func SetInt(key string, value int) {
	blackboard.SetInt(key, value)
}

// IntP ...
func (bb *BB) IntP(key string) *int {
	i, kok := bb.Value(key)
	if !kok {
		return nil
	}
	value, ok := i.(*int)
	if !ok {
		return nil
	}
	return value
}

// IntP ...
func IntP(key string) *int {
	return blackboard.IntP(key)
}

// AllInt ...
func (bb *BB) AllInt() []KI {
	slice := make([]KI, 0)
	for k := range bb.value {
		if iv := bb.IntP(k); iv != nil {
			slice = append(slice, KI{k, iv})
		}
	}
	return slice
}

// AllInt ...
func AllInt() []KI {
	return blackboard.AllInt()
}
