package repository

import (
	"YandexPra/iternal/domain"
	"reflect"
	"testing"
)

func TestNewUrlRepo(t *testing.T) {
	var tests []struct {
		name string
		want *UrlRepo
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUrlRepo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUrlRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUrlRepo_Get(t *testing.T) {

	type args struct {
		key string
	}
	var tests []struct {
		name    string
		args    args
		want    domain.Urls
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ur := &UrlRepo{}
			got, err := ur.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUrlRepo_Post(t *testing.T) {
	type args struct {
		url domain.Urls
	}
	var tests []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ur := &UrlRepo{}
			got, err := ur.Post(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Post() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Post() got = %v, want %v", got, tt.want)
			}
		})
	}
}
