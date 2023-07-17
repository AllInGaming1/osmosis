package types

import (
	fmt "fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys.
var (
	KeyMinimumRiskFactor                         = []byte("MinimumRiskFactor")
	KeyForceSuperfluidUndelegateAllowedAddresses = []byte("ForceSuperfluidUndelegateAllowedAddresses")
	defaultMinimumRiskFactor                     = sdk.NewDecWithPrec(5, 1) // 50%
)

// ParamTable for minting module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(minimumRiskFactor sdk.Dec, forceSuperfluidUndelegateAllowedAddresses []string) Params {
	return Params{
		MinimumRiskFactor:                         minimumRiskFactor,
		ForceSuperfluidUndelegateAllowedAddresses: forceSuperfluidUndelegateAllowedAddresses,
	}
}

// default minting module parameters.
func DefaultParams() Params {
	return Params{
		MinimumRiskFactor:                         defaultMinimumRiskFactor, // 5%
		ForceSuperfluidUndelegateAllowedAddresses: []string{},
	}
}

// validate params.
func (p Params) Validate() error {
	if err := ValidateMinimumRiskFactor(p.MinimumRiskFactor); err != nil {
		return err
	}

	if err := ValidateAddresses(p.ForceSuperfluidUndelegateAllowedAddresses); err != nil {
		return err
	}

	return nil
}

// Implements params.ParamSet.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMinimumRiskFactor, &p.MinimumRiskFactor, ValidateMinimumRiskFactor),
		paramtypes.NewParamSetPair(KeyForceSuperfluidUndelegateAllowedAddresses, &p.ForceSuperfluidUndelegateAllowedAddresses, ValidateAddresses),
	}
}

func ValidateMinimumRiskFactor(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() || v.GT(sdk.NewDec(100)) {
		return fmt.Errorf("minimum risk factor should be between 0 - 100: %s", v.String())
	}

	return nil
}

func ValidateUnbondingDuration(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("unbonding duration should be positive: %s", v.String())
	}

	return nil
}

func ValidateAddresses(i interface{}) error {
	addresses, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	for _, address := range addresses {
		_, err := sdk.AccAddressFromBech32(address)
		if err != nil {
			return err
		}
	}

	return nil
}
