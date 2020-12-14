package protobuf

import (
	"testing"

	"protos/example"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestClean(t *testing.T) {
	type args struct {
		req proto.Message
	}
	type want struct {
		req proto.Message
	}
	tests := []struct {
		msg  string
		args args
		want want
	}{
		{
			"top-level empty OrderEggRequest and zero values should be replaced with nil",
			args{
				&example.OrderEggRequest{
					Name:      proto.String(""),
					WithShell: proto.Bool(false),
				},
			},
			want{
				&example.OrderEggRequest{},
			},
		},
		{
			"nested fields also get cleaned",
			args{
				&example.OrderEggRequest{
					Name: proto.String("nice order"),
					Recipient: &example.Recipient{
						Name:    proto.String("Denton"),
						Address: proto.String(""),
					},
				},
			},
			want{
				&example.OrderEggRequest{
					Name: proto.String("nice order"),
					Recipient: &example.Recipient{
						Name:    proto.String("Denton"),
						Address: nil,
					},
				},
			},
		},
		{
			"fields marked with do_not_clean do not get cleaned",
			args{
				&example.OrderEggRequest{
					Name:        proto.String("im an order"),
					Description: proto.String(""),
				},
			},
			want{
				&example.OrderEggRequest{
					Name:        proto.String("im an order"),
					Description: proto.String(""),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			assert.Equal(t, tt.want.req, Clean(tt.args.req))
		})
	}
}
