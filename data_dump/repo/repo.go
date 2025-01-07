package repo

type Repo interface {
	Dump(data interface{}) error
}
