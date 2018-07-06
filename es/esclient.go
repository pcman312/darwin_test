package es

import (
	E "gopkg.in/olivere/elastic.v3"
)

type ESClient struct {
	ES *E.Client
}

func (e ESClient) DoESStuff() (err error) {
	search := e.ES.Search().
		Index("some-index").
		Type("some type")

	bq := E.NewBoolQuery()

	results, err := search.Query(bq).
		Size(1).
		Do()
	if E.IsNotFound(err) {
		return err
	}
	if results.TotalHits() == 0 {
		return err
	}

	return nil
}
