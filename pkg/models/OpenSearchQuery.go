package models

type OpenSearchQueryStruct struct {
	Size  int `json:"size"`
	Query any `json:"query"`
}

type OpenSearchMultiMatchQuery struct {
	MultiMatch OpenSearchMultiMatch `json:"multi_match"`
}

type OpenSearchFunctionScoreQuery struct {
	FunctionScore OpenSearchFunctionScore `json:"function_score"`
}

type OpenSearchFunctionScore struct {
	Functions []any `json:"functions"`
}

type OpenSearchMultiMatch struct {
	Query  string   `json:"query"`
	Fields []string `json:"fields"`
}

type OpenSearchRandomScoreFunction struct {
	RandomScore OpenSearchRandomScore `json:"random_score"`
}

type OpenSearchRandomScore struct {
	Seed int64 `json:"seed"`
}
