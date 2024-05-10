package commondatabase

type IRepository interface {
	Start() error
}

//type RepositoryConstructor func(config *conf.Config, root *model.Repositories) (IRepository, error)
