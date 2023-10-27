package swfx

import "github.com/x65534/swfx/tagcode"

type Metadata struct {
	Value string
}

func (tag *Metadata) Code() tagcode.TagCode {
	return tagcode.Metadata
}

func (tag *Metadata) readData(r SwfReader, length int) {
	tag.Value = r.ReadString(length)
}
