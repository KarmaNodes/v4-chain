package types

// DefaultGenesis returns the default stats genesis state.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
		UnconditionalRevShareConfig: UnconditionalRevShareConfig{
			Configs: []UnconditionalRevShareConfig_RecipientConfig{},
		},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	return nil
}
