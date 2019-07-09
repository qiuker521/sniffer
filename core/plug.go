package core

import (
	http "sniffer/module/http"
	"github.com/google/gopacket"
	"io"
	"path/filepath"
	"fmt"
)

type Plug struct {
	dir           string
	ResolveStream func(net gopacket.Flow, transport gopacket.Flow, r io.Reader)
	BPF           string

	InternalPlugList map[string]PlugInterface
}

// All internal plug-ins must implement this interface
// ResolvePacket - entry
// BPFFilter     - set BPF, like: mysql(tcp and port 3306)
// SetFlag       - plug-in params
// Version       - plug-in version
type PlugInterface interface {
	ResolveStream(net gopacket.Flow, transport gopacket.Flow, r io.Reader)
	BPFFilter() string
	SetFlag([]string)
	Version() string
}

func NewPlug() *Plug {

	var p Plug

	p.dir, _ = filepath.Abs("./plug/")
	p.LoadInternalPlugList()

	return &p
}

func (p *Plug) LoadInternalPlugList() {

	list := make(map[string]PlugInterface)
	list["http"] = http.NewInstance()

	p.InternalPlugList = list
}

func (p *Plug) ChangePath(dir string) {
	p.dir = dir
}

func (p *Plug) PrintList() {

	//Print Internal Plug
	for inPlugName, _ := range p.InternalPlugList {
		fmt.Println("internal plug : " + inPlugName)
	}

	//split
	fmt.Println("-- --- --")

}

func (p *Plug) SetOption(plugName string, plugParams []string) {

	//Load Internal Plug
	if internalPlug, ok := p.InternalPlugList[plugName]; ok {

		p.ResolveStream = internalPlug.ResolveStream
		internalPlug.SetFlag(plugParams)
		p.BPF = internalPlug.BPFFilter()

		return
	}

}
