package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	// Connect to Cassandra Cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	var sessionErr error
	session, sessionErr = cluster.CreateSession()
	if sessionErr != nil {
		panic(sessionErr)
	}

}

func GetSession() *gocql.Session {
	return session
}

