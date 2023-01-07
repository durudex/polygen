/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package language

import "github.com/durudex/polygen"

type Codegen interface {
	Generate(*polygen.ParsedCollection) error
	Finish() error
}
