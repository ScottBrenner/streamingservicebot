package main

import (
	"testing"

	"github.com/turnage/graw/reddit"
)

func Test_checkURL(t *testing.T) {
	tests := []struct {
		name string
		args *reddit.Post
		want bool
	}{
		{"youtube.com", &reddit.Post{URL: "https://www.youtube.com/watch?v=aDlZckOUHiw"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkURL(tt.args); got != tt.want {
				t.Errorf("checkURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processYouTube(t *testing.T) {
	type args struct {
		p *reddit.Post
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processYouTube(tt.args.p); got != tt.want {
				t.Errorf("processYouTube() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateReply(t *testing.T) {
	type args struct {
		p *reddit.Post
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateReply(tt.args.p); got != tt.want {
				t.Errorf("generateReply() = %v, want %v", got, tt.want)
			}
		})
	}
}
