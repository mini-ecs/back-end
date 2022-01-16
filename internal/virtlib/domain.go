package virtlib

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/digitalocean/go-libvirt"
	"github.com/digitalocean/go-libvirt/socket/dialers"
	"net"
	"strconv"
	"time"
)

/*
	该包提供的功能：
	1. Domain管理。
		- 获取Domain的IP
		- 获取Domain的VNC端口
		- 制作快照（内存快照、磁盘快照）
		- 恢复快照
		- 删除快照
		- 列表快照
		- 获取Domain的信息，如CPU、内存资源状态
		- 创建Domain
		- 删除Domain
	2. 镜像文件管理-可替换（先实现本机存储,对用户透明）
		- iso等镜像
		- qcow2等磁盘镜像
	3. 节点管理（宿主机状态）
		- 查询宿主机的CPU、内存资源及其他状态
*/

type Lib struct {
	ip   net.IP
	port string
	con  *libvirt.Libvirt
}

func New(ip net.IP, port string) (*Lib, error) {
	_, err := strconv.Atoi(port)
	if err != nil {
		return nil, errors.New("port input error")
	}
	return &Lib{
		ip:   ip,
		port: port,
	}, nil
}

func (l *Lib) Connect() error {
	// todo: ip, port check
	c, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", l.ip.String(), l.port), 2*time.Second)
	if err != nil {
		return err
	}

	l.con = libvirt.NewWithDialer(dialers.NewAlreadyConnected(c))

	if err := l.con.ConnectToURI(libvirt.QEMUSystem); err != nil {
		return err
	}
	return nil
}

func (l *Lib) DisConnect() error {
	if l.con != nil {
		return l.con.Disconnect()
	}
	return errors.New("libvirt pointer is nil")
}

func (l *Lib) GetDomainByName(name string) (libvirt.Domain, error) {
	domain, err := l.con.DomainLookupByName(name)
	if err != nil {
		return libvirt.Domain{}, err
	}
	return domain, nil
}

var intToStateMap = map[libvirt.DomainState]string{
	0: "DomainNoState",
	1: "DomainRunning",
	2: "DomainBlocked",
	3: "DomainPaused",
	4: "DomainShutdown",
	5: "DomainShutoff",
	6: "DomainCrashed",
	7: "DomainPmSuspended",
}

func (l *Lib) GetDomainState(domain libvirt.Domain) (DomainState, error) {
	state, maxMemory, mem, vcpu, cpu, err := l.con.DomainGetInfo(domain)
	if err != nil {
		return DomainState{}, err
	}
	ds := libvirt.DomainState(state)
	d := DomainState{
		State:     intToStateMap[ds],
		MaxMemory: uint32(maxMemory),
		Memory:    uint32(mem),
		VirCPU:    uint32(vcpu),
		CPUTime:   uint32(cpu),
	}
	return d, nil
}

func (l *Lib) GetDomainXML(domain libvirt.Domain) Domain {
	desc, err := l.con.DomainGetXMLDesc(domain, 0)
	if err != nil {
		panic(err)
	}
	var descDomain Domain
	err = xml.Unmarshal([]byte(desc), &descDomain)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", descDomain)
	return descDomain
}

func (l *Lib) getDomainMac(domain libvirt.Domain) string {
	d := l.GetDomainXML(domain)
	return d.Devices.Interface.Mac.Address
}
func (l *Lib) getDomainBridgeName(domain libvirt.Domain) string {
	d := l.GetDomainXML(domain)
	return d.Devices.Interface.Source.Bridge
}

func (l *Lib) GetDomainVNCPort(domain libvirt.Domain) string {
	d := l.GetDomainXML(domain)
	return d.Devices.Graphics.Port
}

func (l *Lib) CreateSnapshot(domain libvirt.Domain) {
	//l.con.DomainSnapshotCreateXML(domain)
}

func (l *Lib) CreateDomain(opt DomainCreateOpt) error {
	xmls, err := xml.Marshal(opt)
	if err != nil {
		return err
	}
	//err = ioutil.WriteFile("create.xml", xmls, 0666)
	//if err != nil {
	//	return err
	//}
	domain, err := l.con.DomainCreateXML(string(xmls), 0)
	if err != nil {
		return err
	}
	fmt.Println(domain)
	return nil
}

func (l *Lib) StartDomain(name string) error {
	d, err := l.GetDomainByName(name)
	if err != nil {
		return err
	}
	err = l.con.DomainCreate(d)
	if err != nil {
		return err
	}
	return nil
}

// ShutdownDomain 关机
func (l *Lib) ShutdownDomain(name string) error {
	d, err := l.GetDomainByName(name)
	if err != nil {
		return err
	}
	err = l.con.DomainShutdown(d)
	if err != nil {
		return err
	}
	return nil
}

// DestroyDomain 删除
func (l *Lib) DestroyDomain(name string) error {
	d, err := l.GetDomainByName(name)
	if err != nil {
		return err
	}
	err = l.con.DomainDestroy(d)
	if err != nil {
		return err
	}
	return nil
}

// RebootDomain 重启
func (l *Lib) RebootDomain(name string) error {
	d, err := l.GetDomainByName(name)
	if err != nil {
		return err
	}
	// libvirt.DomainRebootFlagValues
	err = l.con.DomainReboot(d, 0)
	if err != nil {
		return err
	}
	return nil
}

// SuspendDomain 挂起
func (l *Lib) SuspendDomain(name string) error {
	d, err := l.GetDomainByName(name)
	if err != nil {
		return err
	}
	err = l.con.DomainSuspend(d)
	if err != nil {
		return err
	}
	return nil
}

// ResumeDomain 恢复（挂起的反动作）
func (l *Lib) ResumeDomain(name string) error {
	d, err := l.GetDomainByName(name)
	if err != nil {
		return err
	}
	err = l.con.DomainResume(d)
	if err != nil {
		return err
	}
	return nil
}

// GetDomainIP error
func (l *Lib) GetDomainIP(domain libvirt.Domain) {
	mac := l.getDomainBridgeName(domain)
	inter, err := l.con.InterfaceLookupByName(mac)
	if err != nil {
		panic(err)
	}
	desc, err := l.con.InterfaceGetXMLDesc(inter, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(desc)

}

// GetDomainIPAddress 不知道为什么会返回一个数组
func (l *Lib) GetDomainIPAddress(d libvirt.Domain) []libvirt.DomainInterface {
	addresses, err := l.con.DomainInterfaceAddresses(d, uint32(libvirt.DomainInterfaceAddressesSrcLease), 0)
	if err != nil {
		panic(err)
	}
	return addresses
}

func (l *Lib) GetAllInterfaces() {
	interfaces, _, _ := l.con.ConnectListAllInterfaces(10, 0)
	//l.con.DomainInterfaceAddresses()
	listInterfaces, err := l.con.ConnectListInterfaces(10)
	if err != nil {
		return
	}
	d, err := l.GetDomainByName("ubuntu")
	addresses, err := l.con.DomainInterfaceAddresses(d, uint32(libvirt.DomainInterfaceAddressesSrcLease), 0)
	if err != nil {
		panic(err)
	}
	//l.con.InterfaceLookupByMacString()
	//l.con.DomainSnapshotCurrent()
	fmt.Printf("%+v \n\n %+v\n\n %+v", interfaces, listInterfaces, addresses)
}
