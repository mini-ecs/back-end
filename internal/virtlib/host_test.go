package virtlib

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func TestLib_GetCapabilities(t *testing.T) {
	l := GenerateEnv()
	s, _ := l.GetCapabilities()
	//fmt.Println(s)
	c := Capabilities{}
	err := xml.Unmarshal([]byte(s), &c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", c)
}
