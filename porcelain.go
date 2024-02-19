package porcelain

import (
	"errors"

	"github.com/cormik22/porcelain/src/providers/pornhub"
	"github.com/cormik22/porcelain/src/types"
)

type (
	searchProvider  func(query string, page int) ([]*types.SearchVideo, error)
	qualifyProvider func(video *types.SearchVideo) (*types.Video, error)
)

var (
	searchProviders = []searchProvider{
		pornhub.Search,
	}
	qualifyProviders = map[string]qualifyProvider{
		"pornhub": pornhub.Qualify,
	}
)

func Search(query string, page int) ([]*types.SearchVideo, error) {
	videos := []*types.SearchVideo{}

	for _, provider := range searchProviders {
		providerVideos, err := provider(query, page)
		if err != nil {
			return nil, err
		}

		videos = append(videos, providerVideos...)
	}

	return videos, nil
}

func Qualify(video *types.SearchVideo) (*types.Video, error) {
	provider, ok := qualifyProviders[video.Provider]
	if !ok {
		return nil, errors.New("provider not found")
	}

	return provider(video)
}
