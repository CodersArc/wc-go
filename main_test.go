package main

import (
	"os"
	"os/exec"
	"testing"
)

func Test_wcgo(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "Step One -c",
			args: []string{"-c", "test.txt"},
			want: "340235 test.txt",
		},
		{
			name: "Step Two -l",
			args: []string{"-l", "test.txt"},
			want: "7148 test.txt",
		},
		{
			name: "Step Three -w",
			args: []string{"-w", "test.txt"},
			want: "59886 test.txt",
		},
		{
			name: "Step Four -m",
			args: []string{"-m", "test.txt"},
			want: "333921 test.txt",
		},
		{
			name: "Step Five default",
			args: []string{"test.txt"},
			want: "7148   59886  340235 test.txt",
		},
		{
			name: "The Final Step -l from stdin",
			args: []string{"-l"},
			want: "7148",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("wc-go", tt.args...)
			if tt.name == "The Final Step -l from stdin" {
				cmd.Stdin = os.Stdin
			}
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("cmd.Run() failed with %s\n", err)
			}
			if got := string(output); got != tt.want+"\n" {
				t.Errorf("wc-go() = %v, want %v", got, tt.want)
			}
		})
	}
}
