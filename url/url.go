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
	RegistrarClick(id string)
	BuscarClicks(id string) int
}

type URL struct {
	ID      string    `json:"id"`
	Criacao time.Time `json:"criacao"`
	Destino string    `json:"destino"`
}

type Stats struct {
	URL    *URL `json:"url"`
	Clicks int  `json:"clicks"`
}

var repo Repositorio

func init() {
	rand.Seed(time.Now().UnixNano())
}

func ConfigurarRepositorio(r Repositorio) {
	repo = r
}

func RegistrarClick(id string) {
	repo.RegistrarClick(id)
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

func (u *URL) Stats() *Stats {
	clicks := repo.BuscarClicks(u.ID)
	return &Stats{u, clicks}
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
