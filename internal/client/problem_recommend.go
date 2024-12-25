package client

import (
	"os"

	"github.com/synoti21/baekjoon-slack-bot/common/consts"
)

type ProblemRecommendClient interface {
	GetProblemsByUserID(uid string, probCnt int) (*ProbRecommendResponse, error)
	GetProblemsByCategory(category consts.ProbCategory) (*ProbRecommendResponse, error)
	GetSimilarProblemsByProblemIDs(pid int) (*ProbRecommendResponse, error)
	GetSimilarProblemsByContents(content string) (*ProbRecommendResponse, error)
}

type problemRecommendClient struct {
	endpoint string
}

type ProbRecommendRequest struct {
	UserIDList []string            `json:"user_id_list"`
	Category   consts.ProbCategory `json:"category"`
	ProblemNum int                 `json:"problem_num"` // Our model gets a "problem_num" field for the count of problems.
}

type ProbRecommendResponse struct {
	ProblemIDsByUserID           map[string][]int
	SimilarProblemIDsByProblemID map[string][]int
	ProblemIDsByCategory         map[consts.ProbCategory][]int
}

var _ ProblemRecommendClient = (*problemRecommendClient)(nil)

func NewProblemRecommendClient() (ProblemRecommendClient, error) {
	e := os.Getenv("REC_SVC_ENDPOINT")
	if e == "" {
		panic("REC_SVC_ENDPOINT not set.")
	}
	return &problemRecommendClient{
		endpoint: e,
	}, nil
}

func (p *problemRecommendClient) GetProblemsByUserID(uid string, probCnt int) (*ProbRecommendResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (p *problemRecommendClient) GetProblemsByCategory(category consts.ProbCategory) (*ProbRecommendResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (p *problemRecommendClient) GetSimilarProblemsByProblemIDs(pid int) (*ProbRecommendResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (p *problemRecommendClient) GetSimilarProblemsByContents(content string) (*ProbRecommendResponse, error) {
	panic("not implemented") // TODO: Implement
}
