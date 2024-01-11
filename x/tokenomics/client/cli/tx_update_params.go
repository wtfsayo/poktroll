package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/pokt-network/poktroll/x/tokenomics/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdUpdateParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-params",
		Short: "Broadcast message update-params",
		// TODO_IN_THIS_PR: Update `Short` and `Long` to show how parameters could/should be updated
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateParams(
				clientCtx.GetFromAddress().String(),
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