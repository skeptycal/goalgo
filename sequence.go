package goalgo

type (

	// Sequence maintains a sequence of items (order is extrinsic)
	//
	// Set implements the Set interface from
	// MIT 6.006 lecture 3: Sets and Sorting (Spring 2020)
	//
	// • Ex: (x0, x1, x2, . . . , xn-1) (zero indexing)
	//
	// (n is typically used to denote the number of items stored in the data structure)
	Sequence interface {
		Container
		StaticSeq
		DynamicSeq
	}

	// StaticSeq implements the static components of Sequence
	StaticSeq interface {
		IterSeq() (values []Any)                // return the stored items one-by-one in sequence order
		GetAt(index int) (value Any, err error) // return the item at index i
		SetAt(index int, value Any) error       // set the item at index i to value
	}

	// DynamicSeq implements the dynamic components of Sequence
	DynamicSeq interface {
		AtSeq
		FirstSeq
		LastSeq
	}

	AtSeq interface {
		InsertAt(index int, value Any) error       // insert the item at index
		DeleteAt(index int) (value Any, err error) // remove and return the item at index
	}
	FirstSeq interface {
		InsertFirst(value Any) error         // insert the first item in sequence
		DeleteFirst() (value Any, err error) // remove and return the first item in sequence
	}

	LastSeq interface {
		InsertLast(value Any) error         // insert the last item in sequence
		DeleteLast() (value Any, err error) // remove and return the last item in sequence
	}

	// Stack is a special case of Sequence
	// A Linked List held by the tail ...
	Stack interface {
		Container
		LastSeq
	}

	// Queue is a special case of Sequence
	Queue interface {
		Container
		InsertLast(value Any) error          // insert the last item in sequence
		DeleteFirst() (value Any, err error) // remove and return the first item in sequence
	}

	// LinkedList can insert or delete from the front in O(1),
	// ... but Get/Set take O(n)
	//
	// Perhaps best for sorting or manipulating lists of
	// items that will only be read back once at the end
	// of a long set of processing steps.
	//
	// Text analysis? HTML parsing? Security analysis?
	//
	// Possibly image processing? Then convert to Array
	// to do block modifications, save, change formats, etc.
	LinkedList interface {
		Container
		FirstSeq
	}

	// Array is O(1) for static operations (Get, Set, Iter)
	// ... but O(1) for all other operations.
	//
	// Probably the best for lookup tables and references
	// that are only generated once per session and are
	// locked afterwards.
	//
	// Possibly useful for static cache of values, workers, etc.
	Array interface {
		Container
		StaticSeq
	}

	// DynamicArray is an Array with an extra allocation
	// for dynamic operations.
	//
	// Dynamic arrays only support dynamic "last" operations
	// in Θ(1) time
	//
	// Periodic O(n) reallocation to increase size is needed.
	//
	// Inserting into a dynamic array takes O(1) "amortized"
	// time when averaged in with the reallocations used.
	//
	// Deleting "last" from a dynamic array is O(1) but can be
	// wasteful of space. Also requires some reallocation.
	//
	// The Python list data type is a dynamic array.
	DynamicArray interface {
		Container
		Array
		Stack
	}
)
