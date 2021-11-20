/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"gospl/nbs/debtor/GetDebtor"

	"github.com/spf13/cobra"
)

// debtorCmd represents the debtor command
var debtorCmd = &cobra.Command{
	Use:     "debtor",
	Aliases: []string{"d"},
	Short:   "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		nationalId, err := cmd.Flags().GetInt("national-id")
		if err != nil {
			return err
		}
		bankCode, err := cmd.Flags().GetInt("bank-code")
		if err != nil {
			return err
		}

		accountNumber, err := cmd.Flags().GetInt("account-number")
		if err != nil {
			return err
		}

		controlNumber, err := cmd.Flags().GetInt("control-number")
		if err != nil {
			return err
		}

		request := &GetDebtor.Request{
			TaxIdentificationNumber:      cmd.Flag("tax-id").Value.String(),
			NationalIdentificationNumber: nationalId,
			BankCode:                     bankCode,
			AccountNumber:                accountNumber,
			ControlNumber:                controlNumber,
		}
		data, err := client.GetDebtor(request)
		if err != nil {
			return err
		}
		currentData = data
		return nil
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		if currentData != nil {
			outJson := cmd.Flag("out-json").Value.String()
			if outJson != "" {
				if err := client.WriteJson(currentData, outJson); err != nil {
					return err
				}
			} else {
				currentData.WriteOut()
			}

		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(debtorCmd)

	debtorCmd.Flags().StringP("id", "i", "", "Debtor ID")
	debtorCmd.Flags().IntP("bank-code", "b", 0, "Bank code")
	debtorCmd.Flags().IntP("account-number", "a", 0, "Account number")
	debtorCmd.Flags().IntP("control-number", "c", 0, "Control number")
	debtorCmd.Flags().IntP("national-id", "n", 0, "National ID")
	debtorCmd.Flags().StringP("tax-id", "t", "", "Tax ID")
}
