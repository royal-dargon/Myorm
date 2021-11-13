// 用于实现与数据库的交互
package session

import (
	"Myorm/log"
	"database/sql"
	"strings"
)

// Session 结构体目前包含三个成员变量，第一个是sql.Open()方法连接数据库成功后返回的指针
// 第二个和第三个是用来拼接Sql语句的，用户调用raw就可以改变这些值
type Session struct {
	db      *sql.DB
	sql     strings.Builder
	sqlVars []interface{}
}

func New(db *sql.DB) *Session {
	return &Session{db: db}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
}

func (s *Session) DB() *sql.DB {
	return s.db
}

func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

// 下面开始封装Exec()、QueryRow()、QueryRows()三个原生方法
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

// 获取一条记录
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

// 从数据库中获得很多的数据
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

// 封装的目的主要有两个，第一是统一打印日志包括执行的SQL和错误
// 第二是在执行结束了以后，统一清空sql和sqlVals两个变量，复用Session，开启一次会话可以执行多次Sql
