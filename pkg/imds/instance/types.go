package instance

type Metadata struct {
	Compute ComputeMetadata `json:"compute"`
	Network NetworkMetadata `json:"network"`
}

type ComputeMetadata struct {
	Name              string        `json:"name"`
	ResourceGroupName string        `json:"resourceGroupName"`
	ResourceID        string        `json:"resourceId"`
	SubscriptionID    string        `json:"subscriptionId"`
	TagsList          []TagMetadata `json:"tagsList"`
	VMID              string        `json:"vmId"`
	VMScaleSetName    string        `json:"vmScaleSetName"`
	VMSize            string        `json:"vmSize"`
	Zone              string        `json:"zone"`
}

type TagMetadata struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type NetworkMetadata struct {
	Interfaces []InterfaceMetadata `json:"interface"`
}

type InterfaceMetadata struct {
	IPv4       IPv4Metadata `json:"ipv4"`
	MacAddress string       `json:"macAddress"`
}

type IPv4Metadata struct {
	IPAddresses []IPAddressMetadata `json:"ipAddress"`
	Subnets     []SubnetMetadata    `json:"subnet"`
}

type IPAddressMetadata struct {
	PrivateIPAddress string `json:"privateIpAddress"`
	PublicIpAddress  string `json:"publicIpAddress"`
}

type SubnetMetadata struct {
	Address string `json:"address"`
	Prefix  string `json:"prefix"`
}
