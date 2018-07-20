package blackboard

// BB ...
type BB struct {
	value map[string]interface{}
}

// NewBlackboard ...
func NewBlackboard() *BB {
	return &BB{make(map[string]interface{}, 0)}
}

var (
	blackboard = NewBlackboard()
)

// Singleton ...
func Singleton() *BB {
	return blackboard
}

// SetValue ...
func (bb *BB) SetValue(key string, value interface{}) {
	bb.value[key] = value
}

// SetValue ...
func SetValue(key string, value interface{}) {
	blackboard.SetValue(key, value)
}

// Value ...
func (bb *BB) Value(key string) (v interface{}, ok bool) {
	v, ok = bb.value[key]
	return
}

// Value ...
func Value(key string) (interface{}, bool) {
	return blackboard.Value(key)
}
