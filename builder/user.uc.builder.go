package builder

import (
	"gorm.io/gorm"
	"to_do_list/common"
	"to_do_list/module/users/infrastructure/repositories/mysql"
	"to_do_list/module/users/usecase/command"
	"to_do_list/module/users/usecase/query"
)

type simpleBuilder struct {
	db *gorm.DB
}

func NewSimpleBuilder(db *gorm.DB) *simpleBuilder {
	return &simpleBuilder{db: db}
}

func (b *simpleBuilder) BuildUserRepository() query.UserRepository {
	return mysql.NewUserMySQLRepo(b.db)
}

func (b *simpleBuilder) BuildUserQueryRepository() query.UserQueryRepository {
	return mysql.NewUserMySQLRepo(b.db)
}

func (b *simpleBuilder) BuildUserCmdRepository() query.UserCmdRepository {
	return mysql.NewUserMySQLRepo(b.db)
}

func (b *simpleBuilder) BuildHasher() command.Hasher {
	return &common.Hasher{}
}
