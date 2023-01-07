/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package codegen

import (
	"context"

	"github.com/durudex/polygen"

	"github.com/alecthomas/participle/v2"
	"github.com/durudex/go-polybase"
	"github.com/durudex/go-polylang"
	"github.com/durudex/go-polylang/ast"
)

const GenesisCollectionID = "Collection"

type GenesisCollection struct {
	Code string `json:"code"`
}

func (c *codegen) astCollection(ctx context.Context, id string) (*ast.Collection, error) {
	var response polybase.SingleResponse[GenesisCollection]

	if err := c.coll.Record(id).Get(ctx, &response); err != nil {
		return nil, err
	}

	parser := participle.MustBuild[ast.Collection](
		participle.Lexer(polylang.Lexer),
	)

	return parser.ParseString("", response.Data.Code)
}

func (c *codegen) parseCollection(v *ast.Collection) *polygen.ParsedCollection {
	collection := &polygen.ParsedCollection{
		Name:   v.Name,
		Models: make([]*polygen.Model, 1),
	}

	collection.Models[0] = &polygen.Model{Name: v.Name}

	for _, item := range v.Items {
		switch {
		case item.Field != nil:
			if item.Field.Type.Object != nil {
				collection.Models = append(collection.Models, &polygen.Model{
					Name:   item.Field.Name,
					Fields: item.Field.Type.Object,
				})

				continue
			}

			collection.Models[0].Fields = append(collection.Models[0].Fields, item.Field)
		case item.Function != nil:
			collection.Functions = append(collection.Functions, &ast.Function{
				Name:       item.Function.Name,
				Parameters: item.Function.Parameters,
			})
		}
	}

	return collection
}
