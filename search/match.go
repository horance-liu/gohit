package search

type Result struct {
	Field   string
	Content string
}

type Matcher interface {
	Search(feed *Feed, term string) ([]*Result, error)
}

func match(matcher Matcher, feed *Feed, term string, results chan<- *Result) {
	searchResults, err := matcher.Search(feed, term)
	if err != nil {
		return
	}

	for _, result := range searchResults {
		results <- result
	}
}
