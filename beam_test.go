package computebeam

import (
	"testing"

	"github.com/togo-framework/compute"
	"github.com/togo-framework/togo"
)

func TestArgvAppendsRunner(t *testing.T) {
	t.Setenv("BEAM_RUNNER", "FlinkRunner")
	bin, args, err := newBeam(togo.New()).argv(compute.Job{Cmd: []string{"python", "pipe.py", "--in=x"}})
	if err != nil {
		t.Fatal(err)
	}
	if bin != "python" {
		t.Fatalf("bin = %q", bin)
	}
	want := []string{"pipe.py", "--in=x", "--runner", "FlinkRunner"}
	for i := range want {
		if i >= len(args) || args[i] != want[i] {
			t.Fatalf("args = %v, want %v", args, want)
		}
	}
}
