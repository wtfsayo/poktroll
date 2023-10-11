package cli

import (
	"strconv"

	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	sessiontypes "pocket/x/session/types"
	"pocket/x/supplier/types"
)

var _ = strconv.Itoa(0)

func CmdSubmitProof() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-proof [session-header] [proof]",
		Short: "Broadcast message submit-proof",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSessionHeader := new(sessiontypes.SessionHeader)
			err = json.Unmarshal([]byte(args[0]), argSessionHeader)
			if err != nil {
				return err
			}
			argSmstProof := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitProof(
				clientCtx.GetFromAddress().String(),
				argSessionHeader,
				argSmstProof,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
