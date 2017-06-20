package main

import (
	"fmt"
	"github.com/liguoqinjim/behavior3go/config"
	"github.com/liguoqinjim/behavior3go/core"
	"github.com/liguoqinjim/behavior3go/loader"
)

func main() {
	treeConfig, ok := config.LoadTreeCfg("tree.json")
	if ok {
		tree := loader.CreateBevTreeFromConfig(treeConfig, nil)
		fmt.Println("\n\n")

		board := core.NewBlackboard()

		//for i := 0; i < 5; i++ {
		//	tree.Tick(i, board)
		//}
		i := 0
		tree.Tick(i, board)
	} else {
		fmt.Println("LoadTreeCfg error")
	}
}
