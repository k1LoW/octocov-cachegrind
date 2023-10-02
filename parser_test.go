package cachegrind

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/k1LoW/octocov/report"
)

func TestParse(t *testing.T) {
	f, err := os.Open("testdata/cachegrind.out")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		f.Close()
	})
	got, err := Parse(f)
	if err != nil {
		t.Fatal(err)
	}
	want := &report.CustomMetricSet{
		Key: "Cachegrind (./testdata/testbin/testbin 1000)",
		Metadata: []*report.MetadataKV{
			{
				Key:   "I1 cache",
				Value: "32768 B, 64 B, 8-way associative",
			},
			{
				Key:   "D1 cache",
				Value: "32768 B, 64 B, 8-way associative",
			},
			{
				Key:   "LL cache",
				Value: "8388608 B, 64 B, 16-way associative",
			},
			{
				Key:   "cmd",
				Value: "./testdata/testbin/testbin 1000",
			},
		},
		Metrics: []*report.CustomMetric{
			{Key: "Ir", Value: 37543175},
			{Key: "I1mr", Value: 973290},
			{Key: "ILmr", Value: 4579},
			{Key: "Dr", Value: 8266649},
			{Key: "D1mr", Value: 75730},
			{Key: "DLmr", Value: 9913},
			{Key: "Dw", Value: 5082704},
			{Key: "D1mw", Value: 61895},
			{Key: "DLmw", Value: 42493},
		},
	}

	opts := []cmp.Option{
		cmpopts.IgnoreFields(report.CustomMetricSet{}, "report"),
	}
	if diff := cmp.Diff(got, want, opts...); diff != "" {
		t.Error(diff)
	}
}
