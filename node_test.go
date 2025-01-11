package datastructures

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewNode(t *testing.T) {
	type args struct {
		key   string
		value int
	}

	tests := []struct {
		name string
		args args
		want *Node[int]
	}{
		{
			name: "Given a key and a value, it should return a Node",
			args: args{
				key:   "k",
				value: 1,
			},
			want: &Node[int]{
				key:   "k",
				value: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNode(tt.args.key, tt.args.value)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestNode_Print(t *testing.T) {
	// Capture the current os.Stdout value so it can be restablished when the test ends.
	currentStdOut := os.Stdout
	defer func() { os.Stdout = currentStdOut }()

	tests := []struct {
		name string
		n    *Node[int]
		want []byte
	}{
		{
			name: "Given a non-nil Node, it should print its key and value",
			n: &Node[int]{
				key:   "k",
				value: 1,
			},
			want: []byte(`k: 1`),
		},
		{
			name: "Given a nil Node, it should print a nil value",
			n:    nil,
			want: []byte(`<nil>`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Replace os.Stdout with a temporary file with a random name at the current directory
			// until the end of the test scenario.
			r, w, err := os.Pipe()
			require.NoError(t, err)
			os.Stdout = w

			tt.n.Print()

			// Close the writer so the read operation doesn't get blocked.
			w.Close()

			got, err := io.ReadAll(r)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
