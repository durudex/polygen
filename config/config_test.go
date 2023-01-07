/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package config_test

import (
	"reflect"
	"testing"

	"github.com/durudex/polygen/config"
)

func Test_New(t *testing.T) {
	want := &config.Config{
		Collections: []string{"durudex/user", "durudex/post"},
		Language: config.Language{
			Go: &config.Go{
				Package:   "generated",
				Directory: "generated",
			},
		},
	}

	got, err := config.New("fixtures/config.yml")
	if err != nil {
		t.Fatal("error: creating a new config: ", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatal("error: config does not match")
	}
}
