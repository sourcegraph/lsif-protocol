package reader

import (
	"strconv"
	"sync"
	"sync/atomic"
)

// Interner converts strings into unique identifers. Submitting the same byte value to
// the interner will result in the same identifier being produced. Each unique input is
// guaranteed to have a unique output (no two inputs share the same identifier). The
// identifier space of two distinct interner instances may overlap.
//
// Assumption: The output of LSIF indexers will not generally mix types of identifiers.
// If integers are used, they are used for all ids. If strings are used, they are used
// for all ids.
type Interner struct {
	sync.RWMutex
	m map[string]int
}

// NewInterner creates a new empty interner.
func NewInterner() *Interner {
	return &Interner{
		m: map[string]int{},
	}
}

// Intern returns the unique identifier for the given byte value. The byte value should
// be a raw LSIF input identifier, which should be a JSON-encoded number or quoted string.
// This method is safe to call from multiple goroutines.
func (i *Interner) Intern(raw []byte) (int, error) {
	if len(raw) == 0 {
		// No identifier supplied
		return 0, nil
	}

	if raw[0] != '"' {
		// Not a string, expect a number
		return strconv.Atoi(string(raw))
	}

	// Generate a numeric identifier for the de-quoted string
	s := string(raw[1 : len(raw)-1])

	// See if this is an "inty" string (e.g., "1234"). We can use a
	// fast-path here that does not need to lock or stash the string
	// value in a map.
	if v, err := strconv.Atoi(s); err == nil {
		return v, nil
	}

	i.RLock()
	v, ok := i.m[s]
	i.RUnlock()
	if ok {
		return v, nil
	}

	i.Lock()
	defer i.Unlock()

	v, ok = i.m[s]
	if !ok {
		// Generate and stash a new identifier
		v = len(i.m) + 1
		i.m[s] = v
	}

	return v, nil
}

// progressSentinel is used to signal to goroutines that may be attempting to intern a
// currently-being-interned that another goroutine is in the process of interning for a given key
// and to spin waiting for interning to complete and to then use that value. This is necessary as
// we use the length of the underlying structure (here stored in the `count` field) to generate a
// unique identifier for a given key.
const progressSentinel = -1

type LockfreeInterner struct {
	m     sync.Map
	count int64
}

// NewInterner creates a new empty interner.
func NewLockfreeInterner() *LockfreeInterner {
	return &LockfreeInterner{
		m: sync.Map{},
	}
}

// Intern returns the unique identifier for the given byte value. The byte value should
// be a raw LSIF input identifier, which should be a JSON-encoded number or quoted string.
// This method is safe to call from multiple goroutines.
func (i *LockfreeInterner) Intern(raw []byte) (int, error) {
	if len(raw) == 0 {
		// No identifier supplied
		return 0, nil
	}

	if raw[0] != '"' {
		// Not a string, expect a number
		return strconv.Atoi(string(raw))
	}

	// Generate a numeric identifier for the de-quoted string
	s := string(raw[1 : len(raw)-1])

	// See if this is an "inty" string (e.g., "1234"). We can use a
	// fast-path here that does not need to lock or stash the string
	// value in a map.
	if v, err := strconv.Atoi(s); err == nil {
		return v, nil
	}

	// attempt to signal the start of interning a value, checking if
	// the given key is already in the process of being interned
	v, ok := i.m.LoadOrStore(s, progressSentinel)
	if ok {
		// busy-loop waiting for store to complete. Can we avoid this?
		// sync.Cond perhaps?
		for v == progressSentinel {
			v, _ = i.m.Load(s)
		}
		return int(v.(int64)), nil
	}
	// if we're here, we are the winner for grabbing the conceptual "lock"
	r := atomic.AddInt64(&i.count, 1)
	i.m.Store(s, r)

	return int(r), nil
}
