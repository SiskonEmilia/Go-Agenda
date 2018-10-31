// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"Go-Agenda/global"
	"Go-Agenda/model"
	"Go-Agenda/operation"
	"fmt"

	"github.com/spf13/cobra"
)

// deleteUserCmd represents the deleteUser command
var deleteUserCmd = &cobra.Command{
	Use:   "deleteUser",
	Short: "Delete user",
	Long:  `Delete current login user`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Delete current user
		username, err := model.GetCurrentUserName()
		if err != nil {
			global.ErrorLog.Println(err.Error())
			return
		}
		err = operation.DeleteUser(username)
		if err != nil {
			global.ErrorLog.Println(err.Error())
			return
		}
		fmt.Println("Your account has been deleted.")
	},
}

func init() {
	rootCmd.AddCommand(deleteUserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
