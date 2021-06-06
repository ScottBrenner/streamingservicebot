package main

import (
	"testing"

	"github.com/turnage/graw/reddit"
)

func Test_streamingServiceBot_Post(t *testing.T) {
	type fields struct {
		bot reddit.Bot
	}
	type args struct {
		p *reddit.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &streamingServiceBot{
				bot: tt.fields.bot,
			}
			if err := r.Post(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("streamingServiceBot.Post() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
