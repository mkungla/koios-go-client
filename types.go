// Copyright 2022 The Howijd.Network Authors
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//   http://www.apache.org/licenses/LICENSE-2.0
//   or LICENSE file in repository root.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package koios

type (
	// Address defines model for _address.
	Address string

	// AnyAddress defines model for _any_address.
	AnyAddress string

	// AssetName defines model for _asset_name.
	AssetName string

	// AssetPolicy defines model for _asset_policy.
	AssetPolicy string

	// BlockHash defines model for _block_hash.
	BlockHash string

	// EarnedEpochNo defines model for _earned_epoch_no.
	EarnedEpochNo string

	// EpochNo defines model for _epoch_no.
	EpochNo string

	// PoolBech32 defines model for _pool_bech32.
	PoolBech32 string

	// PoolBech32Optional defines model for _pool_bech32_optional.
	PoolBech32Optional string

	// ScriptHash defines model for _script_hash.
	ScriptHash string

	// StakeAddress defines model for _stake_address.
	StakeAddress string
)

type (
	// Tip defines model for tip.
	Tip []struct {
		// Absolute Slot number (slots not divided into epochs)
		AbsSlot int `json:"abs_slot"`

		// Block Height number on chain
		BlockNo int `json:"block_no"`

		// Timestamp for when the block was created
		BlockTime string `json:"block_time"`

		// Epoch number
		Epoch int `json:"epoch"`

		// Slot number within Epoch
		EpochSlot int `json:"epoch_slot"`

		// Block Hash in hex
		Hash string `json:"hash"`
	}

	// TipResponse response of /tip.
	TipResponse struct {
		Response
		Tip Tip `json:"response"`
	}
)

type (
	// Genesis defines model for genesis.
	Genesis []struct {
		// Active Slot Co-Efficient (f) - determines the _probability_ of number of
		// slots in epoch that are expected to have blocks
		// (so mainnet, this would be: 432000 * 0.05 = 21600 estimated blocks).
		Activeslotcoeff string `json:"activeslotcoeff"`

		// A JSON dump of Alonzo Genesis.
		Alonzogenesis string `json:"alonzogenesis"`

		// Number of slots in an epoch.
		Epochlength string `json:"epochlength"`

		// Number of KES key evolutions that will automatically occur before a KES
		// (hot) key is expired. This parameter is for security of a pool,
		// in case an operator had access to his hot(online) machine compromised.
		Maxkesrevolutions string `json:"maxkesrevolutions"`

		// Maximum smallest units (lovelaces) supply for the blockchain.
		Maxlovelacesupply string `json:"maxlovelacesupply"`

		// Network ID used at various CLI identification to distinguish between
		// Mainnet and other networks.
		Networkid string `json:"networkid"`

		// Unique network identifier for chain.
		Networkmagic string `json:"networkmagic"`

		// A unit (k) used to divide epochs to determine stability window
		// (used in security checks like ensuring atleast 1 block was
		// created in 3*k/f period, or to finalize next epoch's nonce
		// at 4*k/f slots before end of epoch).
		Securityparam string `json:"securityparam"`

		// Duration of a single slot (in seconds).
		Slotlength string `json:"slotlength"`

		// Number of slots that represent a single KES period
		// (a unit used for validation of KES key evolutions).
		Slotsperkesperiod string `json:"slotsperkesperiod"`

		// Timestamp for first block (genesis) on chain.
		Systemstart string `json:"systemstart"`

		// Number of BFT members that need to approve
		// (via vote) a Protocol Update Proposal.
		Updatequorum string `json:"updatequorum"`
	}

	// GenesisResponse response of /genesis.
	GenesisResponse struct {
		Response
		Genesis Genesis `json:"response"`
	}
)

type (
	// Totals defines model for totals.
	Totals []struct {

		// Circulating UTxOs for given epoch (in lovelaces).
		Circulation *string `json:"circulation,omitempty"`

		// Epoch number.
		EpochNo *int `json:"epoch_no,omitempty"`

		// Total Reserves yet to be unlocked on chain.
		Reserves *string `json:"reserves,omitempty"`

		// Rewards accumulated as of given epoch (in lovelaces).
		Reward *string `json:"reward,omitempty"`

		// Total Active Supply (sum of treasury funds, rewards,
		// UTxOs, deposits and fees) for given epoch (in lovelaces).
		Supply *string `json:"supply,omitempty"`

		// Funds in treasury for given epoch (in lovelaces).
		Treasury *string `json:"treasury,omitempty"`
	}

	// GenesisResponse response of /genesis.
	TotalsResponse struct {
		Response
		Totals Totals `json:"response"`
	}
)