package pornhub

import (
	"errors"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/cormik22/porcelain/src/request"
	"github.com/cormik22/porcelain/src/types"
)

func Qualify(video *types.SearchVideo) (*types.Video, error) {
	doc, err := request.Request("https://pornhub.com/view_video.php?viewkey=" + video.Id)
	if err != nil {
		return nil, err
	}

	qualifiedVideo := &types.Video{
		Id:        video.Id,
		Provider:  video.Provider,
		Url:       video.Url,
		Title:     video.Title,
		Thumbnail: video.Thumbnail,
		EmbedUrl:  video.EmbedUrl,
		Duration:  video.Duration,
		ViewCount: video.ViewCount,
	}

	// The uploader is always provided for PornHub, so we can safely assume it's there
	if video.Uploader != nil {
		qualifiedVideo.Uploader = *video.Uploader
	} else {
		return nil, errors.New("uploader not found")
	}

	page := doc.Find(".video-wrapper").First()

	likesRaw := page.Find(".votesUp").AttrOr("data-rating", "0")
	likes, _ := strconv.Atoi(likesRaw)
	dislikesRaw := page.Find(".votesDown").AttrOr("data-rating", "0")
	dislikes, _ := strconv.Atoi(dislikesRaw)

	qualifiedVideo.Likes = likes
	qualifiedVideo.Dislikes = dislikes

	categories := []string{}

	page.Find(".categoriesWrapper a").Each(func(i int, s *goquery.Selection) {
		trimmed := strings.TrimSpace(s.Text())
		if trimmed == "" {
			return
		}

		if trimmed == "Suggest" {
			return
		}

		categories = append(categories, trimmed)
	})

	qualifiedVideo.Categories = categories

	tags := []string{}

	page.Find(".tagsWrapper a").Each(func(i int, s *goquery.Selection) {
		trimmed := strings.TrimSpace(s.Text())
		if trimmed == "" {
			return
		}

		if trimmed == "Suggest" {
			return
		}

		tags = append(tags, trimmed)
	})

	qualifiedVideo.Tags = tags

	pornstars := []string{}

	page.Find(".pornstarsWrapper a").Each(func(i int, s *goquery.Selection) {
		trimmed := strings.TrimSpace(s.Text())
		if trimmed == "" {
			return
		}

		if trimmed == "Suggest" {
			return
		}

		pornstars = append(pornstars, trimmed)
	})

	qualifiedVideo.Pornstars = pornstars

	return qualifiedVideo, nil
}
