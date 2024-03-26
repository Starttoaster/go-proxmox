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

	mux.HandleFunc("nodes", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(w, fixture("nodes/get_nodes.json"))
		if err != nil {
			return
		}
	})

	want := GetNodesResponse{
		Data: []GetNodesData{},
	}

	r, resp, err := client.Nodes.GetNodes()
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, *r)
}

func TestGetNode(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("nodes/srv1", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(w, fixture("nodes/get_node_status.json"))
		if err != nil {
			return
		}
	})

	want := GetNodeStatusResponse{
		Data: GetNodeStatusData{},
	}

	r, resp, err := client.Nodes.GetNodeStatus("srv1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, want, *r)
}
