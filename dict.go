package goalgo

import (
	"fmt"
	"sort"
	"strings"

	"github.com/skeptycal/types"
)

const (
	KVFormatString      = "%15v = %15v"
	msgMapKeysProtected = "key changes not allowed in protected map: %v"
	msgKeyNotFound      = "key not found: %v"
)

// Dict is a dictionary used to store keys and values of any type.
// It is not concurrent-safe.
type Dict interface {
	Delete(key Any) error
	IsSortable() bool // always false for basic dict

	types.GetSetter
	types.Protector
	types.Slicer
	types.Stringer
}

func errMapKeysProtected(a Any) error {
	return fmt.Errorf(msgMapKeysProtected, a)
}

func errKeyNotFound(a Any) error {
	return fmt.Errorf(msgKeyNotFound, a)
}

type (
	// AnyMap is a map used to store keys and values of any type.
	AnyMap map[Any]Any

	AnySlice interface {
		sort.Interface
	}

	stringSlice struct {
		list []string
	}

	dict struct {
		name        string
		protected   bool
		sortEnabled bool
		m           AnyMap
	}
)

func (s *stringSlice) Len() int           { return len(s.list) }
func (s *stringSlice) Swap(i, j int)      { s.list[i], s.list[j] = s.list[j], s.list[i] }
func (s *stringSlice) Less(i, j int) bool { return s.list[i] < s.list[j] }

func NewDict(name string, protected bool) Dict {

	// v := ValueOf()

	return &dict{
		name:      name,
		protected: protected,
		m:         AnyMap{},
	}
}

// Len returns the number of elements in the
// underlying map.
// Len is part of sort.Interface.
func (d *dict) Len() int {
	return len(d.m)
}

func (d *dict) IsSortable() bool { return false }

// Swap exchanges the values associated with
// the keys in the underlying map.
// Swap is part of sort.Interface.
func (d *dict) Swap(i, j int) {
	// d.m[i], d.m[j] = d.m[j], d.m[i]
}

// Less compares the values associated with
// the values in the underlying map.
func (d *dict) Less(i, j int) bool {

	if i == j {
		return false
	}

	vi := d.m[i]
	vj := d.m[j]

	if vi == vj {
		return false
	}

	switch ti := vi.(type) {
	case int:
		if tj, ok := vj.(int); ok {
			return ti < tj
		}
	case uint:
		if tj, ok := vj.(uint); ok {
			return ti < tj
		}
	case float32:
		if tj, ok := vj.(float32); ok {
			return ti < tj
		}
	case float64:
		if tj, ok := vj.(float64); ok {
			return ti < tj
		}
	case string:
		if tj, ok := vj.(string); ok {
			return ti < tj
		}
	// case map:

	default:
		return false
	}
	return false
}

// func compare[T constraints.Ordered](a, b T) {
//     // works
//     if a > b {
//         fmt.Printf("%v is bigger than %v", a, b)
//     }
// }

func (d *dict) Keys() []Any {
	return d.keysSet_AllocWithCap()
}

// Keys returns a slice of keys of the underlying
// map. If sorting is enabled with d.Enable(),
// the keys will be sorted.
func (d *dict) keysAppend_DeferAllocationWithCap() []Any {
	if len(d.m) == 0 {
		return []Any{}
	}
	if len(d.m) == 1 {
		return []Any{d.m[0]}
	}

	keys := make([]Any, 0, len(d.m))
	for k := range d.m {
		keys = append(keys, k)
	}

	return keys
}

func (d *dict) keysAppend_PreAllocateRedoWithCap() (keys []Any) {
	if len(d.m) == 0 {
		return
	}
	if len(d.m) == 1 {
		keys = append(keys, d.m[0])
		return
	}

	keys = make([]Any, 0, len(d.m))
	for k := range d.m {
		keys = append(keys, k)
	}

	return keys
}

func (d *dict) keysAppend_NoChecks_PreAllocate() (keys []Any) {
	for k := range d.m {
		keys = append(keys, k)
	}
	return keys
}

func (d *dict) keysAppend_AllocateZero() []Any {
	if len(d.m) == 0 {
		return []Any{}
	}
	if len(d.m) == 1 {
		return []Any{d.m[0]}
	}

	keys := make([]Any, 0)
	for k := range d.m {
		keys = append(keys, k)
	}

	return keys
}

func (d *dict) keysSet_AllocWithCap() []Any {
	if len(d.m) == 0 {
		return []Any{}
	}

	if len(d.m) == 1 {
		return []Any{d.m[0]}
	}

	keys := make([]Any, 0, len(d.m))

	i := 0
	for k := range d.m {
		keys[i] = k
		i++
	}

	return keys
}

func (d *dict) keysSet_AllocZero() []Any {
	if len(d.m) == 0 {
		return []Any{}
	}

	if len(d.m) == 1 {
		return []Any{d.m[0]}
	}

	keys := make([]Any, 0)

	i := 0
	for k := range d.m {
		keys[i] = k
		i++
	}

	return keys
}

func (d *dict) keysSet_NoChecks_AllocCap() []Any {
	keys := make([]Any, 0, len(d.m))

	i := 0
	for k := range d.m {
		keys[i] = k
		i++
	}

	return keys
}

func (d *dict) keysSet_NoChecks_AllocZero() []Any {
	keys := make([]Any, 0)

	i := 0
	for k := range d.m {
		keys[i] = k
		i++
	}

	return keys
}
func (d *dict) keysSet_NoChecks_PreAlloc() (keys []Any) {

	i := 0
	for k := range d.m {
		keys[i] = k
		i++
	}

	return keys
}

// Values returns a slice of values of the underlying
// map. If sorting is enabled with d.Enable(),
// the values will be sorted.
func (d *dict) Values() (values []Any) {
	if len(d.m) == 0 {
		return
	}
	if len(d.m) == 1 {
		values = append(values, d.m[0])
		return
	}
	values = make([]Any, 0, len(d.m))
	for _, v := range d.m {
		values = append(values, v)
	}
	return values
}

func (d *dict) Get(key Any) (Any, error) {
	if v, ok := d.m[key]; ok {
		return v, nil
	}
	return nil, errKeyNotFound(key)
}

func (d *dict) Set(key, value Any) error {
	if key == nil {
		return errKeyNotFound(key)
	}

	if _, ok := d.m[key]; ok && d.protected {
		return errMapKeysProtected(key)
	}

	d.m[key] = value
	return nil
}

func (d *dict) Delete(key Any) error {
	if !d.protected {
		delete(d.m, d.m[key])
		return nil
	}
	return errMapKeysProtected(key)
}

// Protect prevents dictionary keys from being modified.
func (d *dict) Protect() { d.protected = true }

// Unprotect allows dictionary keys to be modified.
func (d *dict) Unprotect() { d.protected = false }

func kvHelper(sb *strings.Builder, key, value Any) {
	sb.WriteString(fmt.Sprintf(KVFormatString, "key", "value"))
}

func (d *dict) String() string {
	sb := &strings.Builder{}
	defer sb.Reset()

	kvHelper(sb, "key", "value")

	for k, v := range d.m {
		kvHelper(sb, k, v)
	}
	return sb.String()
}

// Enable activates sorting of keys and values
func (d *dict) Enable() { d.sortEnabled = true }

// Disable deactivates sorting of keys and values
func (d *dict) Disable() { d.sortEnabled = false }
