package fftool

import (
	"github.com/google/uuid"
	"strings"
	"testing"
)

var sf *StreamFormat

func init() {
	var e error
	sf, e = NewFFProbe().StreamFormat(`d:\video\女大学生的沙龙室.Room.Salon.College.Girls.2018.HD720P.X264.AAC.Korean.CHS.mp4`)
	if e != nil {
		panic(e)
	}
}

// TestConfig_OptimizeWithFormat ...
func TestConfig_OptimizeWithFormat(t *testing.T) {
	type fields struct {
		Scale           Scale
		BitRate         int64
		FrameRate       float64
		Output          string
		VideoFormat     string
		AudioFormat     string
		M3U8Name        string
		SegmentFileName string
		HLSTime         int
	}
	type args struct {
		sfmt *StreamFormat
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "1",
			fields: fields{
				Scale:           2,
				BitRate:         0,
				FrameRate:       0,
				Output:          "",
				VideoFormat:     "",
				AudioFormat:     "",
				M3U8Name:        "",
				SegmentFileName: "",
				HLSTime:         0,
			},
			args: args{
				sfmt: sf,
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "2",
			fields: fields{
				Scale:           1,
				BitRate:         0,
				FrameRate:       0,
				Output:          "",
				VideoFormat:     "",
				AudioFormat:     "",
				M3U8Name:        "",
				SegmentFileName: "",
				HLSTime:         0,
			},
			args: args{
				sfmt: sf,
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "3",
			fields: fields{
				Scale:           1,
				BitRate:         0,
				FrameRate:       0,
				Output:          "",
				VideoFormat:     "",
				AudioFormat:     "",
				M3U8Name:        "",
				SegmentFileName: "",
				HLSTime:         0,
			},
			args: args{
				sfmt: nil,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Scale:           tt.fields.Scale,
				BitRate:         tt.fields.BitRate,
				FrameRate:       tt.fields.FrameRate,
				OutputPath:      tt.fields.Output,
				VideoFormat:     tt.fields.VideoFormat,
				AudioFormat:     tt.fields.AudioFormat,
				M3U8Name:        tt.fields.M3U8Name,
				SegmentFileName: tt.fields.SegmentFileName,
				HLSTime:         tt.fields.HLSTime,
			}
			err := c.OptimizeWithFormat(tt.args.sfmt)
			if (err != nil) != tt.wantErr {
				t.Errorf("OptimizeWithFormat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("%+v", c)
		})
	}
}

// TestConfig_Args ...
func TestConfig_Args(t *testing.T) {
	type fields struct {
		Scale           Scale
		BitRate         int64
		FrameRate       float64
		Output          string
		VideoFormat     string
		AudioFormat     string
		M3U8Name        string
		SegmentFileName string
		HLSTime         int
	}
	type args struct {
		intput string
		output string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "args1",
			fields: fields{
				Scale:           0,
				BitRate:         0,
				FrameRate:       0,
				Output:          "",
				VideoFormat:     "",
				AudioFormat:     "",
				M3U8Name:        "",
				SegmentFileName: "",
				HLSTime:         0,
			},
			args: args{
				intput: "d:\\video\\周杰伦 唱歌贼难听.2019.1080P.h264.aac.Japanese.None.mp4",
				output: "d:\\temp\\",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := DefaultConfig()
			//c := &Config{
			//	Scale:           tt.fields.Scale,
			//	BitRate:         tt.fields.BitRate,
			//	FrameRate:       tt.fields.FrameRate,
			//	OutputPath:          tt.fields.OutputPath,
			//	VideoFormat:     tt.fields.VideoFormat,
			//	AudioFormat:     tt.fields.AudioFormat,
			//	M3U8Name:        tt.fields.M3U8Name,
			//	SegmentFileName: tt.fields.SegmentFileName,
			//	HLSTime:         tt.fields.HLSTime,
			//}
			c.OutputName = uuid.New().String()
			c.NeedSlice = true
			c.ProcessCore = ProcessCPU
			if got := c.Args(tt.args.intput, tt.args.output); got != tt.want {
				t.Errorf("Args() = %v, want %v", strings.ReplaceAll(got, ",", " "), tt.want)
			}
		})
	}
}
