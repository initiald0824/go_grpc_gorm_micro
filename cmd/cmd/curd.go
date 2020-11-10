/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"go_grpc_gorm_micro/service"

	"github.com/spf13/cobra"
	"github.com/gookit/color"
)

// curdCmd represents the curd command
var curdCmd = &cobra.Command{
	Use:   "curd",
	Short: "rapid generation of curd through go+grpc+gorm",
	Long: `rapid generation of curd through go+grpc+gorm,The steps are as follows:

1. Design MySQL data structure table
2. Under project file config.yaml, Configure MySQL and configure the connection
3、cd cmd && go run main.go curd tableName`,
	Run: func(cmd *cobra.Command, args []string) {
		tableName,_ := cmd.Flags().GetString("tableName")
		err := service.Generate(tableName, "")
		if err != nil {
			color.Info.Println(err)
			return
		}

		fmt.Printf("tableName generate success")
	},
}

func init() {
	rootCmd.AddCommand(curdCmd)

	curdCmd.Flags().StringP("tableName","t","", "mysql tableName select")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// curdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// curdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
