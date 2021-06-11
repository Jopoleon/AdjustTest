package main

import (
	"net/url"
	"reflect"
	"testing"
)

func Test_parseArguments(t *testing.T) {
	type args struct {
		args     []string
		parallel int
	}
	tests := []struct {
		name    string
		args    args
		want    Arguments
		wantErr bool
	}{
		{
			name: "success case",
			args: args{
				args:     []string{"http://www.adjust.com", "google.com"},
				parallel: 22,
			},
			want: Arguments{
				urls:     []string{"http://www.adjust.com", "http://google.com"},
				parallel: 22,
			},
			wantErr: false,
		},
		{
			name: "bad url case",
			args: args{
				args: []string{"httpjustcom", "httglecom"},
				//parallel: 22,
			},
			want:    Arguments{},
			wantErr: true,
		},
		{
			name: "no args error",
			args: args{
				args: []string{},
				//parallel: 22,
			},
			want:    Arguments{},
			wantErr: true,
		},
		{
			name: "zero parallel error",
			args: args{
				args:     []string{"http://www.adjust.com", "google.com"},
				parallel: 0,
			},
			want:    Arguments{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseArguments(tt.args.args, tt.args.parallel)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseArguments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseArguments() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMD5(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "success case",
			args:    args{url: "http://www.adjust.com"},
			want:    "5d12e19368b240e49ee9d84aa257b219",
			wantErr: false,
		},
		{
			name:    "bad url case",
			args:    args{url: "httpwwwadjustcom"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getMD5(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("getMD5() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMD5() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseRawURL(t *testing.T) {
	type args struct {
		rawurl string
	}
	tests := []struct {
		name    string
		args    args
		want    *url.URL
		wantErr bool
	}{
		{
			name: "success case",
			args: args{rawurl: "www.google.com"},
			want: &url.URL{
				Scheme: "http",
				Host:   "www.google.com",
			},
			wantErr: false,
		},
		{
			name: "no scheme case",
			args: args{rawurl: "google.com"},
			want: &url.URL{
				Scheme: "http",
				Host:   "google.com",
			},
			wantErr: false,
		},
		{
			name:    "no host case",
			args:    args{rawurl: "wwwgooglecom"},
			want:    &url.URL{},
			wantErr: true,
		},
		{
			name:    "empty rawurl case",
			args:    args{rawurl: ""},
			want:    &url.URL{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseRawURL(tt.args.rawurl)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseRawURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseRawURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}
