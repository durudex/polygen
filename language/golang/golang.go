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

	"github.com/durudex/polygen/config"
	"github.com/durudex/polygen/language"
	"github.com/durudex/polygen/language/golang/template"
	"github.com/durudex/polygen/parser"

	"github.com/iancoleman/strcase"
)

type golang struct {
	cfg   *config.Go
	names []string
}

func New(cfg *config.Go) language.Codegen {
	return &golang{cfg: cfg}
}

func (g *golang) Generate(coll *parser.ParsedCollection) error {
	path := g.cfg.Directory + "/" + strcase.ToSnake(coll.Name) + "_gen.go"

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	g.names = append(g.names, coll.Name)

	template.WriteHeader(f, g.cfg.Package)
	template.WriteImport(f)
	template.WriteModel(f, coll.Models)
	template.WriteCollection(f, coll.ID, coll.Name, coll.Functions)

	for _, fc := range coll.Functions {
		template.WriteInput(f, coll.Name, fc.Name, fc.Parameters)
		template.WriteFunction(f, coll.Name, fc)
	}

	return nil
}

func (g *golang) Finish() error {
	if err := g.genClient(); err != nil {
		return err
	}

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	return exec.Command("go", "fmt", filepath.Join(dir, g.cfg.Directory)).Run()
}

func (g *golang) genClient() error {
	f, err := os.Create(g.cfg.Directory + "/client_gen.go")
	if err != nil {
		return err
	}
	defer f.Close()

	template.WriteHeader(f, g.cfg.Package)
	template.WriteClient(f, g.names)

	return nil
}
