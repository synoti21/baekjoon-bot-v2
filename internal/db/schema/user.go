package schema

// User is a common schema that maps the userID (Discord or Slack (or other else)) with Baekjoon ID
type User struct {
	UserID     string
	BaekjoonID string
}
