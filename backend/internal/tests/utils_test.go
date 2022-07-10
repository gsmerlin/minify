package tests

import (
	"testing"

	"github.com/gsmerlin/minify/backend/internal/utils"
)

func TestRandSeq(t *testing.T) {
	num := utils.RandSeq(5)
	if len(num) != 5 {
		t.Errorf("Expected length of 5, got %d", len(num))
	}
}
