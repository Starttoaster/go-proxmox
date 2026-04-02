package proxmox

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func testStr(s string) *string {
	return &s
}

func testInt(i int) *int {
	return &i
}

func testFloat64(f float64) *float64 {
	return &f
}

func testIntOrString(is IntOrString) *IntOrString {
	return &is
}

func TestGetClusterCephStatus(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api2/json/cluster/ceph/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(w, fixture("clusters/get_cluster_ceph_status.json"))
		if err != nil {
			return
		}
	})

	want := GetClusterCephStatusResponse{
		Data: GetClusterCephStatusData{
			Health: CephHealthStatus{
				Status: "HEALTH_OK",
			},
		},
	}

	r, resp, err := client.Cluster.GetClusterCephStatus()
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, *r)
}

func TestGetClusterResources(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api2/json/cluster/resources", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(w, fixture("clusters/get_cluster_resources.json"))
		if err != nil {
			return
		}
	})

	want := GetClusterResourcesResponse{
		Data: []GetClusterResourcesData{
			{
				ID:        "lxc/100",
				Node:      "node1",
				Status:    "stopped",
				Type:      "lxc",
				CPU:       testFloat64(0),
				Disk:      testInt(0),
				DiskRead:  testInt(0),
				DiskWrite: testInt(0),
				MaxCPU:    testInt(1),
				MaxDisk:   testInt(8589934592),
				MaxMem:    testInt(536870912),
				Mem:       testInt(0),
				Name:      testStr("test"),
				NetIn:     testInt(0),
				NetOut:    testInt(0),
				Template:  testInt(0),
				Uptime:    testInt(0),
				VMID:      testIntOrString("100"),
				MemHost:   testInt(0),
			},
			{
				ID:        "qemu/101",
				Node:      "node2",
				Status:    "running",
				Type:      "qemu",
				CPU:       testFloat64(0.0536056293799136),
				Disk:      testInt(0),
				DiskRead:  testInt(12915866330),
				DiskWrite: testInt(177310137344),
				MaxCPU:    testInt(8),
				MaxDisk:   testInt(137438953472),
				MaxMem:    testInt(16978542592),
				Mem:       testInt(17032423424),
				Name:      testStr("vm1"),
				NetIn:     testInt(610675465429),
				NetOut:    testInt(805175854687),
				Template:  testInt(0),
				Uptime:    testInt(1833355),
				VMID:      testIntOrString("101"),
				Tags:      testStr("cluster;node"),
				MemHost:   testInt(17032423424),
			},
			{
				ID:        "qemu/104",
				Node:      "node1",
				Status:    "stopped",
				Type:      "qemu",
				CPU:       testFloat64(0),
				Disk:      testInt(0),
				DiskRead:  testInt(0),
				DiskWrite: testInt(0),
				MaxCPU:    testInt(2),
				MaxDisk:   testInt(5368709120),
				MaxMem:    testInt(2147483648),
				Mem:       testInt(0),
				Name:      testStr("vm2"),
				NetIn:     testInt(0),
				NetOut:    testInt(0),
				Template:  testInt(1),
				Uptime:    testInt(0),
				VMID:      testIntOrString("104"),
				MemHost:   testInt(0),
			},
			{
				ID:        "qemu/106",
				Node:      "node3",
				Status:    "running",
				Type:      "qemu",
				CPU:       testFloat64(0.0183663020967424),
				Disk:      testInt(0),
				DiskRead:  testInt(1278565474),
				DiskWrite: testInt(12873182208),
				MaxCPU:    testInt(1),
				MaxDisk:   testInt(17179869184),
				MaxMem:    testInt(2147483648),
				Mem:       testInt(2223873024),
				Name:      testStr("vm3"),
				NetIn:     testInt(75861088579),
				NetOut:    testInt(39968681271),
				Template:  testInt(0),
				Uptime:    testInt(1833344),
				VMID:      testIntOrString("106"),
				Tags:      testStr("lb;cluster"),
				MemHost:   testInt(2223873024),
			},
			{
				ID:        "qemu/107",
				Node:      "node3",
				Status:    "running",
				Type:      "qemu",
				CPU:       testFloat64(0.568872041259627),
				Disk:      testInt(0),
				DiskRead:  testInt(37376390774),
				DiskWrite: testInt(551949237248),
				MaxCPU:    testInt(2),
				MaxDisk:   testInt(68719476736),
				MaxMem:    testInt(8489271296),
				Mem:       testInt(8582443008),
				Name:      testStr("vm4"),
				NetIn:     testInt(860952792767),
				NetOut:    testInt(1116581503813),
				Template:  testInt(0),
				Uptime:    testInt(1833330),
				VMID:      testIntOrString("107"),
				Tags:      testStr("ctrl;cluster"),
				MemHost:   testInt(8582443008),
			},
			{
				ID:        "qemu/108",
				Node:      "node1",
				Status:    "running",
				Type:      "qemu",
				CPU:       testFloat64(0.361713948952683),
				Disk:      testInt(0),
				DiskRead:  testInt(37426132838),
				DiskWrite: testInt(545553050624),
				MaxCPU:    testInt(2),
				MaxDisk:   testInt(68719476736),
				MaxMem:    testInt(8489271296),
				Mem:       testInt(8630176768),
				Name:      testStr("vm5"),
				NetIn:     testInt(776140060776),
				NetOut:    testInt(884997933117),
				Template:  testInt(0),
				Uptime:    testInt(1833362),
				VMID:      testIntOrString("108"),
				Tags:      testStr("ctrl;cluster"),
				MemHost:   testInt(8630176768),
			},
			{
				ID:        "qemu/109",
				Node:      "node1",
				Status:    "running",
				Type:      "qemu",
				CPU:       testFloat64(0.105540199677705),
				Disk:      testInt(0),
				DiskRead:  testInt(64818850006),
				DiskWrite: testInt(1713218477056),
				MaxCPU:    testInt(8),
				MaxDisk:   testInt(68719476736),
				MaxMem:    testInt(25467813888),
				Mem:       testInt(25644397568),
				Name:      testStr("vm6"),
				NetIn:     testInt(2152083381007),
				NetOut:    testInt(1088280916253),
				Template:  testInt(0),
				Uptime:    testInt(1833348),
				VMID:      testIntOrString("109"),
				Tags:      testStr("cluster;node"),
				MemHost:   testInt(25644397568),
			},
			{
				ID:        "qemu/110",
				Node:      "node3",
				Status:    "running",
				Type:      "qemu",
				CPU:       testFloat64(0.318993667996052),
				Disk:      testInt(0),
				DiskRead:  testInt(109828920502),
				DiskWrite: testInt(4620150258688),
				MaxCPU:    testInt(8),
				MaxDisk:   testInt(68719476736),
				MaxMem:    testInt(25467813888),
				Mem:       testInt(25618977792),
				Name:      testStr("vm7"),
				NetIn:     testInt(2198201784088),
				NetOut:    testInt(1199842494628),
				Template:  testInt(0),
				Uptime:    testInt(1833316),
				VMID:      testIntOrString("110"),
				Tags:      testStr("cluster;node"),
				MemHost:   testInt(25618977792),
			},
			{
				ID:        "qemu/111",
				Node:      "node2",
				Status:    "running",
				Type:      "qemu",
				CPU:       testFloat64(0.0787374456478033),
				Disk:      testInt(0),
				DiskRead:  testInt(16019653994),
				DiskWrite: testInt(497428134912),
				MaxCPU:    testInt(8),
				MaxDisk:   testInt(68719476736),
				MaxMem:    testInt(25467813888),
				Mem:       testInt(25609127936),
				Name:      testStr("vm8"),
				NetIn:     testInt(280615631906),
				NetOut:    testInt(223444955696),
				Template:  testInt(0),
				Uptime:    testInt(331166),
				VMID:      testIntOrString("111"),
				Tags:      testStr("cluster;node"),
				MemHost:   testInt(25609127936),
			},
			{
				ID:        "qemu/112",
				Node:      "node2",
				Status:    "running",
				Type:      "qemu",
				CPU:       testFloat64(0.276449978946787),
				Disk:      testInt(0),
				DiskRead:  testInt(29948415846),
				DiskWrite: testInt(552639618048),
				MaxCPU:    testInt(2),
				MaxDisk:   testInt(68719476736),
				MaxMem:    testInt(8489271296),
				Mem:       testInt(8588177408),
				Name:      testStr("vm9"),
				NetIn:     testInt(799226699378),
				NetOut:    testInt(911437859584),
				Template:  testInt(0),
				Uptime:    testInt(1833324),
				VMID:      testIntOrString("112"),
				Tags:      testStr("ctrl;cluster"),
				MemHost:   testInt(8588177408),
			},
			{
				ID:        "qemu/113",
				Node:      "node2",
				Status:    "running",
				Type:      "qemu",
				CPU:       testFloat64(0.00187151823271519),
				Disk:      testInt(0),
				DiskRead:  testInt(979514878),
				DiskWrite: testInt(10049477632),
				MaxCPU:    testInt(4),
				MaxDisk:   testInt(17179869184),
				MaxMem:    testInt(4395630592),
				Mem:       testInt(2839129088),
				Name:      testStr("vm10"),
				NetIn:     testInt(25397938405),
				NetOut:    testInt(1416515850),
				Template:  testInt(0),
				Uptime:    testInt(1833309),
				VMID:      testIntOrString("113"),
				Tags:      testStr("vpn"),
				MemHost:   testInt(2839129088),
			},
			{
				ID:         "node/node3",
				Node:       "node3",
				Status:     "online",
				Type:       "node",
				CPU:        testFloat64(0.11512574190785),
				Disk:       testInt(12727828480),
				MaxCPU:     testInt(32),
				MaxDisk:    testInt(100861726720),
				MaxMem:     testInt(98776571904),
				Mem:        testInt(46290939904),
				Uptime:     testInt(1833410),
				HAState:    testStr("online"),
				CgroupMode: testInt(2),
				Level:      testStr(""),
			},
			{
				ID:         "node/node2",
				Node:       "node2",
				Status:     "online",
				Type:       "node",
				CPU:        testFloat64(0.0530744118735907),
				Disk:       testInt(18304585728),
				MaxCPU:     testInt(32),
				MaxDisk:    testInt(100861726720),
				MaxMem:     testInt(98785284096),
				Mem:        testInt(63819849728),
				Uptime:     testInt(1833408),
				HAState:    testStr("online"),
				CgroupMode: testInt(2),
				Level:      testStr(""),
			},
			{
				ID:         "node/node1",
				Node:       "node1",
				Status:     "online",
				Type:       "node",
				CPU:        testFloat64(0.0520396376679433),
				Disk:       testInt(21876637696),
				MaxCPU:     testInt(32),
				MaxDisk:    testInt(100861726720),
				MaxMem:     testInt(98776588288),
				Mem:        testInt(45275222016),
				Uptime:     testInt(1833414),
				HAState:    testStr("online"),
				CgroupMode: testInt(2),
				Level:      testStr(""),
			},
			{
				ID:         "storage/node3/ceph-nvme",
				Node:       "node3",
				Status:     "available",
				Type:       "storage",
				Disk:       testInt(1219904847687),
				MaxDisk:    testInt(1758015471431),
				Content:    testStr("rootdir,images"),
				PluginType: testStr("rbd"),
				Shared:     testInt(1),
				Storage:    testStr("ceph-nvme"),
			},
			{
				ID:         "storage/node2/ceph-nvme",
				Node:       "node2",
				Status:     "available",
				Type:       "storage",
				Disk:       testInt(1219904847687),
				MaxDisk:    testInt(1758015471431),
				Content:    testStr("rootdir,images"),
				PluginType: testStr("rbd"),
				Shared:     testInt(1),
				Storage:    testStr("ceph-nvme"),
			},
			{
				ID:         "storage/node1/ceph-nvme",
				Node:       "node1",
				Status:     "available",
				Type:       "storage",
				Disk:       testInt(1219904847687),
				MaxDisk:    testInt(1758015471431),
				Content:    testStr("rootdir,images"),
				PluginType: testStr("rbd"),
				Shared:     testInt(1),
				Storage:    testStr("ceph-nvme"),
			},
			{
				ID:         "storage/node3/local-lvm",
				Node:       "node3",
				Status:     "available",
				Type:       "storage",
				Disk:       testInt(2716826442),
				MaxDisk:    testInt(876395626496),
				Content:    testStr("rootdir,images"),
				PluginType: testStr("lvmthin"),
				Shared:     testInt(0),
				Storage:    testStr("local-lvm"),
			},
			{
				ID:         "storage/node2/local-lvm",
				Node:       "node2",
				Status:     "available",
				Type:       "storage",
				Disk:       testInt(371078053704),
				MaxDisk:    testInt(1836111101952),
				Content:    testStr("rootdir,images"),
				PluginType: testStr("lvmthin"),
				Shared:     testInt(0),
				Storage:    testStr("local-lvm"),
			},
			{
				ID:         "storage/node1/local-lvm",
				Node:       "node1",
				Status:     "available",
				Type:       "storage",
				Disk:       testInt(130946277834),
				MaxDisk:    testInt(1884119105536),
				Content:    testStr("rootdir,images"),
				PluginType: testStr("lvmthin"),
				Shared:     testInt(0),
				Storage:    testStr("local-lvm"),
			},
			{
				ID:         "storage/node3/local",
				Node:       "node3",
				Status:     "available",
				Type:       "storage",
				Disk:       testInt(12727828480),
				MaxDisk:    testInt(100861726720),
				Content:    testStr("vztmpl,backup,iso"),
				PluginType: testStr("dir"),
				Shared:     testInt(0),
				Storage:    testStr("local"),
			},
			{
				ID:         "storage/node2/local",
				Node:       "node2",
				Status:     "available",
				Type:       "storage",
				Disk:       testInt(18304585728),
				MaxDisk:    testInt(100861726720),
				Content:    testStr("vztmpl,backup,iso"),
				PluginType: testStr("dir"),
				Shared:     testInt(0),
				Storage:    testStr("local"),
			},
			{
				ID:         "storage/node1/local",
				Node:       "node1",
				Status:     "available",
				Type:       "storage",
				Disk:       testInt(21876637696),
				MaxDisk:    testInt(100861726720),
				Content:    testStr("vztmpl,backup,iso"),
				PluginType: testStr("dir"),
				Shared:     testInt(0),
				Storage:    testStr("local"),
			},
			{
				ID:         "storage/node3/cephfs",
				Node:       "node3",
				Status:     "available",
				Type:       "storage",
				Disk:       testInt(6866075648),
				MaxDisk:    testInt(544974307328),
				Content:    testStr("iso,backup,vztmpl"),
				PluginType: testStr("cephfs"),
				Shared:     testInt(1),
				Storage:    testStr("cephfs"),
			},
			{
				ID:         "storage/node2/cephfs",
				Node:       "node2",
				Status:     "available",
				Type:       "storage",
				Disk:       testInt(6866075648),
				MaxDisk:    testInt(544974307328),
				Content:    testStr("iso,backup,vztmpl"),
				PluginType: testStr("cephfs"),
				Shared:     testInt(1),
				Storage:    testStr("cephfs"),
			},
			{
				ID:         "storage/node1/cephfs",
				Node:       "node1",
				Status:     "available",
				Type:       "storage",
				Disk:       testInt(6866075648),
				MaxDisk:    testInt(544974307328),
				Content:    testStr("iso,backup,vztmpl"),
				PluginType: testStr("cephfs"),
				Shared:     testInt(1),
				Storage:    testStr("cephfs"),
			},
			{
				ID:     "sdn/node3/localzone",
				Node:   "node3",
				Status: "ok",
				Type:   "sdn",
				SDN:    testStr("localzone"),
			},
			{
				ID:     "sdn/node2/localzone",
				Node:   "node2",
				Status: "ok",
				Type:   "sdn",
				SDN:    testStr("localzone"),
			},
			{
				ID:     "sdn/node1/localzone",
				Node:   "node1",
				Status: "ok",
				Type:   "sdn",
				SDN:    testStr("localzone"),
			},
		},
	}

	r, resp, err := client.Cluster.GetClusterResources()
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, *r)
}

func TestGetClusterStatus(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api2/json/cluster/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(w, fixture("clusters/get_cluster_status.json"))
		if err != nil {
			return
		}
	})

	want := GetClusterStatusResponse{
		Data: []GetClusterStatusData{
			{
				Type:    "cluster",
				ID:      "cluster",
				Version: testInt(5),
				Quorate: testInt(1),
				Name:    "prd",
			},
			{
				NodeID: testInt(2),
				Type:   "node",
				IP:     testStr("10.0.1.2"),
				Name:   "cmp2",
				Level:  testStr(""),
				Online: testInt(1),
				ID:     "node/cmp2",
				Local:  testInt(0),
			},
			{
				NodeID: testInt(4),
				Type:   "node",
				IP:     testStr("10.0.1.3"),
				Name:   "cmp3",
				Level:  testStr(""),
				Online: testInt(1),
				ID:     "node/cmp3",
				Local:  testInt(0),
			},
			{
				NodeID: testInt(1),
				Type:   "node",
				IP:     testStr("10.0.1.1"),
				Name:   "cmp1",
				Level:  testStr(""),
				Online: testInt(1),
				ID:     "node/cmp1",
				Local:  testInt(1),
			},
		},
	}

	r, resp, err := client.Cluster.GetClusterStatus()
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, *r)
}
