/**
 * Created by shiyi on 2017/11/22.
 * Email: shiyi@fightcoder.com
 */

package models

import (
	"self/commons"
	"sync"
)

var once sync.Once

func InitAllInTest() {
	once.Do(func() {
		commons.InitAll("../../../cfg/cfg.toml.debug")
	})
}
