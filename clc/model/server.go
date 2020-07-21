package model

import "fmt"

type Server struct {
	ID              string        `json:"id"`
	Name            string        `json:"name"`
	DisplayName     string        `json:"displayName"`
	Description     string        `json:"description"`
	GroupID         string        `json:"groupId"`
	IsTemplate      bool          `json:"isTemplate"`
	LocationID      string        `json:"locationId"`
	OSType          string        `json:"osType"`
	IsManagedOS     bool          `json:"isManagedOS"`
	IsManagedBackup bool          `json:"isManagedBackup"`
	Status          string        `json:"status"`
	Details         ServerDetails `json:"details"`
	Type            string        `json:"type"`
	StorageType     string        `json:"storageType"`
	ChangeInfo      ChangeInfo
	Links           []Link
}

func (s Server) String() string {
	return fmt.Sprintf(`%s
[%s]
---------------------------------------------------------
OS:             %s
IPAddresses:    %s
CPU:            %d
Memory:         %d
vSphere:        %s
`, s.Name, s.Description, s.OSType, s.Details.IPAddresses,
		s.Details.CPU, s.Details.MemoryGB(), s.Details.ManagementLinks[0].URI)
}

type ServerDetails struct {
	IPAddresses          []IPAddress
	SecondaryIPAddresses []IPAddress
	AlertPolicies        []AlertPolicy
	CPU                  int      `json:"cpu"`
	CoresPerSocket       int      `json:"corePerSocket"`
	Sockets              int      `json:"sockets"`
	DataStores           []string `json:"datastores"`
	DiskCount            int      `json:"diskCount"`
	HostName             string   `json:"hostName"`
	SourceServerName     string   `json:"sourceServerName"`
	HardwareUUID         string   `json:"hardwareUUID"`
	InMaintenanceMode    bool     `json:"inMaintenanceMode"`
	ManagementAddress    string   `json:"managementAddress"`
	ManagementLinks      []struct {
		Name string `json:"name"`
		URI  string `json:"uri"`
	} `json:"managementLinks"`
	DisksFromMain []struct {
		ID             string   `json:"id"`
		SizeGB         int      `json:"sizeGB"`
		PartitionPaths []string `json:"partitionPaths"`
	} `json:"disksFromMain"`
	MemoryMB             int    `json:"memoryMB"`
	PowerState           string `json:"powerState"`
	StorageGB            int    `json:"storageGB"`
	Disks                []Disk
	Partitions           []Partition
	Snapshots            []Snapshot
	CustomFields         []CustomField
	ProcessorDescription string `json:"processorDescription"`
	StorageDescription   string `json:"storageDescription"`
}

func (sd ServerDetails) MemoryGB() int {
	return sd.MemoryMB / 1024
}

type IPAddress struct {
	Public   string `json:"public"`
	Internal string `json:"internal"`
}

type AlertPolicy struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Links []Link
}

type Disk struct {
	ID             string `json:"id"`
	SizeGB         int    `json:"sizeGB"`
	PartitionPaths []string
}

type Partition struct {
	SizeGB int    `json:"sizeGB"`
	Path   string `json:"path"`
}

type Snapshot struct {
	Name  string `json:"name"`
	Links []Link `json:"links"`
}
