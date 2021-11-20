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

	"gospl/nbs/exchange/GetCurrentExchangeRate"
	"gospl/nbs/exchange/GetExchangeRateByCurrency"
	"gospl/nbs/exchange/GetExchangeRateListType"

	"github.com/spf13/cobra"
)

// exchangeCmd represents the exchange command
var exchangeCmd = &cobra.Command{
	Use:   "exchange",
	Short: "Exchange rate query commands",
	Long:  `Commands used for querying exchange rate data`,
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

var currentRsdEurCmd = &cobra.Command{
	Use:   "current-rsd-eur",
	Short: "Get current RSD to EUR exchange rate",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {

		data, err := client.GetCurrentExchangeRateRsdEur()
		if err != nil {
			return err
		}
		currentData = data
		return nil

	},
}

var exchangeRateByCurrencyCmd = &cobra.Command{
	Use:   "by-currency",
	Short: "Get exchange rates by currency",
	Long: `Allows querying exchange rates list by currency, using date ranges
(remember, format for dates is YYYYMMDD)
`,
	Example: `
	Get exchange rates for GBP (currency code: 826), exchange list type 1 (see help for 'list-type' subcommand), 
date range from Nov. 1st to Nov. 15th, 2021:

gospl exchange by-currency --currency-code 826 --list-type-id 1  --date-from 20211101 --date-to 20211115

`,
	RunE: func(cmd *cobra.Command, args []string) error {

		query := GetExchangeRateByCurrency.Request{}
		listTypeId, err := cmd.Flags().GetInt("list-type-id")
		if err != nil {
			return err
		}
		query.ExchangeRateListTypeID = listTypeId
		currencyCode, err := cmd.Flags().GetInt("currency-code")
		if err != nil {
			return err
		}
		query.CurrencyCode = currencyCode
		query.DateFrom = cmd.Flag("date-from").Value.String()
		query.DateTo = cmd.Flag("date-to").Value.String()

		data, err := client.GetExchangeRateByCurrency(&query)
		if err != nil {
			return err
		}
		currentData = data
		return nil
	},
}

var exchangeRateCurrentRateCmd = &cobra.Command{
	Use:   "current-rate <exchangeListTypeID>",
	Short: "Get current exchange rates",
	Long: `Get current exchange rates for specific exchange list type.
Available exchange list types may be queried using 'list-types' command 
`,
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		listTypeId, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		request := &GetCurrentExchangeRate.Request{ExchangeRateListTypeID: listTypeId}
		data, err := client.GetCurrentExchangeRate(request)
		if err != nil {
			return err
		}
		currentData = data
		return nil

	},
}

var exchangeRateListTypeCmd = &cobra.Command{
	Use:     "list-types",
	Aliases: []string{"lt"},
	Short:   "Get exchange rate list types",
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		listType := 0
		if len(args) == 1 {
			listType, err = strconv.Atoi(args[0])
			if err != nil {
				return err
			}
		}
		request := &GetExchangeRateListType.Request{ExchangeRateListTypeID: listType}
		data, err := client.GetExchangeRateListType(request)
		if err != nil {
			return err
		}
		currentData = data
		return nil
	},
}

func init() {
	rootCmd.AddCommand(exchangeCmd)
	exchangeCmd.PersistentFlags().String("out-json", "", "Write results to JSON file")

	exchangeCmd.AddCommand(currentRsdEurCmd)
	exchangeCmd.AddCommand(exchangeRateByCurrencyCmd)
	exchangeCmd.AddCommand(exchangeRateCurrentRateCmd)
	exchangeCmd.AddCommand(exchangeRateListTypeCmd)

	exchangeRateByCurrencyCmd.Flags().Int("list-type-id", 1, "List type id")
	exchangeRateByCurrencyCmd.Flags().Int("currency-code", 978, "Currency code")
	exchangeRateByCurrencyCmd.Flags().String("date-from", "", "From date (format YYYYMMDD)")
	exchangeRateByCurrencyCmd.Flags().String("date-to", "", "From date format(YYYYMMDD)")

}
