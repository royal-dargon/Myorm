// 这里主要实现的是交互前的准备和交互后的收尾工作
package myorm

import (
	"Myorm/log"
	"Myorm/session"
	"database/sql"
)

type Engine struct {
	db *sql.DB
}

// 这里主要做了两件事情，分别是连接数据库和检查是否正常连接
func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	// 设置一个ping来确定是否连接成功
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	e = &Engine{db: db}
	log.Info("Connect database success!")
	return
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error(err)
		log.Error("Failed to close database")
	}
	log.Info("Close database success")
}

// 通过实例进行交互
func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db)
}
