package virtlib

import "encoding/xml"

type memory struct {
	Text string `xml:",chardata"`
	Unit string `xml:"unit,attr,omitempty"`
}
type currentMemory struct {
	Text string `xml:",chardata"`
	Unit string `xml:"unit,attr,omitempty"`
}
type vcpu struct {
	Text      string `xml:",chardata"`
	Placement string `xml:"placement,attr,omitempty"`
}
type resource struct {
	Text      string `xml:",chardata"`
	Partition string `xml:"partition,omitempty"`
}
type _type struct {
	Text    string `xml:",chardata"`
	Arch    string `xml:"arch,attr,omitempty"`
	Machine string `xml:"machine,attr,omitempty"`
}
type boot struct {
	Text string `xml:",chardata"`
	Dev  string `xml:"dev,attr,omitempty"`
}
type os struct {
	Type _type  `xml:"type,omitempty"`
	Boot []boot `xml:"boot,omitempty"`
}
type model struct {
	Fallback string `xml:"fallback,attr,omitempty"`
}
type feature struct {
	Policy string `xml:"policy,attr,omitempty"`
	Name   string `xml:"name,attr,omitempty"`
}
type cpu struct {
	Mode    string    `xml:"mode,attr,omitempty"`
	Match   string    `xml:"match,attr,omitempty"`
	Check   string    `xml:"check,attr,omitempty"`
	Model   model     `xml:"model,omitempty"`
	Feature []feature `xml:"feature,omitempty"`
}
type clock struct {
	Offset string `xml:"offset,attr,omitempty"`
}
type driver struct {
	Name string `xml:"name,attr,omitempty"`
	Type string `xml:"type,attr,omitempty"`
}
type source struct {
	File  string `xml:"file,attr,omitempty"`
	Index string `xml:"index,attr,omitempty"`
}
type target struct {
	Dev string `xml:"dev,attr,omitempty"`
	Bus string `xml:"bus,attr,omitempty"`
}

type alias struct {
	Name string `xml:"name,attr,omitempty"`
}
type address struct {
	Type       string `xml:"type,attr,omitempty"`
	Controller string `xml:"controller,attr,omitempty"`
	Bus        string `xml:"bus,attr,omitempty"`
	Target     string `xml:"target,attr,omitempty"`
	Unit       string `xml:"unit,attr,omitempty"`
}
type disk struct {
	Type         string  `xml:"type,attr,omitempty"`
	Device       string  `xml:"device,attr,omitempty"`
	Driver       driver  `xml:"driver,omitempty"`
	Source       source  `xml:"source,omitempty"`
	BackingStore string  `xml:"backingStore,omitempty"`
	Target       target  `xml:"target,omitempty"`
	Alias        alias   `xml:"alias,omitempty"`
	Address      address `xml:"address,omitempty"`
	Readonly     string  `xml:"readonly,omitempty"`
}
type controller struct {
	Type    string `xml:"type,attr,omitempty"`
	Index   string `xml:"index,attr,omitempty"`
	Model   string `xml:"model,attr"`
	Alias   alias  `xml:"alias"`
	Address struct {
		Type     string `xml:"type,attr"`
		Domain   string `xml:"domain,attr"`
		Bus      string `xml:"bus,attr"`
		Slot     string `xml:"slot,attr"`
		Function string `xml:"function,attr"`
	} `xml:"address"`
}
type mac struct {
	Address string `xml:"address,attr"`
}
type interfaceSource struct {
	Network string `xml:"network,attr"`
	Portid  string `xml:"portid,attr"`
	Bridge  string `xml:"bridge,attr"`
}
type interfaceModel struct {
	Type string `xml:"type,attr"`
}
type interfaceAdress struct {
	Type     string `xml:"type,attr"`
	Domain   string `xml:"domain,attr"`
	Bus      string `xml:"bus,attr"`
	Slot     string `xml:"slot,attr"`
	Function string `xml:"function,attr"`
}
type _interface struct {
	Type    string          `xml:"type,attr"`
	Mac     mac             `xml:"mac"`
	Source  interfaceSource `xml:"source"`
	Target  target          `xml:"target"`
	Model   interfaceModel  `xml:"model"`
	Alias   alias           `xml:"alias"`
	Address interfaceAdress `xml:"address"`
}
type serialSource struct {
	Path string `xml:"path,attr"`
}
type _model struct {
	Name string `xml:"name,attr"`
}
type serialTarget struct {
	Type  string `xml:"type,attr"`
	Port  string `xml:"port,attr"`
	Model _model `xml:"model"`
}
type serial struct {
	Type   string       `xml:"type,attr"`
	Source serialSource `xml:"source"`
	Target serialTarget `xml:"target"`
	Alias  alias        `xml:"alias"`
}
type console struct {
	Type   string       `xml:"type,attr"`
	Tty    string       `xml:"tty,attr"`
	Source serialSource `xml:"source"`
	Target serialTarget `xml:"target"`
	Alias  alias        `xml:"alias"`
}
type input struct {
	Type  string `xml:"type,attr"`
	Bus   string `xml:"bus,attr"`
	Alias alias  `xml:"alias"`
}
type listen struct {
	Type    string `xml:"type,attr"`
	Address string `xml:"address,attr"`
}
type graphics struct {
	Type       string `xml:"type,attr"`
	Port       string `xml:"port,attr"`
	Autoport   string `xml:"autoport,attr"`
	AttrListen string `xml:"listen,attr"`
	Keymap     string `xml:"keymap,attr"`
	Listen     listen `xml:"listen"`
}
type video struct {
	Model struct {
		Type    string `xml:"type,attr"`
		Vram    string `xml:"vram,attr"`
		Heads   string `xml:"heads,attr"`
		Primary string `xml:"primary,attr"`
	} `xml:"model"`
	Alias struct {
		Name string `xml:"name,attr"`
	} `xml:"alias"`
	Address struct {
		Type     string `xml:"type,attr"`
		Domain   string `xml:"domain,attr"`
		Bus      string `xml:"bus,attr"`
		Slot     string `xml:"slot,attr"`
		Function string `xml:"function,attr"`
	} `xml:"address"`
}
type memballoon struct {
	Model string `xml:"model,attr"`
	Alias struct {
		Name string `xml:"name,attr"`
	} `xml:"alias"`
	Address struct {
		Type     string `xml:"type,attr"`
		Domain   string `xml:"domain,attr"`
		Bus      string `xml:"bus,attr"`
		Slot     string `xml:"slot,attr"`
		Function string `xml:"function,attr"`
	} `xml:"address"`
}
type devices struct {
	Emulator   string       `xml:"emulator,omitempty"`
	Disk       []disk       `xml:"disk,omitempty"`
	Controller []controller `xml:"controller,omitempty"`
	Interface  _interface   `xml:"interface,omitempty"`
	Serial     serial       `xml:"serial,omitempty"`
	Console    console      `xml:"console,omitempty"`
	Input      []input      `xml:"input,omitempty"`
	Graphics   graphics     `xml:"graphics,omitempty"`
	Video      video        `xml:"video,omitempty"`
	Memballoon memballoon   `xml:"memballoon,omitempty"`
}
type seclabel struct {
	Type       string `xml:"type,attr"`
	Model      string `xml:"model,attr"`
	Relabel    string `xml:"relabel,attr"`
	Label      string `xml:"label"`
	Imagelabel string `xml:"imagelabel"`
}
type Domain struct {
	XMLName       xml.Name      `xml:"domain"`
	Type          string        `xml:"type,attr"`
	ID            string        `xml:"id,attr"`
	Name          string        `xml:"name"`
	Uuid          string        `xml:"uuid"`
	Memory        memory        `xml:"memory"`
	CurrentMemory currentMemory `xml:"currentMemory"`
	Vcpu          vcpu          `xml:"vcpu"`
	Resource      resource      `xml:"resource"`
	Os            os            `xml:"os"`
	Cpu           cpu           `xml:"cpu"`
	Clock         clock         `xml:"clock"`
	OnPoweroff    string        `xml:"on_poweroff"`
	OnReboot      string        `xml:"on_reboot"`
	OnCrash       string        `xml:"on_crash"`
	Devices       devices       `xml:"devices"`
	Seclabel      seclabel      `xml:"seclabel"`
}

type DomainState struct {
	State     string
	MaxMemory uint32
	Memory    uint32
	VirCPU    uint32
	CPUTime   uint32
}

type Capability struct {
	XMLName xml.Name `xml:"capabilities"`
	Text    string   `xml:",chardata"`
	Host    struct {
		Text string `xml:",chardata"`
		Uuid string `xml:"uuid"`
		Cpu  struct {
			Text      string `xml:",chardata"`
			Arch      string `xml:"arch"`
			Model     string `xml:"model"`
			Vendor    string `xml:"vendor"`
			Microcode struct {
				Text    string `xml:",chardata"`
				Version string `xml:"version,attr"`
			} `xml:"microcode"`
			Counter struct {
				Text      string `xml:",chardata"`
				Name      string `xml:"name,attr"`
				Frequency string `xml:"frequency,attr"`
				Scaling   string `xml:"scaling,attr"`
			} `xml:"counter"`
			Topology struct {
				Text    string `xml:",chardata"`
				Sockets string `xml:"sockets,attr"`
				Cores   string `xml:"cores,attr"`
				Threads string `xml:"threads,attr"`
			} `xml:"topology"`
			Feature []struct {
				Text string `xml:",chardata"`
				Name string `xml:"name,attr"`
			} `xml:"feature"`
			Pages []struct {
				Text string `xml:",chardata"`
				Unit string `xml:"unit,attr"`
				Size string `xml:"size,attr"`
			} `xml:"pages"`
		} `xml:"cpu"`
		PowerManagement struct {
			Text          string `xml:",chardata"`
			SuspendMem    string `xml:"suspend_mem"`
			SuspendDisk   string `xml:"suspend_disk"`
			SuspendHybrid string `xml:"suspend_hybrid"`
		} `xml:"power_management"`
		Iommu struct {
			Text    string `xml:",chardata"`
			Support string `xml:"support,attr"`
		} `xml:"iommu"`
		MigrationFeatures struct {
			Text          string `xml:",chardata"`
			Live          string `xml:"live"`
			URITransports struct {
				Text         string   `xml:",chardata"`
				URITransport []string `xml:"uri_transport"`
			} `xml:"uri_transports"`
		} `xml:"migration_features"`
		Topology struct {
			Text  string `xml:",chardata"`
			Cells struct {
				Text string `xml:",chardata"`
				Num  string `xml:"num,attr"`
				Cell struct {
					Text   string `xml:",chardata"`
					ID     string `xml:"id,attr"`
					Memory struct {
						Text string `xml:",chardata"`
						Unit string `xml:"unit,attr"`
					} `xml:"memory"`
					Pages []struct {
						Text string `xml:",chardata"`
						Unit string `xml:"unit,attr"`
						Size string `xml:"size,attr"`
					} `xml:"pages"`
					Distances struct {
						Text    string `xml:",chardata"`
						Sibling struct {
							Text  string `xml:",chardata"`
							ID    string `xml:"id,attr"`
							Value string `xml:"value,attr"`
						} `xml:"sibling"`
					} `xml:"distances"`
					Cpus struct {
						Text string `xml:",chardata"`
						Num  string `xml:"num,attr"`
						Cpu  []struct {
							Text     string `xml:",chardata"`
							ID       string `xml:"id,attr"`
							SocketID string `xml:"socket_id,attr"`
							CoreID   string `xml:"core_id,attr"`
							Siblings string `xml:"siblings,attr"`
						} `xml:"cpu"`
					} `xml:"cpus"`
				} `xml:"cell"`
			} `xml:"cells"`
		} `xml:"topology"`
		Cache struct {
			Text string `xml:",chardata"`
			Bank struct {
				Text  string `xml:",chardata"`
				ID    string `xml:"id,attr"`
				Level string `xml:"level,attr"`
				Type  string `xml:"type,attr"`
				Size  string `xml:"size,attr"`
				Unit  string `xml:"unit,attr"`
				Cpus  string `xml:"cpus,attr"`
			} `xml:"bank"`
		} `xml:"cache"`
		Secmodel []struct {
			Text      string `xml:",chardata"`
			Model     string `xml:"model"`
			Doi       string `xml:"doi"`
			Baselabel []struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"baselabel"`
		} `xml:"secmodel"`
	} `xml:"host"`
	Guest []struct {
		Text   string `xml:",chardata"`
		OsType string `xml:"os_type"`
		Arch   struct {
			Text     string `xml:",chardata"`
			Name     string `xml:"name,attr"`
			Wordsize string `xml:"wordsize"`
			Emulator string `xml:"emulator"`
			Machine  []struct {
				Text      string `xml:",chardata"`
				MaxCpus   string `xml:"maxCpus,attr"`
				Canonical string `xml:"canonical,attr"`
			} `xml:"machine"`
			Domain []struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"domain"`
		} `xml:"arch"`
		Features struct {
			Text         string `xml:",chardata"`
			Cpuselection string `xml:"cpuselection"`
			Deviceboot   string `xml:"deviceboot"`
			Disksnapshot struct {
				Text    string `xml:",chardata"`
				Default string `xml:"default,attr"`
				Toggle  string `xml:"toggle,attr"`
			} `xml:"disksnapshot"`
			Acpi struct {
				Text    string `xml:",chardata"`
				Default string `xml:"default,attr"`
				Toggle  string `xml:"toggle,attr"`
			} `xml:"acpi"`
			Pae    string `xml:"pae"`
			Nonpae string `xml:"nonpae"`
			Apic   struct {
				Text    string `xml:",chardata"`
				Default string `xml:"default,attr"`
				Toggle  string `xml:"toggle,attr"`
			} `xml:"apic"`
		} `xml:"features"`
	} `xml:"guest"`
}
type Welcome4 struct {
	Capabilities Capabilities `xml:"capabilities"`
}

type Capabilities struct {
	Host  Host    `xml:"host"`
	Guest []Guest `xml:"guest"`
}

type Guest struct {
	OSType   OSType   `xml:"os_type"`
	Arch     Arch     `xml:"arch"`
	Features Features `xml:"features"`
}

type Arch struct {
	Wordsize string        `xml:"wordsize"`
	Emulator string        `xml:"emulator"`
	Machine  *MachineUnion `xml:"machine"`
	Domain   *DomainUnion  `xml:"domain"`
	Name     string        `xml:"_name"`
}

type DomainElement struct {
	Type Type `xml:"_type"`
}

type MachineElement struct {
	MaxCpus   string  `xml:"_maxCpus"`
	Text      string  `xml:"__text"`
	Canonical *string `xml:"_canonical,omitempty"`
}

type PurpleMachine struct {
	MaxCpus string `xml:"_maxCpus"`
	Text    string `xml:"__text"`
}

type Features struct {
	Cpuselection string  `xml:"cpuselection"`
	Deviceboot   string  `xml:"deviceboot"`
	Disksnapshot ACPI    `xml:"disksnapshot"`
	ACPI         *ACPI   `xml:"acpi,omitempty"`
	Pae          *string `xml:"pae,omitempty"`
	Nonpae       *string `xml:"nonpae,omitempty"`
	APIC         *ACPI   `xml:"apic,omitempty"`
}

type ACPI struct {
	Default Default `xml:"_default"`
	Toggle  Support `xml:"_toggle"`
}

type Host struct {
	UUID              string            `xml:"uuid"`
	CPU               HostCPU           `xml:"cpu"`
	PowerManagement   PowerManagement   `xml:"power_management"`
	Iommu             Iommu             `xml:"iommu"`
	MigrationFeatures MigrationFeatures `xml:"migration_features"`
	Topology          HostTopology      `xml:"topology"`
	Cache             Cache             `xml:"cache"`
	Secmodel          []Secmodel        `xml:"secmodel"`
}

type HostCPU struct {
	Arch      string      `xml:"arch"`
	Model     string      `xml:"model"`
	Vendor    string      `xml:"vendor"`
	Microcode Microcode   `xml:"microcode"`
	Counter   Counter     `xml:"counter"`
	Topology  CPUTopology `xml:"topology"`
	Feature   []Feature   `xml:"feature"`
	Pages     []CPUPage   `xml:"pages"`
}

type Counter struct {
	Name      string  `xml:"_name"`
	Frequency string  `xml:"_frequency"`
	Scaling   Support `xml:"_scaling"`
}

type Feature struct {
	Name string `xml:"_name"`
}

type Microcode struct {
	Version string `xml:"_version"`
}

type CPUPage struct {
	Unit string `xml:"_unit"`
	Size string `xml:"_size"`
}

type CPUTopology struct {
	Sockets string `xml:"_sockets"`
	Cores   string `xml:"_cores"`
	Threads string `xml:"_threads"`
}

type Cache struct {
	Bank Bank `xml:"bank"`
}

type Bank struct {
	ID    string `xml:"_id"`
	Level string `xml:"_level"`
	Type  string `xml:"_type"`
	Size  string `xml:"_size"`
	Unit  string `xml:"_unit"`
	Cpus  string `xml:"_cpus"`
}

type Iommu struct {
	Support Support `xml:"_support"`
}

type MigrationFeatures struct {
	Live          string        `xml:"live"`
	URITransports URITransports `xml:"uri_transports"`
}

type URITransports struct {
	URITransport []string `xml:"uri_transport"`
}

type PowerManagement struct {
	SuspendMem    string `xml:"suspend_mem"`
	SuspendDisk   string `xml:"suspend_disk"`
	SuspendHybrid string `xml:"suspend_hybrid"`
}

type Secmodel struct {
	Model     string      `xml:"model"`
	Doi       string      `xml:"doi"`
	Baselabel []Baselabel `xml:"baselabel,omitempty"`
}

type Baselabel struct {
	Type Type   `xml:"_type"`
	Text string `xml:"__text"`
}

type HostTopology struct {
	Cells Cells `xml:"cells"`
}

type Cells struct {
	Cell Cell   `xml:"cell"`
	Num  string `xml:"_num"`
}

type Cell struct {
	Memory    Memory     `xml:"memory"`
	Pages     []CellPage `xml:"pages"`
	Distances Distances  `xml:"distances"`
	Cpus      Cpus       `xml:"cpus"`
	ID        string     `xml:"_id"`
}

type Cpus struct {
	CPU []CPUElement `xml:"cpu"`
	Num string       `xml:"_num"`
}

type CPUElement struct {
	ID       string `xml:"_id"`
	SocketID string `xml:"_socket_id"`
	CoreID   string `xml:"_core_id"`
	Siblings string `xml:"_siblings"`
}

type Distances struct {
	Sibling Sibling `xml:"sibling"`
}

type Sibling struct {
	ID    string `xml:"_id"`
	Value string `xml:"_value"`
}

type Memory struct {
	Unit string `xml:"_unit"`
	Text string `xml:"__text"`
}

type CellPage struct {
	Unit string `xml:"_unit"`
	Size string `xml:"_size"`
	Text string `xml:"__text"`
}

type Type string

const (
	KVM  Type = "kvm"
	Qemu Type = "qemu"
)

type Default string

const (
	On Default = "on"
)

type Support string

const (
	No  Support = "no"
	Yes Support = "yes"
)

type OSType string

const (
	Hvm OSType = "hvm"
)

type DomainUnion struct {
	DomainElement      *DomainElement
	DomainElementArray []DomainElement
}

type MachineUnion struct {
	MachineElementArray []MachineElement
	PurpleMachine       *PurpleMachine
}
