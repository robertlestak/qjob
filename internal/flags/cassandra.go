package flags

var (
	CassandraHosts          = FlagSet.String("cassandra-hosts", "", "Cassandra hosts")
	CassandraUser           = FlagSet.String("cassandra-user", "", "Cassandra user")
	CassandraPassword       = FlagSet.String("cassandra-password", "", "Cassandra password")
	CassandraKeyspace       = FlagSet.String("cassandra-keyspace", "", "Cassandra keyspace")
	CassandraConsistency    = FlagSet.String("cassandra-consistency", "QUORUM", "Cassandra consistency")
	CassandraRetrieveQuery  = FlagSet.String("cassandra-retrieve-query", "", "Cassandra retrieve query")
	CassandraRetrieveParams = FlagSet.String("cassandra-retrieve-params", "", "Cassandra retrieve params")
	CassandraClearQuery     = FlagSet.String("cassandra-clear-query", "", "Cassandra clear query")
	CassandraClearParams    = FlagSet.String("cassandra-clear-params", "", "Cassandra clear params")
	CassandraFailQuery      = FlagSet.String("cassandra-fail-query", "", "Cassandra fail query")
	CassandraFailParams     = FlagSet.String("cassandra-fail-params", "", "Cassandra fail params")
	CassandraQueryKey       = FlagSet.Bool("cassandra-query-key", false, "Cassandra query returns key as first column and value as second column")
)
