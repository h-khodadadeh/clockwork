package dataSource

import (
	"strings"
)

type QueryLoggerDataSourceInterface interface {
	DataSource
	QueryLoggerInterface
}

type QueryLoggerInterface interface {
	LogQuery(query string, duration float32, bind []string)
}

type mySqlStructure = struct {
	Model string `json:"model"`
	Query string `json:"query"`
	Duration float32 `json:"duration"`
	Connection string `json:"connection"`
	Tags []string `json:"tags"`
}

type MysqlDataSource struct {
	commands []interface{}
	totalDuration float32
}

func (source *MysqlDataSource) LogQuery(query string, duration float32, bind []string)  {
	var tags []string

	if duration > 2 {
		tags = append(tags, "slow")
	}

	structure := mySqlStructure{
		Model: "mysql",
		Query: query + " [" + strings.Join(bind, ", ") + "]",
		Duration: duration,
		Connection: "test-connection",
		Tags: tags,
	}

	source.totalDuration += duration
	source.commands = append(source.commands, &structure)
}

func (source *MysqlDataSource) Resolve(dataBuffer *DataBuffer) {
	dataBuffer.DatabaseQueries = source.commands
	dataBuffer.DatabaseDuration = source.totalDuration
	dataBuffer.DatabaseQueriesCount = len(dataBuffer.DatabaseQueries)
}
