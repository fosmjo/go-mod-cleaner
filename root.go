/*
Copyright © 2023 fosmjo <imefangjie@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var modfilePaths []string

var rootCmd = &cobra.Command{
	Use: "go-mod-cleaner",
	RunE: func(cmd *cobra.Command, args []string) error {
		modCachePath := filepath.Join(os.Getenv("GOPATH"), "pkg", "mod")
		cleaner := NewCleaner(modCachePath, modfilePaths)
		return cleaner.Clean()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringSliceVarP(
		&modfilePaths, "modfile", "m", nil,
		"modfile paths or dirs, mods referenced by these modfiles will not be removed",
	)
	err := rootCmd.MarkFlagRequired("modfile")
	if err != nil {
		panic(err)
	}
}
