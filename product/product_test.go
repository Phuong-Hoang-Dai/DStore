package product_test

import (
	"testing"

	"github.com/Phuong-Hoang-Dai/DStore/product"
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
		{product.Paging{Limit: 100, Offset: 0}, 49, 0},
		{product.Paging{Limit: -1, Offset: -1}, 0, 0},
	}

	for _, test := range tests {
		t.Run("Test Process", func(t *testing.T) {
			test.p.Process()
			if test.p.Limit != test.expectedLimit ||
				test.p.Offset != test.expectedOffset {
				t.Errorf("paging.Process() limit or offset is incorrect."+
					"\n offset: %v & expected: %v\n limit: %v & expected: %v",
					test.p.Limit, test.expectedLimit,
					test.p.Offset, test.expectedOffset)
			}
		})
	}
}
