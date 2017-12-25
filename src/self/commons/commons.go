/**
 * Created by shiyi on 2017/10/6.
 * Email: shiyi@fightcoder.com
 */

package commons

import (
	"self/commons/g"
	"self/commons/store"
)

func InitAll(confFile string) {
	g.InitConfig(confFile)
	g.InitLog()

	//store.InitMinio()
	store.InitStore()
}

func CloseAll() {
	store.CloseStore()
	g.CloseLog()
}
