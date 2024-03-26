package proxmox

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetNodes(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api2/json/nodes", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(w, fixture("nodes/get_nodes.json"))
		if err != nil {
			return
		}
	})

	want := GetNodesResponse{
		Data: []GetNodesData{
			{CPU: 0.0522061746389294, Disk: 5513633792, ID: "node/srv3", Level: "", MaxCPU: 8, MaxDisk: 100861726720, MaxMem: 16367738880, Mem: 14766223360, Node: "srv3", SslFingerprint: "00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00", Status: "online", Type: "node", Uptime: 418333},
			{CPU: 0.0220677146311971, Disk: 5727686656, ID: "node/srv1", Level: "", MaxCPU: 16, MaxDisk: 100861726720, MaxMem: 134850498560, Mem: 45189853184, Node: "srv1", SslFingerprint: "00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00", Status: "online", Type: "node", Uptime: 418326},
			{CPU: 0.0673548074849297, Disk: 5488590848, ID: "node/srv2", Level: "", MaxCPU: 8, MaxDisk: 100861726720, MaxMem: 16367742976, Mem: 13080690688, Node: "srv2", SslFingerprint: "00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00", Status: "online", Type: "node", Uptime: 418322},
		},
	}

	r, resp, err := client.Nodes.GetNodes()
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, *r)
}

func TestGetNode(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api2/json/nodes/srv1/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(w, fixture("nodes/get_node_status.json"))
		if err != nil {
			return
		}
	})

	want := GetNodeStatusResponse{
		Data: GetNodeStatusData{
			BootInfo:      BootInfo{Mode: "efi", SecureBoot: 0},
			CPU:           0.0282238002623366,
			CPUInfo:       CPUInfo{Cores: 50, Cpus: 200, Flags: "fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush dts acpi mmx fxsr sse sse2 ss ht tm pbe syscall nx pdpe1gb rdtscp lm constant_tsc art arch_perfmon pebs bts rep_good nopl xtopology nonstop_tsc cpuid aperfmperf tsc_known_freq pni pclmulqdq dtes64 monitor ds_cpl vmx smx est tm2 ssse3 sdbg fma cx16 xtpr pdcm pcid sse4_1 sse4_2 x2apic movbe popcnt tsc_deadline_timer aes xsave avx f16c rdrand lahf_lm abm 3dnowprefetch cpuid_fault invpcid_single ssbd ibrs ibpb stibp ibrs_enhanced tpr_shadow flexpriority ept vpid ept_ad fsgsbase tsc_adjust bmi1 avx2 smep bmi2 erms invpcid mpx avx512f avx512dq rdseed adx smap avx512ifma clflushopt intel_pt avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves dtherm ida arat pln pts hwp hwp_notify hwp_act_window hwp_epp hwp_pkg_req vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq rdpid fsrm md_clear flush_l1d arch_capabilities", Hvm: "1", Mhz: "4886.225", Model: "99th Gen Intel(R) Core(TM) i19-9000 @ 6.50GHz", Sockets: 1, UserHz: 100},
			CurrentKernel: CurrentKernel{Machine: "x86_64", Release: "6.5.11-8-pve", Sysname: "Linux", Version: "#1 SMP PREEMPT_DYNAMIC PMX 6.5.11-8 (2024-01-30T12:27Z)"},
			Idle:          0,
			Ksm:           Ksm{Shared: 0},
			Kversion:      "Linux 6.5.11-8-pve #1 SMP PREEMPT_DYNAMIC PMX 6.5.11-8 (2024-01-30T12:27Z)",
			LoadAvg:       []string{"0.53", "0.46", "0.43"},
			Memory:        Memory{Free: 89574653952, Total: 134850498560, Used: 45275844608},
			PveVersion:    "pve-manager/0.0.0/0000000000000000",
			RootFs:        RootFs{Avail: 89962983424, Free: 95133720576, Total: 100861726720, Used: 5728006144},
			Swap:          Swap{Free: 8589930496, Total: 8589930496, Used: 0},
			Uptime:        419090,
			Wait:          0.00150768163794533,
		},
	}

	r, resp, err := client.Nodes.GetNodeStatus("srv1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, *r)
}
