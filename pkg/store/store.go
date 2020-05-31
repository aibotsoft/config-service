package store

import (
	"context"
	pb "github.com/aibotsoft/gen/confpb"
	api "github.com/aibotsoft/gen/pinapi"
	"github.com/aibotsoft/micro/cache"
	"github.com/aibotsoft/micro/config"
	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/dgraph-io/ristretto"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
)

type Store struct {
	cfg   *config.Config
	log   *zap.SugaredLogger
	db    *sqlx.DB
	Cache *ristretto.Cache
}

func New(cfg *config.Config, log *zap.SugaredLogger, db *sqlx.DB) *Store {
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

func (s *Store) GetCurrency(ctx context.Context) (cur []pb.Currency, err error) {
	got, b := s.Cache.Get("GetCurrency")
	if b {
		s.log.Info("got from cache")
		return got.([]pb.Currency), nil
	}
	err = s.db.SelectContext(ctx, &cur, "select Code, Value from dbo.Currency")
	if err != nil {
		return nil, err
	}
	s.Cache.SetWithTTL("GetCurrency", cur, 1, time.Hour)
	return
}

func (s *Store) GetServices(ctx context.Context) (ser []pb.BetService, err error) {
	err = s.db.SelectContext(ctx, &ser, "select a.Id, a.ServiceName FortedName, p.ServiceName, p.GrpcPort from dbo.Account a join dbo.Port p on a.Id = p.AccountId")
	return
}

func (s *Store) SaveCurrency(ctx context.Context, currencyList []api.Currency) error {
	var cur []pb.Currency
	//s.log.Infow("", "", currencyList)
	for i := range currencyList {
		cur = append(cur, pb.Currency{Code: currencyList[i].GetCode(), Value: currencyList[i].GetRate()})
	}
	tvp := mssql.TVP{TypeName: "CurrencyType", Value: cur}
	_, err := s.db.ExecContext(ctx, "dbo.uspSaveCurrency", tvp)
	if err != nil {
		return errors.Wrap(err, "uspSaveCurrency error")
	}
	return nil
}
