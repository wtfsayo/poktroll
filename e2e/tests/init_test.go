package e2e

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	"github.com/pokt-network/poktroll/app"
	"github.com/pokt-network/poktroll/x/application/types"
	apptypes "github.com/pokt-network/poktroll/x/application/types"
	sharedtypes "github.com/pokt-network/poktroll/x/shared/types"
)

var (
	addrRe   *regexp.Regexp
	amountRe *regexp.Regexp

	accNameToAddrMap     = make(map[string]string)
	accAddrToNameMap     = make(map[string]string)
	accNameToAppMap      = make(map[string]apptypes.Application)
	accNameToSupplierMap = make(map[string]sharedtypes.Supplier)

	keyRingFlag = "--keyring-backend=test"
)

func init() {
	addrRe = regexp.MustCompile(`address:\s+(\S+)\s+name:\s+(\S+)`)
	amountRe = regexp.MustCompile(`amount:\s+"(.+?)"\s+denom:\s+upokt`)
}

type suite struct {
	gocuke.TestingT
	pocketd       *pocketdBin
	scenarioState map[string]any // temporary state for each scenario
	cdc           codec.Codec
}

func (s *suite) Before() {
	s.pocketd = new(pocketdBin)
	s.scenarioState = make(map[string]any)
	s.cdc = app.MakeEncodingConfig().Marshaler
	s.buildAddrMap()
	s.buildAppMap()
}

// TestFeatures runs the e2e tests specified in any .features files in this directory
// * This test suite assumes that a LocalNet is running
func TestFeatures(t *testing.T) {
	gocuke.NewRunner(t, &suite{}).Path("*.feature").Run()
}

func (s *suite) TheUserHasThePocketdBinaryInstalled() {
	s.TheUserRunsTheCommand("help")
}

func (s *suite) ThePocketdBinaryShouldExitWithoutError() {
	require.NoError(s, s.pocketd.result.Err)
}

func (s *suite) TheUserRunsTheCommand(cmd string) {
	cmds := strings.Split(cmd, " ")
	res, err := s.pocketd.RunCommand(cmds...)
	s.pocketd.result = res
	if err != nil {
		s.Fatalf("error running command %s: %s", cmd, err)
	}
}

func (s *suite) TheUserShouldBeAbleToSeeStandardOutputContaining(arg1 string) {
	if !strings.Contains(s.pocketd.result.Stdout, arg1) {
		s.Fatalf("stdout must contain %s", arg1)
	}
}

func (s *suite) TheUserSendsUpoktFromAccountToAccount(amount int64, accName1, accName2 string) {
	args := []string{
		"tx",
		"bank",
		"send",
		accNameToAddrMap[accName1],
		accNameToAddrMap[accName2],
		fmt.Sprintf("%dupokt", amount),
		keyRingFlag,
		"-y",
	}
	res, err := s.pocketd.RunCommandOnHost("", args...)
	if err != nil {
		s.Fatalf("error sending upokt: %s", err)
	}
	s.pocketd.result = res
}

func (s *suite) TheAccountHasABalanceGreaterThanUpokt(accName string, amount int64) {
	bal := s.getAccBalance(accName)
	if int64(bal) < amount {
		s.Fatalf("account %s does not have enough upokt: %d < %d", accName, bal, amount)
	}
	s.scenarioState[accName] = bal // save the balance for later
}

func (s *suite) AnAccountExistsFor(accName string) {
	bal := s.getAccBalance(accName)
	s.scenarioState[accName] = bal // save the balance for later
}

func (s *suite) TheAccountBalanceOfShouldBeUpoktThanBefore(accName string, amount int64, condition string) {
	prev, ok := s.scenarioState[accName]
	if !ok {
		s.Fatalf("no previous balance found for %s", accName)
	}

	bal := s.getAccBalance(accName)
	switch condition {
	case "more":
		if bal <= prev.(int) {
			s.Fatalf("account %s expected to have more upokt but: %d <= %d", accName, bal, prev)
		}
	case "less":
		if bal >= prev.(int) {
			s.Fatalf("account %s expected to have less upokt but: %d >= %d", accName, bal, prev)
		}
	default:
		s.Fatalf("unknown condition %s", condition)
	}
}

func (s *suite) TheUserShouldWaitForSeconds(dur int64) {
	time.Sleep(time.Duration(dur) * time.Second)
}

func (s *suite) TheUserStakesAWithUpoktFromTheAccount(actorType string, amount int64, accName string) {
	args := []string{
		"tx",
		actorType,
		fmt.Sprintf("stake-%s", actorType),
		fmt.Sprintf("%dupokt", amount),
		"--from",
		accName,
		keyRingFlag,
		"-y",
	}
	res, err := s.pocketd.RunCommandOnHost("", args...)
	if err != nil {
		s.Fatalf("error staking %s: %s", actorType, err)
	}
	s.pocketd.result = res
}

func (s *suite) TheUserUnstakesAFromTheAccount(actorType string, accName string) {
	args := []string{
		"tx",
		actorType,
		fmt.Sprintf("unstake-%s", actorType),
		"--from",
		accName,
		keyRingFlag,
		"-y",
	}
	res, err := s.pocketd.RunCommandOnHost("", args...)
	if err != nil {
		s.Fatalf("error unstaking %s: %s", actorType, err)
	}
	s.pocketd.result = res
}

func (s *suite) TheForAccountIsNotStaked(actorType, accName string) {
	found, _ := s.getStakedAmount(actorType, accName)
	if found {
		s.Fatalf("account %s should not be staked", accName)
	}
}

func (s *suite) TheForAccountIsStakedWithUpokt(actorType, accName string, amount int64) {
	found, stakeAmount := s.getStakedAmount(actorType, accName)
	if !found {
		s.Fatalf("account %s should be staked", accName)
	}
	if int64(stakeAmount) != amount {
		s.Fatalf("account %s stake amount is not %d", accName, amount)
	}
}

func (s *suite) TheApplicationReceivesASuccessfulRelayResponseSignedBy(appName string, supplierName string) {
	// TODO(#126): Implement this step
}

func (s *suite) TheApplicationIsStakedForService(appName string, serviceId string) {
	for _, serviceConfig := range accNameToAppMap[appName].ServiceConfigs {
		if serviceConfig.Service.Id == serviceId {
			fmt.Println("OLSH")
			return
		}
	}
	s.Fatalf("application %s is not staked for service %s", appName, serviceId)
}

func (s *suite) TheSupplierIsStakedForService(supplierName string, serviceId string) {
	for _, serviceConfig := range accNameToAppMap[appName].ServiceConfigs {
		if serviceConfig.Service.Id == serviceId {
			fmt.Println("OLSH")
			return
		}
	}
	s.Fatalf("application %s is not staked for service %s", appName, serviceId)
}

func (s *suite) TheSupplierIsPartOfTheSessionForApplication(supplierName string, appName string) {
	// TODO(#126): Implement this step
}

func (s *suite) TheApplicationSendsTheSupplierARelayRequestForService(appName string, supplierName string, requestName string, serviceId string) {
	// TODO(#126): Implement this step
}

func (s *suite) getStakedAmount(actorType, accName string) (bool, int) {
	s.Helper()
	args := []string{
		"query",
		actorType,
		fmt.Sprintf("list-%s", actorType),
	}
	res, err := s.pocketd.RunCommandOnHost("", args...)
	if err != nil {
		s.Fatalf("error getting %s: %s", actorType, err)
	}
	s.pocketd.result = res
	found := strings.Contains(res.Stdout, accNameToAddrMap[accName])
	amount := 0
	if found {
		escapedAddress := regexp.QuoteMeta(accNameToAddrMap[accName])
		stakedAmountRe := regexp.MustCompile(`address: ` + escapedAddress + `\s+stake:\s+amount: "(\d+)"`)
		matches := stakedAmountRe.FindStringSubmatch(res.Stdout)
		if len(matches) < 2 {
			s.Fatalf("no stake amount found for %s", accName)
		}
		amount, err = strconv.Atoi(matches[1])
		require.NoError(s, err)
	}
	return found, amount
}

func (s *suite) buildAddrMap() {
	s.Helper()
	res, err := s.pocketd.RunCommand(
		"keys", "list", keyRingFlag,
	)
	if err != nil {
		s.Fatalf("error getting keys: %s", err)
	}
	s.pocketd.result = res
	matches := addrRe.FindAllStringSubmatch(res.Stdout, -1)
	for _, match := range matches {
		name := match[2]
		address := match[1]
		accNameToAddrMap[name] = address
		accAddrToNameMap[address] = name
	}
}

func (s *suite) buildAppMap() {
	s.Helper()
	argsAndFlags := []string{
		"query",
		"application",
		"list-application",
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	res, err := s.pocketd.RunCommandOnHost("", argsAndFlags...)
	if err != nil {
		s.Fatalf("error getting application list: %s", err)
	}
	s.pocketd.result = res
	var resp types.QueryAllApplicationResponse
	responseBz := []byte(strings.TrimSpace(res.Stdout))
	s.cdc.MustUnmarshalJSON(responseBz, &resp)
	for _, app := range resp.Application {
		accNameToAppMap[accAddrToNameMap[app.Address]] = app
	}
	fmt.Println(accNameToAppMap)
}

func (s *suite) getAccBalance(accName string) int {
	s.Helper()
	args := []string{
		"query",
		"bank",
		"balances",
		accNameToAddrMap[accName],
	}
	res, err := s.pocketd.RunCommandOnHost("", args...)
	if err != nil {
		s.Fatalf("error getting balance: %s", err)
	}
	s.pocketd.result = res
	match := amountRe.FindStringSubmatch(res.Stdout)
	if len(match) < 2 {
		s.Fatalf("no balance found for %s", accName)
	}
	found, err := strconv.Atoi(match[1])
	require.NoError(s, err)
	return found
}
