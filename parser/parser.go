/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package parser

import (
	"context"

	"github.com/durudex/polygen"

	"github.com/alecthomas/participle/v2"
	"github.com/durudex/go-polybase"
	"github.com/durudex/go-polylang"
	"github.com/durudex/go-polylang/ast"
)

type ParsedCollection struct {
	ID        string
	Name      string
	Models    []*polygen.Model
	Functions []*ast.Function
}

type Parser interface {
	Parse(ctx context.Context, id string) (*ParsedCollection, error)
}

type parser struct {
	coll polybase.Collection[polygen.GenesisCollection]
}

func New() Parser {
	client := polybase.New(&polybase.Config{
		URL: polybase.TestnetURL,
	})
	coll := polybase.NewCollection[polygen.GenesisCollection](
		client,
		polygen.GenesisCollectionID,
	)

	return &parser{coll: coll}
}

func (p *parser) Parse(ctx context.Context, id string) (*ParsedCollection, error) {
	collAst, err := p.astCollection(ctx, id)
	if err != nil {
		return nil, err
	}

	collection := &ParsedCollection{
		ID:     id,
		Name:   collAst.Name,
		Models: make([]*polygen.Model, 1),
	}

	collection.Models[0] = &polygen.Model{Name: collAst.Name}

	for _, item := range collAst.Items {
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

	return collection, nil
}

func (p *parser) astCollection(ctx context.Context, id string) (*ast.Collection, error) {
	response := p.coll.Record(id).Get(ctx)

	parser := participle.MustBuild[ast.Collection](
		participle.Lexer(polylang.Lexer),
	)

	return parser.ParseString("", response.Data.Code)
}
