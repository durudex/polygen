/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package codegen

import (
	"context"
	"os"

	"github.com/durudex/polygen/config"
	"github.com/durudex/polygen/language/golang"

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
	var gn golang.Golang

	if c.cfg.Language.Go != nil {
		gn = golang.New(c.cfg.Language.Go)

		if err := c.checkDir(c.cfg.Language.Go.Package); err != nil {
			return err
		}
	}

	for _, id := range c.cfg.Collections {
		ast, err := c.astCollection(context.Background(), id)
		if err != nil {
			return err
		}

		parsed := c.parseCollection(ast)

		if c.cfg.Language.Go != nil {
			if err := gn.Generate(parsed); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *codegen) checkDir(path string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		if os.IsExist(err) {
			// TODO: add delete old files by option
			return nil
		}

		return err
	}

	return nil
}
