package segmenter

import "math"

// Segmenter is used for loading strings in segment
type Segmenter struct {
	strs          []string
	startIdx      int
	endIdx        int
	segmentLength int
	numOps        int
	opsCounter    int
}

// Configs holds configs for StringsSegmenter
type Configs struct {
	Strings       []string
	SegmentLength int
}

// NewSegmenter returns new instance of StringsSegmenter
func NewSegmenter(configs Configs) *Segmenter {
	segmenter := &Segmenter{
		strs:          configs.Strings,
		segmentLength: configs.SegmentLength,
		startIdx:      0,
		endIdx:        configs.SegmentLength,
		numOps:        int(math.Ceil(float64(len(configs.Strings)) / float64(configs.SegmentLength))),
		opsCounter:    0,
	}
	if segmenter.endIdx > len(configs.Strings) {
		segmenter.endIdx = len(configs.Strings)
	}
	return segmenter
}

// Next returns next strings segment, when there is no next
// segment the returned value is nil
func (s *Segmenter) Next() []string {
	if s.opsCounter == s.numOps {
		return nil
	}
	defer func() {
		s.opsCounter++
		s.startIdx = s.opsCounter * s.segmentLength
		s.endIdx = s.startIdx + s.segmentLength
		if s.endIdx > len(s.strs) {
			s.endIdx = len(s.strs)
		}
	}()
	return s.strs[s.startIdx:s.endIdx]
}

// HasNext returns true when strings still has next segment
func (s *Segmenter) HasNext() bool {
	return s.opsCounter < s.numOps
}
