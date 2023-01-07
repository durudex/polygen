/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package golang

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/durudex/polygen"
	"github.com/durudex/polygen/config"
	"github.com/durudex/polygen/language"
	"github.com/durudex/polygen/language/golang/template"

	"github.com/iancoleman/strcase"
)

type golang struct{ cfg *config.Go }

func New(cfg *config.Go) language.Codegen {
	return &golang{cfg: cfg}
}

func (g *golang) Generate(coll *polygen.ParsedCollection) error {
	path := g.cfg.Directory + "/" + strcase.ToSnake(coll.Name) + "_gen.go"

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	template.WriteHeader(f, coll.Name)
	template.WriteModel(f, coll.Models)

	for _, fc := range coll.Functions {
		template.WriteInput(f, coll.Name, fc.Name, fc.Parameters)
	}

	return nil
}

func (g *golang) Finish() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	return exec.Command("go", "fmt", filepath.Join(dir, g.cfg.Directory)).Run()
}
