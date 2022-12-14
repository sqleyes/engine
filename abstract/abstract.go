package abstract

import "github.com/google/gopacket"

type Installed Msg
type ERROR Msg
type Intact gopacket.Packet

type Config map[string]any
type Command int

const (
	None Command = iota
	Start
	Pause
	Stop
)

type Msg struct {
	Code int
	Text string
}
type Broken struct {
	SrcIP, DstIP     string
	SrcPort, DstPort int
	Payload          []byte
}
type Plugin interface {
	React(any) Command
}
