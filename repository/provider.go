package repository

import (
	"github.com/freeddser/devops/config"
	"github.com/freeddser/devops/logging"
)

var log = logging.MustGetLogger()

var (
	mysqldatasource  *DataSource
	mysqlsqldbname   string
	mysqlsqluser     string
	mysqlsqlpassword string
	mysqlsqlhost     string
	mysqlsqlport     string
	redisMain        string
)

func InitMysqlDbFactory() error {

	mysqlsqluser = config.MustGetString("mysql.user")
	mysqlsqlpassword = config.MustGetString("mysql.password")
	mysqlsqldbname = config.MustGetString("mysql.dbname")
	mysqlsqlhost = config.MustGetString("mysql.host")
	mysqlsqlport = config.MustGetString("mysql.port")
	redisMain = config.MustGetString("cache.redis_main")

	var err error
	mysqldatasource, err = NewDatabaseConnection(mysqlsqlhost, mysqlsqlport, mysqlsqluser, mysqlsqlpassword, mysqlsqldbname)
	if err != nil {
		return err
	}
	log.Info("Init Mysql Connection Started")
	InitRedis(redisMain)
	log.Info("Init Redis Connection Started")
	return nil
}
