# Contributing

## New API services

This library uses a Service pattern for different paths in the PVE API, like `ClusterService` for `/api2/json/cluster`.

It currently supports API methods contained under the following paths (assume an `/api2/json` prefix):

* `/cluster`
* `/nodes`

 If you would like to add an API method from a different API path, you will need to add a new client Service, [like so.](https://github.com/Starttoaster/go-proxmox/blob/fe6f9b739155dcf713694320e790ab945dab6215/client.go#L31) Then you can create a new file for your service that contains your API method.

## New API methods

[For an example of how to add an API method.](https://github.com/Starttoaster/go-proxmox/blob/fe6f9b739155dcf713694320e790ab945dab6215/nodes.go#L37) Important to note the structs just above the method that contain the data the method returns to the user with appropriate types. This library currently does not make use of generics, so optional fields in the API response should be pointers. Also note the comments just above the structs and functions, please follow the same convention in yoru contribution.

Please try to help keep up with this library's tests by contributing a test for your API method contributions. See [this test](https://github.com/Starttoaster/go-proxmox/blob/fe6f9b739155dcf713694320e790ab945dab6215/nodes_test.go#L11) for an example, and the [json text file that the mock API server's router returns.](https://github.com/Starttoaster/go-proxmox/blob/fe6f9b739155dcf713694320e790ab945dab6215/testdata/nodes/get_nodes.json)