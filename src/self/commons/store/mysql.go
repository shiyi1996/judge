/**
 * Created by shiyi on 2017/10/1.
 * Email: shiyi@fightcoder.com
 */

package store

import (
	"self/commons/g"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	log "github.com/sirupsen/logrus"
)

var OrmWeb *xorm.Engine

func initMysql() {
	var err error

	cfg := g.Conf()

	//Web端数据库Orm引擎
	{
		OrmWeb, err = xorm.NewEngine("mysql", cfg.Mysql.WebAddr)
		fmt.Printf("OrmWeb: %#v\n", OrmWeb)

		if err != nil {
			log.Fatalln("fail to connect mysql ", cfg.Mysql.WebAddr, err)
		}
		OrmWeb.SetMaxIdleConns(cfg.Mysql.MaxIdle)
		OrmWeb.SetMaxOpenConns(cfg.Mysql.MaxIdle)
		if cfg.Mysql.Debug {
			OrmWeb.ShowSQL(true)
			OrmWeb.ShowExecTime(true)
			OrmWeb.Logger().SetLevel(core.LOG_DEBUG)
		} else {
			OrmWeb.Logger().SetLevel(core.LOG_WARNING)
		}
	}
}

func closeMysql() {

}
