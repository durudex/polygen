/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package codegen

import (
	"context"

	"github.com/durudex/polygen/config"

	"github.com/durudex/go-polybase"
)

type Codegen interface {
	Generate() error
}

type codegen struct {
	cfg  *config.Config
	coll polybase.Collection
}

func New(cfg *config.Config) Codegen {
	client := polybase.New(polybase.Config{
		URL: polybase.TestnetURL,
	})

	return &codegen{
		cfg:  cfg,
		coll: client.Collection(GenesisCollectionID),
	}
}

func (c *codegen) Generate() error {
	for _, id := range c.cfg.Collections {
		ast, err := c.astCollection(context.Background(), id)
		if err != nil {
			return err
		}

		c.parseCollection(ast)

		// TODO
	}

	return nil
}
