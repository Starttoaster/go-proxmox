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
				ID:        "qemu/101",
				Node:      "node1",
				Status:    "running",
				Type:      "qemu",
				CPU:       testFloat64(0.0441805063517043),
				Disk:      testInt(0),
				DiskRead:  testInt(12915866330),
				DiskWrite: testInt(177175804928),
				MaxCPU:    testInt(8),
				MaxDisk:   testInt(137438953472),
				MaxMem:    testInt(16978542592),
				Mem:       testInt(17032423424),
				Name:      testStr("vm1"),
				NetIn:     testInt(610347185584),
				NetOut:    testInt(804729168397),
				Template:  testInt(0),
				Uptime:    testInt(1830765),
				VMID:      testIntOrString("101"),
				Tags:      testStr("cluster;node"),
				MemHost:   testInt(17032423424),
			},
			{
				ID:        "qemu/104",
				Node:      "node2",
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
				CPU:       testFloat64(0.0235323360893611),
				Disk:      testInt(0),
				DiskRead:  testInt(1278565474),
				DiskWrite: testInt(12868774912),
				MaxCPU:    testInt(1),
				MaxDisk:   testInt(17179869184),
				MaxMem:    testInt(2147483648),
				Mem:       testInt(2223873024),
				Name:      testStr("vm3"),
				NetIn:     testInt(75784079055),
				NetOut:    testInt(39914288931),
				Template:  testInt(0),
				Uptime:    testInt(1830753),
				VMID:      testIntOrString("106"),
				Tags:      testStr("lb;cluster"),
				MemHost:   testInt(2223873024),
			},
			{
				ID:        "qemu/107",
				Node:      "node3",
				Status:    "running",
				Type:      "qemu",
				CPU:       testFloat64(0.327313401970204),
				Disk:      testInt(0),
				DiskRead:  testInt(37364655734),
				DiskWrite: testInt(551237491712),
				MaxCPU:    testInt(2),
				MaxDisk:   testInt(68719476736),
				MaxMem:    testInt(8489271296),
				Mem:       testInt(8582443008),
				Name:      testStr("vm4"),
				NetIn:     testInt(859967944211),
				NetOut:    testInt(1114946233408),
				Template:  testInt(0),
				Uptime:    testInt(1830739),
				VMID:      testIntOrString("107"),
				Tags:      testStr("ctrl;cluster"),
				MemHost:   testInt(8582443008),
			},
			{
				ID:        "qemu/108",
				Node:      "node2",
				Status:    "running",
				Type:      "qemu",
				CPU:       testFloat64(0.244756950901687),
				Disk:      testInt(0),
				DiskRead:  testInt(37424977766),
				DiskWrite: testInt(544875895808),
				MaxCPU:    testInt(2),
				MaxDisk:   testInt(68719476736),
				MaxMem:    testInt(8489271296),
				Mem:       testInt(8630158336),
				Name:      testStr("vm5"),
				NetIn:     testInt(774914908666),
				NetOut:    testInt(883465947580),
				Template:  testInt(0),
				Uptime:    testInt(1830762),
				VMID:      testIntOrString("108"),
				Tags:      testStr("ctrl;cluster"),
				MemHost:   testInt(8630158336),
			},
			{
				ID:        "qemu/109",
				Node:      "node2",
				Status:    "running",
				Type:      "qemu",
				CPU:       testFloat64(0.0960014955123138),
				Disk:      testInt(0),
				DiskRead:  testInt(64818850006),
				DiskWrite: testInt(1712979336192),
				MaxCPU:    testInt(8),
				MaxDisk:   testInt(68719476736),
				MaxMem:    testInt(25467813888),
				Mem:       testInt(25644379136),
				Name:      testStr("vm6"),
				NetIn:     testInt(2151335804623),
				NetOut:    testInt(1085435188576),
				Template:  testInt(0),
				Uptime:    testInt(1830747),
				VMID:      testIntOrString("109"),
				Tags:      testStr("cluster;node"),
				MemHost:   testInt(25644379136),
			},
			{
				ID:        "qemu/110",
				Node:      "node3",
				Status:    "running",
				Type:      "qemu",
				CPU:       testFloat64(0.15015234902473),
				Disk:      testInt(0),
				DiskRead:  testInt(109703759030),
				DiskWrite: testInt(4613273582592),
				MaxCPU:    testInt(8),
				MaxDisk:   testInt(68719476736),
				MaxMem:    testInt(25467813888),
				Mem:       testInt(25642550272),
				Name:      testStr("vm7"),
				NetIn:     testInt(2194421979818),
				NetOut:    testInt(1198693108197),
				Template:  testInt(0),
				Uptime:    testInt(1830725),
				VMID:      testIntOrString("110"),
				Tags:      testStr("cluster;node"),
				MemHost:   testInt(25642550272),
			},
			{
				ID:        "qemu/111",
				Node:      "node1",
				Status:    "running",
				Type:      "qemu",
				CPU:       testFloat64(0.114121851754402),
				Disk:      testInt(0),
				DiskRead:  testInt(16019653994),
				DiskWrite: testInt(497241222144),
				MaxCPU:    testInt(8),
				MaxDisk:   testInt(68719476736),
				MaxMem:    testInt(25467813888),
				Mem:       testInt(25609126912),
				Name:      testStr("vm8"),
				NetIn:     testInt(278917230253),
				NetOut:    testInt(216578420007),
				Template:  testInt(0),
				Uptime:    testInt(328575),
				VMID:      testIntOrString("111"),
				Tags:      testStr("cluster;node"),
				MemHost:   testInt(25609126912),
			},
			{
				ID:        "qemu/112",
				Node:      "node1",
				Status:    "running",
				Type:      "qemu",
				CPU:       testFloat64(0.229044558608836),
				Disk:      testInt(0),
				DiskRead:  testInt(29946933094),
				DiskWrite: testInt(551972510720),
				MaxCPU:    testInt(2),
				MaxDisk:   testInt(68719476736),
				MaxMem:    testInt(8489271296),
				Mem:       testInt(8588177408),
				Name:      testStr("vm9"),
				NetIn:     testInt(798569839796),
				NetOut:    testInt(910633172759),
				Template:  testInt(0),
				Uptime:    testInt(1830733),
				VMID:      testIntOrString("112"),
				Tags:      testStr("ctrl;cluster"),
				MemHost:   testInt(8588177408),
			},
			{
				ID:        "qemu/113",
				Node:      "node1",
				Status:    "running",
				Type:      "qemu",
				CPU:       testFloat64(0.00427122720016477),
				Disk:      testInt(0),
				DiskRead:  testInt(979514878),
				DiskWrite: testInt(10042747904),
				MaxCPU:    testInt(4),
				MaxDisk:   testInt(17179869184),
				MaxMem:    testInt(4395630592),
				Mem:       testInt(2839129088),
				Name:      testStr("vm10"),
				NetIn:     testInt(25391575912),
				NetOut:    testInt(1414584307),
				Template:  testInt(0),
				Uptime:    testInt(1830719),
				VMID:      testIntOrString("113"),
				Tags:      testStr("vpn"),
				MemHost:   testInt(2839129088),
			},
			{
				ID:         "node/node3",
				Node:       "node3",
				Status:     "online",
				Type:       "node",
				CPU:        testFloat64(0.0665251238499646),
				Disk:       testInt(12725993472),
				MaxCPU:     testInt(32),
				MaxDisk:    testInt(100861726720),
				MaxMem:     testInt(98776571904),
				Mem:        testInt(46570885120),
				Uptime:     testInt(1830819),
				HAState:    testStr("online"),
				CgroupMode: testInt(2),
				Level:      testStr(""),
			},
			{
				ID:         "node/node1",
				Node:       "node1",
				Status:     "online",
				Type:       "node",
				CPU:        testFloat64(0.0643182201247261),
				Disk:       testInt(18302795776),
				MaxCPU:     testInt(32),
				MaxDisk:    testInt(100861726720),
				MaxMem:     testInt(98785284096),
				Mem:        testInt(63559094272),
				Uptime:     testInt(1830818),
				HAState:    testStr("online"),
				CgroupMode: testInt(2),
				Level:      testStr(""),
			},
			{
				ID:         "node/node2",
				Node:       "node2",
				Status:     "online",
				Type:       "node",
				CPU:        testFloat64(0.0406366978046066),
				Disk:       testInt(21873754112),
				MaxCPU:     testInt(32),
				MaxDisk:    testInt(100861726720),
				MaxMem:     testInt(98776588288),
				Mem:        testInt(45313855488),
				Uptime:     testInt(1830813),
				HAState:    testStr("online"),
				CgroupMode: testInt(2),
				Level:      testStr(""),
			},
			{
				ID:         "storage/node3/local",
				Node:       "node3",
				Status:     "available",
				Type:       "storage",
				Disk:       testInt(12725993472),
				MaxDisk:    testInt(100861726720),
				Content:    testStr("iso,vztmpl,backup"),
				PluginType: testStr("dir"),
				Shared:     testInt(0),
				Storage:    testStr("local"),
			},
			{
				ID:         "storage/node1/local",
				Node:       "node1",
				Status:     "available",
				Type:       "storage",
				Disk:       testInt(18302799872),
				MaxDisk:    testInt(100861726720),
				Content:    testStr("iso,vztmpl,backup"),
				PluginType: testStr("dir"),
				Shared:     testInt(0),
				Storage:    testStr("local"),
			},
			{
				ID:         "storage/node2/local",
				Node:       "node2",
				Status:     "available",
				Type:       "storage",
				Disk:       testInt(21873754112),
				MaxDisk:    testInt(100861726720),
				Content:    testStr("iso,vztmpl,backup"),
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
				MaxDisk:    testInt(544877838336),
				Content:    testStr("backup,vztmpl,iso"),
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
				MaxDisk:    testInt(544877838336),
				Content:    testStr("backup,vztmpl,iso"),
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
				MaxDisk:    testInt(544877838336),
				Content:    testStr("backup,vztmpl,iso"),
				PluginType: testStr("cephfs"),
				Shared:     testInt(1),
				Storage:    testStr("cephfs"),
			},
			{
				ID:         "storage/node3/local-lvm",
				Node:       "node3",
				Status:     "available",
				Type:       "storage",
				Disk:       testInt(2716826442),
				MaxDisk:    testInt(876395626496),
				Content:    testStr("images,rootdir"),
				PluginType: testStr("lvmthin"),
				Shared:     testInt(0),
				Storage:    testStr("local-lvm"),
			},
			{
				ID:         "storage/node1/local-lvm",
				Node:       "node1",
				Status:     "available",
				Type:       "storage",
				Disk:       testInt(371078053704),
				MaxDisk:    testInt(1836111101952),
				Content:    testStr("images,rootdir"),
				PluginType: testStr("lvmthin"),
				Shared:     testInt(0),
				Storage:    testStr("local-lvm"),
			},
			{
				ID:         "storage/node2/local-lvm",
				Node:       "node2",
				Status:     "available",
				Type:       "storage",
				Disk:       testInt(130946277834),
				MaxDisk:    testInt(1884119105536),
				Content:    testStr("images,rootdir"),
				PluginType: testStr("lvmthin"),
				Shared:     testInt(0),
				Storage:    testStr("local-lvm"),
			},
			{
				ID:         "storage/node3/ceph-nvme",
				Node:       "node3",
				Status:     "available",
				Type:       "storage",
				Disk:       testInt(1219804183551),
				MaxDisk:    testInt(1757817650175),
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
				Disk:       testInt(1219804183551),
				MaxDisk:    testInt(1757817650175),
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
				Disk:       testInt(1219803659263),
				MaxDisk:    testInt(1757817158655),
				Content:    testStr("rootdir,images"),
				PluginType: testStr("rbd"),
				Shared:     testInt(1),
				Storage:    testStr("ceph-nvme"),
			},
			{
				ID:     "sdn/node3/localzone",
				Node:   "node3",
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
			{
				ID:     "sdn/node2/localzone",
				Node:   "node2",
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
