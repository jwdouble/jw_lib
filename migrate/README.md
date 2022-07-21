migrate create -ext sql -dir orm/migration xxx
migrate -database "postgres://postgres:password@150.158.7.96:5432/jwdouble?sslmode=disable" -path "$PATH/jw_base/orm/migration" up 1 #升级一个版本

goto V       Migrate to version V
up [N]       Apply all or N up migrations
down [N] [-all]    Apply all or N down migrations
Use -all to apply all down migrations
drop [-f]    Drop everything inside database
Use -f to bypass confirmation
force V      Set version V but don't run migration (ignores dirty state)
version      Print current migration version

