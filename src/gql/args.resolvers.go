package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"
	"strings"

	"github.com/KenshiTech/unchained/datasets"
	"github.com/KenshiTech/unchained/gql/generated"
)

// Value is the resolver for the value field.
func (r *eventLogArgResolver) Value(ctx context.Context, obj *datasets.EventLogArg) (string, error) {
	switch {
	case strings.HasPrefix(obj.Type, "uint"), strings.HasPrefix(obj.Type, "int"):
		return obj.Value.(string), nil

	case obj.Type == "bool":
		return fmt.Sprintf("%t", obj.Value), nil

	case obj.Type == "string":
		return obj.Value.(string), nil

	case obj.Type == "address":
		return obj.Value.(string), nil

	default:
		return "", fmt.Errorf("unsupported type: %s", obj.Type)
	}
}

// EventLogArg returns generated.EventLogArgResolver implementation.
func (r *Resolver) EventLogArg() generated.EventLogArgResolver { return &eventLogArgResolver{r} }

type eventLogArgResolver struct{ *Resolver }
