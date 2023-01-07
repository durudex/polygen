/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package adapter

import "github.com/durudex/go-polylang/ast"

var TypeToString = map[ast.BasicType]string{
	ast.String: "string", ast.Number: "int", ast.Boolean: "bool",
}
