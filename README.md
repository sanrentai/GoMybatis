# SQL mapper framework for Golang
# 中文 https://github.com/zhuxiujia/GoMybatis/blob/master/README.md
[![Go Report Card](https://goreportcard.com/badge/github.com/zhuxiujia/GoMybatis)](https://goreportcard.com/report/github.com/zhuxiujia/GoMybatis)
[![Build Status](https://travis-ci.com/zhuxiujia/GoMybatis.svg?branch=master)](https://travis-ci.com/zhuxiujia/GoMybatis)
[![GoDoc](https://godoc.org/github.com/zhuxiujia/GoMybatis?status.svg)](https://godoc.org/github.com/zhuxiujia/GoMybatis)
[![Coverage Status](https://coveralls.io/repos/github/zhuxiujia/GoMybatis/badge.svg?branch=master)](https://coveralls.io/github/zhuxiujia/GoMybatis?branch=master)
[![codecov](https://codecov.io/gh/zhuxiujia/GoMybatis/branch/master/graph/badge.svg)](https://codecov.io/gh/zhuxiujia/GoMybatis)


![Image text](https://zhuxiujia.github.io/gomybatis.io/assets/vuetify.png)
### website https://zhuxiujia.github.io/gomybatis.io/info.html
# advantage
<a href="https://zhuxiujia.github.io/gomybatis.io/info.html">-High concurrency</a>Assuming that the response time of the database is 0, the number of transactions per second in the framework on a 6-core 16GB PC can reach 246982 Tps/s, which takes only 0.4s.<br>
<a href="https://zhuxiujia.github.io/gomybatis.io/info.html">-Transaction support</a>Session flexible plug-in, compatible with transitional micro-services<br>
<a href="https://zhuxiujia.github.io/gomybatis.io/info.html">-dynamic sql</a>In xml, you can flexibly use`<if>`，`<foreach>`，`<resultMap>,<bind>`and so on. The 15 practical functions of Mybatis are included in the Java framework.<br>
`<select>,<update>,<insert>,<delete>,<trim>,<if>,<set>,<where>,<foreach>,<resultMap>,<bind>,<choose><when><otherwise>,<sql><include>`<br>
<a href="https://zhuxiujia.github.io/gomybatis.io/info.html">-Multi-database support Mysql,Postgres,Tidb,SQLite,Oracle....and more</a><br>
<a href="https://zhuxiujia.github.io/gomybatis.io/info.html">-Quick Start</a>Based on Reflective Dynamic Agent, no intermediate code such as *.go generated by go generator is needed. After reading xml, the function can be called directly.<br>
<a href="https://zhuxiujia.github.io/gomybatis.io/info.html">-Expression and XML object-oriented design</a>If foo. Bar is a pointer, then calling foo. Bar in XML will take the actual value and completely avoid using & and * symbolic pointer operations.<br>
#### Asynchronous Message Queue Logging System
![Image text](https://zhuxiujia.github.io/gomybatis.io/assets/log_system.png)
#### Over-remote substitution for micro-service members to handle transactions supports micro-services during the transition from single database (Mysql, postgresql) to distributed database (TiDB, cockroachdb...)
![Image text](https://zhuxiujia.github.io/gomybatis.io/assets/tx.png)

Database Driven List
```
 Mysql: github.com/go-sql-driver/mysql
 MyMysql: github.com/ziutek/mymysql/godrv
 Postgres: github.com/lib/pq
 Tidb: github.com/pingcap/tidb
 SQLite: github.com/mattn/go-sqlite3
 MsSql: github.com/denisenkom/go-mssqldb
 MsSql: github.com/lunny/godbc
 Oracle: github.com/mattn/go-oci8
 CockroachDB(Postgres): github.com/lib/pq
 ```
 
## Using tutorials

> Sample source code https://github.com/zhuxiujia/GoMybatis/tree/master/example

Set up GoPath and download GoMybatis and the corresponding database driver with the go get command
```
go get github.com/zhuxiujia/GoMybatis
go get github.com/go-sql-driver/mysql
```
Actual use of mapper
```
import (
	_ "github.com/go-sql-driver/mysql" //导入mysql驱动
	"github.com/zhuxiujia/GoMybatis"
	"fmt"
	"time"
)

//Define the content of xml. It is recommended that the * Mapper. XML file be stored in the project directory. When editing xml, you can enjoy IDE rendering and intelligent prompts such as GoLand. Production environments can use statikFS to package XML files into processes

var xmlBytes = []byte(`
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
"https://raw.githubusercontent.com/zhuxiujia/GoMybatis/master/mybatis-3-mapper.dtd">
<mapper namespace="ActivityMapperImpl">
    <!--SelectAll(result *[]Activity)error-->
    <select id="selectAll">
        select * from biz_activity where delete_flag=1 order by create_time desc
    </select>
</mapper>
`)

type ExampleActivityMapperImpl struct {
     SelectAll  func() ([]Activity, error)
}

func main() {
	var err error
	//Mysql link format user name: password @ (database link address: port)/database name, such as root: 123456 @(***.com: 3306)/test
	engine, err := GoMybatis.Open("mysql", "*?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
	   panic(err)
	}
	var exampleActivityMapperImpl ExampleActivityMapperImpl
	
	//Loading XML implementation logic to ExampleActivity Mapper Impl
	GoMybatis.WriteMapperPtrByEngine(&exampleActivityMapperImpl, xmlBytes, engine,true)

	//Using mapper
	result, err := exampleActivityMapperImpl.SelectAll(&result)
        if err != nil {
	   panic(err)
	}
	fmt.Println(result)
}
```
## Welcome to the top right-hand corner of star.
