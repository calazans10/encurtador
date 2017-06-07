package url

type repositorioMemoria struct {
	urls map[string]*URL
}

func NovoRepositorioMemoria() *repositorioMemoria {
	return &repositorioMemoria{make(map[string]*URL)}
}

func (r *repositorioMemoria) IDExiste(id string) bool {
	_, existe := r.urls[id]
	return existe
}

func (r *repositorioMemoria) BuscarPorID(id string) *URL {
	return r.urls[id]
}

func (r *repositorioMemoria) BuscarPorURL(url string) *URL {
	for _, u := range r.urls {
		if u.Destino == url {
			return u
		}
	}
	return nil
}

func (r *repositorioMemoria) Salvar(url URL) error {
	r.urls[url.ID] = &url
	return nil
}
