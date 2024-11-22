package client

import (
	"os"

	"github.com/synoti21/baekjoon-slack-bot/internal/common/consts"
)

type ProbRecommandSvc interface {
	GetProbsByUserID(uid string)
	GetProbsByCategory(category consts.ProbCategory)
	GetSimilarProbsByProbIDs(pid string)
	GetSimilarProbsByProbContents(pctnt string)
}

type probRecommandSvc struct {
	endpoint string
}

var _ ProbRecommandSvc = (*probRecommandSvc)(nil)

func NewProbRecommandSvc() (ProbRecommandSvc, error) {
	e := os.Getenv("REC_SVC_ENDPOINT")
	if e == "" {
		panic("REC_SVC_ENDPOINT not set.")
	}
	return &probRecommandSvc{
		endpoint: e,
	}, nil
}

func (p *probRecommandSvc) GetProbsByUserID(uid string) {
	panic("not implemented") // TODO: Implement
}

func (p *probRecommandSvc) GetProbsByCategory(category consts.ProbCategory) {
	panic("not implemented") // TODO: Implement
}

func (p *probRecommandSvc) GetSimilarProbsByProbIDs(pid string) {
	panic("not implemented") // TODO: Implement
}

func (p *probRecommandSvc) GetSimilarProbsByProbContents(pctnt string) {
	panic("not implemented") // TODO: Implement
}
