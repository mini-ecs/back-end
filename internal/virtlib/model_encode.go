package virtlib

import "encoding/xml"

type __memo struct {
	Text string `xml:",chardata"`
	Unit string `xml:"unit,attr"`
}
type __currentMemory struct {
	Text string `xml:",chardata"`
	Unit string `xml:"unit,attr"`
}
type __type struct {
	Text    string `xml:",chardata"`
	Arch    string `xml:"arch,attr"`
	Machine string `xml:"machine,attr"`
}
type __boot struct {
	Text string `xml:",chardata"`
	Dev  string `xml:"dev,attr"`
}
type __os struct {
	Text string   `xml:",chardata"`
	Type __type   `xml:"type"`
	Boot []__boot `xml:"boot"`
}
type __source struct {
	Text string `xml:",chardata"`
	File string `xml:"file,attr"`
}
type __target struct {
	Text string `xml:",chardata"`
	Dev  string `xml:"dev,attr"`
}
type __driver struct {
	Text string `xml:",chardata"`
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
}
type __disk struct {
	Text     string    `xml:",chardata"`
	Type     string    `xml:"type,attr"`
	Device   string    `xml:"device,attr"`
	Source   __source  `xml:"source"`
	Target   __target  `xml:"target"`
	Readonly string    `xml:"readonly,omitempty"`
	Driver   *__driver `xml:"driver,omitempty"`
}
type __interSource struct {
	Text    string `xml:",chardata"`
	Network string `xml:"network,attr"`
}
type __interface struct {
	Text   string        `xml:",chardata"`
	Type   string        `xml:"type,attr"`
	Source __interSource `xml:"source"`
}
type __graphics struct {
	Text   string `xml:",chardata"`
	Type   string `xml:"type,attr"`
	Port   string `xml:"port,attr"`
	Listen string `xml:"listen,attr"`
	Passwd string `xml:"passwd,attr"`
	Keymap string `xml:"keymap,attr"`
}
type __serTarget struct {
	Text string `xml:",chardata"`
	Port string `xml:"port,attr"`
}
type __serial struct {
	Text   string      `xml:",chardata"`
	Type   string      `xml:"type,attr"`
	Target __serTarget `xml:"target"`
}
type __conTarget struct {
	Text string `xml:",chardata"`
	Type string `xml:"type,attr"`
	Port string `xml:"port,attr"`
}
type __console struct {
	Text   string      `xml:",chardata"`
	Type   string      `xml:"type,attr"`
	Target __conTarget `xml:"target"`
}
type __devices struct {
	Text      string      `xml:",chardata"`
	Emulator  string      `xml:"emulator"`
	Disk      []__disk    `xml:"disk"`
	Interface __interface `xml:"interface"`
	Graphics  __graphics  `xml:"graphics"`
	Serial    __serial    `xml:"serial"`
	Console   __console   `xml:"console"`
}
type DomainCreateOpt struct {
	XMLName       xml.Name        `xml:"domain"`
	Text          string          `xml:",chardata"`
	Type          string          `xml:"type,attr"`
	Name          string          `xml:"name"`
	Uuid          string          `xml:"uuid"`
	Memory        __memo          `xml:"memory"`
	CurrentMemory __currentMemory `xml:"currentMemory"`
	Vcpu          string          `xml:"vcpu"`
	Os            __os            `xml:"os"`
	Devices       __devices       `xml:"devices"`
}

var DefaultOpt = DomainCreateOpt{
	Type: "kvm",
	Name: "ubuntu",
	Uuid: "c7a5fdbd-cdaf-9455-926a-d65c16db1809",
	Memory: struct {
		Text string `xml:",chardata"`
		Unit string `xml:"unit,attr"`
	}{
		Text: "4",
		Unit: "GiB",
	},
	Vcpu: "4",
	Os: __os{
		Type: struct {
			Text    string `xml:",chardata"`
			Arch    string `xml:"arch,attr"`
			Machine string `xml:"machine,attr"`
		}{
			Text:    "hvm",
			Arch:    "x86_64",
			Machine: "pc",
		},
		Boot: []__boot{
			{Dev: "hd"},
			{Dev: "cdrom"},
		},
	},
	Devices: __devices{
		Emulator: "/usr/bin/qemu-system-x86_64",
		Disk: []__disk{{
			Type:   "file",
			Device: "cdrom",
			Source: __source{
				File: "/home/fangaoyang/work/ubuntu-20.04.3-desktop-amd64.iso",
			},
			Target: __target{
				Dev: "hdc",
			},
		}, {
			Type:   "file",
			Device: "disk",
			Driver: &__driver{
				Name: "qemu",
				Type: "qcow2",
			},
			Source: __source{
				File: "/home/fangaoyang/work/desktop.qcow2",
			},
			Target: __target{
				Dev: "hda",
			},
		}},
		Interface: __interface{
			Type:   "network",
			Source: __interSource{Network: "default"},
		},
		Graphics: __graphics{
			Type:   "vnc",
			Port:   "-1",
			Listen: "0.0.0.0",
			Passwd: "123",
			Keymap: "en-us",
		},
		Serial: __serial{
			Type: "pty",
			Target: __serTarget{
				Port: "0",
			},
		},
		Console: __console{
			Type: "pty",
			Target: __conTarget{
				Type: "serial",
				Port: "0",
			},
		},
	},
}
