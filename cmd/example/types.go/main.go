package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"unsafe"

	"github.com/skeptycal/goalgo"
	"github.com/skeptycal/types"
	// . "github.com/skeptycal/types"
)

const (
	DefaultColumnSpacing = 20
	DefaultRowSpacing    = 5
)

type Any = goalgo.Any

type table struct {
	columnCount     int
	rowCounter      int
	columnSpacing   int
	rowSpacing      int
	rowFormatString string

	Data goalgo.Sequence
}

var (
	NL          = types.NL
	NewAnyValue = types.NewAnyValue
)

func NewTable(columns int) *table {
	t := new(table)
	t.columnCount = columns
	t.columnSpacing = DefaultColumnSpacing
	t.rowSpacing = DefaultRowSpacing
	t.rowCounter = 0

	t.setFormatString()
	return t
}

type Iterable interface {
	First() Any
	Next() Any
	Len() int
}

type iterSlice struct {
	counter  int
	isSorted bool
	slc      []Any
}

func (i *iterSlice) Len() int { return len(i.slc) }

func (i *iterSlice) Swap(ii, jj int) {
	i.slc[ii], i.slc[jj] = i.slc[jj], i.slc[ii]
}

func (i *iterSlice) Less(ii, jj int) bool {
	return AnyLess(ii, jj)
}

func AnyLess(i, j Any) bool {
	vi := NewAnyValue(i)
	vj := NewAnyValue(j)

	if vi.IsOrdered() {
		switch tj := vj.Interface().(type) {
		case int:
			switch ti := vi.Interface().(type) {
			case int:
				return ti < tj
			}
		case string:
			switch ti := vi.Interface().(type) {
			case string:
				return ti < tj
			}
		case uint:
			switch ti := vi.Interface().(type) {
			case uint:
				return ti < tj
			}
		case float64:
			switch ti := vi.Interface().(type) {
			case float64:
				return ti < tj
			}
		case float32:
			switch ti := vi.Interface().(type) {
			case float32:
				return ti < tj
			}
		default:
			return false
		}
	}
	return false
}

func (i *iterSlice) First() Any {
	if !sort.IsSorted(i) {
		sort.Slice(i.slc, i.Less)
	}
	i.counter = 0
	return i.Next()
}
func (i *iterSlice) Next() Any {
	defer i.inc()
	return i.slc[i.counter]
}
func (i *iterSlice) inc() {
	i.counter += 1
}

type iterMap struct {
	currentKey Any
	slc        Iterable
	isSorted   bool
	m          map[Any]Any
}

func (i *iterMap) Len() int { return len(i.m) }

// First returns the first element in the sequence.
//
// Calls to first will reset the internal counter
// for Next() calls to zero.
func (i *iterMap) First() Any {
	if i.slc.Len() < 1 {
		i.Build()
	}
	return i.m[i.slc.First()]
}

// Next returns the next element in the sequence.
func (i *iterMap) Next() Any {
	return i.m[i.slc.Next()]
}

// Build creates a slice of keys that are indexed by
// the counter in an iterSlice object.
func (i *iterMap) Build() {
	keys := make([]Any, 0, i.Len())

	for k := range i.m {
		keys = append(keys, k)
	}

	// TODO: keys could be sorted here ...
	// keys are sorted in the slice ...

	i.slc = NewIterable(keys, i.isSorted)
}

func NewIterable(a Any, sorted bool) Iterable {

	v := NewAnyValue(a)

	switch v.Kind() {
	case reflect.Map:
		return &iterMap{m: v.Interface().(map[Any]Any), isSorted: sorted}
	case reflect.Slice:
		return &iterSlice{slc: v.Interface().([]Any), isSorted: sorted}
	case reflect.String:
		return &iterSlice{slc: v.Interface().([]Any), isSorted: sorted}
	default:
		return nil
	}

}

func (t *table) String() string {
	var args []interface{}
	// for _, v := range t.Rows() {
	// 	if IsIterable(v) {
	// 		while Iterable(v).Next() != nil {

	// 			s := fmt.Sprint(v)
	// 			if len(s) > DefaultColumnSpacing {
	// 				s = s[:DefaultColumnSpacing-3] + "..."
	// 			}
	// 			args = append(args, s)
	// 		}
	// 	}
	//
	// }
	return fmt.Sprintf(t.rowFormatString, args...)
}

func (t *table) Rows(a ...interface{}) []Any {
	return t.Data.IterSeq()
}

var columnspacingString = fmt.Sprintf("%%-%d.%dv ", DefaultColumnSpacing, DefaultColumnSpacing)

func tablePrint(a ...interface{}) (n int, err error) {
	fmtString := strings.Repeat(columnspacingString, len(a)) + NL
	var args []interface{}
	for _, v := range a {
		s := fmt.Sprint(v)
		if len(s) > DefaultColumnSpacing {
			s = s[:DefaultColumnSpacing-3] + "..."
		}
		args = append(args, s)
	}
	return fmt.Printf(fmtString, args...)
}

func main() {

	tablePrint("Name:", "Value:", "Type:", "Kind: ")
	for _, tt := range reflectTests {

		A := NewAnyValue(tt.a)
		tablePrint(tt.name, A.ValueOf(), A.TypeOf(), A.Kind())
	}
}

func (t *table) setFormatString() {
	t.rowFormatString = strings.Repeat(columnspacingString, t.columnCount) + NL
}

func (t *table) makeFormatSubString() string {
	return fmt.Sprintf("%%-%d.%dv ", t.columnSpacing, t.columnSpacing)
}

var iii = 42
var ptrSample = &iii

var reflectTests = []struct {
	name string
	a    Any
	want reflect.Kind
}{
	{"invalid", nil, reflect.Invalid},
	{"bool", true, reflect.Bool},
	{"int", 42, reflect.Int},
	{"uint", uint(42), reflect.Uint},
	{"int", 42, reflect.Int},
	{"Int8", int8(42), reflect.Int8},
	{"Int16", int16(42), reflect.Int16},
	{"Int32", int32(42), reflect.Int32},
	{"Int64", int64(42), reflect.Int64},
	{"Uint", uint(42), reflect.Uint},
	{"Uint8", uint8(42), reflect.Uint8},
	{"Uint16", uint16(42), reflect.Uint16},
	{"Uint32", uint32(42), reflect.Uint32},
	{"Uint64", uint64(42), reflect.Uint64},
	{"Uintptr", uintptr(42), reflect.Uintptr},
	{"Float32", float32(42), reflect.Float32},
	{"Float64", float64(42), reflect.Float64},
	{"Complex64", complex64(42), reflect.Complex64},
	{"Complex128", complex128(42), reflect.Complex128},
	{"Array", [4]int{42, 42, 42, 42}, reflect.Array},
	{"Chan", make(chan int, 1), reflect.Chan},
	{"Func", types.IsComparable, reflect.Func},
	{"Map", make(map[string]interface{}), reflect.Map},
	{"Ptr", ptrSample, reflect.Ptr},
	{"Slice", []int{42}, reflect.Slice},
	{"String", "42", reflect.String},
	{"UnsafePointer", unsafe.Pointer(nil), reflect.UnsafePointer},
	// {"Interface", nil, reflect.Interface},
	// {"ValueOf(42)", ValueOf(42), ValueOf(ValueOf(42)).Kind()},
}
