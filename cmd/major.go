package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var majorCmd = &cobra.Command{
	Use:   "major",
	Short: "メジャーバージョンを更新します",
	Long:  `メジャーバージョンを更新します。例としてinputが 3.5.1 だった場合 4.0.0 に変更されます`,
	RunE: func(cmd *cobra.Command, args []string) error {
		oldVersion, err := cmd.Flags().GetString("version")
		if oldVersion == "" {
			return fmt.Errorf("フラグ `version` が指定されていません")
		}
		if err != nil {
			return nil
		}

		splitOldVersion := strings.Split(oldVersion, ".")
		oldMajor, err := strconv.Atoi(splitOldVersion[0])
		if err != nil {
			return err
		}
		newMajor := strconv.Itoa(oldMajor + 1)

		newVersion := newMajor + ".0.0"
		fmt.Println(newVersion)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(majorCmd)
	majorCmd.Flags().StringP("version", "v", "", "変更前のバージョンを設定します。")
}
