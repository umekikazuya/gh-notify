package app

type githubClient interface {
	Exec() ([]Model, error)
}
