package common

type AccessTokenLoader interface {
	LoadAccessToken(username string) (*string, error)
	SaveAccessToken(token any, username string) error
}
