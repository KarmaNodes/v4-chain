package types

import (
	"math/big"

	errorsmod "cosmossdk.io/errors"

	"github.com/dydxprotocol/v4-chain/protocol/dtypes"
)

// PositionSize is an interface for expressing the size of a position
type PositionSize interface {
	// Returns true if and only if the position size is positive.
	GetIsLong() bool
	// Returns the signed position size in big.Int.
	GetBigQuantums() *big.Int
	GetId() uint32
}

type PositionUpdate struct {
	Id          uint32
	BigQuantums *big.Int
}

func NewPositionUpdate(id uint32) PositionUpdate {
	return PositionUpdate{
		Id:          id,
		BigQuantums: big.NewInt(0),
	}
}

// Both updates and positions should conform to this interface
var _ PositionSize = AssetUpdate{}
var _ PositionSize = PerpetualUpdate{}
var _ PositionSize = PositionUpdate{}

// AssetPositions and PerpetualPositions use pointer receivers
// due to the way proto-gen generates them.
var _ PositionSize = &AssetPosition{}
var _ PositionSize = &PerpetualPosition{}

func (m *AssetPosition) GetId() uint32 {
	return m.GetAssetId()
}

// Get the asset position quantum size in big.Int. Panics if the size is zero.
func (m *AssetPosition) GetBigQuantums() *big.Int {
	if m == nil {
		return new(big.Int)
	}

	if m.Quantums.Sign() == 0 {
		panic(errorsmod.Wrapf(
			ErrAssetPositionZeroQuantum,
			"asset position (asset Id: %v) has zero quantum",
			m.AssetId,
		))
	}

	return m.Quantums.BigInt()
}

func (m *AssetPosition) GetIsLong() bool {
	if m == nil {
		return false
	}
	return m.GetBigQuantums().Sign() > 0
}

func (m *PerpetualPosition) GetId() uint32 {
	return m.GetPerpetualId()
}

func (m *PerpetualPosition) SetQuantums(sizeQuantums int64) {
	m.Quantums = dtypes.NewInt(sizeQuantums)
}

// Get the perpetual position quantum size in big.Int. Panics if the size is zero.
func (m *PerpetualPosition) GetBigQuantums() *big.Int {
	if m == nil {
		return new(big.Int)
	}

	if m.Quantums.Sign() == 0 {
		panic(errorsmod.Wrapf(
			ErrPerpPositionZeroQuantum,
			"perpetual position (perpetual Id: %v) has zero quantum",
			m.PerpetualId,
		))
	}

	return m.Quantums.BigInt()
}

func (m *PerpetualPosition) GetIsLong() bool {
	if m == nil {
		return false
	}
	return m.GetBigQuantums().Sign() > 0
}

func (au AssetUpdate) GetIsLong() bool {
	return au.GetBigQuantums().Sign() > 0
}

func (au AssetUpdate) GetBigQuantums() *big.Int {
	return au.BigQuantumsDelta
}

func (au AssetUpdate) GetId() uint32 {
	return au.AssetId
}

func (pu PerpetualUpdate) GetBigQuantums() *big.Int {
	return pu.BigQuantumsDelta
}

func (pu PerpetualUpdate) GetId() uint32 {
	return pu.PerpetualId
}

func (pu PerpetualUpdate) GetIsLong() bool {
	return pu.GetBigQuantums().Sign() > 0
}

func (pu PositionUpdate) GetId() uint32 {
	return pu.Id
}

func (pu PositionUpdate) GetIsLong() bool {
	return pu.BigQuantums.Sign() > 0
}

func (pu PositionUpdate) SetBigQuantums(bigQuantums *big.Int) {
	pu.BigQuantums.Set(bigQuantums)
}

func (pu PositionUpdate) GetBigQuantums() *big.Int {
	return pu.BigQuantums
}
