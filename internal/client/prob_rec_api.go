package client

import (
	"os"

	"github.com/synoti21/baekjoon-slack-bot/common/consts"
)

type ProbRecommendAPI interface {
	GetProbsByUserID(uid string, probCnt int) (*ProbRecommendAPIResponse, error)
	GetProbsByCategory(category consts.ProbCategory) (*ProbRecommendAPIResponse, error)
	GetSimilarProbsByProbIDs(pid int) (*ProbRecommendAPIResponse, error)
	GetSimilarProbsByProbContents(pctnt string) (*ProbRecommendAPIResponse, error)
}

type probRecommendAPI struct {
	endpoint string
}

type ProbRecommendAPIRequest struct {
	UserIDs  []string            `json:"user_id_list"`
	Category consts.ProbCategory `json:"category"`
	ProbNum  int                 `json:"problem_num"`
}

type ProbRecommendAPIResponse struct {
	ProbIDsByUserID        map[string][]int
	SimilarProbIDsByProbID map[string][]int
	ProbIDsByCategory      map[consts.ProbCategory][]int
}

var _ ProbRecommendAPI = (*probRecommendAPI)(nil)

func NewProbRecommandSvc() (ProbRecommendAPI, error) {
	e := os.Getenv("REC_SVC_ENDPOINT")
	if e == "" {
		panic("REC_SVC_ENDPOINT not set.")
	}
	return &probRecommendAPI{
		endpoint: e,
	}, nil
}

func (p *probRecommendAPI) GetProbsByUserID(uid string, probCnt int) (*ProbRecommendAPIResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (p *probRecommendAPI) GetProbsByCategory(category consts.ProbCategory) (*ProbRecommendAPIResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (p *probRecommendAPI) GetSimilarProbsByProbIDs(pid int) (*ProbRecommendAPIResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (p *probRecommendAPI) GetSimilarProbsByProbContents(pctnt string) (*ProbRecommendAPIResponse, error) {
	panic("not implemented") // TODO: Implement
}
