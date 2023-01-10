/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		Collections []string `yaml:"collection"`
		PolyGen     PolyGen  `yaml:"polygen"`
		Language    Language `yaml:"language"`
	}

	PolyGen struct {
		ClearDirectory bool `yaml:"clear-directory"`
	}

	Language struct {
		Go *Go `yaml:"go"`
	}

	Go struct {
		Package   string `yaml:"package"`
		Directory string `yaml:"directory"`
	}
)

func New(filename string) (*Config, error) {
	var cfg Config

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
