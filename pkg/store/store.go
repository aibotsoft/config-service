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

func (s *Store) GetAccountByName(ctx context.Context, name string) (pb.Account, error) {
	var acc pb.Account
	err := s.db.GetContext(ctx, &acc, "select Id, AccountType, CurrencyCode, ServiceName, Username, Password, Commission, Share from dbo.Account where ServiceName=@p1", name)
	return acc, err
}

func NewStore(cfg *config.Config, log *zap.SugaredLogger, db *sqlx.DB) *Store {
	return &Store{log: log, db: db, Cache: cache.NewCache(cfg)}
}
