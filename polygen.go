/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package polygen

import "github.com/durudex/go-polylang/ast"

const GenesisCollectionID = "Collection"

type GenesisCollection struct {
	Code string `json:"code"`
}

type Model struct {
	Name   string
	Fields []*ast.Field
}
