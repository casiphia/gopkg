package compress

import (
	"reflect"
	"testing"
)

func Test_Snappy_Compress(t *testing.T) {
	type args struct {
		src []byte
	}
	tests := []struct {
		name string
		s    *_Snappy
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			name: "good case",
			args: args{
				src: []byte("test"),
			},
			want: []byte{4, 12, 116, 101, 115, 116},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &_Snappy{}
			if got := s.Compress(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("_Snappy.Compress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Snappy_DeCompress(t *testing.T) {
	type args struct {
		src []byte
	}
	tests := []struct {
		name    string
		s       *_Snappy
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "good case",
			args: args{
				src: []byte{4, 12, 116, 101, 115, 116},
			},
			want:    []byte("test"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &_Snappy{}
			got, err := s.DeCompress(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("_Snappy.DeCompress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("_Snappy.DeCompress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Snappy_CompressWithPrefix(t *testing.T) {
	type args struct {
		src    []byte
		prefix []byte
	}
	tests := []struct {
		name string
		s    *_Snappy
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			name: "good case",
			args: args{
				src:    []byte("test"),
				prefix: []byte("p_"),
			},
			want: []byte{112, 95, 4, 12, 116, 101, 115, 116},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &_Snappy{}
			if got := s.CompressWithPrefix(tt.args.src, tt.args.prefix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("_Snappy.CompressWithPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
