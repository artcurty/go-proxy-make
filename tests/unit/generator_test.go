package unit

import (
	"github.com/artcurty/go-proxy-make/internal"
	"testing"
)

func TestGenerateProxyFunctionForInput(t *testing.T) {
	type args struct {
		inputFile string
		outputDir string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				inputFile: "testdata/valid_openapi.yaml",
				outputDir: "testdata/generatedtest",
			},
			wantErr: false,
		},
		{
			name: "FileNotFound",
			args: args{
				inputFile: "testdata/nonexistent.yaml",
				outputDir: "testdata/generatedtest",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := internal.GenerateProxyFunctionForInput(tt.args.inputFile, tt.args.outputDir); (err != nil) != tt.wantErr {
				t.Errorf("GenerateProxyFunctionForInput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
