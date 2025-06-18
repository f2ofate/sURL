package storage

type Repository interface {
	AddShortURL(originUrl string) (string, bool)
	GetOriginURL(shortUrl string) string
}
