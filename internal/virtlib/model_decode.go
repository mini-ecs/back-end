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
