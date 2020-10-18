package service

import (
	"fmt"
	"math"
	"net/url"

	"github.com/evilmonkeyinc/account-manager/gen/server"
)

func buildPagingLinks(host string, currentURL *url.URL, page, limit, total int) server.PagingLinks {
	lastPage := int(math.Ceil(float64(total) / float64(limit)))

	firstURL, _ := url.Parse(currentURL.String())
	firstQuery := firstURL.Query()
	firstQuery.Set("page", "0")
	firstQuery.Set("limit", fmt.Sprintf("%d", limit))
	firstURL.RawQuery = firstQuery.Encode()

	lastURL, _ := url.Parse(currentURL.String())
	lastQuery := lastURL.Query()
	lastQuery.Set("page", fmt.Sprintf("%d", lastPage))
	lastQuery.Set("limit", fmt.Sprintf("%d", limit))
	lastURL.RawQuery = lastQuery.Encode()

	var next, previous *string
	if page < lastPage {
		parsedURL, _ := url.Parse(currentURL.String())
		query := parsedURL.Query()
		query.Set("page", fmt.Sprintf("%d", page+1))
		query.Set("limit", fmt.Sprintf("%d", limit))
		parsedURL.RawQuery = query.Encode()
		str := fmt.Sprintf("%s%s", host, parsedURL.String())
		next = &str
	}
	if page > 0 {
		prev := page - 1
		if page > lastPage {
			prev = lastPage
		}
		parsedURL, _ := url.Parse(currentURL.String())
		query := parsedURL.Query()
		query.Set("page", fmt.Sprintf("%d", prev))
		query.Set("limit", fmt.Sprintf("%d", limit))
		parsedURL.RawQuery = query.Encode()
		str := fmt.Sprintf("%s%s", host, parsedURL.String())
		previous = &str
	}

	return server.PagingLinks{
		Current:  fmt.Sprintf("%s%s", host, currentURL.String()),
		First:    fmt.Sprintf("%s%s", host, firstURL.String()),
		Last:     fmt.Sprintf("%s%s", host, lastURL.String()),
		Next:     next,
		Previous: previous,
	}
}
