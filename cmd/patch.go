package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var patchCmd = &cobra.Command{
	Use:   "patch",
	Short: "パッチバージョンを更新します",
	Long:  `パッチバージョンを更新します。例としてinputが 3.5.1 だった場合 3.5.2 に変更されます`,
	RunE: func(cmd *cobra.Command, args []string) error {
		oldVersion, err := cmd.Flags().GetString("version")
		if oldVersion == "" {
			return fmt.Errorf("フラグ `version` が指定されていません")
		}
		if err != nil {
			return nil
		}

		splitOldVersion := strings.Split(oldVersion, ".")
		patch := splitOldVersion[2]

		if strings.Contains(patch, "-") {
			patchSplit := strings.Split(patch, "-")
			patch = patchSplit[0]
		}

		if strings.Contains(patch, "+") {
			patchSplit := strings.Split(patch, "+")
			patch = patchSplit[0]
		}

		intPatch, err := strconv.Atoi(patch)
		if err != nil {
			return err
		}
		newPatch := strconv.Itoa(intPatch + 1)

		newVersion := splitOldVersion[0] + "." + splitOldVersion[1] + "." + newPatch
		fmt.Println(newVersion)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(patchCmd)
	patchCmd.Flags().StringP("version", "v", "", "変更前のバージョンを設定します。")
}
