package goalgo

type (

	// Sequence maintains a sequence of items (order is extrinsic)
	//
	// Set implements the Set interface from
	// MIT 6.006 lecture 3: Sets and Sorting
	//
	// • Sequence about extrinsic order, set is about intrinsic order
	//
	// • Maintain a set of items having unique keys (e.g., item x has key x.key)
	//
	// • (Set or multi-set? We restrict to unique keys for now.)
	//
	// • Often we let key of an item be the item itself, but may want to store more info than just key
	Set interface {
		Dictionary
		OrderSet

		// GetSetter

	}

	StaticSet interface {
		Find(k Any) (v Any, ok bool) // return the stored item with key k
	}

	DynamicSet interface {
		Delete(k Any) (v Any, err error) // remove and return the item with key k
		Insert(k, v Any) (err error)     // insert the item with key k
	}

	OrderSet interface {
		IterOrd() []Any // IterOrd returns the stored items one by one in key order.
		Min() Any       // Min returns the item with the smallest key.
		Max() Any       // Max returns the item with the largest key.
		Next(k Any) Any // Prev returns the item with the smallest key larger than k.
		Prev(k Any) Any // Next returns the item with the largest key smaller than k.
	}

	// Dictionary is a special case of Set without the Order operations
	Dictionary interface {
		Container
		StaticSet
		DynamicSet
	}
)
