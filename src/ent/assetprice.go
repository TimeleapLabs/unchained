// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/KenshiTech/unchained/ent/assetprice"
	"github.com/KenshiTech/unchained/ent/helpers"
)

// AssetPrice is the model entity for the AssetPrice schema.
type AssetPrice struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Block holds the value of the "block" field.
	Block uint64 `json:"block,omitempty"`
	// SignersCount holds the value of the "signersCount" field.
	SignersCount *uint64 `json:"signersCount,omitempty"`
	// Price holds the value of the "price" field.
	Price *helpers.BigInt `json:"price,omitempty"`
	// Signature holds the value of the "signature" field.
	Signature []byte `json:"signature,omitempty"`
	// Asset holds the value of the "asset" field.
	Asset string `json:"asset,omitempty"`
	// Chain holds the value of the "chain" field.
	Chain string `json:"chain,omitempty"`
	// Pair holds the value of the "pair" field.
	Pair string `json:"pair,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AssetPriceQuery when eager-loading is set.
	Edges        AssetPriceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AssetPriceEdges holds the relations/edges for other nodes in the graph.
type AssetPriceEdges struct {
	// Signers holds the value of the signers edge.
	Signers []*Signer `json:"signers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedSigners map[string][]*Signer
}

// SignersOrErr returns the Signers value or an error if the edge
// was not loaded in eager-loading.
func (e AssetPriceEdges) SignersOrErr() ([]*Signer, error) {
	if e.loadedTypes[0] {
		return e.Signers, nil
	}
	return nil, &NotLoadedError{edge: "signers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AssetPrice) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case assetprice.FieldSignature:
			values[i] = new([]byte)
		case assetprice.FieldPrice:
			values[i] = new(helpers.BigInt)
		case assetprice.FieldID, assetprice.FieldBlock, assetprice.FieldSignersCount:
			values[i] = new(sql.NullInt64)
		case assetprice.FieldAsset, assetprice.FieldChain, assetprice.FieldPair:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AssetPrice fields.
func (ap *AssetPrice) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case assetprice.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ap.ID = int(value.Int64)
		case assetprice.FieldBlock:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field block", values[i])
			} else if value.Valid {
				ap.Block = uint64(value.Int64)
			}
		case assetprice.FieldSignersCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field signersCount", values[i])
			} else if value.Valid {
				ap.SignersCount = new(uint64)
				*ap.SignersCount = uint64(value.Int64)
			}
		case assetprice.FieldPrice:
			if value, ok := values[i].(*helpers.BigInt); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value != nil {
				ap.Price = value
			}
		case assetprice.FieldSignature:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field signature", values[i])
			} else if value != nil {
				ap.Signature = *value
			}
		case assetprice.FieldAsset:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field asset", values[i])
			} else if value.Valid {
				ap.Asset = value.String
			}
		case assetprice.FieldChain:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field chain", values[i])
			} else if value.Valid {
				ap.Chain = value.String
			}
		case assetprice.FieldPair:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pair", values[i])
			} else if value.Valid {
				ap.Pair = value.String
			}
		default:
			ap.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AssetPrice.
// This includes values selected through modifiers, order, etc.
func (ap *AssetPrice) Value(name string) (ent.Value, error) {
	return ap.selectValues.Get(name)
}

// QuerySigners queries the "signers" edge of the AssetPrice entity.
func (ap *AssetPrice) QuerySigners() *SignerQuery {
	return NewAssetPriceClient(ap.config).QuerySigners(ap)
}

// Update returns a builder for updating this AssetPrice.
// Note that you need to call AssetPrice.Unwrap() before calling this method if this AssetPrice
// was returned from a transaction, and the transaction was committed or rolled back.
func (ap *AssetPrice) Update() *AssetPriceUpdateOne {
	return NewAssetPriceClient(ap.config).UpdateOne(ap)
}

// Unwrap unwraps the AssetPrice entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ap *AssetPrice) Unwrap() *AssetPrice {
	_tx, ok := ap.config.driver.(*txDriver)
	if !ok {
		panic("ent: AssetPrice is not a transactional entity")
	}
	ap.config.driver = _tx.drv
	return ap
}

// String implements the fmt.Stringer.
func (ap *AssetPrice) String() string {
	var builder strings.Builder
	builder.WriteString("AssetPrice(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ap.ID))
	builder.WriteString("block=")
	builder.WriteString(fmt.Sprintf("%v", ap.Block))
	builder.WriteString(", ")
	if v := ap.SignersCount; v != nil {
		builder.WriteString("signersCount=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", ap.Price))
	builder.WriteString(", ")
	builder.WriteString("signature=")
	builder.WriteString(fmt.Sprintf("%v", ap.Signature))
	builder.WriteString(", ")
	builder.WriteString("asset=")
	builder.WriteString(ap.Asset)
	builder.WriteString(", ")
	builder.WriteString("chain=")
	builder.WriteString(ap.Chain)
	builder.WriteString(", ")
	builder.WriteString("pair=")
	builder.WriteString(ap.Pair)
	builder.WriteByte(')')
	return builder.String()
}

// NamedSigners returns the Signers named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ap *AssetPrice) NamedSigners(name string) ([]*Signer, error) {
	if ap.Edges.namedSigners == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ap.Edges.namedSigners[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ap *AssetPrice) appendNamedSigners(name string, edges ...*Signer) {
	if ap.Edges.namedSigners == nil {
		ap.Edges.namedSigners = make(map[string][]*Signer)
	}
	if len(edges) == 0 {
		ap.Edges.namedSigners[name] = []*Signer{}
	} else {
		ap.Edges.namedSigners[name] = append(ap.Edges.namedSigners[name], edges...)
	}
}

// AssetPrices is a parsable slice of AssetPrice.
type AssetPrices []*AssetPrice
