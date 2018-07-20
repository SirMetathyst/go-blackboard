package blackboard

// KSS ...
type KSS struct {
	Key   string
	Value *[]string
}

// SetStringSliceP ...
func (bb *BB) SetStringSliceP(key string, value *[]string) {
	bb.SetValue(key, value)
}

// SetStringSliceP ...
func SetStringSliceP(key string, value *[]string) {
	blackboard.SetStringSliceP(key, value)
}

// SetStringSlice ...
func (bb *BB) SetStringSlice(key string, value []string) {
	bb.SetValue(key, &value)
}

// SetStringSlice ...
func SetStringSlice(key string, value []string) {
	blackboard.SetStringSlice(key, value)
}

// StringSliceP ...
func (bb *BB) StringSliceP(key string) *[]string {
	i, ok := bb.Value(key)
	if !ok {
		return nil
	}
	return i.(*[]string)
}

// StringSliceP ...
func StringSliceP(key string) *[]string {
	return blackboard.StringSliceP(key)
}

// AllStringSlice ...
func (bb *BB) AllStringSlice() []KSS {
	slice := make([]KSS, 0)
	for k, v := range bb.value {
		if ssv, ok := v.(*[]string); ok {
			slice = append(slice, KSS{k, ssv})
		}
	}
	return slice
}

// AllStringSlice ...
func AllStringSlice() []KSS {
	return blackboard.AllStringSlice()
}
