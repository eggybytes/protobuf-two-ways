package protobuf

import (
	"protos/annotations"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Clean replaces every zero-valued primitive field with a nil value. It recurses into nested
// messages, so cleans nested primitives also
func Clean(pb proto.Message) proto.Message {
	m := pb.ProtoReflect()

	m.Range(cleanTopLevel(m))

	return pb
}

func cleanTopLevel(m protoreflect.Message) func(protoreflect.FieldDescriptor, protoreflect.Value) bool {
	return func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		// Skip cleaning any fields that are annotated with do_not_clean
		opts := fd.Options().(*descriptorpb.FieldOptions)
		if proto.GetExtension(opts, annotations.E_DoNotClean).(bool) {
			return true
		}

		// Otherwise, set any empty primitive fields to nil. For non-primitive fields, recurse down
		// one level with this function
		switch kind := fd.Kind(); kind {
		case protoreflect.BoolKind:
			if fd.Default().Bool() == v.Bool() {
				m.Clear(fd)
			}
		case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind, protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
			if fd.Default().Int() == v.Int() {
				m.Clear(fd)
			}
		case protoreflect.Uint32Kind, protoreflect.Fixed32Kind, protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
			if fd.Default().Uint() == v.Uint() {
				m.Clear(fd)
			}
		case protoreflect.FloatKind, protoreflect.DoubleKind:
			if fd.Default().Float() == v.Float() {
				m.Clear(fd)
			}
		case protoreflect.StringKind:
			if fd.Default().String() == v.String() {
				m.Clear(fd)
			}
		case protoreflect.BytesKind:
			if len(v.Bytes()) == 0 {
				m.Clear(fd)
			}
		case protoreflect.EnumKind:
			if fd.Default().Enum() == v.Enum() {
				m.Clear(fd)
			}
		case protoreflect.GroupKind:
			panic("groups are deprecated 🍳")
		case protoreflect.MessageKind:
			nested := v.Message()
			nested.Range(cleanTopLevel(nested))
		}

		return true
	}
}
