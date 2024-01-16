package v19

import (
	"encoding/json"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	banktestutil "github.com/cosmos/cosmos-sdk/x/bank/testutil"

	"github.com/CosmosContracts/juno/v19/app/keepers"
)

const (
	// based off some mintscan data, this is JUST for testing. (ujuno)
	AvaliableCoins     = 2176229301554
	DelegatedCoins     = 6803271425657
	StakingRewardCoins = 909613203388
)

// junod q auth account juno190g5j8aszqhvtg7cprmev8xcxs6csra7xnk3n3 --output=json
// returns the account and remaining unvested
func CreateMainnetVestingAccount(ctx sdk.Context, k keepers.AppKeepers) (*vestingtypes.PeriodicVestingAccount, math.Int) {
	// manually convert the account_number, sequence, start_time, end_time to a non string
	// regex: "length":"([0-9]+)"  -> "length":$1

	str := `{"@type":"/cosmos.vesting.v1beta1.PeriodicVestingAccount","base_vesting_account":{"base_account":{"address":"juno190g5j8aszqhvtg7cprmev8xcxs6csra7xnk3n3","pub_key":{"@type":"/cosmos.crypto.multisig.LegacyAminoPubKey","threshold":3,"public_keys":[{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A+hZFIZ3i+5jD83trhShTmD/bAAXpIyb6tHJDJtiOAn6"},{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"ApcEus7fRKwSRNNs4nlOy62fFH9Ep7lg9DQRsnx9Ht0H"},{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A1SgSrlikj83agLUJPYDuWTjPkw4rPzkWgMMy/5RxANy"},{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A5fjV581bOBJSuHpBkY7ve3uZJFAv4JI14+K+RiHgr30"},{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"Akr1IlYcH7p08W8pdYNkmAEvDRoPOXdOefg56cm3paKm"}]},"account_number":46159,"sequence":33},"original_vesting":[{"denom":"ujuno","amount":"12457737000000"}],"delegated_free":[{"denom":"ujuno","amount":"398745473523"}],"delegated_vesting":[{"denom":"ujuno","amount":"6404764327386"}],"end_time":2006348400},"start_time":1633100400,"vesting_periods":[{"length":0,"amount":[{"denom":"ujuno","amount":"2373341000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"294128000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"294128000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"294128000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"294128000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"294128000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"294128000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"294128000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"294128000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"294128000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"294128000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"294128000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"294128000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"147064000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"147064000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"147064000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"147064000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"147064000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"147064000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"147064000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"147064000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"147064000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"147064000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"147064000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"147064000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"73952000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"73952000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"73952000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"73952000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"73952000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"73952000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"73952000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"73952000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"73952000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"73952000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"73952000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"73952000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"66388000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"66388000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"66388000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"66388000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"66388000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"66388000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"66388000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"66388000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"66388000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"66388000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"66388000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"66388000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"58825000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"58825000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"58825000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"58825000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"58825000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"58825000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"58825000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"58825000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"58825000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"58825000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"58825000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"58825000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"51262000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"51262000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"51262000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"51262000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"51262000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"51262000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"51262000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"51262000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"51262000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"51262000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"51262000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"51262000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"43867000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"43867000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"43867000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"43867000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"43867000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"43867000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"43867000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"43867000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"43867000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"43867000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"43867000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"43867000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"36556000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"36556000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"36556000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"36556000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"36556000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"36556000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"36556000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"36556000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"36556000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"36556000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"36556000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"36556000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"29245000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"29245000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"29245000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"29245000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"29245000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"29245000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"29245000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"29245000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"29245000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"29245000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"29245000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"29245000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"21934000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"21934000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"21934000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"21934000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"21934000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"21934000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"21934000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"21934000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"21934000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"21934000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"21934000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"21934000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"14623000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"14623000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"14623000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"14623000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"14623000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"14623000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"14623000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"14623000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"14623000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"14623000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"14623000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"14623000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"2522000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"2522000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"2522000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"2522000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"2522000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"2522000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"2522000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"2522000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"2522000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"2522000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"2522000000"}]},{"length":2592000,"amount":[{"denom":"ujuno","amount":"2526000000"}]}]}`

	var acc vestingtypes.PeriodicVestingAccount
	if err := json.Unmarshal([]byte(str), &acc); err != nil {
		panic(err)
	}

	unvested := SumPeriodVestingAccountsUnvestedTokensAmount(ctx, &acc)

	err := banktestutil.FundAccount(k.BankKeeper, ctx, acc.BaseAccount.GetAddress(), sdk.NewCoins(sdk.NewCoin("ujuno", math.NewInt(AvaliableCoins+DelegatedCoins+StakingRewardCoins))))
	if err != nil {
		panic(err)
	}

	k.AccountKeeper.SetAccount(ctx, &acc)
	return &acc, unvested
}
