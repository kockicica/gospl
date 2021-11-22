/*
Copyright © 2021 kockicica

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate few things",
}

var generateDocCmd = &cobra.Command{
	Use: "markdown [directory]",
	Long: `Create markdown usage documentation.
If directory argument is specified, markdown files will be generated there
`,
	Aliases: []string{"md", "doc"},
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var docPath string
		if len(args) == 0 {
			docPath = "./doc"
		} else {
			docPath = args[0]
		}
		if !path.IsAbs(docPath) {
			current, err := os.Getwd()
			if err != nil {
				return err
			}
			docPath = path.Join(current, docPath)
			if err := os.MkdirAll(docPath, 0644); err != nil {
				return err
			}
		}
		fmt.Println("Generate in:", docPath)
		err := doc.GenMarkdownTree(rootCmd, docPath)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.AddCommand(generateDocCmd)

}
