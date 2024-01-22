package api

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"testing"
)

func TestNewShortUrl(t *testing.T) {
	var tests []struct {
		name string
		want *ShortUrl
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewShortUrl(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewShortUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShortUrl_Get(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ShortUrl{}
			a.Get(tt.args.c)
		})
	}

}

func TestShortUrl_Post(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ShortUrl{}
			a.Post(tt.args.c)
		})
	}
}
