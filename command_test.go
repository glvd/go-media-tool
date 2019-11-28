package fftool

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

// TestFFProbeStreamFormat ...
func TestFFProbeStreamFormat(t *testing.T) {
	format, _ := FFProbeStreamFormat("D:\\video\\周杰伦唱歌贼难听.mp4")
	v, _ := json.Marshal(format)
	ioutil.WriteFile("d:\\test.json", v, os.ModePerm)
	t.Log(string(v), format.Resolution())
	format1, _ := FFProbeStreamFormat("D:\\video\\[BT天堂btbttt.com]我的女友.My.Girlfriend.2018.HD720P.X264.AAC.Korean.中文字幕.mp4")
	v1, _ := json.Marshal(format1)
	ioutil.WriteFile("d:\\test1.json", v, os.ModePerm)
	t.Log(string(v1), format1.Resolution())
}

// TestCommand_RunContext ...
func TestCommand_RunContext(t *testing.T) {
	type fields struct {
		path string
		Name string
		Args []string
	}
	type args struct {
		ctx  context.Context
		info chan<- string
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
				Args: nil,
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
			if err := c.RunContext(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("RunContext() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
