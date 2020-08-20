package model

import (
	"fmt"
	"strings"
)

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
	internalIPAddressBuilder := strings.Builder{}
	publicIPAddressBuilder := strings.Builder{}

	internalIPAddressBuilder.WriteString("[")
	publicIPAddressBuilder.WriteString("[")
	for i, ip := range s.Details.IPAddresses {
		if len(ip.Internal) != 0 {
			internalIPAddressBuilder.WriteString(ip.Internal)
			if i < len(s.Details.IPAddresses)-1 {
				internalIPAddressBuilder.WriteString(", ")
			}
		}

		if len(ip.Public) != 0 {
			publicIPAddressBuilder.WriteString(ip.Public)
			if i < len(s.Details.IPAddresses)-1 {
				publicIPAddressBuilder.WriteString(", ")
			}
		}
	}
	internalIPAddressBuilder.WriteString("]")
	publicIPAddressBuilder.WriteString("]")

	diskBuilder := strings.Builder{}
	for _, d := range s.Details.DisksFromMain {
		path := "unknown"
		if len(d.PartitionPaths) > 0 {
			path = d.PartitionPaths[0]
		}
		diskBuilder.WriteString(fmt.Sprintf("\t\tDisk ID: %s\tSizeGB %5d\tPath %s\n", d.ID, d.SizeGB, path))
	}

	vSphere := "Unknown"
	if len(s.Details.ManagementLinks) != 0 {
		vSphere = s.Details.ManagementLinks[0].URI
	}

	return fmt.Sprintf(`%s
[%s]
-----------------------------------------------------------------------
OS:                          %s
Internal IP:                 %s
Public IP:                   %s
CPU:                         %d Cores
Memory:                      %d GB
PowerState:                  %s
Created By:                  %s
Created Date:                %s
vSphere:                     %s

Disks
%s
`, s.Name, s.Description, s.OSType, internalIPAddressBuilder.String(), publicIPAddressBuilder.String(),
		s.Details.CPU, s.Details.MemoryGB(), strings.ToTitle(s.Details.PowerState), s.ChangeInfo.CreatedBy,
		s.ChangeInfo.CreatedDate, vSphere, diskBuilder.String())
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
	DisksFromMain        []Disk `json:"disksFromMain"`
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
