package proxmox

import "net/http"

// ClusterService is the service that encapsulates node API methods
type ClusterService struct {
	client *Client
}

// GetClusterStatusResponse contains the response for the /cluster/status endpoint
type GetClusterStatusResponse struct {
	Data []GetClusterStatusData `json:"data"`
}

// GetClusterStatusData contains data of a cluster's status from GetClusterStatus
type GetClusterStatusData struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Type    string  `json:"type"`
	IP      *string `json:"ip"`
	Level   *string `json:"level"`
	Local   *int    `json:"local"`
	NodeID  *int    `json:"nodeid"`
	Online  *int    `json:"online"`
	Quorate *int    `json:"quorate"`
	Version *int    `json:"version"`
}

// GetClusterStatus makes a GET request to the /cluster/status endpoint
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/status
func (s *ClusterService) GetClusterStatus() (*GetClusterStatusResponse, *http.Response, error) {
	u := "cluster/status"
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	d := new(GetClusterStatusResponse)
	resp, err := s.client.Do(req, d)
	if err != nil {
		return nil, resp, err
	}

	return d, resp, nil
}

// GetClusterResourcesResponse contains the response for the /cluster/resources endpoint
type GetClusterResourcesResponse struct {
	Data []GetClusterResourcesData `json:"data"`
}

// GetClusterResourcesData contains data of a cluster's resources from GetClusterResources
type GetClusterResourcesData struct {
	ID         string       `json:"id"`
	Node       string       `json:"node"`
	Status     string       `json:"status"`
	Type       string       `json:"type"`
	CPU        *float64     `json:"cpu"`
	Disk       *int         `json:"disk"`
	DiskRead   *int         `json:"diskread"`
	DiskWrite  *int         `json:"diskwrite"`
	MaxCPU     *int         `json:"maxcpu"`
	MaxDisk    *int         `json:"maxdisk"`
	MaxMem     *int         `json:"maxmem"`
	Mem        *int         `json:"mem"`
	Name       *string      `json:"name"`
	NetIn      *int         `json:"netin"`
	NetOut     *int         `json:"netout"`
	Template   *int         `json:"template"`
	Uptime     *int         `json:"uptime"`
	VMID       *IntOrString `json:"vmid"`
	HAState    *string      `json:"hastate"`
	CgroupMode *int         `json:"cgroup-mode"`
	Level      *string      `json:"level"`
	Content    *string      `json:"content"`
	PluginType *string      `json:"plugintype"`
	Shared     *int         `json:"shared"`
	Storage    *string      `json:"storage"`
	SDN        *string      `json:"sdn"`
}

// GetClusterResources makes a GET request to the /cluster/resources endpoint
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/resources
func (s *ClusterService) GetClusterResources() (*GetClusterResourcesResponse, *http.Response, error) {
	u := "cluster/resources"
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	d := new(GetClusterResourcesResponse)
	resp, err := s.client.Do(req, d)
	if err != nil {
		return nil, resp, err
	}

	return d, resp, nil
}

// GetClusterCephStatusResponse contains the response for the /cluster/ceph/status endpoint
type GetClusterCephStatusResponse struct {
	Data GetClusterCephStatusData `json:"data"`
}

// GetClusterCephStatusData contains data of a ceph cluster's status from GetClusterCephStatus
type GetClusterCephStatusData struct {
	Health CephHealthStatus `json:"health"`
}

// GetClusterCephStatus makes a GET request to the /cluster/resources endpoint
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/cluster/resources
func (s *ClusterService) GetClusterCephStatus() (*GetClusterCephStatusResponse, *http.Response, error) {
	u := "cluster/ceph/status"
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	d := new(GetClusterCephStatusResponse)
	resp, err := s.client.Do(req, d)
	if err != nil {
		return nil, resp, err
	}

	return d, resp, nil
}
