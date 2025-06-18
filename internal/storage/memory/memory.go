package memory

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
)

type MemStorage map[string]string

// AddShortURL добавляет оригинальную ссылку в хранилище в формате [shortURL]originURL
func (m MemStorage) AddShortURL(originUrl string) (string, bool) {
	var urlRegex = regexp.MustCompile(`^https?://.+$`)

	hash := md5.Sum([]byte(originUrl))
	hashStr := hex.EncodeToString(hash[:4])

	if !urlRegex.MatchString(originUrl) {
		return "", false
	}

	m[hashStr] = originUrl
	return hashStr, true
}

// GetOriginURL получает оригинальную ссылку из сокращённой
func (m MemStorage) GetOriginURL(shortUrl string) string {
	originUrl := m[shortUrl]
	return originUrl
}
