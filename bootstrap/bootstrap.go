package bootstrap

import (
	"context"
	"github.com/caarlos0/env/v6"
	"github.com/google/wire"
	"github.com/weplanx/openapi/common"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Provides = wire.NewSet(
	UseMongoDB,
	UseDatabase,
)

// SetValues 初始化配置
func SetValues() (values *common.Values, err error) {
	values = new(common.Values)
	if err = env.Parse(values); err != nil {
		return
	}
	return
}

func UseMongoDB(values *common.Values) (*mongo.Client, error) {
	return mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(values.Database.Uri),
	)
}

func UseDatabase(client *mongo.Client, values *common.Values) (db *mongo.Database) {
	return client.Database(values.Database.DbName)
}
