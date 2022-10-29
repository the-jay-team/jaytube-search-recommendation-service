package models

type OpenSearchQueryResult struct {
	Hits OpenSearchQueryResultHit `json:"hits"`
}

type OpenSearchQueryResultHit struct {
	HitsArray []OpenSearchQueryResultHitArray `json:"hits"`
}

type OpenSearchQueryResultHitArray struct {
	Id     string              `json:"_id"`
	Source OpenSearchVideoData `json:"_source"`
}
