package pokeapi

import (
	"net/http"
	"time"

	"github.com/almuiz021/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	myCache    *pokecache.Cache
}

func NewClient(timeout, interval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		myCache: pokecache.NewCache(interval),
	}
}
