package pornhub

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/cormik22/porcelain/src/request"
	"github.com/cormik22/porcelain/src/types"
)

func parseDuration(duration string) int {
	parts := strings.Split(strings.TrimSpace(duration), ":")
	if len(parts) != 2 {
		return 0
	}

	minutes, _ := strconv.Atoi(parts[0])
	seconds, _ := strconv.Atoi(parts[1])

	return minutes*60 + seconds
}

func parseViews(views string) int {
	views = strings.ToLower(strings.TrimSpace(views))

	switch views[len(views)-1] {
	case 'm':
		v, _ := strconv.ParseFloat(views[:len(views)-1], 64)
		return int(v * 1_000_000)
	case 'k':
		v, _ := strconv.ParseFloat(views[:len(views)-1], 64)
		return int(v * 1_000)
	default:
		v, _ := strconv.Atoi(views)
		return v
	}
}

func Search(query string, page int) ([]*types.SearchVideo, error) {
	doc, err := request.Request("https://pornhub.com/video/search?search=" + url.QueryEscape(query) + "&page=" + strconv.Itoa(page))
	if err != nil {
		return nil, err
	}

	videos := []*types.SearchVideo{}

	doc.Find("#videoSearchResult .pcVideoListItem").Each(func(i int, s *goquery.Selection) {
		title := strings.TrimSpace(s.Find(".title").Text())
		url := "https://pornhub.com" + s.Find("a").AttrOr("href", "")
		if url == "" {
			return
		}
		urlParts := strings.Split(url, "=")
		if len(urlParts) != 2 {
			return
		}
		id := urlParts[1]
		channelUrl := s.Find(".usernameWrap a").AttrOr("href", "")
		channelId := strings.Split(channelUrl, "/")[2]
		thumbUrl := s.Find("img").AttrOr("data-mediumthumb", "")
		duration := parseDuration(strings.TrimSpace(s.Find(".duration").Text()))
		views := parseViews(strings.TrimSpace(s.Find(".views var").Text()))

		videos = append(videos, &types.SearchVideo{
			Id:        id,
			Provider:  "pornhub",
			Url:       url,
			Title:     title,
			Thumbnail: thumbUrl,
			EmbedUrl:  "https://www.pornhub.com/embed/" + id,
			Duration:  duration,
			ViewCount: views,
			Uploader:  &channelId,
		})
	})

	return videos, nil
}
