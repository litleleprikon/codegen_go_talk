package model

type LineCoverageType int

const (
	LineCoverageTypeUnspecified = iota
	LineCoverageTypeUncovered
	LineCoverageTypeCovered
)

func (c LineCoverageType) String() string {
	switch c {
	case LineCoverageTypeCovered:
		return "covered"
	case LineCoverageTypeUncovered:
		return "uncovered"
	default:
		return "unspecified"
	}
}

func NewLineCoverageType(t string) LineCoverageType {
	switch t {
	case "covered":
		return LineCoverageTypeCovered
	case "uncovered":
		return LineCoverageTypeUncovered
	default:
		return LineCoverageTypeUnspecified
	}
}

type CommitHash string
type LineNumber int
type FileName string

type LineCoverage struct {
	Line     LineNumber
	Coverage LineCoverageType
}
