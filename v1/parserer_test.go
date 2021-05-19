package v1

import (
	"context"
	"encoding/json"
	mlog "github.com/IvanWhisper/michelangelo/log"
	"github.com/IvanWhisper/vtl/expr"
	"testing"
)

var count int64

func TestParse2Tree(t *testing.T) {
	mlog.New(nil)
	ctx := context.Background()
	src := `
BOOL XIV-LXJ = (XIV-LXJ2 * .N.XIV-XJCUT + XIV-LXJ * .N.XIV-LXJT *
                .N.XIV-LQJ * .N.XIV-LZQA * .N.XIV-LZRAJ *
                .N.XIV-XJCUT)
`
	tree := Parse2Tree(ctx, src)
	f := make([]*expr.ExprItem, 0)
	expr.Flattening(tree, &f, 0, 0, func() int64 {
		count++
		return count
	})
	var rawJson string
	if bs, err := json.Marshal(tree); err == nil {
		rawJson = string(bs)
		println("raw-Json: ", string(bs))
	} else {
		println(err)
	}
	var lastJson string
	root := expr.Tree(f)
	if bs, err := json.Marshal(root); err == nil {
		lastJson = string(bs)
		println("T-Json: ", string(bs))
	} else {
		println(err)
	}
	if rawJson != lastJson {
		t.Error("结果不一致")
	}

}
