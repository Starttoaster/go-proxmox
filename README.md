# go-proxmox

[![Go Report Card](https://goreportcard.com/badge/github.com/starttoaster/go-proxmox)](https://goreportcard.com/report/github.com/starttoaster/go-proxmox) [![codecov](https://codecov.io/gh/Starttoaster/go-proxmox/graph/badge.svg?token=UFMXVXNKL8)](https://codecov.io/gh/Starttoaster/go-proxmox) [![Go Reference](https://pkg.go.dev/badge/github.com/starttoaster/go-proxmox.svg)](https://pkg.go.dev/github.com/starttoaster/go-proxmox)

This is an API client library for Proxmox VE servers. It aims to be simple to use and consume in your own Go programs, make very little assumptions about how the user would consume it, and use as few non-stdlib dependencies to do so as possible.

This is currently tested against Proxmox VE 8.x systems.

This library is in its early development phase. Minor changes may be made in its usage, and only a portion of API methods are currently supported. See the CONTRIBUTING.md for details on contributing new methods to this library. Or make an Issue to discuss it.

## Usage

This API client library currently supports API tokens for authentication.

```go
import proxmox "github.com/starttoaster/go-proxmox"

// Create a new API client using a Proxmox API token
c, _ := proxmox.NewClient(tokenID, token, proxmox.WithBaseURL("https://10.0.0.10:8006/"))

// Retrieve cluster nodes
nodes, _, _ := c.Nodes.GetNodes()

// Retrieve the status of a node named server1
node, _, _ := c.Nodes.GetNodeStatus("server1")

// Retrieve the version of a node named server1
version, _, _ := c.Nodes.GetNodeVersion("server1")
```

### Insecure API servers

If your PVE server's TLS can't be verified, you can pass an insecure HTTP client to the library.

```go
httpClient := http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
}

c, _ := proxmox.NewClient(tokenID, token, 
    proxmox.WithBaseURL("https://10.0.0.10:8006/"), 
    proxmox.WithHTTPClient(&httpClient),
)
```
