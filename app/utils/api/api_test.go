package api_test

import (
	"fmt"
	"testing"

	"github.com/gonebot-dev/gonebuilder-tui/app/utils/api"
)

func TestAPI(t *testing.T) {
	api.SyncRepo()
	fmt.Printf("Plugins: %#v\n", api.Plugins)
	fmt.Printf("Adapters: %#v\n", api.Adapters)
}
