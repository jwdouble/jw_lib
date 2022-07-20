migrate create -ext sql -dir orm/migration xxx
migrate -database "postgres://postgres:password@150.158.7.96:5432/jwdouble?sslmode=disable" -path "$PATH/jw_base/orm/migration" down

// 看文档得看清楚版本和更新日期
// go mod默认使用v1版本，需要自己确认模块的最新版本是什么