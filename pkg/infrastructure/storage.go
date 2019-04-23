package infrastructure

import "github.com/krzysztofSkolimowski/url-shrinker/pkg/app"

type Storage interface {
	Save(app.Key, app.URL) error

	Load(app.Key) (app.URL, error)
	Key(app.URL) (app.Key, error)
}
