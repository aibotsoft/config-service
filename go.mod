module github.com/aibotsoft/config-service

go 1.14

require (
	github.com/aibotsoft/gen v0.0.0-20200510070600-b4df9155b7be
	github.com/aibotsoft/micro v0.0.0-20200507184600-261b9f247278
	github.com/denisenkom/go-mssqldb v0.0.0-20200206145737-bbfc9a55622e
	github.com/dgraph-io/ristretto v0.0.2
	github.com/jmoiron/sqlx v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.5.1
	go.uber.org/zap v1.14.1
	google.golang.org/grpc v1.28.0
)

replace github.com/aibotsoft/micro => ../micro

replace github.com/aibotsoft/gen => ../gen
