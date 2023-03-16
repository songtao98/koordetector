//go:build 386 || amd64 || amd64p32 || arm || arm64 || mips64le || mips64p32le || mipsle || ppc64le || riscv64
// +build 386 amd64 amd64p32 arm arm64 mips64le mips64p32le mipsle ppc64le riscv64

/*
Copyright 2022 The Koordinator Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by bpf2go; DO NOT EDIT.

package cpu_schedule_latency

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

// loadBpf returns the embedded CollectionSpec for bpf.
func loadBpf() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_BpfBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf: %w", err)
	}

	return spec, err
}

// loadBpfObjects loads bpf and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpfObjects
//	*bpfPrograms
//	*bpfMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpfObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpfSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfSpecs struct {
	bpfProgramSpecs
	bpfMapSpecs
}

// bpfSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfProgramSpecs struct {
	HandleSchedWakeup    *ebpf.ProgramSpec `ebpf:"handle__sched_wakeup"`
	HandleSchedWakeupNew *ebpf.ProgramSpec `ebpf:"handle__sched_wakeup_new"`
	HandleSwitch         *ebpf.ProgramSpec `ebpf:"handle_switch"`
}

// bpfMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfMapSpecs struct {
	OutputCgroupCounter *ebpf.MapSpec `ebpf:"output_cgroup_counter"`
	OutputCgroupDelay   *ebpf.MapSpec `ebpf:"output_cgroup_delay"`
	PidStartTime        *ebpf.MapSpec `ebpf:"pid_start_time"`
}

// bpfObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfObjects struct {
	bpfPrograms
	bpfMaps
}

func (o *bpfObjects) Close() error {
	return _BpfClose(
		&o.bpfPrograms,
		&o.bpfMaps,
	)
}

// bpfMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfMaps struct {
	OutputCgroupCounter *ebpf.Map `ebpf:"output_cgroup_counter"`
	OutputCgroupDelay   *ebpf.Map `ebpf:"output_cgroup_delay"`
	PidStartTime        *ebpf.Map `ebpf:"pid_start_time"`
}

func (m *bpfMaps) Close() error {
	return _BpfClose(
		m.OutputCgroupCounter,
		m.OutputCgroupDelay,
		m.PidStartTime,
	)
}

// bpfPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfPrograms struct {
	HandleSchedWakeup    *ebpf.Program `ebpf:"handle__sched_wakeup"`
	HandleSchedWakeupNew *ebpf.Program `ebpf:"handle__sched_wakeup_new"`
	HandleSwitch         *ebpf.Program `ebpf:"handle_switch"`
}

func (p *bpfPrograms) Close() error {
	return _BpfClose(
		p.HandleSchedWakeup,
		p.HandleSchedWakeupNew,
		p.HandleSwitch,
	)
}

func _BpfClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_bpfel.o
var _BpfBytes []byte
