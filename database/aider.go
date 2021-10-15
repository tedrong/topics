package database

type ClientType int

const (
	Desktop ClientType = iota
	Mobile
)

func ClientIndex() []string {
	return []string{"desktop", "mobile"}
}
