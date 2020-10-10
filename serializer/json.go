package serializer

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func ProtoBufToJSON(message proto.Message) (string, error) {
	marshler := jsonpb.Marshaler{
		// OrigName specifies whether to use the original protobuf name for fields.
		OrigName: true,

		// EnumsAsInts specifies whether to render enum values as integers,
		// as opposed to string values.
		EnumsAsInts: false,

		// EmitDefaults specifies whether to render fields with zero values.
		EmitDefaults: true,

		// Indent controls whether the output is compact or not.
		// If empty, the output is compact JSON. Otherwise, every JSON object
		// entry and JSON array value will be on its own line.
		// Each line will be preceded by repeated copies of Indent, where the
		// number of copies is the current indentation depth.
		Indent: " ",
	}

	return marshler.MarshalToString(message)
}
