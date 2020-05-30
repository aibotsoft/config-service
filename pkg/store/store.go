package store

import (
	"context"
	pb "github.com/aibotsoft/gen/confpb"
	"github.com/aibotsoft/micro/cache"
	"github.com/aibotsoft/micro/config"
	"github.com/dgraph-io/ristretto"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Store struct {
	cfg   *config.Config
	log   *zap.SugaredLogger
	db    *sqlx.DB
	Cache *ristretto.Cache
}

func NewStore(cfg *config.Config, log *zap.SugaredLogger, db *sqlx.DB) *Store {
	return &Store{log: log, db: db, Cache: cache.NewCache(cfg)}
}

func (s *Store) Close() {
	err := s.db.Close()
	if err != nil {
		s.log.Error(err)
	}
	s.Cache.Close()
}

func (s *Store) GetPortByName(ctx context.Context, serviceName string) (int64, error) {
	get, b := s.Cache.Get(serviceName)
	if b {
		return get.(int64), nil
	}
	var port int64
	err := s.db.GetContext(ctx, &port, "select GrpcPort from dbo.Port where ServiceName = @p1", serviceName)
	if err != nil {
		return 0, err
	}
	s.Cache.Set(serviceName, port, 1)
	return port, nil
}

//func (s *Store) GetAccountByName(ctx context.Context, name string) (pb.Account, error) {
//	var acc pb.Account
//	err := s.db.GetContext(ctx, &acc, "select Id, AccountType, CurrencyCode, ServiceName, Username, Password, Commission, Share from dbo.Account where ServiceName=@p1", name)
//	return acc, err
//}
const accQ = `
select a.Id,
       AccountType,
       CurrencyCode,
       a.ServiceName,
       Username,
       Password,
       Commission,
       Share
from dbo.Account a
join dbo.Port P on a.Id = P.AccountId
where p.ServiceName = @p1
`

func (s *Store) GetAccountByServiceName(ctx context.Context, name string) (acc pb.Account, err error) {
	err = s.db.GetContext(ctx, &acc, accQ, name)
	return
}

func (s *Store) GetCurrency(ctx context.Context) ([]pb.Currency, error) {
	var cur []pb.Currency
	err := s.db.SelectContext(ctx, &cur, "select Id, Code, Value from dbo.Currency")
	return cur, err
}

func (s *Store) GetServices(ctx context.Context) ([]pb.BetService, error) {
	var ser []pb.BetService
	err := s.db.SelectContext(ctx, &ser, "select a.Id, a.ServiceName FortedName, p.ServiceName, p.GrpcPort from dbo.Account a join dbo.Port p on a.Id = p.AccountId")
	return ser, err
}
