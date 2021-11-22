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
	"strconv"

	"gospl/nbs/account/GetCompanyAccount"
	"gospl/nbs/account/GetCompanyAccountStatus"
	"gospl/nbs/account/GetCompanyAccountType"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// accountCmd represents the company command
var accountCmd = &cobra.Command{
	Use:     "account",
	Aliases: []string{"company", "a"},
	Short:   "Get account info",
	Long: `Search for company account info using one or more of the following fields:
company name, national identification number, tax identification number, bank code, account number,
control number or city.
`,
	Example: `
gospl account --name somename
`,
	Args: cobra.NoArgs,
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
	RunE: func(cmd *cobra.Command, args []string) error {

		query := &GetCompanyAccount.Request{}
		query.CompanyName = viper.GetString("name")
		query.NationalIdentificationNumber = viper.GetInt("national-id")
		query.TaxIdentificationNumber = viper.GetString("tax-id")
		query.AccountNumber = viper.GetInt("account-number")
		query.ControlNumber = viper.GetInt("control-number")
		query.City = viper.GetString("city")
		query.BankCode = viper.GetInt("bank-code")
		client = client.WithContext(cmd.Context())
		client.SetVerbose(verbose)
		data, err := client.GetCompanyAccount(query)
		if err != nil {
			return err
		}
		//fmt.Printf("%#v\n", data)
		outJson := cmd.Flag("out-json").Value.String()
		if outJson != "" {
			if err := client.WriteJson(data, outJson); err != nil {
				return err
			}
		} else {
			data.WriteOut()
		}

		return nil
	},
}

var accountStatusCmd = &cobra.Command{
	Use:                   "status",
	Aliases:               []string{"statuses", "s"},
	Short:                 "Get existing account statuses",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var statusId = 0
		var err error
		if len(args) == 1 {
			statusId, err = strconv.Atoi(args[0])
			if err != nil {
				return err
			}
		}
		request := &GetCompanyAccountStatus.Request{CompanyAccountStatusID: statusId}
		data, err := client.GetCompanyAccountStatus(request)
		if err != nil {
			return err
		}
		currentData = data
		return nil
	},
}

var accountTypeCmd = &cobra.Command{
	Use:                   "type",
	Aliases:               []string{"types", "t"},
	Short:                 "Get existing account types",
	DisableFlagsInUseLine: true,
	Args:                  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var typeId = 0
		var err error
		if len(args) == 1 {
			typeId, err = strconv.Atoi(args[0])
			if err != nil {
				return err
			}
		}
		request := &GetCompanyAccountType.Request{CompanyAccountTypeID: typeId}
		data, err := client.GetCompanyAccountType(request)
		if err != nil {
			return err
		}
		currentData = data
		return nil
	},
}

func init() {
	rootCmd.AddCommand(accountCmd)

	accountCmd.Flags().StringP("name", "n", "", "Search for company name")
	_ = viper.BindPFlag("name", accountCmd.Flags().Lookup("name"))
	accountCmd.Flags().Int("national-id", 0, "Search for company national id")
	_ = viper.BindPFlag("national-id", accountCmd.Flags().Lookup("national-id"))
	accountCmd.Flags().StringP("tax-id", "t", "", "Search for company tax id")
	_ = viper.BindPFlag("tax-id", accountCmd.Flags().Lookup("tax-id"))
	accountCmd.Flags().Int("bank-code", 0, "Search for bank code")
	_ = viper.BindPFlag("bank-code", accountCmd.Flags().Lookup("bank-code"))
	accountCmd.Flags().Int("account-number", 0, "Search for account number")
	_ = viper.BindPFlag("account-number", accountCmd.Flags().Lookup("account-number"))
	accountCmd.Flags().Int("control-number", 0, "Search for control number")
	_ = viper.BindPFlag("control-number", accountCmd.Flags().Lookup("control-number"))
	accountCmd.Flags().String("city", "", "Search for city")
	_ = viper.BindPFlag("city", accountCmd.Flags().Lookup("city"))

	accountCmd.PersistentFlags().String("out-json", "", "Write results to JSON file")

	accountCmd.AddCommand(accountStatusCmd)
	accountCmd.AddCommand(accountTypeCmd)
}
