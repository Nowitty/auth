package user

import (
	"auth/internal/client/db"
	"auth/internal/model"
	"auth/internal/repository"
	"auth/internal/repository/user/converter"
	modelRepo "auth/internal/repository/user/model"
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
)

const (
	tableName       = "auth"
	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	passwordColumn  = "password"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	dbClient db.Client
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{dbClient: db}
}

func (r *repo) Create(ctx context.Context, info *model.UserInfo) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, passwordColumn).
		Values(info.Name, info.Email, "some-password").
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	var id int64
	
	q := db.Query{
		Name: "Create",
		QueryRaw: query,
	}

	err = r.dbClient.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		log.Fatalf("failed to scan auth id: %v", err)
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	builder := sq.Select(idColumn, nameColumn, emailColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name: "Create",
		QueryRaw: query,
	}

	var user modelRepo.User
	err = r.dbClient.DB().ScanOneContext(ctx, &user, q, args...)

	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}
