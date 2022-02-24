package enum

type Uf int

const (
	Acre Uf = iota + 1
)

func (u Uf) String() string {
	return []string{"AC"}[u-1]
}

func (u Uf) EnumIndex() int {
	return int(u)
}
