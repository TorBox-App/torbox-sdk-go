// TORBOX_API_KEY=... go run ./examples/quickstart
package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	torbox "github.com/TorBox-App/torbox-sdk-go/v9"
)

func main() {
	ctx := context.Background()
	c := torbox.New(os.Getenv("TORBOX_API_KEY"))

	me, err := torbox.Result(c.UserAPI.GetUserData(ctx).Execute())
	if err != nil {
		panic(err)
	}
	data := me.GetData()
	fmt.Println("signed in as", data.GetEmail())

	_, err = torbox.Result(c.TorrentsAPI.CreateTorrent(ctx).Magnet("magnet:?xt=urn:btih:0000000000000000000000000000000000000000").Execute())
	var apiErr *torbox.Error
	if errors.As(err, &apiErr) {
		fmt.Println("could not add:", apiErr.Status, apiErr.Code, apiErr.Detail)
	}
}
