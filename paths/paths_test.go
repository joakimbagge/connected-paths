package paths

import (
	"testing"
)

func TestConnectedPathsCount(t *testing.T) {
	testCases := []struct {
		cols      int
		rows      int
		wantCount int
	}{
		{cols: 0, rows: 0, wantCount: 0},
		{cols: 1, rows: 1, wantCount: 1},
		{cols: 1, rows: 2, wantCount: 2},
		{cols: 2, rows: 1, wantCount: 1},
		{cols: 2, rows: 2, wantCount: 4},
		{cols: 3, rows: 3, wantCount: 17},
		{cols: 5, rows: 3, wantCount: 99},
	}
	for _, tc := range testCases {
		got := ConnectedPaths(tc.cols, tc.rows)
		if len(got) != tc.wantCount {
			t.Errorf("For grid %dx%d: expected %d paths, got %d",
				tc.cols, tc.rows, tc.wantCount, len(got))
		}
	}
}
