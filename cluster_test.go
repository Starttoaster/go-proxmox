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
				CPU:       testFloat64(.00215395696684676),
				Disk:      testInt(0),
				DiskRead:  testInt(1899047936),
				DiskWrite: testInt(2697581568),
				ID:        "qemu/101",
				MaxCPU:    testInt(4),
				MaxDisk:   testInt(549755813888),
				MaxMem:    testInt(17179869184),
				Mem:       testInt(3865654169),
				Name:      testStr("my-vm"),
				NetIn:     testInt(554461212),
				NetOut:    testInt(13830445),
				Node:      "node1",
				Status:    "running",
				Template:  testInt(0),
				Type:      "qemu",
				Uptime:    testInt(234806),
				VMID:      testIntOrString("101"),
			},
			{
				CgroupMode: testInt(2),
				CPU:        testFloat64(0.0496424063946151),
				Disk:       testInt(5996113920),
				ID:         "node/node1",
				Level:      testStr(""),
				MaxCPU:     testInt(16),
				MaxDisk:    testInt(100861726720),
				MaxMem:     testInt(134850514944),
				Mem:        testInt(64139268096),
				Node:       "node1",
				Status:     "online",
				Type:       "node",
				Uptime:     testInt(234832),
			},
			{
				Content:    testStr("iso,backup,vztmpl"),
				Disk:       testInt(5996113920),
				ID:         "storage/node1/local",
				MaxDisk:    testInt(100861726720),
				Node:       "node1",
				PluginType: testStr("dir"),
				Shared:     testInt(0),
				Status:     "available",
				Storage:    testStr("local"),
				Type:       "storage",
			},
			{
				ID:     "sdn/node1/localnetwork",
				Node:   "node1",
				SDN:    testStr("localnetwork"),
				Status: "ok",
				Type:   "sdn",
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
