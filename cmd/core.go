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
	"strconv"

	"gospl/nbs/core/GetBank"
	"gospl/nbs/core/GetBankStatus"
	"gospl/nbs/core/GetBankType"
	"gospl/nbs/core/GetCompanyStatus"
	"gospl/nbs/core/GetCompanyType"
	"gospl/nbs/core/GetCurrency"

	"github.com/spf13/cobra"
)

// coreCmd represents the core command
var coreCmd = &cobra.Command{
	Use:     "core",
	Aliases: []string{"c"},
	Short:   "Query nbs core services",
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

var coreBankCmd = &cobra.Command{
	Use:     "bank",
	Aliases: []string{"b"},
	Short:   "Get bank informations",
	RunE: func(cmd *cobra.Command, args []string) error {

		bankCode, err := cmd.Flags().GetInt("bank-code")
		if err != nil {
			return err
		}

		nationalId, err := cmd.Flags().GetInt("national-id")
		if err != nil {
			return err
		}
		request := &GetBank.Request{
			BankID:                       cmd.Flag("id").Value.String(),
			BankCode:                     bankCode,
			NationalIdentificationNumber: nationalId,
		}
		data, err := client.GetBank(request)
		if err != nil {
			return err
		}
		currentData = data
		return nil
	},
}

var coreBankStatusCmd = &cobra.Command{
	Use:                   "status",
	Aliases:               []string{"s", "statuses"},
	Short:                 "Get existing bank statuses",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		request := &GetBankStatus.Request{BankStatusID: 0}
		data, err := client.GetBankStatus(request)
		if err != nil {
			return err
		}
		currentData = data
		return nil
	},
}

var coreBankTypeCmd = &cobra.Command{
	Use:                   "type",
	Aliases:               []string{"types", "t"},
	Short:                 "Get existing bank types",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		request := &GetBankType.Request{BankTypeID: 0}
		data, err := client.GetBankType(request)
		if err != nil {
			return err
		}
		currentData = data
		return nil
	},
}

var coreCurrencyCmd = &cobra.Command{
	Use:                   "currency",
	Aliases:               []string{"c"},
	Short:                 "Get existing currencies",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		request := &GetCurrency.Request{
			CurrencyID:           cmd.Flag("currency-id").Value.String(),
			CurrencyCodeAlfaChar: cmd.Flag("currency-code-alpha").Value.String(),
		}
		if v := cmd.Flag("currency-code").Value.String(); v != "" {
			currencyCode, err := strconv.Atoi(v)
			if err != nil {
				return err
			}
			request.CurrencyCode = currencyCode
		}

		data, err := client.GetCurrency(request)
		if err != nil {
			return err
		}
		currentData = data
		return nil
	},
}

var coreCompanyStatusCmd = &cobra.Command{
	Use:                   "status",
	Aliases:               []string{"s"},
	Short:                 "Get company status",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		request := &GetCompanyStatus.Request{
			CompanyStatusID: 0,
		}

		data, err := client.GetCompanyStatus(request)
		if err != nil {
			return err
		}
		currentData = data
		return nil
	},
}

var coreCompanyTypeCmd = &cobra.Command{
	Use:                   "type",
	Aliases:               []string{"t"},
	Short:                 "Get company type",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		request := &GetCompanyType.Request{
			CompanyTypeID: 0,
		}

		data, err := client.GetCompanyType(request)
		if err != nil {
			return err
		}
		currentData = data
		return nil
	},
}

func init() {
	rootCmd.AddCommand(coreCmd)
	coreCmd.AddCommand(coreBankCmd)
	coreBankCmd.AddCommand(coreBankStatusCmd)
	coreBankCmd.AddCommand(coreBankTypeCmd)
	coreCmd.AddCommand(coreCurrencyCmd)
	coreCmd.AddCommand(coreCompanyStatusCmd)
	coreCmd.AddCommand(coreCompanyTypeCmd)

	coreBankCmd.Flags().StringP("id", "i", "", "Bank ID")
	coreBankCmd.Flags().IntP("bank-code", "c", 0, "Bank code")
	coreBankCmd.Flags().IntP("national-id", "n", 0, "Bank national ID number")
	coreCmd.PersistentFlags().String("out-json", "", "Write results to JSON file")

	coreCurrencyCmd.Flags().StringP("currency-id", "i", "", "Currency ID")
	coreCurrencyCmd.Flags().IntP("currency-code", "c", 0, "Currency code")
	coreCurrencyCmd.Flags().StringP("currency-code-alpha", "a", "", "Currency code alpha (3 chars)")
}
