// Package computebeam is a togo compute backend that runs Apache Beam pipelines.
// A Beam pipeline is a self-contained program, so this runs job.Cmd and appends
// `--runner <BEAM_RUNNER>`. Select with `togo provider:use compute beam`.
//
//   BEAM_RUNNER  the Beam runner (default "DirectRunner"; e.g. DataflowRunner,
//                FlinkRunner, SparkRunner)
package computebeam

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os/exec"

	"github.com/togo-framework/compute"
	"github.com/togo-framework/providers"
	"github.com/togo-framework/togo"
)

func init() {
	togo.RegisterProviderFunc("compute-beam", togo.PriorityService+1, func(k *togo.Kernel) error {
		providers.Use(k, providers.CapCompute, "beam", newBeam(k), false)
		if k.Log != nil {
			k.Log.Info("plugin active", "plugin", "compute-beam")
		}
		return nil
	})
}

type beam struct{ runner string }

func newBeam(k *togo.Kernel) *beam {
	return &beam{runner: providers.Value(k, providers.CapCompute, "beam", "runner", "DirectRunner", false)}
}

// argv returns the program invocation with the Beam runner appended.
func (b *beam) argv(job compute.Job) (string, []string, error) {
	if len(job.Cmd) == 0 {
		return "", nil, fmt.Errorf("compute-beam: job %q has no Cmd (the pipeline program)", job.Name)
	}
	args := append(append([]string{}, job.Cmd[1:]...), "--runner", b.runner)
	return job.Cmd[0], args, nil
}

func (b *beam) Submit(ctx context.Context, job compute.Job) (compute.Run, error) {
	bin, args, err := b.argv(job)
	if err != nil {
		return compute.Run{}, err
	}
	cmd := exec.CommandContext(ctx, bin, args...)
	var buf bytes.Buffer
	cmd.Stdout, cmd.Stderr = &buf, &buf
	run := compute.Run{ID: id(), Status: "succeeded"}
	if e := cmd.Run(); e != nil {
		run.Status = "failed"
	}
	run.Output = buf.String()
	return run, nil
}

func id() string { b := make([]byte, 8); _, _ = rand.Read(b); return hex.EncodeToString(b) }
