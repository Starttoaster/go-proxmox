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
