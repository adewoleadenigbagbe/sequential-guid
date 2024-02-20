package sequentialguid

import (
	"testing"
)

const UUID_SIZE int = 36

func TestNew(t *testing.T) {

	got := NewSequentialGuid()

	if len(got.String()) != UUID_SIZE {
		t.Errorf("got %d, wanted %d", len(got), UUID_SIZE)
	}
}
