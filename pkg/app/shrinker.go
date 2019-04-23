package app

type Key string
type URL string

type Shrinker interface {
	ShrinkAndSave(URL, Key) error
}

type Retriever interface {
	Key(URL) (Key, error)
	Decode(Key) (URL, error)
}
