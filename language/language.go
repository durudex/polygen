/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package language

import "github.com/durudex/polygen/parser"

type Codegen interface {
	Generate(*parser.ParsedCollection) error
	Finish() error
}
