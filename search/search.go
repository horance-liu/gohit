package search

import (
	"log"
	"sync"
)

func Run(term string) {
	feeds, err := retrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(feeds))

	results := make(chan *Result)
	for _, feed := range feeds {
		go func(matcher Matcher, feed *Feed) {
			match(matcher, feed, term, results)
			waitGroup.Done()
		}(findMatcher(feed.Type), feed)
	}

	go func() {
		waitGroup.Wait()
		close(results)
	}()

	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
