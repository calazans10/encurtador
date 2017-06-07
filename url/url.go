package url

import (
	"math/rand"
	"net/url"
	"time"
)

const (
	tamanho  = 5
	simbolos = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_-+"
)

type Repositorio interface {
	IDExiste(id string) bool
	BuscarPorID(id string) *URL
	BuscarPorURL(id string) *URL
	Salvar(url URL) error
}

type URL struct {
	ID      string
	Criacao time.Time
	Destino string
}

var repo Repositorio

func init() {
	rand.Seed(time.Now().UnixNano())
}

func ConfigurarRepositorio(r Repositorio) {
	repo = r
}

func BuscarOuCriarNovaURL(destino string) (u *URL, nova bool, err error) {
	if u = repo.BuscarPorURL(destino); u != nil {
		return u, false, nil
	}

	if _, err = url.ParseRequestURI(destino); err != nil {
		return nil, false, err
	}

	url := URL{gerarID(), time.Now(), destino}
	repo.Salvar(url)
	return &url, true, nil
}

func Buscar(id string) *URL {
	return repo.BuscarPorID(id)
}

func gerarID() string {
	novoID := func() string {
		id := make([]byte, tamanho, tamanho)
		for i := range id {
			id[i] = simbolos[rand.Intn(len(simbolos))]
		}
		return string(id)
	}

	for {
		if id := novoID(); !repo.IDExiste(id) {
			return id
		}
	}
}
