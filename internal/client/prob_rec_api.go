package client

import (
	"os"

	"github.com/synoti21/baekjoon-slack-bot/common/consts"
)

type ProbRecommandAPI interface {
	GetProbsByUserID(uid string, probCnt int) (*ProbRecommandAPIResponse, error)
	GetProbsByCategory(category consts.ProbCategory) (*ProbRecommandAPIResponse, error)
	GetSimilarProbsByProbIDs(pid string) (*ProbRecommandAPIResponse, error)
	GetSimilarProbsByProbContents(pctnt string) (*ProbRecommandAPIResponse, error)
}

type probRecommandAPI struct {
	endpoint string
}

type ProbRecommandAPIRequest struct {
	UserIDs  []string            `json:"user_id_list"`
	Category consts.ProbCategory `json:"category"`
	ProbNum  int                 `json:"problem_num"`
}

type ProbRecommandAPIResponse struct {
	ProbIDsByUserID        map[string][]int
	SimilarProbIDsByProbID map[string][]int
	ProbIDsByCategory      map[consts.ProbCategory][]int
}

var _ ProbRecommandAPI = (*probRecommandAPI)(nil)

func NewProbRecommandSvc() (ProbRecommandAPI, error) {
	e := os.Getenv("REC_SVC_ENDPOINT")
	if e == "" {
		panic("REC_SVC_ENDPOINT not set.")
	}
	return &probRecommandAPI{
		endpoint: e,
	}, nil
}

func (p *probRecommandAPI) GetProbsByUserID(uid string, probCnt int) (*ProbRecommandAPIResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (p *probRecommandAPI) GetProbsByCategory(category consts.ProbCategory) (*ProbRecommandAPIResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (p *probRecommandAPI) GetSimilarProbsByProbIDs(pid string) (*ProbRecommandAPIResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (p *probRecommandAPI) GetSimilarProbsByProbContents(pctnt string) (*ProbRecommandAPIResponse, error) {
	panic("not implemented") // TODO: Implement
}
