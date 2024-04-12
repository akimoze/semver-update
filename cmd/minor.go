package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var minorCmd = &cobra.Command{
	Use:   "minor",
	Short: "マイナーバージョンを更新します",
	Long:  `マイナーバージョンを更新します。例としてinputが 3.5.1 だった場合 3.6.0 に変更されます`,
	RunE: func(cmd *cobra.Command, args []string) error {
		oldVersion, err := cmd.Flags().GetString("version")
		if oldVersion == "" {
			return fmt.Errorf("フラグ `version` が指定されていません")
		}
		if err != nil {
			return nil
		}

		splitOldVersion := strings.Split(oldVersion, ".")
		oldMinor, err := strconv.Atoi(splitOldVersion[1])
		if err != nil {
			return err
		}
		newMajor := strconv.Itoa(oldMinor + 1)

		newVersion := splitOldVersion[0] + "." + newMajor + ".0"
		fmt.Println(newVersion)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(minorCmd)
	minorCmd.Flags().StringP("version", "v", "", "変更前のバージョンを設定します。")
}
