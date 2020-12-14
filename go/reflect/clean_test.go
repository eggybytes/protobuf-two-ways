package protobuf

import (
	"testing"

	"protos/todo"

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
			"top-level empty strings and zero values should be replaced with nil",
			args{
				&todo.CreateTodoRequest{
					Name:       proto.String(""),
					IsComplete: proto.Bool(false),
				},
			},
			want{
				&todo.CreateTodoRequest{},
			},
		},
		{
			"fields marked with do_not_clean do not get cleaned",
			args{
				&todo.CreateTodoRequest{
					Name:        proto.String("do me"),
					Description: proto.String(""),
				},
			},
			want{
				&todo.CreateTodoRequest{
					Name:        proto.String("do me"),
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
