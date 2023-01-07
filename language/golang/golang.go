/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package golang

import (
	"os"

	"github.com/durudex/polygen"
	"github.com/durudex/polygen/config"
	"github.com/durudex/polygen/language/golang/template"

	"github.com/iancoleman/strcase"
)

type Golang interface {
	Generate(*polygen.ParsedCollection) error
}

type golang struct{ cfg *config.Go }

func New(cfg *config.Go) Golang {
	return &golang{cfg: cfg}
}

func (g *golang) Generate(coll *polygen.ParsedCollection) error {
	path := g.cfg.Directory + "/" + strcase.ToSnake(coll.Name) + "_gen.go"

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	template.WriteModel(f, coll.Models)

	return nil
}
