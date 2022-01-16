package virtlib

import (
	"encoding/xml"
	"fmt"
	"github.com/google/uuid"
	"github.com/magiconair/properties/assert"
	"net"
	"testing"
)

var domainName = "testlib"

func GenerateEnv() *Lib {
	ip := net.ParseIP("10.249.46.250")
	l, err := New(ip, "16509")
	if err != nil {
		panic("generate env error: " + err.Error())
	}
	err = l.Connect()
	if err != nil {
		panic(err)
	}
	return l
}
func TestLib_GetInterface(t *testing.T) {
	l := GenerateEnv()
	l.GetAllInterfaces()
}

func TestLib_GetDomainIP(t *testing.T) {
	l := GenerateEnv()
	d, err := l.GetDomainByName("ubuntu")
	assert.Equal(t, err, nil)

	i1 := l.GetDomainIPAddress(d)
	fmt.Println(i1)
	l.GetDomainIP(d)

}

func TestLib_NewDomainOpt(t *testing.T) {
	const opt = `<domain type='kvm'>
  <name>ubuntu</name>
  <uuid>c7a5fdbd-cdaf-9455-926a-d65c16db1809</uuid>
  <memory unit="GiB">4</memory>
  <currentMemor unit="GiB">4</currentMemor>
  <vcpu>4</vcpu>
  <os>
    <type arch='x86_64' machine='pc'>hvm</type>
    <boot dev='hd'/>
    <boot dev='cdrom'/>
  </os>
  <devices>
    <emulator>/usr/bin/qemu-system-x86_64</emulator>
    <disk type='file' device='cdrom'>
      <source file='/home/fangaoyang/work/ubuntu-20.04.3-desktop-amd64.iso'/>
      <target dev='hdc'/>
      <readonly/>
    </disk>
    <disk type='file' device='disk'>
      <driver name="qemu" type="qcow2" />
      <source file='/home/fangaoyang/work/desktop.qcow2'/>
      <target dev='hda'/>
    </disk>
    <interface type='network'>
      <source network='default'/>
    </interface>
    <graphics type="vnc" port="-1" listen="0.0.0.0" passwd="123" keymap="en-us"/>
    <serial type="pty">
    	<target port="0" />
	</serial>
  	<console type="pty">
   		<target type="serial" port="0" />
  	</console>
  </devices>
</domain>`
	var d Domain
	err := xml.Unmarshal([]byte(opt), &d)
	if err != nil {
		t.Error(err)
	}

	res, err := xml.Marshal(DefaultOpt)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", string(res))

	res2, _ := xml.Marshal(d)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", string(res2))

	assert.Equal(t, res, res2)
}

func TestLib_CreateDomain(t *testing.T) {
	l := GenerateEnv()
	d := DefaultOpt
	u, _ := uuid.NewUUID()
	d.Uuid = u.String()
	d.Name = "testlib"
	d.Devices.Disk[1].Source.File = "/home/fangaoyang/work/testlib.qcow2"
	fmt.Printf("%+v\n", d)
	err := l.CreateDomain(d)
	assert.Equal(t, err, nil)
}

func TestLib_ShutdownDomain(t *testing.T) {
	l := GenerateEnv()
	err := l.ShutdownDomain(domainName)
	assert.Equal(t, err, nil)
}

func TestLib_RebootDomain(t *testing.T) {
	l := GenerateEnv()
	err := l.RebootDomain(domainName)
	assert.Equal(t, err, nil)
}
func TestLib_SuspendDomain(t *testing.T) {
	l := GenerateEnv()
	err := l.SuspendDomain(domainName)
	assert.Equal(t, err, nil)
}
func TestLib_ResumeDomain(t *testing.T) {
	l := GenerateEnv()
	err := l.ResumeDomain(domainName)
	assert.Equal(t, err, nil)
}
func TestLib_DestroyDomain(t *testing.T) {
	l := GenerateEnv()
	err := l.DestroyDomain(domainName)
	assert.Equal(t, err, nil)
}
