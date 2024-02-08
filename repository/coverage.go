package repository

import "github.com/litleleprikon/codegen_go_talk/model"

type Coverage map[model.CommitHash]map[model.FileName][]model.LineCoverage

func New() Coverage {
	return Coverage{}
}

func (c Coverage) AddCommitInfo(commit model.CommitHash, coverage map[model.FileName][]model.LineCoverage) {
	if _, found := c[commit]; !found {
		c[commit] = map[model.FileName][]model.LineCoverage{}
	}
	for file, cov := range coverage {
		if _, found := c[commit][file]; !found {
			c[commit][file] = cov
		} else {
			c[commit][file] = append(c[commit][file], cov...)
		}
	}
}

func (c Coverage) GetCommitInfo(commit model.CommitHash) map[model.FileName][]model.LineCoverage {
	return c[commit]
}
