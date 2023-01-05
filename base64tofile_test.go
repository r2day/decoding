package decoding

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestBase64ToFile(t *testing.T) {
	
	f, err := ioutil.TempFile("./", "test.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(f.Name())

	type args struct {
		b64 string
		f   *os.File
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"test",
			args{
				"aGVsbG8=",
				f,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Base64ToFile(tt.args.b64, tt.args.f); (err != nil) != tt.wantErr {
				t.Errorf("Base64ToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
