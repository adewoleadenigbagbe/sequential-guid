package sequentialguid

import (
	"testing"

	"github.com/google/uuid"
)

const UUID_SIZE int = 36

func TestNew(t *testing.T) {

	randomId := uuid.New()
	got := SequentialGuid(randomId).New()

	if len(got.String()) != UUID_SIZE {
		t.Errorf("got %d, wanted %d", len(got), UUID_SIZE)
	}
}
