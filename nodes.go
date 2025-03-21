package proxmox

import (
	"fmt"
	"net/http"
)

// NodeService is the service that encapsulates node API methods
type NodeService struct {
	client *Client
}

// GetNodesResponse contains the response for the /nodes endpoint
type GetNodesResponse struct {
	Data []GetNodesData `json:"data"`
}

// GetNodesData contains data of one node from a GetNodes response
type GetNodesData struct {
	CPU            float64 `json:"cpu"`
	Disk           int     `json:"disk"`
	ID             string  `json:"id"`
	Level          string  `json:"level"`
	MaxCPU         int     `json:"maxcpu"`
	MaxDisk        int     `json:"maxdisk"`
	MaxMem         int     `json:"maxmem"`
	Mem            int     `json:"mem"`
	Node           string  `json:"node"`
	SSLFingerprint string  `json:"ssl_fingerprint"`
	Status         string  `json:"status"`
	Type           string  `json:"type"`
	Uptime         int     `json:"uptime"`
}

// GetNodes makes a GET request to the /nodes endpoint
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes
func (s *NodeService) GetNodes() (*GetNodesResponse, *http.Response, error) {
	u := "nodes"
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	d := new(GetNodesResponse)
	resp, err := s.client.Do(req, d)
	if err != nil {
		return nil, resp, err
	}

	return d, resp, nil
}

// GetNodeStatusResponse contains the response for the /nodes/{node}/status endpoint
type GetNodeStatusResponse struct {
	Data GetNodeStatusData `json:"data"`
}

// GetNodeStatusData contains data of one node from a GetNode response
type GetNodeStatusData struct {
	BootInfo      BootInfo      `json:"boot-info"`
	CPU           float64       `json:"cpu"`
	CPUInfo       CPUInfo       `json:"cpuinfo"`
	CurrentKernel CurrentKernel `json:"current-kernel"`
	Idle          int           `json:"idle"`
	KSM           KSM           `json:"ksm"`
	Kversion      string        `json:"kversion"`
	LoadAvg       []string      `json:"loadavg"`
	Memory        Memory        `json:"memory"`
	PveVersion    string        `json:"pveversion"`
	RootFs        RootFS        `json:"rootfs"`
	Swap          Swap          `json:"swap"`
	Uptime        int           `json:"uptime"`
	Wait          float64       `json:"wait"`
}

// GetNodeStatus makes a GET request to the /nodes/{node}/status endpoint
// This returns more information about a node than the /nodes endpoint
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/status
func (s *NodeService) GetNodeStatus(name string) (*GetNodeStatusResponse, *http.Response, error) {
	u := fmt.Sprintf("nodes/%s/status", name)
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	d := new(GetNodeStatusResponse)
	resp, err := s.client.Do(req, d)
	if err != nil {
		return nil, resp, err
	}

	return d, resp, nil
}

// GetNodeVersionResponse contains the response for the /nodes/{node}/version endpoint
type GetNodeVersionResponse struct {
	Data GetNodeVersionData `json:"data"`
}

// GetNodeVersionData contains the version data for one node from a GetNodeVersion request
type GetNodeVersionData struct {
	Release string `json:"release"`
	RepoID  string `json:"repoid"`
	Version string `json:"version"`
}

// GetNodeVersion makes a GET request to the /nodes/{node}/version endpoint
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/version
func (s *NodeService) GetNodeVersion(name string) (*GetNodeVersionResponse, *http.Response, error) {
	u := fmt.Sprintf("nodes/%s/version", name)
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	d := new(GetNodeVersionResponse)
	resp, err := s.client.Do(req, d)
	if err != nil {
		return nil, resp, err
	}

	return d, resp, nil
}

// GetNodeQemuResponse contains the response for the /nodes/{node}/qemu endpoint
type GetNodeQemuResponse struct {
	Data []GetNodeQemuData `json:"data"`
}

// GetNodeQemuData contains data of one VM from a GetNodeQemu response
type GetNodeQemuData struct {
	CPU       float64     `json:"cpu"`
	CPUs      int         `json:"cpus"`
	Disk      int         `json:"disk"`
	DiskRead  int         `json:"diskread"`
	DiskWrite int         `json:"diskwrite"`
	MaxDisk   int         `json:"maxdisk"`
	MaxMem    int         `json:"maxmem"`
	Mem       int         `json:"mem"`
	Name      string      `json:"name"`
	NetIn     int         `json:"netin"`
	NetOut    int         `json:"netout"`
	PID       int         `json:"pid"`
	Status    string      `json:"status"`
	Uptime    int         `json:"uptime"`
	VMID      IntOrString `json:"vmid"`
}

// GetNodeQemu makes a GET request to the /nodes/{node}/qemu endpoint
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/qemu
func (s *NodeService) GetNodeQemu(name string) (*GetNodeQemuResponse, *http.Response, error) {
	u := fmt.Sprintf("nodes/%s/qemu", name)
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	d := new(GetNodeQemuResponse)
	resp, err := s.client.Do(req, d)
	if err != nil {
		return nil, resp, err
	}

	return d, resp, nil
}

// GetNodeLxcResponse contains the response for the /nodes/{node}/lxc endpoint
type GetNodeLxcResponse struct {
	Data []GetNodeLxcData `json:"data"`
}

// GetNodeLxcData contains data of one VM from a GetNodeLxc response
type GetNodeLxcData struct {
	CPU       float64     `json:"cpu"`
	CPUs      int         `json:"cpus"`
	Disk      int         `json:"disk"`
	DiskRead  int         `json:"diskread"`
	DiskWrite int         `json:"diskwrite"`
	MaxDisk   int         `json:"maxdisk"`
	MaxMem    int         `json:"maxmem"`
	MaxSwap   int         `json:"maxswap"`
	Mem       int         `json:"mem"`
	Name      string      `json:"name"`
	NetIn     int         `json:"netin"`
	NetOut    int         `json:"netout"`
	Status    string      `json:"status"`
	Type      string      `json:"type"`
	Uptime    int         `json:"uptime"`
	VMID      IntOrString `json:"vmid"`
}

// GetNodeLxc makes a GET request to the /nodes/{node}/lxc endpoint
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/lxc
func (s *NodeService) GetNodeLxc(name string) (*GetNodeLxcResponse, *http.Response, error) {
	u := fmt.Sprintf("nodes/%s/lxc", name)
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	d := new(GetNodeLxcResponse)
	resp, err := s.client.Do(req, d)
	if err != nil {
		return nil, resp, err
	}

	return d, resp, nil
}

// GetNodeDisksListResponse contains the response for the /nodes/{node}/disks/list endpoint
type GetNodeDisksListResponse struct {
	Data []GetNodeDisksListData `json:"data"`
}

// GetNodeDisksListData contains data of disks from a GetNodeDisksList response
type GetNodeDisksListData struct {
	ByIDLink     string      `json:"by_id_link"`
	DevPath      string      `json:"devpath"`
	GPT          int         `json:"gpt"`
	Health       string      `json:"health"`
	Model        string      `json:"model"`
	RPM          IntOrString `json:"rpm"`
	Serial       string      `json:"serial"`
	Size         int         `json:"size"`
	Type         string      `json:"type"`
	Used         string      `json:"used"`
	Vendor       string      `json:"vendor"`
	WWN          string      `json:"wwn"`
	Wearout      IntOrString `json:"wearout"`
	Bluestore    int         `json:"bluestore,omitempty"`
	OSDEncrypted int         `json:"osdencrypted,omitempty"`
}

// GetNodeDisksList makes a GET request to the /nodes/{node}/disks/list endpoint
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/disks/list
func (s *NodeService) GetNodeDisksList(name string) (*GetNodeDisksListResponse, *http.Response, error) {
	u := fmt.Sprintf("nodes/%s/disks/list", name)
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	d := new(GetNodeDisksListResponse)
	resp, err := s.client.Do(req, d)
	if err != nil {
		return nil, resp, err
	}

	return d, resp, nil
}

// GetNodeCertificatesInfoResponse contains the response for the /nodes/{node}/disks/list endpoint
type GetNodeCertificatesInfoResponse struct {
	Data []GetNodeCertificatesInfoData `json:"data"`
}

// GetNodeCertificatesInfoData contains data of certificates from a GetNodeCertificatesInfo response
type GetNodeCertificatesInfoData struct {
	Filename      string   `json:"filename"`
	Fingerprint   string   `json:"fingerprint"`
	Issuer        string   `json:"issuer"`
	NotAfter      int      `json:"notafter"`
	NotBefore     int      `json:"notbefore"`
	PEM           string   `json:"pem"`
	PublicKeyBits int      `json:"public-key-bits"`
	PublicKeyType string   `json:"public-key-type"`
	San           []string `json:"san"`
	Subject       string   `json:"subject"`
}

// GetNodeCertificatesInfo makes a GET request to the /nodes/{node}/certificates/info endpoint
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/certificates/info
func (s *NodeService) GetNodeCertificatesInfo(name string) (*GetNodeCertificatesInfoResponse, *http.Response, error) {
	u := fmt.Sprintf("nodes/%s/certificates/info", name)
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	d := new(GetNodeCertificatesInfoResponse)
	resp, err := s.client.Do(req, d)
	if err != nil {
		return nil, resp, err
	}

	return d, resp, nil
}

// GetNodeStorageResponse contains the response for the /nodes/{node}/storage endpoint
type GetNodeStorageResponse struct {
	Data []GetNodeStorageData `json:"data"`
}

// GetNodeStorageData contains data of certificates from a GetNodeStorage response
type GetNodeStorageData struct {
	Active       int     `json:"active"`
	Avail        int     `json:"avail"`
	Content      string  `json:"content"`
	Enabled      int     `json:"enabled"`
	Shared       int     `json:"shared"`
	Storage      string  `json:"storage"`
	Total        int     `json:"total"`
	Type         string  `json:"type"`
	Used         int     `json:"used"`
	UsedFraction float64 `json:"used_fraction"`
}

// GetNodeStorage makes a GET request to the /nodes/{node}/storage endpoint
// https://pve.proxmox.com/pve-docs/api-viewer/index.html#/nodes/{node}/storage
func (s *NodeService) GetNodeStorage(name string) (*GetNodeStorageResponse, *http.Response, error) {
	u := fmt.Sprintf("nodes/%s/storage", name)
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	d := new(GetNodeStorageResponse)
	resp, err := s.client.Do(req, d)
	if err != nil {
		return nil, resp, err
	}

	return d, resp, nil
}
