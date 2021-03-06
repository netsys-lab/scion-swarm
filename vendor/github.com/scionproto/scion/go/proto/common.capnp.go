// Code generated by capnpc-go. DO NOT EDIT.

package proto

import (
	capnp "zombiezen.com/go/capnproto2"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type LinkType uint16

// LinkType_TypeID is the unique identifier for the type LinkType.
const LinkType_TypeID = 0x916c98f48c9bbb64

// Values of LinkType.
const (
	LinkType_unset  LinkType = 0
	LinkType_core   LinkType = 1
	LinkType_parent LinkType = 2
	LinkType_child  LinkType = 3
	LinkType_peer   LinkType = 4
)

// String returns the enum's constant name.
func (c LinkType) String() string {
	switch c {
	case LinkType_unset:
		return "unset"
	case LinkType_core:
		return "core"
	case LinkType_parent:
		return "parent"
	case LinkType_child:
		return "child"
	case LinkType_peer:
		return "peer"

	default:
		return ""
	}
}

// LinkTypeFromString returns the enum value with a name,
// or the zero value if there's no such value.
func LinkTypeFromString(c string) LinkType {
	switch c {
	case "unset":
		return LinkType_unset
	case "core":
		return LinkType_core
	case "parent":
		return LinkType_parent
	case "child":
		return LinkType_child
	case "peer":
		return LinkType_peer

	default:
		return 0
	}
}

type LinkType_List struct{ capnp.List }

func NewLinkType_List(s *capnp.Segment, sz int32) (LinkType_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return LinkType_List{l.List}, err
}

func (l LinkType_List) At(i int) LinkType {
	ul := capnp.UInt16List{List: l.List}
	return LinkType(ul.At(i))
}

func (l LinkType_List) Set(i int, v LinkType) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type ServiceType uint16

// ServiceType_TypeID is the unique identifier for the type ServiceType.
const ServiceType_TypeID = 0xe2522d291bd06774

// Values of ServiceType.
const (
	ServiceType_unset ServiceType = 0
	ServiceType_bs    ServiceType = 1
	ServiceType_ps    ServiceType = 2
	ServiceType_cs    ServiceType = 3
	ServiceType_sb    ServiceType = 4
	ServiceType_ds    ServiceType = 5
	ServiceType_br    ServiceType = 6
	ServiceType_sig   ServiceType = 7
)

// String returns the enum's constant name.
func (c ServiceType) String() string {
	switch c {
	case ServiceType_unset:
		return "unset"
	case ServiceType_bs:
		return "bs"
	case ServiceType_ps:
		return "ps"
	case ServiceType_cs:
		return "cs"
	case ServiceType_sb:
		return "sb"
	case ServiceType_ds:
		return "ds"
	case ServiceType_br:
		return "br"
	case ServiceType_sig:
		return "sig"

	default:
		return ""
	}
}

// ServiceTypeFromString returns the enum value with a name,
// or the zero value if there's no such value.
func ServiceTypeFromString(c string) ServiceType {
	switch c {
	case "unset":
		return ServiceType_unset
	case "bs":
		return ServiceType_bs
	case "ps":
		return ServiceType_ps
	case "cs":
		return ServiceType_cs
	case "sb":
		return ServiceType_sb
	case "ds":
		return ServiceType_ds
	case "br":
		return ServiceType_br
	case "sig":
		return ServiceType_sig

	default:
		return 0
	}
}

type ServiceType_List struct{ capnp.List }

func NewServiceType_List(s *capnp.Segment, sz int32) (ServiceType_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return ServiceType_List{l.List}, err
}

func (l ServiceType_List) At(i int) ServiceType {
	ul := capnp.UInt16List{List: l.List}
	return ServiceType(ul.At(i))
}

func (l ServiceType_List) Set(i int, v ServiceType) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

const schema_fa01960eced2b529 = "x\xda\x12hu`2d\x15gb`\x08\x94`e" +
	"\xfb\x9f\xb2{v\xcf\x97\x199\x13\x19\x04y\x99\xfek" +
	"n\xbdt\x8eo\x1a\xe3/\x06\x06F\xc1\x8f\x9b\x04\x7f" +
	"\xb230\x08~\xadg`\xfc_\x92~AZS7" +
	"\xe8\x11\xba*aM\xc6S\xc2\xa6\x8c\xec\x0c\x0c\xc2\x86" +
	"\x8c\xc7\x19t\xff'\xe7\xe7\xe6\xe6\xe7\xe9%3&\x16" +
	"\xe4\x15X\xf9d\xe6\xc9g\x87T\x16\xa4\x0602\x06" +
	"\x8a0210\x08\x9a\x1a1002\x0a\xeaj1" +
	"002\x09\xaaZ1002\x0b\xca\x82\x04Y\x04" +
	"E\xb5\x18\x18\xe4K\xf3\x8aSK\xf8\x93\xf3\x8bR\xed" +
	"\x0b\x12\x8bR\xf3J\xe4\x9332sR\xf8\x0bRS" +
	"\x8b\xe0\xc63\x81\x8d\x0fN-*\xcbLN\x05Y\xc0" +
	"\xc0\x00\xb2B\x01lE$\xc4\x8a@)\xb0\x15\x9eR" +
	"`+\x1c\xa5\xc0VX\x82(VAC\x10\xc5&\xa8" +
	"\x09\xa2\xd8\x05\x15\x95`\xf62'\x153\x17\x143'" +
	"\x173\x17'1\xa7\x143'\x15\xb1\x17g\xa6\x03\x02" +
	"\x00\x00\xff\xff\xafaKL"

func init() {
	schemas.Register(schema_fa01960eced2b529,
		0x916c98f48c9bbb64,
		0xe2522d291bd06774)
}
