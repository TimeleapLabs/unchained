// Code generated by ent, DO NOT EDIT.

package correctnessreport

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the correctnessreport type in the database.
	Label = "correctness_report"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSignersCount holds the string denoting the signerscount field in the database.
	FieldSignersCount = "signers_count"
	// FieldTimestamp holds the string denoting the timestamp field in the database.
	FieldTimestamp = "timestamp"
	// FieldSignature holds the string denoting the signature field in the database.
	FieldSignature = "signature"
	// FieldHash holds the string denoting the hash field in the database.
	FieldHash = "hash"
	// FieldTopic holds the string denoting the topic field in the database.
	FieldTopic = "topic"
	// FieldCorrect holds the string denoting the correct field in the database.
	FieldCorrect = "correct"
	// FieldConsensus holds the string denoting the consensus field in the database.
	FieldConsensus = "consensus"
	// FieldVoted holds the string denoting the voted field in the database.
	FieldVoted = "voted"
	// EdgeSigners holds the string denoting the signers edge name in mutations.
	EdgeSigners = "signers"
	// Table holds the table name of the correctnessreport in the database.
	Table = "correctness_reports"
	// SignersTable is the table that holds the signers relation/edge. The primary key declared below.
	SignersTable = "correctness_report_signers"
	// SignersInverseTable is the table name for the Signer entity.
	// It exists in this package in order to avoid circular dependency with the "signer" package.
	SignersInverseTable = "signers"
)

// Columns holds all SQL columns for correctnessreport fields.
var Columns = []string{
	FieldID,
	FieldSignersCount,
	FieldTimestamp,
	FieldSignature,
	FieldHash,
	FieldTopic,
	FieldCorrect,
	FieldConsensus,
	FieldVoted,
}

var (
	// SignersPrimaryKey and SignersColumn2 are the table columns denoting the
	// primary key for the signers relation (M2M).
	SignersPrimaryKey = []string{"correctness_report_id", "signer_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// SignatureValidator is a validator for the "signature" field. It is called by the builders before save.
	SignatureValidator func([]byte) error
	// HashValidator is a validator for the "hash" field. It is called by the builders before save.
	HashValidator func([]byte) error
	// TopicValidator is a validator for the "topic" field. It is called by the builders before save.
	TopicValidator func([]byte) error
	// DefaultConsensus holds the default value on creation for the "consensus" field.
	DefaultConsensus bool
)

// OrderOption defines the ordering options for the CorrectnessReport queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySignersCountField orders the results by the signersCount field.
func BySignersCountField(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSignersCount, opts...).ToFunc()
}

// ByTimestamp orders the results by the timestamp field.
func ByTimestamp(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimestamp, opts...).ToFunc()
}

// ByCorrect orders the results by the correct field.
func ByCorrect(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCorrect, opts...).ToFunc()
}

// ByConsensus orders the results by the consensus field.
func ByConsensus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConsensus, opts...).ToFunc()
}

// ByVoted orders the results by the voted field.
func ByVoted(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVoted, opts...).ToFunc()
}

// BySignersCount orders the results by signers count.
func BySignersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSignersStep(), opts...)
	}
}

// BySigners orders the results by signers terms.
func BySigners(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSignersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newSignersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SignersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, SignersTable, SignersPrimaryKey...),
	)
}
