package metrics

import "testing"

func TestLocalStoreDrain(t *testing.T) {
	original := DefaultDrain
	DefaultDrain = NewLocalStoreDrain()
	defer func() {
		DefaultDrain = original
	}()

	store := DefaultDrain.(*LocalStoreDrain).store
	const key = "user.signup"

	// increment our key twice
	for i := 0; i < 2; i++ {
		Count(key, 1)
	}
	if len(store[key]) != 1 {
		t.Error(
			"For", key,
			"expected length", 1,
			"got", len(store[key]),
		)
	}

	// incrementing our key by a value other than 1 should generate a new entry
	// in the key map
	Count(key, 2)
	if len(store[key]) != 2 {
		t.Error(
			"For", key,
			"expected length", 1,
			"got", len(store[key]),
		)
	}
}
