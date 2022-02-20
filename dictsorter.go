package goalgo

import (
	"reflect"

	"github.com/skeptycal/types"
)

// Dict is a dictionary used to store keys and values of any
// type that is comparable and sortable.
// It is not concurrent-safe.
type dictSorter struct {
	name        string
	protected   bool
	sortEnabled bool
	m           AnyMap
	kKind       reflect.Kind
	vKind       reflect.Kind

	keysort   AnySlice
	valuesort AnySlice
}

// DictSorter is a dictionary for types that are sortable and
// and implements the standard library sort.Interface methods.
type DictSorter interface {
	Dict
	types.Sorter
}

// Enable activates sorting of keys and values
func (d *dictSorter) Enable() { d.sortEnabled = true }

// Disable deactivates sorting of keys and values
func (d *dictSorter) Disable() { d.sortEnabled = false }
