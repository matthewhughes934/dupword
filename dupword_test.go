// MIT License
//
// Copyright (c) 2022 Abirdcfly
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package dupword_test

import (
	"testing"

	"github.com/abirdcfly/dupword"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	analyzer, err := dupword.NewAnalyzer(true, "")
	if err != nil {
		t.Fatalf("analyzer init failed:%s", err)
	}
	tests := []string{"a", "good"}
	analysistest.Run(t, analysistest.TestData(), analyzer, tests...)
}

func Test_checkOneKey(t *testing.T) {
	type args struct {
		raw string
		key string
	}
	tests := []struct {
		name     string
		args     args
		wantNew  string
		wantFind bool
	}{
		{
			name: "one word",
			args: args{
				raw: "Done",
				key: "the",
			},
			wantNew:  "",
			wantFind: false,
		},
		{
			name: "one word with space",
			args: args{
				raw: " Done \n \t",
				key: "the",
			},
			wantNew:  "",
			wantFind: false,
		},
		{
			name: "one line without key word",
			args: args{
				raw: "hello word",
				key: "the",
			},
			wantNew:  "",
			wantFind: false,
		},
		{
			name: "one line with key word only once",
			args: args{
				raw: "hello the word",
				key: "the",
			},
			wantNew:  "",
			wantFind: false,
		},
		{
			name: "one line with key word twice",
			args: args{
				raw: "hello the the world",
				key: "the",
			},
			wantNew:  "hello the world",
			wantFind: true,
		},
		{
			name: "one line with key word multi times",
			args: args{
				raw: "hello the the the world",
				key: "the",
			},
			wantNew:  "hello the world",
			wantFind: true,
		},
		{
			name: "multi line with key word once",
			args: args{
				raw: "hello \t the \nworld",
				key: "the",
			},
			wantNew:  "",
			wantFind: false,
		},
		{
			name: "multi line with key word twice",
			args: args{
				raw: "hello \t the \n   the world",
				key: "the",
			},
			wantNew:  "hello \t the \n   world",
			wantFind: true,
		},
		{
			name: "multi line with key word multi times",
			args: args{
				raw: "hello \t the \n   the the world",
				key: "the",
			},
			wantNew:  "hello \t the \n   world",
			wantFind: true,
		},
		{
			name: "multi line with key word multi times",
			args: args{
				raw: "print the\nthe line, print the\n\t the line.",
				key: "the",
			},
			wantNew:  "print the\nline, print the\n\t line.",
			wantFind: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNew, gotFind := dupword.CheckOneKey(tt.args.raw, tt.args.key)
			if gotNew != tt.wantNew {
				t.Errorf("CheckOneKey() gotNew = %q, want %q", gotNew, tt.wantNew)
			}
			if gotFind != tt.wantFind {
				t.Errorf("CheckOneKey() gotFind = %v, want %v", gotFind, tt.wantFind)
			}
		})
	}
}
