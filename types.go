package proxmox

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// BootInfo info about host boot
type BootInfo struct {
	Mode       string `json:"mode"`
	SecureBoot int    `json:"secureboot"`
}

// CPUInfo info about host CPU
type CPUInfo struct {
	Cores   int    `json:"cores"`
	CPUs    int    `json:"cpus"`
	Flags   string `json:"flags"`
	HVM     string `json:"hvm"`
	MHz     string `json:"mhz"`
	Model   string `json:"model"`
	Sockets int    `json:"sockets"`
	UserHz  int    `json:"user_hz"`
}

// CurrentKernel info about host kernel
type CurrentKernel struct {
	Machine string `json:"machine"`
	Release string `json:"release"`
	SysName string `json:"sysname"`
	Version string `json:"version"`
}

// KSM info about Kernel same-page merging
type KSM struct {
	Shared int `json:"shared"`
}

// Memory info about host memory
type Memory struct {
	Free  int `json:"free"`
	Total int `json:"total"`
	Used  int `json:"used"`
}

// RootFS info about the host root filesystem
type RootFS struct {
	Avail int `json:"avail"`
	Free  int `json:"free"`
	Total int `json:"total"`
	Used  int `json:"used"`
}

// Swap info about swap
type Swap struct {
	Free  int `json:"free"`
	Total int `json:"total"`
	Used  int `json:"used"`
}

// IntOrString is an alias for some returns from the Proxmox API where we've identified that some versions return a string, and others return an integer
// For example, LXC VMIDs were a string return in PVE 8.1.x, and are an integer in 8.2.x
// Since it can be either depending on the version of Proxmox queried, we're going to return the looser type. String in this case.
type IntOrString string

// UnmarshalJSON implements the json.Unmarshaler interface for IntOrString types
func (is *IntOrString) UnmarshalJSON(data []byte) error {
	// attempt to unmarshal into an integer
	var i int
	var intErr error
	if intErr = json.Unmarshal(data, &i); intErr == nil {
		*is = IntOrString(strconv.Itoa(i))
		return nil
	}

	// attempt to unmarshal into a string
	var str string
	var strErr error
	if strErr = json.Unmarshal(data, &str); strErr == nil {
		*is = IntOrString(str)
		return nil
	}

	return fmt.Errorf("failed to unmarshal as either int or string: int err: %s | string err: %s", intErr.Error(), strErr.Error())
}
