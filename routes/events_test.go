package routes

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_updateEvent(t *testing.T) {
	type args struct {
		context *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateEvent(tt.args.context)
		})
	}
}
