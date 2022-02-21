package virtlib

import (
	"encoding/xml"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLib_GetCapabilities(t *testing.T) {
	l := GenerateEnv()
	s, _ := l.GetCapabilities()
	//fmt.Println(s)
	c := Capability{}
	err := xml.Unmarshal([]byte(s), &c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", c)
}

func TestLib_GetDomMemStats(t *testing.T) {
	l := GetConnectedLib()
	stat, err := l.GetDomMemStats("ubuntu-m")
	assert.Equal(t, err, nil)
	fmt.Printf("%+v", stat)
}
func TestLib_GetDomMemUsage(t *testing.T) {
	l := GetConnectedLib()
	stat, err := l.GetDomMemUsage("ubuntu-m")
	assert.Equal(t, err, nil)
	fmt.Printf("%+v", stat)
}
