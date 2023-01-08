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
	"github.com/durudex/polygen/language"
	"github.com/durudex/polygen/language/golang"
	"github.com/durudex/polygen/parser"
)

type Codegen interface {
	Generate() error
}

type codegen struct {
	cfg    *config.Config
	parser parser.Parser
	langs  []language.Codegen
}

func New(cfg *config.Config) Codegen {
	return &codegen{cfg: cfg, parser: parser.New()}
}

func (c *codegen) Generate() error {
	if c.cfg.Language.Go != nil {
		c.langs = append(c.langs, golang.New(c.cfg.Language.Go))

		if err := c.checkDir(c.cfg.Language.Go.Directory); err != nil {
			return err
		}
	}

	for _, id := range c.cfg.Collections {
		parsed, err := c.parser.Parse(context.Background(), id)
		if err != nil {
			return err
		}

		for _, lang := range c.langs {
			if err := lang.Generate(parsed); err != nil {
				return err
			}
		}
	}

	for _, lang := range c.langs {
		if err := lang.Finish(); err != nil {
			return err
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
