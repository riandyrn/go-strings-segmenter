package segmenter_test

import (
	"strings"
	"testing"

	segmenter "github.com/riandyrn/go-strings-segmenter"
)

func TestNext(t *testing.T) {
	testCases := []struct {
		Name          string
		Strings       []string
		SegmentLength int
		ExpResults    [][]string
	}{
		{
			Name:          "Segment Length Evenly Divisible With List Length",
			Strings:       []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
			SegmentLength: 5,
			ExpResults: [][]string{
				[]string{"0", "1", "2", "3", "4"},
				[]string{"5", "6", "7", "8", "9"},
			},
		},
		{
			Name:          "Segment Length Not Evenly Divisible With List Length",
			Strings:       []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
			SegmentLength: 3,
			ExpResults: [][]string{
				[]string{"0", "1", "2"},
				[]string{"3", "4", "5"},
				[]string{"6", "7", "8"},
				[]string{"9"},
			},
		},
		{
			Name:          "Segment Length Larger Than List Length",
			Strings:       []string{"0", "1", "2", "3"},
			SegmentLength: 5,
			ExpResults: [][]string{
				[]string{"0", "1", "2", "3"},
			},
		},
		{
			Name:          "Zero Segment Length",
			Strings:       []string{"0", "1", "2", "3"},
			SegmentLength: 0,
			ExpResults:    nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			sgmntr := segmenter.NewSegmenter(segmenter.Configs{
				Strings:       testCase.Strings,
				SegmentLength: testCase.SegmentLength,
			})
			var results [][]string
			for sgmntr.HasNext() {
				results = append(results, sgmntr.Next())
			}
			if len(results) != len(testCase.ExpResults) {
				t.Fatalf("unexpected results, expected: %v, got: %v", testCase.ExpResults, results)
			}
			for i := 0; i < len(results); i++ {
				resultStr := strings.Join(results[i], ",")
				expResultStr := strings.Join(testCase.ExpResults[i], ",")
				if resultStr != expResultStr {
					t.Fatalf("unexpected result, expected: %v, got: %v", expResultStr, resultStr)
				}
			}
		})
	}
}
