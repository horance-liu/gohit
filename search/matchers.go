package search

var matchers = make(map[string]Matcher)

func Register(feed string, matcher Matcher) {
	if _, exists := matchers[feed]; !exists {
		matchers[feed] = matcher
	}
}

func findMatcher(feed string) Matcher {
	if matcher, exists := matchers[feed]; exists {
		return matcher
	}
	return matchers["default"]
}

type defaultMatcher struct{}

func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}

func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}