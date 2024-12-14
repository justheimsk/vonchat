package domain

type CachePersistence interface {
	GetUserStatus(string) string
	SetUserStatus(string, string)
}
