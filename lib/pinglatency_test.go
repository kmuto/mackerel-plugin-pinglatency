package pinglatency

import (
	"testing"
)

func TestFetchMetrics(t *testing.T) {
	sample := &Plugin{}

	_, err := sample.FetchMetrics()

	if err != nil {
		t.Errorf("FetchMetrics returns error")
	}
}
