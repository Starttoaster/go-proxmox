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
			{CPU: 0.0522061746389294, Disk: 5513633792, ID: "node/srv3", Level: "", MaxCPU: 8, MaxDisk: 100861726720, MaxMem: 16367738880, Mem: 14766223360, Node: "srv3", SSLFingerprint: "00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00", Status: "online", Type: "node", Uptime: 418333},
			{CPU: 0.0220677146311971, Disk: 5727686656, ID: "node/srv1", Level: "", MaxCPU: 16, MaxDisk: 100861726720, MaxMem: 134850498560, Mem: 45189853184, Node: "srv1", SSLFingerprint: "00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00", Status: "online", Type: "node", Uptime: 418326},
			{CPU: 0.0673548074849297, Disk: 5488590848, ID: "node/srv2", Level: "", MaxCPU: 8, MaxDisk: 100861726720, MaxMem: 16367742976, Mem: 13080690688, Node: "srv2", SSLFingerprint: "00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00:00", Status: "online", Type: "node", Uptime: 418322},
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
			CPUInfo:       CPUInfo{Cores: 50, CPUs: 200, Flags: "fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush dts acpi mmx fxsr sse sse2 ss ht tm pbe syscall nx pdpe1gb rdtscp lm constant_tsc art arch_perfmon pebs bts rep_good nopl xtopology nonstop_tsc cpuid aperfmperf tsc_known_freq pni pclmulqdq dtes64 monitor ds_cpl vmx smx est tm2 ssse3 sdbg fma cx16 xtpr pdcm pcid sse4_1 sse4_2 x2apic movbe popcnt tsc_deadline_timer aes xsave avx f16c rdrand lahf_lm abm 3dnowprefetch cpuid_fault invpcid_single ssbd ibrs ibpb stibp ibrs_enhanced tpr_shadow flexpriority ept vpid ept_ad fsgsbase tsc_adjust bmi1 avx2 smep bmi2 erms invpcid mpx avx512f avx512dq rdseed adx smap avx512ifma clflushopt intel_pt avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 xsaves dtherm ida arat pln pts hwp hwp_notify hwp_act_window hwp_epp hwp_pkg_req vnmi avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntdq rdpid fsrm md_clear flush_l1d arch_capabilities", HVM: "1", MHz: "4886.225", Model: "99th Gen Intel(R) Core(TM) i19-9000 @ 6.50GHz", Sockets: 1, UserHz: 100},
			CurrentKernel: CurrentKernel{Machine: "x86_64", Release: "6.5.11-8-pve", SysName: "Linux", Version: "#1 SMP PREEMPT_DYNAMIC PMX 6.5.11-8 (2024-01-30T12:27Z)"},
			Idle:          0,
			KSM:           KSM{Shared: 0},
			Kversion:      "Linux 6.5.11-8-pve #1 SMP PREEMPT_DYNAMIC PMX 6.5.11-8 (2024-01-30T12:27Z)",
			LoadAvg:       []string{"0.53", "0.46", "0.43"},
			Memory:        Memory{Free: 89574653952, Total: 134850498560, Used: 45275844608},
			PveVersion:    "pve-manager/0.0.0/0000000000000000",
			RootFs:        RootFS{Avail: 89962983424, Free: 95133720576, Total: 100861726720, Used: 5728006144},
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

func TestGetNodeLxc(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api2/json/nodes/srv1/lxc", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(w, fixture("nodes/get_node_lxc.json"))
		if err != nil {
			return
		}
	})

	want := GetNodeLxcResponse{
		Data: []GetNodeLxcData{
			{
				CPU:       0,
				CPUs:      1,
				Disk:      0,
				DiskRead:  0,
				DiskWrite: 0,
				MaxDisk:   8589934592,
				MaxMem:    536870912,
				MaxSwap:   536870912,
				Mem:       0,
				Name:      "CT103",
				NetIn:     0,
				NetOut:    0,
				Status:    "stopped",
				Type:      "lxc",
				Uptime:    0,
				VMID:      IntOrString("103"),
			},
			{
				CPU:       0,
				CPUs:      1,
				Disk:      0,
				DiskRead:  0,
				DiskWrite: 0,
				MaxDisk:   8589934592,
				MaxMem:    536870912,
				MaxSwap:   536870912,
				Mem:       0,
				Name:      "CT104",
				NetIn:     0,
				NetOut:    0,
				Status:    "stopped",
				Type:      "lxc",
				Uptime:    0,
				VMID:      IntOrString("104"),
			},
		},
	}

	r, resp, err := client.Nodes.GetNodeLxc("srv1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, *r)
}

func TestGetNodeQemu(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api2/json/nodes/srv1/qemu", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(w, fixture("nodes/get_node_qemu.json"))
		if err != nil {
			return
		}
	})

	want := GetNodeQemuResponse{
		Data: []GetNodeQemuData{
			{
				CPU:       0.0156071608339279,
				CPUs:      5,
				Disk:      0,
				DiskRead:  0,
				DiskWrite: 0,
				MaxDisk:   274877906944,
				MaxMem:    8589934592,
				Mem:       3072665958,
				PID:       1551,
				Name:      "test",
				NetIn:     294680188,
				NetOut:    200064110,
				Status:    "running",
				Uptime:    28661,
				VMID:      IntOrString("104"),
			},
		},
	}

	r, resp, err := client.Nodes.GetNodeQemu("srv1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, *r)
}

func TestGetNodeDisksList(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api2/json/nodes/srv1/disks/list", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(w, fixture("nodes/get_node_disks_list.json"))
		if err != nil {
			return
		}
	})

	want := GetNodeDisksListResponse{
		Data: []GetNodeDisksListData{
			{
				DevPath:  "/dev/nvme0n1",
				WWN:      "testwwn",
				Serial:   "testser",
				ByIDLink: "/dev/disk/by-id/nvme-drive",
				Size:     512110190592,
				Used:     "BIOS boot",
				RPM:      IntOrString("0"),
				Wearout:  IntOrString("99"),
				GPT:      1,
				Health:   "PASSED",
				Model:    "Test Model",
				Vendor:   "unknown",
				Type:     "nvme",
			},
			{
				Serial:       "testser",
				ByIDLink:     "/dev/disk/by-id/ata-TOSHIBA_THNS",
				Bluestore:    1,
				DevPath:      "/dev/sda",
				OSDEncrypted: 0,
				Type:         "ssd",
				Health:       "PASSED",
				RPM:          IntOrString("0"),
				Wearout:      IntOrString("N/A"),
				GPT:          0,
				Vendor:       "unknown",
				Size:         1920383410176,
				Model:        "Test Model",
				WWN:          "testwwn",
				Used:         "LVM",
			},
			{
				DevPath:  "/dev/sda",
				WWN:      "testwwn",
				Serial:   "testser",
				ByIDLink: "/dev/disk/by-id/ata-ST2000",
				Size:     2000398934016,
				Used:     "LVM",
				RPM:      IntOrString("5400"),
				Wearout:  IntOrString("N/A"),
				GPT:      0,
				Health:   "PASSED",
				Model:    "Test Model",
				Vendor:   "unknown",
				Type:     "hdd",
			},
			{
				DevPath:  "/dev/sdc",
				WWN:      "testwwn",
				Serial:   "testser",
				ByIDLink: "/dev/disk/by-id/ata-WDC_DRIVE2",
				Size:     5000947302400,
				Used:     "ext4",
				RPM:      IntOrString("4800"),
				Wearout:  IntOrString("N/A"),
				GPT:      0,
				Health:   "PASSED",
				Model:    "Test Model",
				Vendor:   "unknown",
				Type:     "hdd",
			},
			{
				DevPath:  "/dev/sde",
				WWN:      "testwwn",
				Serial:   "testser",
				ByIDLink: "/dev/disk/by-id/ata-WDC_DRIVE1",
				Size:     5000947302400,
				Used:     "ZFS",
				RPM:      IntOrString("4800"),
				Wearout:  IntOrString("N/A"),
				GPT:      0,
				Health:   "PASSED",
				Model:    "Test Model",
				Vendor:   "unknown",
				Type:     "hdd",
			},
		},
	}

	r, resp, err := client.Nodes.GetNodeDisksList("srv1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, *r)
}
