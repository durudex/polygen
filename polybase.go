/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package polygen

import "github.com/durudex/go-polylang/ast"

type Model struct {
	Name   string
	Fields []*ast.Field
}

type ParsedCollection struct {
	Name      string
	Models    []*Model
	Functions []*ast.Function
}
