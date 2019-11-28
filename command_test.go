package fftool

import (
	"context"
	"testing"

	"github.com/goextension/log"
)

// TestCommand_RunContext ...
func TestCommand_RunContext(t *testing.T) {
	type fields struct {
		path string
		Name string
		Args []string
	}
	type args struct {
		ctx  context.Context
		info chan string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				path: DefaultCommandPath,
				Name: "ffmpeg",
				Args: []string{"-version"},
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		}, {
			name: "test2",
			fields: fields{
				path: DefaultCommandPath,
				Name: "ffmpeg",
				Args: []string{"-version"},
			},
			args: args{
				ctx:  context.Background(),
				info: make(chan string),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Command{
				path: tt.fields.path,
				Name: tt.fields.Name,
				Args: tt.fields.Args,
			}
			go func() {
				if err := c.RunContext(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
					t.Errorf("RunContext() error = %v, wantErr %v", err, tt.wantErr)
				}
			}()
			if tt.args.info != nil {
				for v := range tt.args.info {
					log.Info(v)
				}
			}
		})
	}
}

// TestCommand_Run ...
func TestCommand_Run(t *testing.T) {
	type fields struct {
		path string
		Name string
		Args []string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				path: DefaultCommandPath,
				Name: "ffmpeg",
				Args: []string{"-version"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Command{
				path: tt.fields.path,
				Name: tt.fields.Name,
				Args: tt.fields.Args,
			}
			got, err := c.Run()
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("Run() got = %v", got)
		})
	}
}
