package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var tests = []string{
	"simple/",
}

func Test(t *testing.T) {
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			var res bytes.Buffer
			req, _ := os.Open(filepath.Join("tests", test, "code_generator_request.pb.bin"))
			err := run(protogen.Options{ParamFunc: flag.Set}, req, &res, GenerateFiles)
			require.NoError(t, err)
			response := &pluginpb.CodeGeneratorResponse{}
			err = proto.Unmarshal(res.Bytes(), response)
			require.NoError(t, err)
			fmt.Println(response.String())
		})
	}
}

func run(opts protogen.Options, input io.Reader, output io.Writer, f func(*protogen.Plugin) error) error {
	in, err := ioutil.ReadAll(input)
	if err != nil {
		return err
	}
	req := &pluginpb.CodeGeneratorRequest{}
	if err := proto.Unmarshal(in, req); err != nil {
		return err
	}
	this := ""
	req.Parameter = &this
	gen, err := opts.New(req)
	if err != nil {
		return err
	}
	if err := f(gen); err != nil {
		gen.Error(err)
	}
	resp := gen.Response()
	out, err := proto.Marshal(resp)
	if err != nil {
		return err
	}
	if _, err := output.Write(out); err != nil {
		return err
	}
	return nil
}
