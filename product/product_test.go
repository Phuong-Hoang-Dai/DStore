package product_test

import (
	"testing"

	"github.com/Phuong-Hoang-Dai/DStore/product"
	"github.com/stretchr/testify/assert"
)

func TestPagingProcess(t *testing.T) {
	tests := []struct {
		p              product.Paging
		expectedLimit  int
		expectedOffset int
	}{
		{product.Paging{Limit: 0, Offset: 1}, 0, 1},
		{product.Paging{Limit: -1, Offset: 1}, 0, 1},
		{product.Paging{Limit: 0, Offset: -1}, 0, 0},
		{product.Paging{Limit: 100, Offset: 0}, 50, 0},
		{product.Paging{Limit: -1, Offset: -1}, 0, 0},
	}

	for _, test := range tests {
		t.Run("TestCase", func(t *testing.T) {

			test.p.Process()
			assert.Equal(t, test.expectedLimit, test.p.Limit)
			assert.Equal(t, test.expectedOffset, test.p.Offset)
		})
	}
}
