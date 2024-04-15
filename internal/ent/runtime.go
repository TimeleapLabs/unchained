// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/KenshiTech/unchained/internal/ent/assetprice"
	"github.com/KenshiTech/unchained/internal/ent/correctnessreport"
	"github.com/KenshiTech/unchained/internal/ent/eventlog"
	"github.com/KenshiTech/unchained/internal/ent/schema"
	"github.com/KenshiTech/unchained/internal/ent/signer"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	assetpriceFields := schema.AssetPrice{}.Fields()
	_ = assetpriceFields
	// assetpriceDescSignature is the schema descriptor for signature field.
	assetpriceDescSignature := assetpriceFields[3].Descriptor()
	// assetprice.SignatureValidator is a validator for the "signature" field. It is called by the builders before save.
	assetprice.SignatureValidator = assetpriceDescSignature.Validators[0].(func([]byte) error)
	// assetpriceDescConsensus is the schema descriptor for consensus field.
	assetpriceDescConsensus := assetpriceFields[7].Descriptor()
	// assetprice.DefaultConsensus holds the default value on creation for the consensus field.
	assetprice.DefaultConsensus = assetpriceDescConsensus.Default.(bool)
	correctnessreportFields := schema.CorrectnessReport{}.Fields()
	_ = correctnessreportFields
	// correctnessreportDescSignature is the schema descriptor for signature field.
	correctnessreportDescSignature := correctnessreportFields[2].Descriptor()
	// correctnessreport.SignatureValidator is a validator for the "signature" field. It is called by the builders before save.
	correctnessreport.SignatureValidator = correctnessreportDescSignature.Validators[0].(func([]byte) error)
	// correctnessreportDescHash is the schema descriptor for hash field.
	correctnessreportDescHash := correctnessreportFields[3].Descriptor()
	// correctnessreport.HashValidator is a validator for the "hash" field. It is called by the builders before save.
	correctnessreport.HashValidator = correctnessreportDescHash.Validators[0].(func([]byte) error)
	// correctnessreportDescTopic is the schema descriptor for topic field.
	correctnessreportDescTopic := correctnessreportFields[4].Descriptor()
	// correctnessreport.TopicValidator is a validator for the "topic" field. It is called by the builders before save.
	correctnessreport.TopicValidator = correctnessreportDescTopic.Validators[0].(func([]byte) error)
	eventlogFields := schema.EventLog{}.Fields()
	_ = eventlogFields
	// eventlogDescSignature is the schema descriptor for signature field.
	eventlogDescSignature := eventlogFields[2].Descriptor()
	// eventlog.SignatureValidator is a validator for the "signature" field. It is called by the builders before save.
	eventlog.SignatureValidator = eventlogDescSignature.Validators[0].(func([]byte) error)
	// eventlogDescTransaction is the schema descriptor for transaction field.
	eventlogDescTransaction := eventlogFields[7].Descriptor()
	// eventlog.TransactionValidator is a validator for the "transaction" field. It is called by the builders before save.
	eventlog.TransactionValidator = eventlogDescTransaction.Validators[0].(func([]byte) error)
	// eventlogDescConsensus is the schema descriptor for consensus field.
	eventlogDescConsensus := eventlogFields[9].Descriptor()
	// eventlog.DefaultConsensus holds the default value on creation for the consensus field.
	eventlog.DefaultConsensus = eventlogDescConsensus.Default.(bool)
	signerFields := schema.Signer{}.Fields()
	_ = signerFields
	// signerDescName is the schema descriptor for name field.
	signerDescName := signerFields[0].Descriptor()
	// signer.NameValidator is a validator for the "name" field. It is called by the builders before save.
	signer.NameValidator = signerDescName.Validators[0].(func(string) error)
	// signerDescKey is the schema descriptor for key field.
	signerDescKey := signerFields[2].Descriptor()
	// signer.KeyValidator is a validator for the "key" field. It is called by the builders before save.
	signer.KeyValidator = func() func([]byte) error {
		validators := signerDescKey.Validators
		fns := [...]func([]byte) error{
			validators[0].(func([]byte) error),
			validators[1].(func([]byte) error),
		}
		return func(key []byte) error {
			for _, fn := range fns {
				if err := fn(key); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// signerDescShortkey is the schema descriptor for shortkey field.
	signerDescShortkey := signerFields[3].Descriptor()
	// signer.ShortkeyValidator is a validator for the "shortkey" field. It is called by the builders before save.
	signer.ShortkeyValidator = func() func([]byte) error {
		validators := signerDescShortkey.Validators
		fns := [...]func([]byte) error{
			validators[0].(func([]byte) error),
			validators[1].(func([]byte) error),
		}
		return func(shortkey []byte) error {
			for _, fn := range fns {
				if err := fn(shortkey); err != nil {
					return err
				}
			}
			return nil
		}
	}()
}
