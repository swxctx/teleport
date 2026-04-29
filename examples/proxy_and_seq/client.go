package main

import (
	"fmt"

	"github.com/swxctx/teleport/lib/goutil"
	tp "github.com/swxctx/teleport"
	"github.com/swxctx/teleport/socket"
)

func main() {
	tp.SetLoggerLevel("ERROR")

	cli := tp.NewPeer(
		tp.PeerConfig{},
	)
	defer cli.Close()

	sess, err := cli.Dial(":8080")
	if err != nil {
		tp.Fatalf("%v", err)
	}

	var result int
	rerr := sess.Pull("/math/add",
		[]int{1, 2, 3, 4, 5},
		&result,
		tp.WithSeq(newRequestId()),
	).Rerror()

	if rerr != nil {
		tp.Fatalf("%v", rerr)
	}
	tp.Printf("result: %d", result)

	rerr = sess.Push(
		"/chat/say",
		fmt.Sprintf("I get result %d", result),
		socket.WithSetMeta("X-ID", "client-001"),
		tp.WithSeq(newRequestId()),
	)
	if rerr != nil {
		tp.Fatalf("%v", rerr)
	}
}

func newRequestId() string {
	return "uid@" + goutil.URLRandomString(8)
}
