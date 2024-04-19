// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/accounts/abi"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IRollupBatchData is an auto generated low-level Go binding around an user-defined struct.
type IRollupBatchData struct {
	Version                uint8
	ParentBatchHeader      []byte
	Chunks                 [][]byte
	SkippedL1MessageBitmap []byte
	PrevStateRoot          [32]byte
	PostStateRoot          [32]byte
	WithdrawalRoot         [32]byte
	Signature              IRollupBatchSignature
}

// IRollupBatchSignature is an auto generated low-level Go binding around an user-defined struct.
type IRollupBatchSignature struct {
	Version   *big.Int
	Signers   []*big.Int
	Signature []byte
}

// RollupMetaData contains all meta data concerning the Rollup contract.
var RollupMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_chainId\",\"type\":\"uint64\"},{\"internalType\":\"addresspayable\",\"name\":\"_messenger\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ErrZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorIncorrectChunkLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorNoBlockInChunk\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"res\",\"type\":\"string\"}],\"name\":\"ChallengeRes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challengerReceiveAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengeDeposit\",\"type\":\"uint256\"}],\"name\":\"ChallengeState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"name\":\"CommitBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawRoot\",\"type\":\"bytes32\"}],\"name\":\"FinalizeBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"name\":\"RevertBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"UpdateChallenger\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPeriod\",\"type\":\"uint256\"}],\"name\":\"UpdateFinalizationPeriodSeconds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMaxNumTxInChunk\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxNumTxInChunk\",\"type\":\"uint256\"}],\"name\":\"UpdateMaxNumTxInChunk\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldWindow\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newWindow\",\"type\":\"uint256\"}],\"name\":\"UpdateProofWindow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"UpdateProver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldVerifier\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newVerifier\",\"type\":\"address\"}],\"name\":\"UpdateVerifier\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FINALIZATION_PERIOD_SECONDS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MESSENGER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROOF_WINDOW\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addChallenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addProver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"batchChallengeReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"batchChallengedSuccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"batchExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"batchInChallenge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"batchInsideChallengeWindow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"challengerReceiver\",\"type\":\"address\"}],\"name\":\"challengeState\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"challengerDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"challenges\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"challengerReceiveAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proverReceiveAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"challengeDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"challengeSuccess\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"finished\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"parentBatchHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"chunks\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"skippedL1MessageBitmap\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"prevStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"postStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"withdrawalRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"signers\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIRollup.BatchSignature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"internalType\":\"structIRollup.BatchData\",\"name\":\"batchData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"sequencers\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"commitBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"committedBatchStores\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"batchVersion\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"originTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalizeTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"prevStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"postStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"withdrawalRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l1DataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"l1MessagePopped\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalL1MessagePopped\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"skippedL1MessageBitmap\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blobVersionedHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"committedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchIndex\",\"type\":\"uint256\"}],\"name\":\"finalizeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"finalizedStateRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_batchHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_postStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_withdrawalRoot\",\"type\":\"bytes32\"}],\"name\":\"importGenesisBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inChallenge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1SequencerContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1StakingContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messageQueue\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_verifier\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxNumTxInChunk\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_finalizationPeriodSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_proofWindow\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchIndex\",\"type\":\"uint256\"}],\"name\":\"isBatchFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isChallenger\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isProver\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1SequencerContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1StakingContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastCommittedBatchIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastFinalizedBatchIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestL2BlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"layer2ChainId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxNumTxInChunk\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageQueue\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_batchIndex\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_proverReceiveAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_aggrProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_kzgDataProof\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"}],\"name\":\"proveState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeChallenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeProver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_batchHeader\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"revertBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revertReqIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPeriod\",\"type\":\"uint256\"}],\"name\":\"updateFinalizePeriodSeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxNumTxInChunk\",\"type\":\"uint256\"}],\"name\":\"updateMaxNumTxInChunk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newWindow\",\"type\":\"uint256\"}],\"name\":\"updateProofWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newVerifier\",\"type\":\"address\"}],\"name\":\"updateVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"withdrawalRoots\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c060405234801562000010575f80fd5b506040516200510538038062005105833981016040819052620000339162000053565b6001600160401b039091166080526001600160a01b031660a052620000a4565b5f806040838503121562000065575f80fd5b82516001600160401b03811681146200007c575f80fd5b60208401519092506001600160a01b038116811462000099575f80fd5b809150509250929050565b60805160a051615038620000cd5f395f61086101525f818161032d0152613a1601526150385ff3fe608060405260043610610311575f3560e01c806368ea42791161019c578063b31a77d3116100e7578063e33491a711610092578063eb1ec18f1161006d578063eb1ec18f146109fd578063ef6602ba14610a12578063f2fde38b14610a27578063f4daa29114610a46575f80fd5b8063e33491a7146109a0578063e3fff1dd146109bf578063e77fc7a4146109de575f80fd5b8063bedb86fb116100c2578063bedb86fb14610943578063ddd8a3dc14610962578063de8b303514610981575f80fd5b8063b31a77d3146108fb578063b571d3dd14610910578063b88a802f1461092f575f80fd5b80638f1d37761161014757806397fc007c1161012257806397fc007c14610883578063a415d8dc146108a2578063abc8d68d146108d0575f80fd5b80638f1d377614610749578063910129d41461081f578063927ede2d14610850575f80fd5b806388b1ea091161017757806388b1ea09146106f45780638c3968941461070d5780638da5cb5b1461072c575f80fd5b806368ea4279146106965780636c578c1d146106c1578063715018a6146106e0575f80fd5b80631e2283021161025c5780633e9e82ca11610207578063536deda5116101e2578063536deda51461064d57806357e0af6c146106605780635c975abb1461067f575f80fd5b80633e9e82ca146106065780634c4b9e4f1461061b5780634db4d9641461063a575f80fd5b80632571098d116102375780632571098d1461059d5780632b7ac3f3146105c85780633b70c18a146105e7575f80fd5b80631e2283021461053157806321e2f9e0146105505780632362f03e1461056f575f80fd5b80630ebb818d116102bc578063121dcd5011610297578063121dcd50146104cd57806318af3b2b146104e25780631d49e45714610512575f80fd5b80630ebb818d1461045557806310d445831461048c578063116a1f42146104ab575f80fd5b80630a245924116102ec5780630a245924146103ce5780630b79cdda146103fc5780630ceb678014610434575f80fd5b806303c7f4af1461031c57806304d772151461036d578063059def61146103ab575f80fd5b3661031857005b5f80fd5b348015610327575f80fd5b5061034f7f000000000000000000000000000000000000000000000000000000000000000081565b60405167ffffffffffffffff90911681526020015b60405180910390f35b348015610378575f80fd5b5061039b610387366004614585565b60a36020525f908152604090205460ff1681565b6040519015158152602001610364565b3480156103b6575f80fd5b506103c0609e5481565b604051908152602001610364565b3480156103d9575f80fd5b5061039b6103e83660046145b7565b609c6020525f908152604090205460ff1681565b348015610407575f80fd5b5061041b610416366004614585565b610a5b565b6040516103649d9c9b9a99989796959493929190614624565b34801561043f575f80fd5b5061045361044e3660046145b7565b610b43565b005b348015610460575f80fd5b50609754610474906001600160a01b031681565b6040516001600160a01b039091168152602001610364565b348015610497575f80fd5b506104536104a63660046146da565b610ba8565b3480156104b6575f80fd5b5061039b6104c5366004614585565b609e54101590565b3480156104d8575f80fd5b506103c0609f5481565b3480156104ed575f80fd5b5061039b6104fc366004614585565b5f90815260a16020526040902060030154421090565b34801561051d575f80fd5b5061045361052c3660046145b7565b610ec7565b34801561053c575f80fd5b5061045361054b366004614585565b610f7d565b34801561055b575f80fd5b5061039b61056a366004614585565b610fca565b34801561057a575f80fd5b506103c0610589366004614585565b5f90815260a1602052604090206001015490565b3480156105a8575f80fd5b506103c06105b7366004614585565b60a26020525f908152604090205481565b3480156105d3575f80fd5b50609b54610474906001600160a01b031681565b3480156105f2575f80fd5b50609a54610474906001600160a01b031681565b348015610611575f80fd5b506103c060a05481565b348015610626575f80fd5b50610453610635366004614722565b610ffe565b610453610648366004614839565b6115a3565b61045361065b36600461494c565b611b19565b34801561066b575f80fd5b5061045361067a366004614585565b6120fa565b34801561068a575f80fd5b5060655460ff1661039b565b3480156106a1575f80fd5b506103c06106b03660046145b7565b60a66020525f908152604090205481565b3480156106cc575f80fd5b506104536106db3660046145b7565b612143565b3480156106eb575f80fd5b50610453612199565b3480156106ff575f80fd5b5060a95461039b9060ff1681565b348015610718575f80fd5b5061045361072736600461497d565b6121ac565b348015610737575f80fd5b506033546001600160a01b0316610474565b348015610754575f80fd5b506107c8610763366004614585565b60a76020525f908152604090208054600182015460028301546003840154600485015460059095015467ffffffffffffffff851695680100000000000000009095046001600160a01b0390811695948116949316929060ff8082169161010090041688565b6040805167ffffffffffffffff90991689526001600160a01b0397881660208a015295871695880195909552949092166060860152608085015260a084015290151560c0830152151560e082015261010001610364565b34801561082a575f80fd5b5061039b610839366004614585565b5f90815260a7602052604090206005015460ff1690565b34801561085b575f80fd5b506104747f000000000000000000000000000000000000000000000000000000000000000081565b34801561088e575f80fd5b5061045361089d3660046145b7565b6124a9565b3480156108ad575f80fd5b5061039b6108bc3660046145b7565b609d6020525f908152604090205460ff1681565b3480156108db575f80fd5b506103c06108ea3660046145b7565b60a86020525f908152604090205481565b348015610906575f80fd5b506103c060aa5481565b34801561091b575f80fd5b5061045361092a3660046145b7565b61251a565b34801561093a575f80fd5b50610453612570565b34801561094e575f80fd5b5061045361095d366004614a30565b6125eb565b34801561096d575f80fd5b50609854610474906001600160a01b031681565b34801561098c575f80fd5b5061039b61099b366004614585565b612609565b3480156109ab575f80fd5b506104536109ba366004614585565b612653565b3480156109ca575f80fd5b506104536109d9366004614585565b612cb0565b3480156109e9575f80fd5b506104536109f8366004614a4b565b612cf9565b348015610a08575f80fd5b506103c060a55481565b348015610a1d575f80fd5b506103c060995481565b348015610a32575f80fd5b50610453610a413660046145b7565b61302e565b348015610a51575f80fd5b506103c060a45481565b60a1602052805f5260405f205f91509050805f01549080600101549080600201549080600301549080600401549080600501549080600601549080600701549080600901549080600a01549080600b018054610ab690614ab5565b80601f0160208091040260200160405190810160405280929190818152602001828054610ae290614ab5565b8015610b2d5780601f10610b0457610100808354040283529160200191610b2d565b820191905f5260205f20905b815481529060010190602001808311610b1057829003601f168201915b50505050509080600c01549080600d015490508d565b610b4b6130bb565b6001600160a01b0381165f818152609d6020908152604091829020805460ff1916600190811790915591519182527f7feb653c7b1f0d23daeed494225b3f28851cdc8973fcc653866d9b6e205fc00991015b60405180910390a250565b610bb06130bb565b5f8111610c045760405162461bcd60e51b815260206004820152601560248201527f636f756e74206d757374206265206e6f6e7a65726f000000000000000000000060448201526064015b60405180910390fd5b5f80610c108585613115565b915091505f610c23836001015160c01c90565b5f81815260a160205260409020600101549091508214610c855760405162461bcd60e51b815260206004820152601460248201527f696e636f727265637420626174636820686173680000000000000000000000006044820152606401610bfb565b5f60a181610c938785614b33565b81526020019081526020015f206001015414610d165760405162461bcd60e51b8152602060048201526024808201527f726576657274696e67206d7573742073746172742066726f6d2074686520656e60448201527f64696e67000000000000000000000000000000000000000000000000000000006064820152608401610bfb565b609e548111610d8d5760405162461bcd60e51b815260206004820152602160248201527f63616e206f6e6c792072657665727420756e46696e616c697a6564206261746360448201527f68000000000000000000000000000000000000000000000000000000000000006064820152608401610bfb565b610d98600182614b46565b609f555b8315610ebf575f81815260a1602052604081206001015560aa5415801590610dc5575060aa5481145b15610e70575f81815260a76020526040812080547fffffffff000000000000000000000000000000000000000000000000000000001681556001810180547fffffffffffffffffffffffff00000000000000000000000000000000000000009081169091556002820180549091169055600381018290556004810182905560050180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016905560aa555b604051829082907ecae2739091badfd91c373f0a16cede691e0cd25bb80cff77dd5caeb4710146905f90a360019081015f81815260a160205260409020909101545f1990940193915081610d9c575b505050505050565b610ecf6130bb565b6001600160a01b0381163b15610f275760405162461bcd60e51b815260206004820152600760248201527f6e6f7420454f41000000000000000000000000000000000000000000000000006044820152606401610bfb565b6001600160a01b0381165f818152609c6020908152604091829020805460ff1916600190811790915591519182527f967f99d5d403870e4356ff46556df3a6b6ba1f50146639aaedfb9f248eb8661e9101610b9d565b610f856130bb565b609980549082905560408051828152602081018490527f6d0f49971e462a2f78a25906f145cb29cd5e7bd01ebf681ac8f58cb814e5877a910160405180910390a15050565b5f81815260a1602052604081206002015415801590610ff857505f82815260a1602052604090206001015415155b92915050565b8161104b5760405162461bcd60e51b815260206004820152600f60248201527f7a65726f20737461746520726f6f7400000000000000000000000000000000006044820152606401610bfb565b5f805260a26020527f8fa6f0ecb9543d3381552e96aa43533c7f086066e38816919ea0cfae371b342a54156110c25760405162461bcd60e51b815260206004820152601660248201527f67656e6573697320626174636820696d706f72746564000000000000000000006044820152606401610bfb565b5f806110ce8686613115565b915091505f6110de836019015190565b90505f6110ec845160f81c90565b90505f6110fd856011015160c01c90565b600986015160c01c600187015160c01c875160f81c010101905080156111655760405162461bcd60e51b815260206004820152601760248201527f6e6f7420616c6c206669656c647320617265207a65726f0000000000000000006044820152606401610bfb565b505f611172856019015190565b036111bf5760405162461bcd60e51b815260206004820152600e60248201527f7a65726f206461746120686173680000000000000000000000000000000000006044820152606401610bfb565b5f6111cb856059015190565b146112185760405162461bcd60e51b815260206004820152601960248201527f6e6f6e7a65726f20706172656e742062617463682068617368000000000000006044820152606401610bfb565b5f611224856039015190565b90507f010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c44401481146112955760405162461bcd60e51b815260206004820152601660248201527f696e76616c69642076657273696f6e65642068617368000000000000000000006044820152606401610bfb565b604051806101c001604052808381526020018581526020014281526020014281526020015f801b81526020018881526020018781526020018481526020015f67ffffffffffffffff8111156112ec576112ec61476f565b604051908082528060200260200182016040528015611315578160200160208202803683370190505b5081525f6020808301829052604080840183905280518083018252838152606080860191909152608080860185905260a095860188905293805260a1835285517f32ae1b88a7d4f92d7e214b63c8ea04cd13e2faaa60c50f499f2254336d98f882908155868401517f32ae1b88a7d4f92d7e214b63c8ea04cd13e2faaa60c50f499f2254336d98f88355918601517f32ae1b88a7d4f92d7e214b63c8ea04cd13e2faaa60c50f499f2254336d98f884558501517f32ae1b88a7d4f92d7e214b63c8ea04cd13e2faaa60c50f499f2254336d98f88555918401517f32ae1b88a7d4f92d7e214b63c8ea04cd13e2faaa60c50f499f2254336d98f88655918301517f32ae1b88a7d4f92d7e214b63c8ea04cd13e2faaa60c50f499f2254336d98f8875560c08301517f32ae1b88a7d4f92d7e214b63c8ea04cd13e2faaa60c50f499f2254336d98f8885560e08301517f32ae1b88a7d4f92d7e214b63c8ea04cd13e2faaa60c50f499f2254336d98f88955610100830151805191926114bd927f32ae1b88a7d4f92d7e214b63c8ea04cd13e2faaa60c50f499f2254336d98f88a92909101906144a8565b506101208201516009820155610140820151600a820155610160820151600b8201906114e99082614ba4565b50610180820151600c8201556101a090910151600d909101555f80805260a26020527f8fa6f0ecb9543d3381552e96aa43533c7f086066e38816919ea0cfae371b342a8890556040518591907f2c32d4ae151744d0bf0b9464a3e897a1d17ed2f1af71f7c9a75f12ce0d28238f908290a3604080518881525f60208201819052869290917f26ba82f907317eedc97d0cbef23de76a43dd6edb563bdb6e9407645b950a7a2d910160405180910390a3505050505050505050565b60975483906001600160a01b03166326c9973d336040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b03909116600482015260248101849052604401602060405180830381865afa158015611618573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061163c9190614c60565b6116885760405162461bcd60e51b815260206004820152601460248201527f63616c6c6572206e6f742073657175656e6365720000000000000000000000006044820152606401610bfb565b611690613130565b61169d6020860186614c7b565b60ff16156116ed5760405162461bcd60e51b815260206004820152600f60248201527f696e76616c69642076657273696f6e00000000000000000000000000000000006044820152606401610bfb565b5f6116fb6040870187614c9b565b9150508061174b5760405162461bcd60e51b815260206004820152600e60248201527f626174636820697320656d7074790000000000000000000000000000000000006044820152606401610bfb565b608086013561179c5760405162461bcd60e51b815260206004820152601b60248201527f70726576696f757320737461746520726f6f74206973207a65726f00000000006044820152606401610bfb565b60a08601356117ed5760405162461bcd60e51b815260206004820152601660248201527f6e657720737461746520726f6f74206973207a65726f000000000000000000006044820152606401610bfb565b5f806118046117ff60208a018a614cff565b613115565b915091505f611817836001015160c01c90565b5f81815260a1602052604090206001015490915082146118795760405162461bcd60e51b815260206004820152601b60248201527f696e636f727265637420706172656e74206261746368206861736800000000006044820152606401610bfb565b5f60a181611888846001614b33565b81526020019081526020015f2060010154146118e65760405162461bcd60e51b815260206004820152601760248201527f626174636820616c726561647920636f6d6d69747465640000000000000000006044820152606401610bfb565b609f5481146119375760405162461bcd60e51b815260206004820152601560248201527f696e636f727265637420626174636820696e64657800000000000000000000006044820152606401610bfb565b5f81815260a1602052604090206005015460808a01351461199a5760405162461bcd60e51b815260206004820152601d60248201527f696e636f72726563742070726576696f757320737461746520726f6f740000006044820152606401610bfb565b6040805160018082528183019092525f916020808301908036833701905050905033815f815181106119ce576119ce614d60565b60200260200101906001600160a01b031690816001600160a01b0316815250506119fb848385848e613183565b6097545f83815260a16020526040908190206001015490517fe89ff5430000000000000000000000000000000000000000000000000000000081526001600160a01b039092169163e89ff54391611a5b918d918d918d9190600401614dd0565b6020604051808303815f875af1158015611a77573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611a9b9190614c60565b611b0d5760405162461bcd60e51b815260206004820152602160248201527f746865207369676e617475726520766572696669636174696f6e206661696c6560448201527f64000000000000000000000000000000000000000000000000000000000000006064820152608401610bfb565b50505050505050505050565b335f908152609d602052604090205460ff16611b775760405162461bcd60e51b815260206004820152601560248201527f63616c6c6572206e6f74206368616c6c656e67657200000000000000000000006044820152606401610bfb565b60aa5415611bc75760405162461bcd60e51b815260206004820152600b60248201527f6e656564207265766572740000000000000000000000000000000000000000006044820152606401610bfb565b60a95460ff1615611c1a5760405162461bcd60e51b815260206004820152601460248201527f616c726561647920696e206368616c6c656e67650000000000000000000000006044820152606401610bfb565b8167ffffffffffffffff16609e5410611c755760405162461bcd60e51b815260206004820152601760248201527f626174636820616c72656164792066696e616c697a65640000000000000000006044820152606401610bfb565b67ffffffffffffffff82165f90815260a160205260408120600101549003611cdf5760405162461bcd60e51b815260206004820152600f60248201527f6261746368206e6f7420657869737400000000000000000000000000000000006044820152606401610bfb565b67ffffffffffffffff82165f90815260a760205260409020546801000000000000000090046001600160a01b031615611d5a5760405162461bcd60e51b815260206004820152601260248201527f616c7265616479206368616c6c656e67656400000000000000000000000000006044820152606401610bfb565b67ffffffffffffffff82165f90815260a160205260409020600301544210611dea5760405162461bcd60e51b815260206004820152603360248201527f63616e6e6f74206368616c6c656e6765206261746368206f757473696465207460448201527f6865206368616c6c656e67652077696e646f77000000000000000000000000006064820152608401610bfb565b60985f9054906101000a90046001600160a01b03166001600160a01b031663a4d66daf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611e3a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e5e9190614e0c565b341015611ead5760405162461bcd60e51b815260206004820152601260248201527f696e73756666696369656e742076616c756500000000000000000000000000006044820152606401610bfb565b335f90815260a6602052604081208054349290611ecb908490614b33565b909155505060408051610100808201835267ffffffffffffffff8581168084523360208086018281526001600160a01b03898116888a019081525f60608a018181523460808c019081524260a08d0190815260c08d0184815260e08e018581528b865260a79099529d9093209b518c549651861668010000000000000000027fffffffff000000000000000000000000000000000000000000000000000000009097169a1699909917949094178a55905160018a0180549184167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179055925160028a0180549190931693169290921790559351600387015592516004860155945160059094018054925115159093027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff941515949094167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090921691909117929092179055907f961c03af6cc5d073982308a314a1bfb67a460342e8d5468e4d6385172cc617b190604080516001600160a01b0392831681529185166020830152349082015260600160405180910390a25f609e5460016120909190614b33565b90505b609f5481116120e8578267ffffffffffffffff1681146120d65760a5545f82815260a16020526040812060030180549091906120d0908490614b33565b90915550505b806120e081614e23565b915050612093565b505060a9805460ff1916600117905550565b6121026130bb565b60a55460408051918252602082018390527f1e3a2094feb4b696dd3d7caea38ad2f41dbdcac3fa3943c7a693aff8a64b0a61910160405180910390a160a555565b61214b6130bb565b6001600160a01b0381165f818152609d60209081526040808320805460ff19169055519182527f7feb653c7b1f0d23daeed494225b3f28851cdc8973fcc653866d9b6e205fc0099101610b9d565b6121a16130bb565b6121aa5f6135da565b565b60aa54156121fc5760405162461bcd60e51b815260206004820152600b60248201527f6e656564207265766572740000000000000000000000000000000000000000006044820152606401610bfb565b67ffffffffffffffff87165f90815260a760205260409020546801000000000000000090046001600160a01b03166122765760405162461bcd60e51b815260206004820152601860248201527f4368616c6c656e676520646f6573206e6f7420657869737400000000000000006044820152606401610bfb565b67ffffffffffffffff87165f90815260a76020526040902060050154610100900460ff16156122e75760405162461bcd60e51b815260206004820152601a60248201527f4368616c6c656e676520616c72656164792066696e69736865640000000000006044820152606401610bfb565b67ffffffffffffffff87165f90815260a7602052604090206005810180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790556002810180546001600160a01b0389167fffffffffffffffffffffffff000000000000000000000000000000000000000090911617905560a9805460ff1916905560a554600490910154429161238291614b33565b116124535767ffffffffffffffff87165f90815260a760209081526040808320600501805460ff1916600117905560a1825291829020600801805483518184028101840190945280845261244e938b939092919083018282801561240d57602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116123ef575b50505050506040518060400160405280600781526020017f54696d656f75740000000000000000000000000000000000000000000000000081525084613643565b6124a0565b61246087868686866137ba565b6124a087336040518060400160405280600d81526020017f50726f6f66207375636365737300000000000000000000000000000000000000815250613b24565b50505050505050565b6124b16130bb565b609b80546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f728af3d16a5760405e27a082c98ab272e9f0a1d02f0085d41532a26093aedd96905f90a35050565b6125226130bb565b6001600160a01b0381165f818152609c60209081526040808320805460ff19169055519182527f967f99d5d403870e4356ff46556df3a6b6ba1f50146639aaedfb9f248eb8661e9101610b9d565b335f90815260a86020526040812054908190036125cf5760405162461bcd60e51b815260206004820152601c60248201527f696e76616c69642062617463684368616c6c656e6765526577617264000000006044820152606401610bfb565b335f81815260a860205260408120556125e89082613c14565b50565b6125f36130bb565b8015612601576125e8613cbf565b6125e8613d19565b5f81815260a760205260408120546801000000000000000090046001600160a01b031615801590610ff85750505f90815260a76020526040902060050154610100900460ff161590565b60aa54156126a35760405162461bcd60e51b815260206004820152600b60248201527f6e656564207265766572740000000000000000000000000000000000000000006044820152606401610bfb565b6126ab613130565b6126b481610fca565b6127005760405162461bcd60e51b815260206004820152600f60248201527f6261746368206e6f7420657869737400000000000000000000000000000000006044820152606401610bfb565b61270981612609565b156127565760405162461bcd60e51b815260206004820152601260248201527f626174636820696e206368616c6c656e676500000000000000000000000000006044820152606401610bfb565b5f81815260a7602052604090206005015460ff16156127b75760405162461bcd60e51b815260206004820152601660248201527f62617463682073686f756c6420626520726576657274000000000000000000006044820152606401610bfb565b5f81815260a160205260409020600301544210156128175760405162461bcd60e51b815260206004820152601960248201527f626174636820696e206368616c6c656e67652077696e646f77000000000000006044820152606401610bfb565b5f81815260a160205260408120600401549060a290612837600185614b46565b81526020019081526020015f2054146128925760405162461bcd60e51b815260206004820152601d60248201527f696e636f72726563742070726576696f757320737461746520726f6f740000006044820152606401610bfb565b5f81815260a26020526040902054156128ed5760405162461bcd60e51b815260206004820152601660248201527f626174636820616c7265616479207665726966696564000000000000000000006044820152606401610bfb565b80609e54600101146129415760405162461bcd60e51b815260206004820152601560248201527f696e636f727265637420626174636820696e64657800000000000000000000006044820152606401610bfb565b609e8190555f81815260a1602081815260408084206006810154855260a38352818520805460ff19166001179055858552600581015460a28452919094205552600901548015612b1157609a545f83815260a160205260408120600b0180546001600160a01b03909316926129b590614ab5565b80601f01602080910402602001604051908101604052809291908181526020018280546129e190614ab5565b8015612a2c5780601f10612a0357610100808354040283529160200191612a2c565b820191905f5260205f20905b815481529060010190602001808311612a0f57829003601f168201915b5050505f87815260a1602090815260408220600a0154949550850193879003925090505b85811015612b0b57610100818703811115612a6a57508086035b6101008204602081028501516040517f55f613ce0000000000000000000000000000000000000000000000000000000081526004810186905260248101849052604481018290529091906001600160a01b038916906355f613ce906064015f604051808303815f87803b158015612adf575f80fd5b505af1158015612af1573d5f803e3d5ffd5b505050506101008501945050505061010081019050612a50565b50505050505b60a15f612b1f600185614b46565b815260208101919091526040015f9081208181556001810182905560028101829055600381018290556004810182905560058101829055600681018290556007810182905590612b726008830182614523565b600982015f9055600a82015f9055600b82015f612b8f919061453e565b505f600c8201819055600d90910181905560a790612bae600185614b46565b815260208082019290925260409081015f90812080547fffffffff00000000000000000000000000000000000000000000000000000000168155600180820180547fffffffffffffffffffffffff000000000000000000000000000000000000000090811690915560028301805490911690556003820183905560048201839055600591820180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016905586835260a1855291839020918201549082015460069092015483519283529382019390935284917f26ba82f907317eedc97d0cbef23de76a43dd6edb563bdb6e9407645b950a7a2d910160405180910390a35050565b612cb86130bb565b60a45460408051918252602082018390527fa577f4223f91f74e2dad65bbb8c30807587ae95d0d34288057bb3ec0d398a437910160405180910390a160a455565b5f54610100900460ff1615808015612d1757505f54600160ff909116105b80612d305750303b158015612d3057505f5460ff166001145b612da25760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610bfb565b5f805460ff191660011790558015612de0575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6001600160a01b0386161580612dfd57506001600160a01b038516155b15612e34576040517fecc6fdf000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612e3c613d52565b612e44613dd6565b6001600160a01b038816612e9a5760405162461bcd60e51b815260206004820152601d60248201527f696e76616c6964206c312073657175656e63657220636f6e74726163740000006044820152606401610bfb565b6001600160a01b038716612ef05760405162461bcd60e51b815260206004820152601b60248201527f696e76616c6964206c31207374616b696e6720636f6e747261637400000000006044820152606401610bfb565b609780546001600160a01b03808b167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617909255609880548a8416908316179055609a8054898416908316179055609b80549288169290911682179055609985905560a484905560a58390556040515f907f728af3d16a5760405e27a082c98ab272e9f0a1d02f0085d41532a26093aedd96908290a3604080515f8152602081018690527f6d0f49971e462a2f78a25906f145cb29cd5e7bd01ebf681ac8f58cb814e5877a910160405180910390a18015613024575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050505050565b6130366130bb565b6001600160a01b0381166130b25760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610bfb565b6125e8816135da565b6033546001600160a01b031633146121aa5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610bfb565b5f805f6131228585613e5a565b812090969095509350505050565b60655460ff16156121aa5760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610bfb565b5f6131916040830183614c9b565b905090505f6131a4876011015160c01c90565b604080516020850281019091529091505f805b848110156132d3575f61323d846131d160408a018a614c9b565b858181106131e1576131e1614d60565b90506020028101906131f39190614cff565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508892508a9150613238905060608d018d614cff565b613f31565b905061324a600187614b46565b82036132bb576132bb6132606040890189614c9b565b8481811061327057613270614d60565b90506020028101906132829190614cff565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506140f692505050565b938401936020939093019291909101906001016131b7565b506132e16060860186614cff565b905061010060ff8301046020021461333b5760405162461bcd60e51b815260206004820152601360248201527f77726f6e67206269746d6170206c656e677468000000000000000000000000006044820152606401610bfb565b6020848102808403206040519a50600199909901989061336b908b9061336390890189614c7b565b60ff16614142565b60c089811b60018c015282811b60098c015284901b60118b015260198a0181905260598a018890526133a98a6133a46060890189614cff565b614149565b5f49806133d357507f010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c4440145b60398b018190525f6133fd8c6133ec60608b018b614cff565b6133f891506079614b33565b902090565b604080516101c081019091529091508061341a60208b018b614c7b565b60ff16815260200182815260200142815260200160a4544261343c9190614b33565b8152602001896080013581526020018960a0013581526020018960c0013581526020018481526020018a81526020018581526020018781526020018980606001906134879190614cff565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200160a05481526020018381525060a15f8d81526020019081526020015f205f820151815f01556020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e082015181600701556101008201518160080190805190602001906135569291906144a8565b506101208201516009820155610140820151600a820155610160820151600b8201906135829082614ba4565b50610180820151600c8201556101a090910151600d90910155609f8b905560405181908c907f2c32d4ae151744d0bf0b9464a3e897a1d17ed2f1af71f7c9a75f12ce0d28238f905f90a3505050505050505050505050565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b67ffffffffffffffff841660aa8190555f90815260a760205260408082205460985491517fd339b988000000000000000000000000000000000000000000000000000000008152680100000000000000009091046001600160a01b039081169392169063d339b988906136bc9088908790600401614e3b565b6020604051808303815f875af11580156136d8573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906136fc9190614e0c565b67ffffffffffffffff87165f90815260a76020526040902060030154909150613726908290614b33565b67ffffffffffffffff87165f90815260a760209081526040808320600101546001600160a01b0316835260a890915281208054909190613767908490614b33565b925050819055508567ffffffffffffffff167f1e66d5dca70bf28588ef2f5cb3c299e65e2e7bdef2767823d3ae47a9caff95c683866040516137aa929190614e62565b60405180910390a2505050505050565b826138075760405162461bcd60e51b815260206004820152601960248201527f496e76616c6964206167677265676174696f6e2070726f6f66000000000000006044820152606401610bfb565b60a081146138575760405162461bcd60e51b815260206004820152601660248201527f496e76616c6964204b5a4720646174612070726f6f66000000000000000000006044820152606401610bfb565b67ffffffffffffffff85165f90815260a160209081526040808320600d015490518392600a9261388d9290918891889101614e8b565b60408051601f19818403018152908290526138a791614ea4565b5f60405180830381855afa9150503d805f81146138df576040519150601f19603f3d011682016040523d82523d5f602084013e6138e4565b606091505b50915091508161395c5760405162461bcd60e51b815260206004820152602a60248201527f6661696c656420746f2063616c6c20706f696e74206576616c756174696f6e2060448201527f707265636f6d70696c65000000000000000000000000000000000000000000006064820152608401610bfb565b5f818060200190518101906139719190614ebf565b9150507f73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff0000000181146139e35760405162461bcd60e51b815260206004820152601c60248201527f707265636f6d70696c6520756e6578706563746564206f7574707574000000006044820152606401610bfb565b50505067ffffffffffffffff85165f90815260a160205260408082206004810154600582015460068301546007909301547f00000000000000000000000000000000000000000000000000000000000000009492939192613a469087898b614ee1565b60a15f8e67ffffffffffffffff1681526020019081526020015f20600d0154604051602001613a7c989796959493929190614f08565b60408051601f198184030181528282528051602091820120609b5467ffffffffffffffff8b165f90815260a190935292909120547f2c09a8480000000000000000000000000000000000000000000000000000000084529093506001600160a01b0390911691632c09a84891613afc918a908a908a908890600401614f64565b5f6040518083038186803b158015613b12575f80fd5b505afa158015611b0d573d5f803e3d5ffd5b67ffffffffffffffff83165f90815260a7602090815260408083208054600390910154680100000000000000009091046001600160a01b031680855260a6909352908320805492939192839290613b7c908490614b46565b909155505067ffffffffffffffff85165f90815260a760209081526040808320600201546001600160a01b0316835260a890915281208054839290613bc2908490614b33565b925050819055508467ffffffffffffffff167f1e66d5dca70bf28588ef2f5cb3c299e65e2e7bdef2767823d3ae47a9caff95c68585604051613c05929190614e62565b60405180910390a25050505050565b8015613cbb575f826001600160a01b0316826040515f6040518083038185875af1925050503d805f8114613c63576040519150601f19603f3d011682016040523d82523d5f602084013e613c68565b606091505b5050905080613cb95760405162461bcd60e51b815260206004820152601b60248201527f526f6c6c75703a20455448207472616e73666572206661696c656400000000006044820152606401610bfb565b505b5050565b613cc7613130565b6065805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258613cfc3390565b6040516001600160a01b03909116815260200160405180910390a1565b613d21614155565b6065805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33613cfc565b5f54610100900460ff16613dce5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610bfb565b6121aa6141a7565b5f54610100900460ff16613e525760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610bfb565b6121aa61422f565b5f816079811015613ead5760405162461bcd60e51b815260206004820152601d60248201527f626174636820686561646572206c656e67746820746f6f20736d616c6c0000006044820152606401610bfb565b6040519150808483378082016040525f613ecb836009015160c01c90565b905061010060ff8201046020026079018214613f295760405162461bcd60e51b815260206004820152601360248201527f77726f6e67206269746d6170206c656e677468000000000000000000000000006044820152606401610bfb565b509250929050565b60405185515f916020880191819060218a01908590613f519086906142b4565b90505f805b82811015613fc057603c8102870160018101518652602101517fffffffffffffffffffffffffffffffffffffffffffffffffffff000000000000166020860152603a850194505f613fab856038015160f01c90565b603c9590950194929092019150600101613f56565b506020028301604052600185019150825b8115614079575f613fe684603a015160f01c90565b9050613ff685828e8e8e8e614349565b94505f614007856038015160f01c90565b9050818110156140595760405162461bcd60e51b815260206004820152601d60248201527f6e756d20747873206c657373207468616e206e756d204c31206d7367730000006044820152606401610bfb565b509a8b019a998a01999690960195603c92909201915f1990910190613fd1565b60995460206140888387614b46565b6140929190614fb3565b11156140e05760405162461bcd60e51b815260206004820152601960248201527f746f6f206d616e792074787320696e206f6e65206368756e6b000000000000006044820152606401610bfb565b50505081900390208852505b9695505050505050565b8051602182019060208301905f9061410f9083906142b4565b90505f5b61411e600183614b46565b81101561413457603c9390930192600101614113565b5050905160c01c60a0555050565b8082535050565b80826079850137505050565b60655460ff166121aa5760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610bfb565b5f54610100900460ff166142235760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610bfb565b6065805460ff19169055565b5f54610100900460ff166142ab5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610bfb565b6121aa336135da565b5f6142c0835160f81c90565b9050805f036142fb576040517f7c62f08c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b614306603c82614feb565b614311906001614b33565b8214610ff8576040517fb8f4bbc000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f855f036143585750856140ec565b609a546001600160a01b03165f80805b8981101561443e5760ff89169150600889901c811580614386575082155b15614395578060200288013593505b600184841c165f0361442d576040517fae453cd5000000000000000000000000000000000000000000000000000000008152600481018a90525f906001600160a01b0387169063ae453cd590602401602060405180830381865afa1580156143ff573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906144239190614e0c565b8d52506020909b019a5b506001988901989788019701614368565b505060ff5f1988011681811c6001161561449a5760405162461bcd60e51b815260206004820152601b60248201527f63616e6e6f7420736b6970206c617374204c31206d65737361676500000000006044820152606401610bfb565b509798975050505050505050565b828054828255905f5260205f20908101928215614513579160200282015b8281111561451357825182547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039091161782556020909201916001909101906144c6565b5061451f929150614571565b5090565b5080545f8255905f5260205f20908101906125e89190614571565b50805461454a90614ab5565b5f825580601f10614559575050565b601f0160209004905f5260205f20908101906125e891905b5b8082111561451f575f8155600101614572565b5f60208284031215614595575f80fd5b5035919050565b80356001600160a01b03811681146145b2575f80fd5b919050565b5f602082840312156145c7575f80fd5b6145d08261459c565b9392505050565b5f5b838110156145f15781810151838201526020016145d9565b50505f910152565b5f81518084526146108160208601602086016145d7565b601f01601f19169290920160200192915050565b5f6101a08f83528e60208401528d60408401528c60608401528b60808401528a60a08401528960c08401528860e0840152876101008401528661012084015280610140840152614676818401876145f9565b610160840195909552505061018001529b9a5050505050505050505050565b5f8083601f8401126146a5575f80fd5b50813567ffffffffffffffff8111156146bc575f80fd5b6020830191508360208285010111156146d3575f80fd5b9250929050565b5f805f604084860312156146ec575f80fd5b833567ffffffffffffffff811115614702575f80fd5b61470e86828701614695565b909790965060209590950135949350505050565b5f805f8060608587031215614735575f80fd5b843567ffffffffffffffff81111561474b575f80fd5b61475787828801614695565b90989097506020870135966040013595509350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff811182821017156147c5576147c561476f565b604052919050565b5f82601f8301126147dc575f80fd5b813567ffffffffffffffff8111156147f6576147f661476f565b6148096020601f19601f8401160161479c565b81815284602083860101111561481d575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f806080858703121561484c575f80fd5b843567ffffffffffffffff80821115614863575f80fd5b908601906101008289031215614877575f80fd5b909450602086810135945090604087013581811115614894575f80fd5b8701601f810189136148a4575f80fd5b8035828111156148b6576148b661476f565b8060051b6148c585820161479c565b918252828101850191858101908c8411156148de575f80fd5b938601935b83851015614903576148f48561459c565b825293860193908601906148e3565b975050505060608801359250508082111561491c575f80fd5b50614929878288016147cd565b91505092959194509250565b803567ffffffffffffffff811681146145b2575f80fd5b5f806040838503121561495d575f80fd5b61496683614935565b91506149746020840161459c565b90509250929050565b5f805f805f805f60a0888a031215614993575f80fd5b61499c88614935565b96506149aa6020890161459c565b9550604088013567ffffffffffffffff808211156149c6575f80fd5b6149d28b838c01614695565b909750955060608a01359150808211156149ea575f80fd5b506149f78a828b01614695565b909450925050608088013563ffffffff81168114614a13575f80fd5b8091505092959891949750929550565b80151581146125e8575f80fd5b5f60208284031215614a40575f80fd5b81356145d081614a23565b5f805f805f805f60e0888a031215614a61575f80fd5b614a6a8861459c565b9650614a786020890161459c565b9550614a866040890161459c565b9450614a946060890161459c565b9699959850939660808101359560a0820135955060c0909101359350915050565b600181811c90821680614ac957607f821691505b602082108103614b00577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b80820180821115610ff857610ff8614b06565b81810381811115610ff857610ff8614b06565b601f821115613cb957805f5260205f20601f840160051c81016020851015614b7e5750805b601f840160051c820191505b81811015614b9d575f8155600101614b8a565b5050505050565b815167ffffffffffffffff811115614bbe57614bbe61476f565b614bd281614bcc8454614ab5565b84614b59565b602080601f831160018114614c05575f8415614bee5750858301515b5f19600386901b1c1916600185901b178555610ebf565b5f85815260208120601f198616915b82811015614c3357888601518255948401946001909101908401614c14565b5085821015614c5057878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b5f60208284031215614c70575f80fd5b81516145d081614a23565b5f60208284031215614c8b575f80fd5b813560ff811681146145d0575f80fd5b5f8083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112614cce575f80fd5b83018035915067ffffffffffffffff821115614ce8575f80fd5b6020019150600581901b36038213156146d3575f80fd5b5f8083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112614d32575f80fd5b83018035915067ffffffffffffffff821115614d4c575f80fd5b6020019150368190038213156146d3575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f815180845260208085019450602084015f5b83811015614dc55781516001600160a01b031687529582019590820190600101614da0565b509495945050505050565b848152608060208201525f614de86080830186614d8d565b8281036040840152614dfa81866145f9565b91505082606083015295945050505050565b5f60208284031215614e1c575f80fd5b5051919050565b5f5f198203614e3457614e34614b06565b5060010190565b604081525f614e4d6040830185614d8d565b905063ffffffff831660208301529392505050565b6001600160a01b0383168152604060208201525f614e8360408301846145f9565b949350505050565b838152818360208301375f910160200190815292915050565b5f8251614eb58184602087016145d7565b9190910192915050565b5f8060408385031215614ed0575f80fd5b505080516020909101519092909150565b5f8085851115614eef575f80fd5b83861115614efb575f80fd5b5050820193919092039150565b7fffffffffffffffff0000000000000000000000000000000000000000000000008960c01b16815287600882015286602882015285604882015284606882015282846088830137608892019182015260a8019695505050505050565b85815267ffffffffffffffff8516602082015260806040820152826080820152828460a08301375f60a084830101525f60a0601f19601f86011683010190508260608301529695505050505050565b5f82614fe6577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b8082028115828204841417610ff857610ff8614b0656fea26469706673582212207ca1c8a25053e6f2fb22ae119806939d3c887fd4f755608de6940661d495519a64736f6c63430008180033",
}

// RollupABI is the input ABI used to generate the binding from.
// Deprecated: Use RollupMetaData.ABI instead.
var RollupABI = RollupMetaData.ABI

// RollupBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RollupMetaData.Bin instead.
var RollupBin = RollupMetaData.Bin

// DeployRollup deploys a new Ethereum contract, binding an instance of Rollup to it.
func DeployRollup(auth *bind.TransactOpts, backend bind.ContractBackend, _chainId uint64, _messenger common.Address) (common.Address, *types.Transaction, *Rollup, error) {
	parsed, err := RollupMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RollupBin), backend, _chainId, _messenger)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Rollup{RollupCaller: RollupCaller{contract: contract}, RollupTransactor: RollupTransactor{contract: contract}, RollupFilterer: RollupFilterer{contract: contract}}, nil
}

// Rollup is an auto generated Go binding around an Ethereum contract.
type Rollup struct {
	RollupCaller     // Read-only binding to the contract
	RollupTransactor // Write-only binding to the contract
	RollupFilterer   // Log filterer for contract events
}

// RollupCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupSession struct {
	Contract     *Rollup           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupCallerSession struct {
	Contract *RollupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RollupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupTransactorSession struct {
	Contract     *RollupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupRaw struct {
	Contract *Rollup // Generic contract binding to access the raw methods on
}

// RollupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupCallerRaw struct {
	Contract *RollupCaller // Generic read-only contract binding to access the raw methods on
}

// RollupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupTransactorRaw struct {
	Contract *RollupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollup creates a new instance of Rollup, bound to a specific deployed contract.
func NewRollup(address common.Address, backend bind.ContractBackend) (*Rollup, error) {
	contract, err := bindRollup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rollup{RollupCaller: RollupCaller{contract: contract}, RollupTransactor: RollupTransactor{contract: contract}, RollupFilterer: RollupFilterer{contract: contract}}, nil
}

// NewRollupCaller creates a new read-only instance of Rollup, bound to a specific deployed contract.
func NewRollupCaller(address common.Address, caller bind.ContractCaller) (*RollupCaller, error) {
	contract, err := bindRollup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCaller{contract: contract}, nil
}

// NewRollupTransactor creates a new write-only instance of Rollup, bound to a specific deployed contract.
func NewRollupTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupTransactor, error) {
	contract, err := bindRollup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupTransactor{contract: contract}, nil
}

// NewRollupFilterer creates a new log filterer instance of Rollup, bound to a specific deployed contract.
func NewRollupFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupFilterer, error) {
	contract, err := bindRollup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupFilterer{contract: contract}, nil
}

// bindRollup binds a generic wrapper to an already deployed contract.
func bindRollup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RollupMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rollup *RollupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rollup.Contract.RollupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rollup *RollupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.Contract.RollupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rollup *RollupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rollup.Contract.RollupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rollup *RollupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rollup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rollup *RollupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rollup *RollupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rollup.Contract.contract.Transact(opts, method, params...)
}

// FINALIZATIONPERIODSECONDS is a free data retrieval call binding the contract method 0xf4daa291.
//
// Solidity: function FINALIZATION_PERIOD_SECONDS() view returns(uint256)
func (_Rollup *RollupCaller) FINALIZATIONPERIODSECONDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "FINALIZATION_PERIOD_SECONDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FINALIZATIONPERIODSECONDS is a free data retrieval call binding the contract method 0xf4daa291.
//
// Solidity: function FINALIZATION_PERIOD_SECONDS() view returns(uint256)
func (_Rollup *RollupSession) FINALIZATIONPERIODSECONDS() (*big.Int, error) {
	return _Rollup.Contract.FINALIZATIONPERIODSECONDS(&_Rollup.CallOpts)
}

// FINALIZATIONPERIODSECONDS is a free data retrieval call binding the contract method 0xf4daa291.
//
// Solidity: function FINALIZATION_PERIOD_SECONDS() view returns(uint256)
func (_Rollup *RollupCallerSession) FINALIZATIONPERIODSECONDS() (*big.Int, error) {
	return _Rollup.Contract.FINALIZATIONPERIODSECONDS(&_Rollup.CallOpts)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_Rollup *RollupCaller) MESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_Rollup *RollupSession) MESSENGER() (common.Address, error) {
	return _Rollup.Contract.MESSENGER(&_Rollup.CallOpts)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_Rollup *RollupCallerSession) MESSENGER() (common.Address, error) {
	return _Rollup.Contract.MESSENGER(&_Rollup.CallOpts)
}

// PROOFWINDOW is a free data retrieval call binding the contract method 0xeb1ec18f.
//
// Solidity: function PROOF_WINDOW() view returns(uint256)
func (_Rollup *RollupCaller) PROOFWINDOW(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "PROOF_WINDOW")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PROOFWINDOW is a free data retrieval call binding the contract method 0xeb1ec18f.
//
// Solidity: function PROOF_WINDOW() view returns(uint256)
func (_Rollup *RollupSession) PROOFWINDOW() (*big.Int, error) {
	return _Rollup.Contract.PROOFWINDOW(&_Rollup.CallOpts)
}

// PROOFWINDOW is a free data retrieval call binding the contract method 0xeb1ec18f.
//
// Solidity: function PROOF_WINDOW() view returns(uint256)
func (_Rollup *RollupCallerSession) PROOFWINDOW() (*big.Int, error) {
	return _Rollup.Contract.PROOFWINDOW(&_Rollup.CallOpts)
}

// BatchChallengeReward is a free data retrieval call binding the contract method 0xabc8d68d.
//
// Solidity: function batchChallengeReward(address ) view returns(uint256)
func (_Rollup *RollupCaller) BatchChallengeReward(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "batchChallengeReward", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchChallengeReward is a free data retrieval call binding the contract method 0xabc8d68d.
//
// Solidity: function batchChallengeReward(address ) view returns(uint256)
func (_Rollup *RollupSession) BatchChallengeReward(arg0 common.Address) (*big.Int, error) {
	return _Rollup.Contract.BatchChallengeReward(&_Rollup.CallOpts, arg0)
}

// BatchChallengeReward is a free data retrieval call binding the contract method 0xabc8d68d.
//
// Solidity: function batchChallengeReward(address ) view returns(uint256)
func (_Rollup *RollupCallerSession) BatchChallengeReward(arg0 common.Address) (*big.Int, error) {
	return _Rollup.Contract.BatchChallengeReward(&_Rollup.CallOpts, arg0)
}

// BatchChallengedSuccess is a free data retrieval call binding the contract method 0x910129d4.
//
// Solidity: function batchChallengedSuccess(uint256 batchIndex) view returns(bool)
func (_Rollup *RollupCaller) BatchChallengedSuccess(opts *bind.CallOpts, batchIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "batchChallengedSuccess", batchIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BatchChallengedSuccess is a free data retrieval call binding the contract method 0x910129d4.
//
// Solidity: function batchChallengedSuccess(uint256 batchIndex) view returns(bool)
func (_Rollup *RollupSession) BatchChallengedSuccess(batchIndex *big.Int) (bool, error) {
	return _Rollup.Contract.BatchChallengedSuccess(&_Rollup.CallOpts, batchIndex)
}

// BatchChallengedSuccess is a free data retrieval call binding the contract method 0x910129d4.
//
// Solidity: function batchChallengedSuccess(uint256 batchIndex) view returns(bool)
func (_Rollup *RollupCallerSession) BatchChallengedSuccess(batchIndex *big.Int) (bool, error) {
	return _Rollup.Contract.BatchChallengedSuccess(&_Rollup.CallOpts, batchIndex)
}

// BatchExist is a free data retrieval call binding the contract method 0x21e2f9e0.
//
// Solidity: function batchExist(uint256 batchIndex) view returns(bool)
func (_Rollup *RollupCaller) BatchExist(opts *bind.CallOpts, batchIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "batchExist", batchIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BatchExist is a free data retrieval call binding the contract method 0x21e2f9e0.
//
// Solidity: function batchExist(uint256 batchIndex) view returns(bool)
func (_Rollup *RollupSession) BatchExist(batchIndex *big.Int) (bool, error) {
	return _Rollup.Contract.BatchExist(&_Rollup.CallOpts, batchIndex)
}

// BatchExist is a free data retrieval call binding the contract method 0x21e2f9e0.
//
// Solidity: function batchExist(uint256 batchIndex) view returns(bool)
func (_Rollup *RollupCallerSession) BatchExist(batchIndex *big.Int) (bool, error) {
	return _Rollup.Contract.BatchExist(&_Rollup.CallOpts, batchIndex)
}

// BatchInChallenge is a free data retrieval call binding the contract method 0xde8b3035.
//
// Solidity: function batchInChallenge(uint256 batchIndex) view returns(bool)
func (_Rollup *RollupCaller) BatchInChallenge(opts *bind.CallOpts, batchIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "batchInChallenge", batchIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BatchInChallenge is a free data retrieval call binding the contract method 0xde8b3035.
//
// Solidity: function batchInChallenge(uint256 batchIndex) view returns(bool)
func (_Rollup *RollupSession) BatchInChallenge(batchIndex *big.Int) (bool, error) {
	return _Rollup.Contract.BatchInChallenge(&_Rollup.CallOpts, batchIndex)
}

// BatchInChallenge is a free data retrieval call binding the contract method 0xde8b3035.
//
// Solidity: function batchInChallenge(uint256 batchIndex) view returns(bool)
func (_Rollup *RollupCallerSession) BatchInChallenge(batchIndex *big.Int) (bool, error) {
	return _Rollup.Contract.BatchInChallenge(&_Rollup.CallOpts, batchIndex)
}

// BatchInsideChallengeWindow is a free data retrieval call binding the contract method 0x18af3b2b.
//
// Solidity: function batchInsideChallengeWindow(uint256 batchIndex) view returns(bool)
func (_Rollup *RollupCaller) BatchInsideChallengeWindow(opts *bind.CallOpts, batchIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "batchInsideChallengeWindow", batchIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BatchInsideChallengeWindow is a free data retrieval call binding the contract method 0x18af3b2b.
//
// Solidity: function batchInsideChallengeWindow(uint256 batchIndex) view returns(bool)
func (_Rollup *RollupSession) BatchInsideChallengeWindow(batchIndex *big.Int) (bool, error) {
	return _Rollup.Contract.BatchInsideChallengeWindow(&_Rollup.CallOpts, batchIndex)
}

// BatchInsideChallengeWindow is a free data retrieval call binding the contract method 0x18af3b2b.
//
// Solidity: function batchInsideChallengeWindow(uint256 batchIndex) view returns(bool)
func (_Rollup *RollupCallerSession) BatchInsideChallengeWindow(batchIndex *big.Int) (bool, error) {
	return _Rollup.Contract.BatchInsideChallengeWindow(&_Rollup.CallOpts, batchIndex)
}

// ChallengerDeposits is a free data retrieval call binding the contract method 0x68ea4279.
//
// Solidity: function challengerDeposits(address ) view returns(uint256)
func (_Rollup *RollupCaller) ChallengerDeposits(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "challengerDeposits", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengerDeposits is a free data retrieval call binding the contract method 0x68ea4279.
//
// Solidity: function challengerDeposits(address ) view returns(uint256)
func (_Rollup *RollupSession) ChallengerDeposits(arg0 common.Address) (*big.Int, error) {
	return _Rollup.Contract.ChallengerDeposits(&_Rollup.CallOpts, arg0)
}

// ChallengerDeposits is a free data retrieval call binding the contract method 0x68ea4279.
//
// Solidity: function challengerDeposits(address ) view returns(uint256)
func (_Rollup *RollupCallerSession) ChallengerDeposits(arg0 common.Address) (*big.Int, error) {
	return _Rollup.Contract.ChallengerDeposits(&_Rollup.CallOpts, arg0)
}

// Challenges is a free data retrieval call binding the contract method 0x8f1d3776.
//
// Solidity: function challenges(uint256 ) view returns(uint64 batchIndex, address challenger, address challengerReceiveAddress, address proverReceiveAddress, uint256 challengeDeposit, uint256 startTime, bool challengeSuccess, bool finished)
func (_Rollup *RollupCaller) Challenges(opts *bind.CallOpts, arg0 *big.Int) (struct {
	BatchIndex               uint64
	Challenger               common.Address
	ChallengerReceiveAddress common.Address
	ProverReceiveAddress     common.Address
	ChallengeDeposit         *big.Int
	StartTime                *big.Int
	ChallengeSuccess         bool
	Finished                 bool
}, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "challenges", arg0)

	outstruct := new(struct {
		BatchIndex               uint64
		Challenger               common.Address
		ChallengerReceiveAddress common.Address
		ProverReceiveAddress     common.Address
		ChallengeDeposit         *big.Int
		StartTime                *big.Int
		ChallengeSuccess         bool
		Finished                 bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BatchIndex = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Challenger = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ChallengerReceiveAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.ProverReceiveAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.ChallengeDeposit = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.ChallengeSuccess = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.Finished = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// Challenges is a free data retrieval call binding the contract method 0x8f1d3776.
//
// Solidity: function challenges(uint256 ) view returns(uint64 batchIndex, address challenger, address challengerReceiveAddress, address proverReceiveAddress, uint256 challengeDeposit, uint256 startTime, bool challengeSuccess, bool finished)
func (_Rollup *RollupSession) Challenges(arg0 *big.Int) (struct {
	BatchIndex               uint64
	Challenger               common.Address
	ChallengerReceiveAddress common.Address
	ProverReceiveAddress     common.Address
	ChallengeDeposit         *big.Int
	StartTime                *big.Int
	ChallengeSuccess         bool
	Finished                 bool
}, error) {
	return _Rollup.Contract.Challenges(&_Rollup.CallOpts, arg0)
}

// Challenges is a free data retrieval call binding the contract method 0x8f1d3776.
//
// Solidity: function challenges(uint256 ) view returns(uint64 batchIndex, address challenger, address challengerReceiveAddress, address proverReceiveAddress, uint256 challengeDeposit, uint256 startTime, bool challengeSuccess, bool finished)
func (_Rollup *RollupCallerSession) Challenges(arg0 *big.Int) (struct {
	BatchIndex               uint64
	Challenger               common.Address
	ChallengerReceiveAddress common.Address
	ProverReceiveAddress     common.Address
	ChallengeDeposit         *big.Int
	StartTime                *big.Int
	ChallengeSuccess         bool
	Finished                 bool
}, error) {
	return _Rollup.Contract.Challenges(&_Rollup.CallOpts, arg0)
}

// CommittedBatchStores is a free data retrieval call binding the contract method 0x0b79cdda.
//
// Solidity: function committedBatchStores(uint256 ) view returns(uint256 batchVersion, bytes32 batchHash, uint256 originTimestamp, uint256 finalizeTimestamp, bytes32 prevStateRoot, bytes32 postStateRoot, bytes32 withdrawalRoot, bytes32 l1DataHash, uint256 l1MessagePopped, uint256 totalL1MessagePopped, bytes skippedL1MessageBitmap, uint256 blockNumber, bytes32 blobVersionedHash)
func (_Rollup *RollupCaller) CommittedBatchStores(opts *bind.CallOpts, arg0 *big.Int) (struct {
	BatchVersion           *big.Int
	BatchHash              [32]byte
	OriginTimestamp        *big.Int
	FinalizeTimestamp      *big.Int
	PrevStateRoot          [32]byte
	PostStateRoot          [32]byte
	WithdrawalRoot         [32]byte
	L1DataHash             [32]byte
	L1MessagePopped        *big.Int
	TotalL1MessagePopped   *big.Int
	SkippedL1MessageBitmap []byte
	BlockNumber            *big.Int
	BlobVersionedHash      [32]byte
}, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "committedBatchStores", arg0)

	outstruct := new(struct {
		BatchVersion           *big.Int
		BatchHash              [32]byte
		OriginTimestamp        *big.Int
		FinalizeTimestamp      *big.Int
		PrevStateRoot          [32]byte
		PostStateRoot          [32]byte
		WithdrawalRoot         [32]byte
		L1DataHash             [32]byte
		L1MessagePopped        *big.Int
		TotalL1MessagePopped   *big.Int
		SkippedL1MessageBitmap []byte
		BlockNumber            *big.Int
		BlobVersionedHash      [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BatchVersion = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BatchHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.OriginTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FinalizeTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.PrevStateRoot = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.PostStateRoot = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.WithdrawalRoot = *abi.ConvertType(out[6], new([32]byte)).(*[32]byte)
	outstruct.L1DataHash = *abi.ConvertType(out[7], new([32]byte)).(*[32]byte)
	outstruct.L1MessagePopped = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.TotalL1MessagePopped = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.SkippedL1MessageBitmap = *abi.ConvertType(out[10], new([]byte)).(*[]byte)
	outstruct.BlockNumber = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	outstruct.BlobVersionedHash = *abi.ConvertType(out[12], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// CommittedBatchStores is a free data retrieval call binding the contract method 0x0b79cdda.
//
// Solidity: function committedBatchStores(uint256 ) view returns(uint256 batchVersion, bytes32 batchHash, uint256 originTimestamp, uint256 finalizeTimestamp, bytes32 prevStateRoot, bytes32 postStateRoot, bytes32 withdrawalRoot, bytes32 l1DataHash, uint256 l1MessagePopped, uint256 totalL1MessagePopped, bytes skippedL1MessageBitmap, uint256 blockNumber, bytes32 blobVersionedHash)
func (_Rollup *RollupSession) CommittedBatchStores(arg0 *big.Int) (struct {
	BatchVersion           *big.Int
	BatchHash              [32]byte
	OriginTimestamp        *big.Int
	FinalizeTimestamp      *big.Int
	PrevStateRoot          [32]byte
	PostStateRoot          [32]byte
	WithdrawalRoot         [32]byte
	L1DataHash             [32]byte
	L1MessagePopped        *big.Int
	TotalL1MessagePopped   *big.Int
	SkippedL1MessageBitmap []byte
	BlockNumber            *big.Int
	BlobVersionedHash      [32]byte
}, error) {
	return _Rollup.Contract.CommittedBatchStores(&_Rollup.CallOpts, arg0)
}

// CommittedBatchStores is a free data retrieval call binding the contract method 0x0b79cdda.
//
// Solidity: function committedBatchStores(uint256 ) view returns(uint256 batchVersion, bytes32 batchHash, uint256 originTimestamp, uint256 finalizeTimestamp, bytes32 prevStateRoot, bytes32 postStateRoot, bytes32 withdrawalRoot, bytes32 l1DataHash, uint256 l1MessagePopped, uint256 totalL1MessagePopped, bytes skippedL1MessageBitmap, uint256 blockNumber, bytes32 blobVersionedHash)
func (_Rollup *RollupCallerSession) CommittedBatchStores(arg0 *big.Int) (struct {
	BatchVersion           *big.Int
	BatchHash              [32]byte
	OriginTimestamp        *big.Int
	FinalizeTimestamp      *big.Int
	PrevStateRoot          [32]byte
	PostStateRoot          [32]byte
	WithdrawalRoot         [32]byte
	L1DataHash             [32]byte
	L1MessagePopped        *big.Int
	TotalL1MessagePopped   *big.Int
	SkippedL1MessageBitmap []byte
	BlockNumber            *big.Int
	BlobVersionedHash      [32]byte
}, error) {
	return _Rollup.Contract.CommittedBatchStores(&_Rollup.CallOpts, arg0)
}

// CommittedBatches is a free data retrieval call binding the contract method 0x2362f03e.
//
// Solidity: function committedBatches(uint256 batchIndex) view returns(bytes32)
func (_Rollup *RollupCaller) CommittedBatches(opts *bind.CallOpts, batchIndex *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "committedBatches", batchIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CommittedBatches is a free data retrieval call binding the contract method 0x2362f03e.
//
// Solidity: function committedBatches(uint256 batchIndex) view returns(bytes32)
func (_Rollup *RollupSession) CommittedBatches(batchIndex *big.Int) ([32]byte, error) {
	return _Rollup.Contract.CommittedBatches(&_Rollup.CallOpts, batchIndex)
}

// CommittedBatches is a free data retrieval call binding the contract method 0x2362f03e.
//
// Solidity: function committedBatches(uint256 batchIndex) view returns(bytes32)
func (_Rollup *RollupCallerSession) CommittedBatches(batchIndex *big.Int) ([32]byte, error) {
	return _Rollup.Contract.CommittedBatches(&_Rollup.CallOpts, batchIndex)
}

// FinalizedStateRoots is a free data retrieval call binding the contract method 0x2571098d.
//
// Solidity: function finalizedStateRoots(uint256 ) view returns(bytes32)
func (_Rollup *RollupCaller) FinalizedStateRoots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "finalizedStateRoots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FinalizedStateRoots is a free data retrieval call binding the contract method 0x2571098d.
//
// Solidity: function finalizedStateRoots(uint256 ) view returns(bytes32)
func (_Rollup *RollupSession) FinalizedStateRoots(arg0 *big.Int) ([32]byte, error) {
	return _Rollup.Contract.FinalizedStateRoots(&_Rollup.CallOpts, arg0)
}

// FinalizedStateRoots is a free data retrieval call binding the contract method 0x2571098d.
//
// Solidity: function finalizedStateRoots(uint256 ) view returns(bytes32)
func (_Rollup *RollupCallerSession) FinalizedStateRoots(arg0 *big.Int) ([32]byte, error) {
	return _Rollup.Contract.FinalizedStateRoots(&_Rollup.CallOpts, arg0)
}

// InChallenge is a free data retrieval call binding the contract method 0x88b1ea09.
//
// Solidity: function inChallenge() view returns(bool)
func (_Rollup *RollupCaller) InChallenge(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "inChallenge")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InChallenge is a free data retrieval call binding the contract method 0x88b1ea09.
//
// Solidity: function inChallenge() view returns(bool)
func (_Rollup *RollupSession) InChallenge() (bool, error) {
	return _Rollup.Contract.InChallenge(&_Rollup.CallOpts)
}

// InChallenge is a free data retrieval call binding the contract method 0x88b1ea09.
//
// Solidity: function inChallenge() view returns(bool)
func (_Rollup *RollupCallerSession) InChallenge() (bool, error) {
	return _Rollup.Contract.InChallenge(&_Rollup.CallOpts)
}

// IsBatchFinalized is a free data retrieval call binding the contract method 0x116a1f42.
//
// Solidity: function isBatchFinalized(uint256 _batchIndex) view returns(bool)
func (_Rollup *RollupCaller) IsBatchFinalized(opts *bind.CallOpts, _batchIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "isBatchFinalized", _batchIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBatchFinalized is a free data retrieval call binding the contract method 0x116a1f42.
//
// Solidity: function isBatchFinalized(uint256 _batchIndex) view returns(bool)
func (_Rollup *RollupSession) IsBatchFinalized(_batchIndex *big.Int) (bool, error) {
	return _Rollup.Contract.IsBatchFinalized(&_Rollup.CallOpts, _batchIndex)
}

// IsBatchFinalized is a free data retrieval call binding the contract method 0x116a1f42.
//
// Solidity: function isBatchFinalized(uint256 _batchIndex) view returns(bool)
func (_Rollup *RollupCallerSession) IsBatchFinalized(_batchIndex *big.Int) (bool, error) {
	return _Rollup.Contract.IsBatchFinalized(&_Rollup.CallOpts, _batchIndex)
}

// IsChallenger is a free data retrieval call binding the contract method 0xa415d8dc.
//
// Solidity: function isChallenger(address ) view returns(bool)
func (_Rollup *RollupCaller) IsChallenger(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "isChallenger", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsChallenger is a free data retrieval call binding the contract method 0xa415d8dc.
//
// Solidity: function isChallenger(address ) view returns(bool)
func (_Rollup *RollupSession) IsChallenger(arg0 common.Address) (bool, error) {
	return _Rollup.Contract.IsChallenger(&_Rollup.CallOpts, arg0)
}

// IsChallenger is a free data retrieval call binding the contract method 0xa415d8dc.
//
// Solidity: function isChallenger(address ) view returns(bool)
func (_Rollup *RollupCallerSession) IsChallenger(arg0 common.Address) (bool, error) {
	return _Rollup.Contract.IsChallenger(&_Rollup.CallOpts, arg0)
}

// IsProver is a free data retrieval call binding the contract method 0x0a245924.
//
// Solidity: function isProver(address ) view returns(bool)
func (_Rollup *RollupCaller) IsProver(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "isProver", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProver is a free data retrieval call binding the contract method 0x0a245924.
//
// Solidity: function isProver(address ) view returns(bool)
func (_Rollup *RollupSession) IsProver(arg0 common.Address) (bool, error) {
	return _Rollup.Contract.IsProver(&_Rollup.CallOpts, arg0)
}

// IsProver is a free data retrieval call binding the contract method 0x0a245924.
//
// Solidity: function isProver(address ) view returns(bool)
func (_Rollup *RollupCallerSession) IsProver(arg0 common.Address) (bool, error) {
	return _Rollup.Contract.IsProver(&_Rollup.CallOpts, arg0)
}

// L1SequencerContract is a free data retrieval call binding the contract method 0x0ebb818d.
//
// Solidity: function l1SequencerContract() view returns(address)
func (_Rollup *RollupCaller) L1SequencerContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "l1SequencerContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1SequencerContract is a free data retrieval call binding the contract method 0x0ebb818d.
//
// Solidity: function l1SequencerContract() view returns(address)
func (_Rollup *RollupSession) L1SequencerContract() (common.Address, error) {
	return _Rollup.Contract.L1SequencerContract(&_Rollup.CallOpts)
}

// L1SequencerContract is a free data retrieval call binding the contract method 0x0ebb818d.
//
// Solidity: function l1SequencerContract() view returns(address)
func (_Rollup *RollupCallerSession) L1SequencerContract() (common.Address, error) {
	return _Rollup.Contract.L1SequencerContract(&_Rollup.CallOpts)
}

// L1StakingContract is a free data retrieval call binding the contract method 0xddd8a3dc.
//
// Solidity: function l1StakingContract() view returns(address)
func (_Rollup *RollupCaller) L1StakingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "l1StakingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1StakingContract is a free data retrieval call binding the contract method 0xddd8a3dc.
//
// Solidity: function l1StakingContract() view returns(address)
func (_Rollup *RollupSession) L1StakingContract() (common.Address, error) {
	return _Rollup.Contract.L1StakingContract(&_Rollup.CallOpts)
}

// L1StakingContract is a free data retrieval call binding the contract method 0xddd8a3dc.
//
// Solidity: function l1StakingContract() view returns(address)
func (_Rollup *RollupCallerSession) L1StakingContract() (common.Address, error) {
	return _Rollup.Contract.L1StakingContract(&_Rollup.CallOpts)
}

// LastCommittedBatchIndex is a free data retrieval call binding the contract method 0x121dcd50.
//
// Solidity: function lastCommittedBatchIndex() view returns(uint256)
func (_Rollup *RollupCaller) LastCommittedBatchIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "lastCommittedBatchIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastCommittedBatchIndex is a free data retrieval call binding the contract method 0x121dcd50.
//
// Solidity: function lastCommittedBatchIndex() view returns(uint256)
func (_Rollup *RollupSession) LastCommittedBatchIndex() (*big.Int, error) {
	return _Rollup.Contract.LastCommittedBatchIndex(&_Rollup.CallOpts)
}

// LastCommittedBatchIndex is a free data retrieval call binding the contract method 0x121dcd50.
//
// Solidity: function lastCommittedBatchIndex() view returns(uint256)
func (_Rollup *RollupCallerSession) LastCommittedBatchIndex() (*big.Int, error) {
	return _Rollup.Contract.LastCommittedBatchIndex(&_Rollup.CallOpts)
}

// LastFinalizedBatchIndex is a free data retrieval call binding the contract method 0x059def61.
//
// Solidity: function lastFinalizedBatchIndex() view returns(uint256)
func (_Rollup *RollupCaller) LastFinalizedBatchIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "lastFinalizedBatchIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastFinalizedBatchIndex is a free data retrieval call binding the contract method 0x059def61.
//
// Solidity: function lastFinalizedBatchIndex() view returns(uint256)
func (_Rollup *RollupSession) LastFinalizedBatchIndex() (*big.Int, error) {
	return _Rollup.Contract.LastFinalizedBatchIndex(&_Rollup.CallOpts)
}

// LastFinalizedBatchIndex is a free data retrieval call binding the contract method 0x059def61.
//
// Solidity: function lastFinalizedBatchIndex() view returns(uint256)
func (_Rollup *RollupCallerSession) LastFinalizedBatchIndex() (*big.Int, error) {
	return _Rollup.Contract.LastFinalizedBatchIndex(&_Rollup.CallOpts)
}

// LatestL2BlockNumber is a free data retrieval call binding the contract method 0x3e9e82ca.
//
// Solidity: function latestL2BlockNumber() view returns(uint256)
func (_Rollup *RollupCaller) LatestL2BlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "latestL2BlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestL2BlockNumber is a free data retrieval call binding the contract method 0x3e9e82ca.
//
// Solidity: function latestL2BlockNumber() view returns(uint256)
func (_Rollup *RollupSession) LatestL2BlockNumber() (*big.Int, error) {
	return _Rollup.Contract.LatestL2BlockNumber(&_Rollup.CallOpts)
}

// LatestL2BlockNumber is a free data retrieval call binding the contract method 0x3e9e82ca.
//
// Solidity: function latestL2BlockNumber() view returns(uint256)
func (_Rollup *RollupCallerSession) LatestL2BlockNumber() (*big.Int, error) {
	return _Rollup.Contract.LatestL2BlockNumber(&_Rollup.CallOpts)
}

// Layer2ChainId is a free data retrieval call binding the contract method 0x03c7f4af.
//
// Solidity: function layer2ChainId() view returns(uint64)
func (_Rollup *RollupCaller) Layer2ChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "layer2ChainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Layer2ChainId is a free data retrieval call binding the contract method 0x03c7f4af.
//
// Solidity: function layer2ChainId() view returns(uint64)
func (_Rollup *RollupSession) Layer2ChainId() (uint64, error) {
	return _Rollup.Contract.Layer2ChainId(&_Rollup.CallOpts)
}

// Layer2ChainId is a free data retrieval call binding the contract method 0x03c7f4af.
//
// Solidity: function layer2ChainId() view returns(uint64)
func (_Rollup *RollupCallerSession) Layer2ChainId() (uint64, error) {
	return _Rollup.Contract.Layer2ChainId(&_Rollup.CallOpts)
}

// MaxNumTxInChunk is a free data retrieval call binding the contract method 0xef6602ba.
//
// Solidity: function maxNumTxInChunk() view returns(uint256)
func (_Rollup *RollupCaller) MaxNumTxInChunk(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "maxNumTxInChunk")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxNumTxInChunk is a free data retrieval call binding the contract method 0xef6602ba.
//
// Solidity: function maxNumTxInChunk() view returns(uint256)
func (_Rollup *RollupSession) MaxNumTxInChunk() (*big.Int, error) {
	return _Rollup.Contract.MaxNumTxInChunk(&_Rollup.CallOpts)
}

// MaxNumTxInChunk is a free data retrieval call binding the contract method 0xef6602ba.
//
// Solidity: function maxNumTxInChunk() view returns(uint256)
func (_Rollup *RollupCallerSession) MaxNumTxInChunk() (*big.Int, error) {
	return _Rollup.Contract.MaxNumTxInChunk(&_Rollup.CallOpts)
}

// MessageQueue is a free data retrieval call binding the contract method 0x3b70c18a.
//
// Solidity: function messageQueue() view returns(address)
func (_Rollup *RollupCaller) MessageQueue(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "messageQueue")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageQueue is a free data retrieval call binding the contract method 0x3b70c18a.
//
// Solidity: function messageQueue() view returns(address)
func (_Rollup *RollupSession) MessageQueue() (common.Address, error) {
	return _Rollup.Contract.MessageQueue(&_Rollup.CallOpts)
}

// MessageQueue is a free data retrieval call binding the contract method 0x3b70c18a.
//
// Solidity: function messageQueue() view returns(address)
func (_Rollup *RollupCallerSession) MessageQueue() (common.Address, error) {
	return _Rollup.Contract.MessageQueue(&_Rollup.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rollup *RollupCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rollup *RollupSession) Owner() (common.Address, error) {
	return _Rollup.Contract.Owner(&_Rollup.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rollup *RollupCallerSession) Owner() (common.Address, error) {
	return _Rollup.Contract.Owner(&_Rollup.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Rollup *RollupCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Rollup *RollupSession) Paused() (bool, error) {
	return _Rollup.Contract.Paused(&_Rollup.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Rollup *RollupCallerSession) Paused() (bool, error) {
	return _Rollup.Contract.Paused(&_Rollup.CallOpts)
}

// RevertReqIndex is a free data retrieval call binding the contract method 0xb31a77d3.
//
// Solidity: function revertReqIndex() view returns(uint256)
func (_Rollup *RollupCaller) RevertReqIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "revertReqIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RevertReqIndex is a free data retrieval call binding the contract method 0xb31a77d3.
//
// Solidity: function revertReqIndex() view returns(uint256)
func (_Rollup *RollupSession) RevertReqIndex() (*big.Int, error) {
	return _Rollup.Contract.RevertReqIndex(&_Rollup.CallOpts)
}

// RevertReqIndex is a free data retrieval call binding the contract method 0xb31a77d3.
//
// Solidity: function revertReqIndex() view returns(uint256)
func (_Rollup *RollupCallerSession) RevertReqIndex() (*big.Int, error) {
	return _Rollup.Contract.RevertReqIndex(&_Rollup.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Rollup *RollupCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Rollup *RollupSession) Verifier() (common.Address, error) {
	return _Rollup.Contract.Verifier(&_Rollup.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Rollup *RollupCallerSession) Verifier() (common.Address, error) {
	return _Rollup.Contract.Verifier(&_Rollup.CallOpts)
}

// WithdrawalRoots is a free data retrieval call binding the contract method 0x04d77215.
//
// Solidity: function withdrawalRoots(bytes32 ) view returns(bool)
func (_Rollup *RollupCaller) WithdrawalRoots(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "withdrawalRoots", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawalRoots is a free data retrieval call binding the contract method 0x04d77215.
//
// Solidity: function withdrawalRoots(bytes32 ) view returns(bool)
func (_Rollup *RollupSession) WithdrawalRoots(arg0 [32]byte) (bool, error) {
	return _Rollup.Contract.WithdrawalRoots(&_Rollup.CallOpts, arg0)
}

// WithdrawalRoots is a free data retrieval call binding the contract method 0x04d77215.
//
// Solidity: function withdrawalRoots(bytes32 ) view returns(bool)
func (_Rollup *RollupCallerSession) WithdrawalRoots(arg0 [32]byte) (bool, error) {
	return _Rollup.Contract.WithdrawalRoots(&_Rollup.CallOpts, arg0)
}

// AddChallenger is a paid mutator transaction binding the contract method 0x0ceb6780.
//
// Solidity: function addChallenger(address _account) returns()
func (_Rollup *RollupTransactor) AddChallenger(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "addChallenger", _account)
}

// AddChallenger is a paid mutator transaction binding the contract method 0x0ceb6780.
//
// Solidity: function addChallenger(address _account) returns()
func (_Rollup *RollupSession) AddChallenger(_account common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.AddChallenger(&_Rollup.TransactOpts, _account)
}

// AddChallenger is a paid mutator transaction binding the contract method 0x0ceb6780.
//
// Solidity: function addChallenger(address _account) returns()
func (_Rollup *RollupTransactorSession) AddChallenger(_account common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.AddChallenger(&_Rollup.TransactOpts, _account)
}

// AddProver is a paid mutator transaction binding the contract method 0x1d49e457.
//
// Solidity: function addProver(address _account) returns()
func (_Rollup *RollupTransactor) AddProver(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "addProver", _account)
}

// AddProver is a paid mutator transaction binding the contract method 0x1d49e457.
//
// Solidity: function addProver(address _account) returns()
func (_Rollup *RollupSession) AddProver(_account common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.AddProver(&_Rollup.TransactOpts, _account)
}

// AddProver is a paid mutator transaction binding the contract method 0x1d49e457.
//
// Solidity: function addProver(address _account) returns()
func (_Rollup *RollupTransactorSession) AddProver(_account common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.AddProver(&_Rollup.TransactOpts, _account)
}

// ChallengeState is a paid mutator transaction binding the contract method 0x536deda5.
//
// Solidity: function challengeState(uint64 batchIndex, address challengerReceiver) payable returns()
func (_Rollup *RollupTransactor) ChallengeState(opts *bind.TransactOpts, batchIndex uint64, challengerReceiver common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "challengeState", batchIndex, challengerReceiver)
}

// ChallengeState is a paid mutator transaction binding the contract method 0x536deda5.
//
// Solidity: function challengeState(uint64 batchIndex, address challengerReceiver) payable returns()
func (_Rollup *RollupSession) ChallengeState(batchIndex uint64, challengerReceiver common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.ChallengeState(&_Rollup.TransactOpts, batchIndex, challengerReceiver)
}

// ChallengeState is a paid mutator transaction binding the contract method 0x536deda5.
//
// Solidity: function challengeState(uint64 batchIndex, address challengerReceiver) payable returns()
func (_Rollup *RollupTransactorSession) ChallengeState(batchIndex uint64, challengerReceiver common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.ChallengeState(&_Rollup.TransactOpts, batchIndex, challengerReceiver)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_Rollup *RollupTransactor) ClaimReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "claimReward")
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_Rollup *RollupSession) ClaimReward() (*types.Transaction, error) {
	return _Rollup.Contract.ClaimReward(&_Rollup.TransactOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_Rollup *RollupTransactorSession) ClaimReward() (*types.Transaction, error) {
	return _Rollup.Contract.ClaimReward(&_Rollup.TransactOpts)
}

// CommitBatch is a paid mutator transaction binding the contract method 0x4db4d964.
//
// Solidity: function commitBatch((uint8,bytes,bytes[],bytes,bytes32,bytes32,bytes32,(uint256,uint256[],bytes)) batchData, uint256 version, address[] sequencers, bytes signature) payable returns()
func (_Rollup *RollupTransactor) CommitBatch(opts *bind.TransactOpts, batchData IRollupBatchData, version *big.Int, sequencers []common.Address, signature []byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "commitBatch", batchData, version, sequencers, signature)
}

// CommitBatch is a paid mutator transaction binding the contract method 0x4db4d964.
//
// Solidity: function commitBatch((uint8,bytes,bytes[],bytes,bytes32,bytes32,bytes32,(uint256,uint256[],bytes)) batchData, uint256 version, address[] sequencers, bytes signature) payable returns()
func (_Rollup *RollupSession) CommitBatch(batchData IRollupBatchData, version *big.Int, sequencers []common.Address, signature []byte) (*types.Transaction, error) {
	return _Rollup.Contract.CommitBatch(&_Rollup.TransactOpts, batchData, version, sequencers, signature)
}

// CommitBatch is a paid mutator transaction binding the contract method 0x4db4d964.
//
// Solidity: function commitBatch((uint8,bytes,bytes[],bytes,bytes32,bytes32,bytes32,(uint256,uint256[],bytes)) batchData, uint256 version, address[] sequencers, bytes signature) payable returns()
func (_Rollup *RollupTransactorSession) CommitBatch(batchData IRollupBatchData, version *big.Int, sequencers []common.Address, signature []byte) (*types.Transaction, error) {
	return _Rollup.Contract.CommitBatch(&_Rollup.TransactOpts, batchData, version, sequencers, signature)
}

// FinalizeBatch is a paid mutator transaction binding the contract method 0xe33491a7.
//
// Solidity: function finalizeBatch(uint256 _batchIndex) returns()
func (_Rollup *RollupTransactor) FinalizeBatch(opts *bind.TransactOpts, _batchIndex *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "finalizeBatch", _batchIndex)
}

// FinalizeBatch is a paid mutator transaction binding the contract method 0xe33491a7.
//
// Solidity: function finalizeBatch(uint256 _batchIndex) returns()
func (_Rollup *RollupSession) FinalizeBatch(_batchIndex *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.FinalizeBatch(&_Rollup.TransactOpts, _batchIndex)
}

// FinalizeBatch is a paid mutator transaction binding the contract method 0xe33491a7.
//
// Solidity: function finalizeBatch(uint256 _batchIndex) returns()
func (_Rollup *RollupTransactorSession) FinalizeBatch(_batchIndex *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.FinalizeBatch(&_Rollup.TransactOpts, _batchIndex)
}

// ImportGenesisBatch is a paid mutator transaction binding the contract method 0x4c4b9e4f.
//
// Solidity: function importGenesisBatch(bytes _batchHeader, bytes32 _postStateRoot, bytes32 _withdrawalRoot) returns()
func (_Rollup *RollupTransactor) ImportGenesisBatch(opts *bind.TransactOpts, _batchHeader []byte, _postStateRoot [32]byte, _withdrawalRoot [32]byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "importGenesisBatch", _batchHeader, _postStateRoot, _withdrawalRoot)
}

// ImportGenesisBatch is a paid mutator transaction binding the contract method 0x4c4b9e4f.
//
// Solidity: function importGenesisBatch(bytes _batchHeader, bytes32 _postStateRoot, bytes32 _withdrawalRoot) returns()
func (_Rollup *RollupSession) ImportGenesisBatch(_batchHeader []byte, _postStateRoot [32]byte, _withdrawalRoot [32]byte) (*types.Transaction, error) {
	return _Rollup.Contract.ImportGenesisBatch(&_Rollup.TransactOpts, _batchHeader, _postStateRoot, _withdrawalRoot)
}

// ImportGenesisBatch is a paid mutator transaction binding the contract method 0x4c4b9e4f.
//
// Solidity: function importGenesisBatch(bytes _batchHeader, bytes32 _postStateRoot, bytes32 _withdrawalRoot) returns()
func (_Rollup *RollupTransactorSession) ImportGenesisBatch(_batchHeader []byte, _postStateRoot [32]byte, _withdrawalRoot [32]byte) (*types.Transaction, error) {
	return _Rollup.Contract.ImportGenesisBatch(&_Rollup.TransactOpts, _batchHeader, _postStateRoot, _withdrawalRoot)
}

// Initialize is a paid mutator transaction binding the contract method 0xe77fc7a4.
//
// Solidity: function initialize(address _l1SequencerContract, address _l1StakingContract, address _messageQueue, address _verifier, uint256 _maxNumTxInChunk, uint256 _finalizationPeriodSeconds, uint256 _proofWindow) returns()
func (_Rollup *RollupTransactor) Initialize(opts *bind.TransactOpts, _l1SequencerContract common.Address, _l1StakingContract common.Address, _messageQueue common.Address, _verifier common.Address, _maxNumTxInChunk *big.Int, _finalizationPeriodSeconds *big.Int, _proofWindow *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "initialize", _l1SequencerContract, _l1StakingContract, _messageQueue, _verifier, _maxNumTxInChunk, _finalizationPeriodSeconds, _proofWindow)
}

// Initialize is a paid mutator transaction binding the contract method 0xe77fc7a4.
//
// Solidity: function initialize(address _l1SequencerContract, address _l1StakingContract, address _messageQueue, address _verifier, uint256 _maxNumTxInChunk, uint256 _finalizationPeriodSeconds, uint256 _proofWindow) returns()
func (_Rollup *RollupSession) Initialize(_l1SequencerContract common.Address, _l1StakingContract common.Address, _messageQueue common.Address, _verifier common.Address, _maxNumTxInChunk *big.Int, _finalizationPeriodSeconds *big.Int, _proofWindow *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.Initialize(&_Rollup.TransactOpts, _l1SequencerContract, _l1StakingContract, _messageQueue, _verifier, _maxNumTxInChunk, _finalizationPeriodSeconds, _proofWindow)
}

// Initialize is a paid mutator transaction binding the contract method 0xe77fc7a4.
//
// Solidity: function initialize(address _l1SequencerContract, address _l1StakingContract, address _messageQueue, address _verifier, uint256 _maxNumTxInChunk, uint256 _finalizationPeriodSeconds, uint256 _proofWindow) returns()
func (_Rollup *RollupTransactorSession) Initialize(_l1SequencerContract common.Address, _l1StakingContract common.Address, _messageQueue common.Address, _verifier common.Address, _maxNumTxInChunk *big.Int, _finalizationPeriodSeconds *big.Int, _proofWindow *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.Initialize(&_Rollup.TransactOpts, _l1SequencerContract, _l1StakingContract, _messageQueue, _verifier, _maxNumTxInChunk, _finalizationPeriodSeconds, _proofWindow)
}

// ProveState is a paid mutator transaction binding the contract method 0x8c396894.
//
// Solidity: function proveState(uint64 _batchIndex, address _proverReceiveAddress, bytes _aggrProof, bytes _kzgDataProof, uint32 _minGasLimit) returns()
func (_Rollup *RollupTransactor) ProveState(opts *bind.TransactOpts, _batchIndex uint64, _proverReceiveAddress common.Address, _aggrProof []byte, _kzgDataProof []byte, _minGasLimit uint32) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "proveState", _batchIndex, _proverReceiveAddress, _aggrProof, _kzgDataProof, _minGasLimit)
}

// ProveState is a paid mutator transaction binding the contract method 0x8c396894.
//
// Solidity: function proveState(uint64 _batchIndex, address _proverReceiveAddress, bytes _aggrProof, bytes _kzgDataProof, uint32 _minGasLimit) returns()
func (_Rollup *RollupSession) ProveState(_batchIndex uint64, _proverReceiveAddress common.Address, _aggrProof []byte, _kzgDataProof []byte, _minGasLimit uint32) (*types.Transaction, error) {
	return _Rollup.Contract.ProveState(&_Rollup.TransactOpts, _batchIndex, _proverReceiveAddress, _aggrProof, _kzgDataProof, _minGasLimit)
}

// ProveState is a paid mutator transaction binding the contract method 0x8c396894.
//
// Solidity: function proveState(uint64 _batchIndex, address _proverReceiveAddress, bytes _aggrProof, bytes _kzgDataProof, uint32 _minGasLimit) returns()
func (_Rollup *RollupTransactorSession) ProveState(_batchIndex uint64, _proverReceiveAddress common.Address, _aggrProof []byte, _kzgDataProof []byte, _minGasLimit uint32) (*types.Transaction, error) {
	return _Rollup.Contract.ProveState(&_Rollup.TransactOpts, _batchIndex, _proverReceiveAddress, _aggrProof, _kzgDataProof, _minGasLimit)
}

// RemoveChallenger is a paid mutator transaction binding the contract method 0x6c578c1d.
//
// Solidity: function removeChallenger(address _account) returns()
func (_Rollup *RollupTransactor) RemoveChallenger(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "removeChallenger", _account)
}

// RemoveChallenger is a paid mutator transaction binding the contract method 0x6c578c1d.
//
// Solidity: function removeChallenger(address _account) returns()
func (_Rollup *RollupSession) RemoveChallenger(_account common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.RemoveChallenger(&_Rollup.TransactOpts, _account)
}

// RemoveChallenger is a paid mutator transaction binding the contract method 0x6c578c1d.
//
// Solidity: function removeChallenger(address _account) returns()
func (_Rollup *RollupTransactorSession) RemoveChallenger(_account common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.RemoveChallenger(&_Rollup.TransactOpts, _account)
}

// RemoveProver is a paid mutator transaction binding the contract method 0xb571d3dd.
//
// Solidity: function removeProver(address _account) returns()
func (_Rollup *RollupTransactor) RemoveProver(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "removeProver", _account)
}

// RemoveProver is a paid mutator transaction binding the contract method 0xb571d3dd.
//
// Solidity: function removeProver(address _account) returns()
func (_Rollup *RollupSession) RemoveProver(_account common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.RemoveProver(&_Rollup.TransactOpts, _account)
}

// RemoveProver is a paid mutator transaction binding the contract method 0xb571d3dd.
//
// Solidity: function removeProver(address _account) returns()
func (_Rollup *RollupTransactorSession) RemoveProver(_account common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.RemoveProver(&_Rollup.TransactOpts, _account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rollup *RollupTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rollup *RollupSession) RenounceOwnership() (*types.Transaction, error) {
	return _Rollup.Contract.RenounceOwnership(&_Rollup.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rollup *RollupTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Rollup.Contract.RenounceOwnership(&_Rollup.TransactOpts)
}

// RevertBatch is a paid mutator transaction binding the contract method 0x10d44583.
//
// Solidity: function revertBatch(bytes _batchHeader, uint256 _count) returns()
func (_Rollup *RollupTransactor) RevertBatch(opts *bind.TransactOpts, _batchHeader []byte, _count *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "revertBatch", _batchHeader, _count)
}

// RevertBatch is a paid mutator transaction binding the contract method 0x10d44583.
//
// Solidity: function revertBatch(bytes _batchHeader, uint256 _count) returns()
func (_Rollup *RollupSession) RevertBatch(_batchHeader []byte, _count *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.RevertBatch(&_Rollup.TransactOpts, _batchHeader, _count)
}

// RevertBatch is a paid mutator transaction binding the contract method 0x10d44583.
//
// Solidity: function revertBatch(bytes _batchHeader, uint256 _count) returns()
func (_Rollup *RollupTransactorSession) RevertBatch(_batchHeader []byte, _count *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.RevertBatch(&_Rollup.TransactOpts, _batchHeader, _count)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _status) returns()
func (_Rollup *RollupTransactor) SetPause(opts *bind.TransactOpts, _status bool) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "setPause", _status)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _status) returns()
func (_Rollup *RollupSession) SetPause(_status bool) (*types.Transaction, error) {
	return _Rollup.Contract.SetPause(&_Rollup.TransactOpts, _status)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _status) returns()
func (_Rollup *RollupTransactorSession) SetPause(_status bool) (*types.Transaction, error) {
	return _Rollup.Contract.SetPause(&_Rollup.TransactOpts, _status)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rollup *RollupTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rollup *RollupSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.TransferOwnership(&_Rollup.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rollup *RollupTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.TransferOwnership(&_Rollup.TransactOpts, newOwner)
}

// UpdateFinalizePeriodSeconds is a paid mutator transaction binding the contract method 0xe3fff1dd.
//
// Solidity: function updateFinalizePeriodSeconds(uint256 _newPeriod) returns()
func (_Rollup *RollupTransactor) UpdateFinalizePeriodSeconds(opts *bind.TransactOpts, _newPeriod *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "updateFinalizePeriodSeconds", _newPeriod)
}

// UpdateFinalizePeriodSeconds is a paid mutator transaction binding the contract method 0xe3fff1dd.
//
// Solidity: function updateFinalizePeriodSeconds(uint256 _newPeriod) returns()
func (_Rollup *RollupSession) UpdateFinalizePeriodSeconds(_newPeriod *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.UpdateFinalizePeriodSeconds(&_Rollup.TransactOpts, _newPeriod)
}

// UpdateFinalizePeriodSeconds is a paid mutator transaction binding the contract method 0xe3fff1dd.
//
// Solidity: function updateFinalizePeriodSeconds(uint256 _newPeriod) returns()
func (_Rollup *RollupTransactorSession) UpdateFinalizePeriodSeconds(_newPeriod *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.UpdateFinalizePeriodSeconds(&_Rollup.TransactOpts, _newPeriod)
}

// UpdateMaxNumTxInChunk is a paid mutator transaction binding the contract method 0x1e228302.
//
// Solidity: function updateMaxNumTxInChunk(uint256 _maxNumTxInChunk) returns()
func (_Rollup *RollupTransactor) UpdateMaxNumTxInChunk(opts *bind.TransactOpts, _maxNumTxInChunk *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "updateMaxNumTxInChunk", _maxNumTxInChunk)
}

// UpdateMaxNumTxInChunk is a paid mutator transaction binding the contract method 0x1e228302.
//
// Solidity: function updateMaxNumTxInChunk(uint256 _maxNumTxInChunk) returns()
func (_Rollup *RollupSession) UpdateMaxNumTxInChunk(_maxNumTxInChunk *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.UpdateMaxNumTxInChunk(&_Rollup.TransactOpts, _maxNumTxInChunk)
}

// UpdateMaxNumTxInChunk is a paid mutator transaction binding the contract method 0x1e228302.
//
// Solidity: function updateMaxNumTxInChunk(uint256 _maxNumTxInChunk) returns()
func (_Rollup *RollupTransactorSession) UpdateMaxNumTxInChunk(_maxNumTxInChunk *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.UpdateMaxNumTxInChunk(&_Rollup.TransactOpts, _maxNumTxInChunk)
}

// UpdateProofWindow is a paid mutator transaction binding the contract method 0x57e0af6c.
//
// Solidity: function updateProofWindow(uint256 _newWindow) returns()
func (_Rollup *RollupTransactor) UpdateProofWindow(opts *bind.TransactOpts, _newWindow *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "updateProofWindow", _newWindow)
}

// UpdateProofWindow is a paid mutator transaction binding the contract method 0x57e0af6c.
//
// Solidity: function updateProofWindow(uint256 _newWindow) returns()
func (_Rollup *RollupSession) UpdateProofWindow(_newWindow *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.UpdateProofWindow(&_Rollup.TransactOpts, _newWindow)
}

// UpdateProofWindow is a paid mutator transaction binding the contract method 0x57e0af6c.
//
// Solidity: function updateProofWindow(uint256 _newWindow) returns()
func (_Rollup *RollupTransactorSession) UpdateProofWindow(_newWindow *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.UpdateProofWindow(&_Rollup.TransactOpts, _newWindow)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x97fc007c.
//
// Solidity: function updateVerifier(address _newVerifier) returns()
func (_Rollup *RollupTransactor) UpdateVerifier(opts *bind.TransactOpts, _newVerifier common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "updateVerifier", _newVerifier)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x97fc007c.
//
// Solidity: function updateVerifier(address _newVerifier) returns()
func (_Rollup *RollupSession) UpdateVerifier(_newVerifier common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.UpdateVerifier(&_Rollup.TransactOpts, _newVerifier)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x97fc007c.
//
// Solidity: function updateVerifier(address _newVerifier) returns()
func (_Rollup *RollupTransactorSession) UpdateVerifier(_newVerifier common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.UpdateVerifier(&_Rollup.TransactOpts, _newVerifier)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Rollup *RollupTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Rollup *RollupSession) Receive() (*types.Transaction, error) {
	return _Rollup.Contract.Receive(&_Rollup.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Rollup *RollupTransactorSession) Receive() (*types.Transaction, error) {
	return _Rollup.Contract.Receive(&_Rollup.TransactOpts)
}

// RollupChallengeResIterator is returned from FilterChallengeRes and is used to iterate over the raw logs and unpacked data for ChallengeRes events raised by the Rollup contract.
type RollupChallengeResIterator struct {
	Event *RollupChallengeRes // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupChallengeResIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupChallengeRes)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupChallengeRes)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupChallengeResIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupChallengeResIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupChallengeRes represents a ChallengeRes event raised by the Rollup contract.
type RollupChallengeRes struct {
	BatchIndex uint64
	Winner     common.Address
	Res        string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChallengeRes is a free log retrieval operation binding the contract event 0x1e66d5dca70bf28588ef2f5cb3c299e65e2e7bdef2767823d3ae47a9caff95c6.
//
// Solidity: event ChallengeRes(uint64 indexed batchIndex, address winner, string res)
func (_Rollup *RollupFilterer) FilterChallengeRes(opts *bind.FilterOpts, batchIndex []uint64) (*RollupChallengeResIterator, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "ChallengeRes", batchIndexRule)
	if err != nil {
		return nil, err
	}
	return &RollupChallengeResIterator{contract: _Rollup.contract, event: "ChallengeRes", logs: logs, sub: sub}, nil
}

// WatchChallengeRes is a free log subscription operation binding the contract event 0x1e66d5dca70bf28588ef2f5cb3c299e65e2e7bdef2767823d3ae47a9caff95c6.
//
// Solidity: event ChallengeRes(uint64 indexed batchIndex, address winner, string res)
func (_Rollup *RollupFilterer) WatchChallengeRes(opts *bind.WatchOpts, sink chan<- *RollupChallengeRes, batchIndex []uint64) (event.Subscription, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "ChallengeRes", batchIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupChallengeRes)
				if err := _Rollup.contract.UnpackLog(event, "ChallengeRes", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChallengeRes is a log parse operation binding the contract event 0x1e66d5dca70bf28588ef2f5cb3c299e65e2e7bdef2767823d3ae47a9caff95c6.
//
// Solidity: event ChallengeRes(uint64 indexed batchIndex, address winner, string res)
func (_Rollup *RollupFilterer) ParseChallengeRes(log types.Log) (*RollupChallengeRes, error) {
	event := new(RollupChallengeRes)
	if err := _Rollup.contract.UnpackLog(event, "ChallengeRes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupChallengeStateIterator is returned from FilterChallengeState and is used to iterate over the raw logs and unpacked data for ChallengeState events raised by the Rollup contract.
type RollupChallengeStateIterator struct {
	Event *RollupChallengeState // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupChallengeStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupChallengeState)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupChallengeState)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupChallengeStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupChallengeStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupChallengeState represents a ChallengeState event raised by the Rollup contract.
type RollupChallengeState struct {
	BatchIndex               uint64
	Challenger               common.Address
	ChallengerReceiveAddress common.Address
	ChallengeDeposit         *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterChallengeState is a free log retrieval operation binding the contract event 0x961c03af6cc5d073982308a314a1bfb67a460342e8d5468e4d6385172cc617b1.
//
// Solidity: event ChallengeState(uint64 indexed batchIndex, address challenger, address challengerReceiveAddress, uint256 challengeDeposit)
func (_Rollup *RollupFilterer) FilterChallengeState(opts *bind.FilterOpts, batchIndex []uint64) (*RollupChallengeStateIterator, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "ChallengeState", batchIndexRule)
	if err != nil {
		return nil, err
	}
	return &RollupChallengeStateIterator{contract: _Rollup.contract, event: "ChallengeState", logs: logs, sub: sub}, nil
}

// WatchChallengeState is a free log subscription operation binding the contract event 0x961c03af6cc5d073982308a314a1bfb67a460342e8d5468e4d6385172cc617b1.
//
// Solidity: event ChallengeState(uint64 indexed batchIndex, address challenger, address challengerReceiveAddress, uint256 challengeDeposit)
func (_Rollup *RollupFilterer) WatchChallengeState(opts *bind.WatchOpts, sink chan<- *RollupChallengeState, batchIndex []uint64) (event.Subscription, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "ChallengeState", batchIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupChallengeState)
				if err := _Rollup.contract.UnpackLog(event, "ChallengeState", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChallengeState is a log parse operation binding the contract event 0x961c03af6cc5d073982308a314a1bfb67a460342e8d5468e4d6385172cc617b1.
//
// Solidity: event ChallengeState(uint64 indexed batchIndex, address challenger, address challengerReceiveAddress, uint256 challengeDeposit)
func (_Rollup *RollupFilterer) ParseChallengeState(log types.Log) (*RollupChallengeState, error) {
	event := new(RollupChallengeState)
	if err := _Rollup.contract.UnpackLog(event, "ChallengeState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupCommitBatchIterator is returned from FilterCommitBatch and is used to iterate over the raw logs and unpacked data for CommitBatch events raised by the Rollup contract.
type RollupCommitBatchIterator struct {
	Event *RollupCommitBatch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupCommitBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupCommitBatch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupCommitBatch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupCommitBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupCommitBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupCommitBatch represents a CommitBatch event raised by the Rollup contract.
type RollupCommitBatch struct {
	BatchIndex *big.Int
	BatchHash  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCommitBatch is a free log retrieval operation binding the contract event 0x2c32d4ae151744d0bf0b9464a3e897a1d17ed2f1af71f7c9a75f12ce0d28238f.
//
// Solidity: event CommitBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_Rollup *RollupFilterer) FilterCommitBatch(opts *bind.FilterOpts, batchIndex []*big.Int, batchHash [][32]byte) (*RollupCommitBatchIterator, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "CommitBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return &RollupCommitBatchIterator{contract: _Rollup.contract, event: "CommitBatch", logs: logs, sub: sub}, nil
}

// WatchCommitBatch is a free log subscription operation binding the contract event 0x2c32d4ae151744d0bf0b9464a3e897a1d17ed2f1af71f7c9a75f12ce0d28238f.
//
// Solidity: event CommitBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_Rollup *RollupFilterer) WatchCommitBatch(opts *bind.WatchOpts, sink chan<- *RollupCommitBatch, batchIndex []*big.Int, batchHash [][32]byte) (event.Subscription, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "CommitBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupCommitBatch)
				if err := _Rollup.contract.UnpackLog(event, "CommitBatch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCommitBatch is a log parse operation binding the contract event 0x2c32d4ae151744d0bf0b9464a3e897a1d17ed2f1af71f7c9a75f12ce0d28238f.
//
// Solidity: event CommitBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_Rollup *RollupFilterer) ParseCommitBatch(log types.Log) (*RollupCommitBatch, error) {
	event := new(RollupCommitBatch)
	if err := _Rollup.contract.UnpackLog(event, "CommitBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupFinalizeBatchIterator is returned from FilterFinalizeBatch and is used to iterate over the raw logs and unpacked data for FinalizeBatch events raised by the Rollup contract.
type RollupFinalizeBatchIterator struct {
	Event *RollupFinalizeBatch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupFinalizeBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupFinalizeBatch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupFinalizeBatch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupFinalizeBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupFinalizeBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupFinalizeBatch represents a FinalizeBatch event raised by the Rollup contract.
type RollupFinalizeBatch struct {
	BatchIndex   *big.Int
	BatchHash    [32]byte
	StateRoot    [32]byte
	WithdrawRoot [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFinalizeBatch is a free log retrieval operation binding the contract event 0x26ba82f907317eedc97d0cbef23de76a43dd6edb563bdb6e9407645b950a7a2d.
//
// Solidity: event FinalizeBatch(uint256 indexed batchIndex, bytes32 indexed batchHash, bytes32 stateRoot, bytes32 withdrawRoot)
func (_Rollup *RollupFilterer) FilterFinalizeBatch(opts *bind.FilterOpts, batchIndex []*big.Int, batchHash [][32]byte) (*RollupFinalizeBatchIterator, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "FinalizeBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return &RollupFinalizeBatchIterator{contract: _Rollup.contract, event: "FinalizeBatch", logs: logs, sub: sub}, nil
}

// WatchFinalizeBatch is a free log subscription operation binding the contract event 0x26ba82f907317eedc97d0cbef23de76a43dd6edb563bdb6e9407645b950a7a2d.
//
// Solidity: event FinalizeBatch(uint256 indexed batchIndex, bytes32 indexed batchHash, bytes32 stateRoot, bytes32 withdrawRoot)
func (_Rollup *RollupFilterer) WatchFinalizeBatch(opts *bind.WatchOpts, sink chan<- *RollupFinalizeBatch, batchIndex []*big.Int, batchHash [][32]byte) (event.Subscription, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "FinalizeBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupFinalizeBatch)
				if err := _Rollup.contract.UnpackLog(event, "FinalizeBatch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFinalizeBatch is a log parse operation binding the contract event 0x26ba82f907317eedc97d0cbef23de76a43dd6edb563bdb6e9407645b950a7a2d.
//
// Solidity: event FinalizeBatch(uint256 indexed batchIndex, bytes32 indexed batchHash, bytes32 stateRoot, bytes32 withdrawRoot)
func (_Rollup *RollupFilterer) ParseFinalizeBatch(log types.Log) (*RollupFinalizeBatch, error) {
	event := new(RollupFinalizeBatch)
	if err := _Rollup.contract.UnpackLog(event, "FinalizeBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Rollup contract.
type RollupInitializedIterator struct {
	Event *RollupInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupInitialized represents a Initialized event raised by the Rollup contract.
type RollupInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Rollup *RollupFilterer) FilterInitialized(opts *bind.FilterOpts) (*RollupInitializedIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RollupInitializedIterator{contract: _Rollup.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Rollup *RollupFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RollupInitialized) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupInitialized)
				if err := _Rollup.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Rollup *RollupFilterer) ParseInitialized(log types.Log) (*RollupInitialized, error) {
	event := new(RollupInitialized)
	if err := _Rollup.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Rollup contract.
type RollupOwnershipTransferredIterator struct {
	Event *RollupOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupOwnershipTransferred represents a OwnershipTransferred event raised by the Rollup contract.
type RollupOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rollup *RollupFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RollupOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RollupOwnershipTransferredIterator{contract: _Rollup.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rollup *RollupFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RollupOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupOwnershipTransferred)
				if err := _Rollup.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rollup *RollupFilterer) ParseOwnershipTransferred(log types.Log) (*RollupOwnershipTransferred, error) {
	event := new(RollupOwnershipTransferred)
	if err := _Rollup.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Rollup contract.
type RollupPausedIterator struct {
	Event *RollupPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupPaused represents a Paused event raised by the Rollup contract.
type RollupPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Rollup *RollupFilterer) FilterPaused(opts *bind.FilterOpts) (*RollupPausedIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RollupPausedIterator{contract: _Rollup.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Rollup *RollupFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RollupPaused) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupPaused)
				if err := _Rollup.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Rollup *RollupFilterer) ParsePaused(log types.Log) (*RollupPaused, error) {
	event := new(RollupPaused)
	if err := _Rollup.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupRevertBatchIterator is returned from FilterRevertBatch and is used to iterate over the raw logs and unpacked data for RevertBatch events raised by the Rollup contract.
type RollupRevertBatchIterator struct {
	Event *RollupRevertBatch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupRevertBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupRevertBatch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupRevertBatch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupRevertBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupRevertBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupRevertBatch represents a RevertBatch event raised by the Rollup contract.
type RollupRevertBatch struct {
	BatchIndex *big.Int
	BatchHash  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRevertBatch is a free log retrieval operation binding the contract event 0x00cae2739091badfd91c373f0a16cede691e0cd25bb80cff77dd5caeb4710146.
//
// Solidity: event RevertBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_Rollup *RollupFilterer) FilterRevertBatch(opts *bind.FilterOpts, batchIndex []*big.Int, batchHash [][32]byte) (*RollupRevertBatchIterator, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "RevertBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return &RollupRevertBatchIterator{contract: _Rollup.contract, event: "RevertBatch", logs: logs, sub: sub}, nil
}

// WatchRevertBatch is a free log subscription operation binding the contract event 0x00cae2739091badfd91c373f0a16cede691e0cd25bb80cff77dd5caeb4710146.
//
// Solidity: event RevertBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_Rollup *RollupFilterer) WatchRevertBatch(opts *bind.WatchOpts, sink chan<- *RollupRevertBatch, batchIndex []*big.Int, batchHash [][32]byte) (event.Subscription, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "RevertBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupRevertBatch)
				if err := _Rollup.contract.UnpackLog(event, "RevertBatch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevertBatch is a log parse operation binding the contract event 0x00cae2739091badfd91c373f0a16cede691e0cd25bb80cff77dd5caeb4710146.
//
// Solidity: event RevertBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_Rollup *RollupFilterer) ParseRevertBatch(log types.Log) (*RollupRevertBatch, error) {
	event := new(RollupRevertBatch)
	if err := _Rollup.contract.UnpackLog(event, "RevertBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Rollup contract.
type RollupUnpausedIterator struct {
	Event *RollupUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupUnpaused represents a Unpaused event raised by the Rollup contract.
type RollupUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Rollup *RollupFilterer) FilterUnpaused(opts *bind.FilterOpts) (*RollupUnpausedIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &RollupUnpausedIterator{contract: _Rollup.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Rollup *RollupFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RollupUnpaused) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupUnpaused)
				if err := _Rollup.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Rollup *RollupFilterer) ParseUnpaused(log types.Log) (*RollupUnpaused, error) {
	event := new(RollupUnpaused)
	if err := _Rollup.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupUpdateChallengerIterator is returned from FilterUpdateChallenger and is used to iterate over the raw logs and unpacked data for UpdateChallenger events raised by the Rollup contract.
type RollupUpdateChallengerIterator struct {
	Event *RollupUpdateChallenger // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupUpdateChallengerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupUpdateChallenger)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupUpdateChallenger)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupUpdateChallengerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupUpdateChallengerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupUpdateChallenger represents a UpdateChallenger event raised by the Rollup contract.
type RollupUpdateChallenger struct {
	Account common.Address
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateChallenger is a free log retrieval operation binding the contract event 0x7feb653c7b1f0d23daeed494225b3f28851cdc8973fcc653866d9b6e205fc009.
//
// Solidity: event UpdateChallenger(address indexed account, bool status)
func (_Rollup *RollupFilterer) FilterUpdateChallenger(opts *bind.FilterOpts, account []common.Address) (*RollupUpdateChallengerIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "UpdateChallenger", accountRule)
	if err != nil {
		return nil, err
	}
	return &RollupUpdateChallengerIterator{contract: _Rollup.contract, event: "UpdateChallenger", logs: logs, sub: sub}, nil
}

// WatchUpdateChallenger is a free log subscription operation binding the contract event 0x7feb653c7b1f0d23daeed494225b3f28851cdc8973fcc653866d9b6e205fc009.
//
// Solidity: event UpdateChallenger(address indexed account, bool status)
func (_Rollup *RollupFilterer) WatchUpdateChallenger(opts *bind.WatchOpts, sink chan<- *RollupUpdateChallenger, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "UpdateChallenger", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupUpdateChallenger)
				if err := _Rollup.contract.UnpackLog(event, "UpdateChallenger", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateChallenger is a log parse operation binding the contract event 0x7feb653c7b1f0d23daeed494225b3f28851cdc8973fcc653866d9b6e205fc009.
//
// Solidity: event UpdateChallenger(address indexed account, bool status)
func (_Rollup *RollupFilterer) ParseUpdateChallenger(log types.Log) (*RollupUpdateChallenger, error) {
	event := new(RollupUpdateChallenger)
	if err := _Rollup.contract.UnpackLog(event, "UpdateChallenger", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupUpdateFinalizationPeriodSecondsIterator is returned from FilterUpdateFinalizationPeriodSeconds and is used to iterate over the raw logs and unpacked data for UpdateFinalizationPeriodSeconds events raised by the Rollup contract.
type RollupUpdateFinalizationPeriodSecondsIterator struct {
	Event *RollupUpdateFinalizationPeriodSeconds // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupUpdateFinalizationPeriodSecondsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupUpdateFinalizationPeriodSeconds)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupUpdateFinalizationPeriodSeconds)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupUpdateFinalizationPeriodSecondsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupUpdateFinalizationPeriodSecondsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupUpdateFinalizationPeriodSeconds represents a UpdateFinalizationPeriodSeconds event raised by the Rollup contract.
type RollupUpdateFinalizationPeriodSeconds struct {
	OldPeriod *big.Int
	NewPeriod *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateFinalizationPeriodSeconds is a free log retrieval operation binding the contract event 0xa577f4223f91f74e2dad65bbb8c30807587ae95d0d34288057bb3ec0d398a437.
//
// Solidity: event UpdateFinalizationPeriodSeconds(uint256 oldPeriod, uint256 newPeriod)
func (_Rollup *RollupFilterer) FilterUpdateFinalizationPeriodSeconds(opts *bind.FilterOpts) (*RollupUpdateFinalizationPeriodSecondsIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "UpdateFinalizationPeriodSeconds")
	if err != nil {
		return nil, err
	}
	return &RollupUpdateFinalizationPeriodSecondsIterator{contract: _Rollup.contract, event: "UpdateFinalizationPeriodSeconds", logs: logs, sub: sub}, nil
}

// WatchUpdateFinalizationPeriodSeconds is a free log subscription operation binding the contract event 0xa577f4223f91f74e2dad65bbb8c30807587ae95d0d34288057bb3ec0d398a437.
//
// Solidity: event UpdateFinalizationPeriodSeconds(uint256 oldPeriod, uint256 newPeriod)
func (_Rollup *RollupFilterer) WatchUpdateFinalizationPeriodSeconds(opts *bind.WatchOpts, sink chan<- *RollupUpdateFinalizationPeriodSeconds) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "UpdateFinalizationPeriodSeconds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupUpdateFinalizationPeriodSeconds)
				if err := _Rollup.contract.UnpackLog(event, "UpdateFinalizationPeriodSeconds", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateFinalizationPeriodSeconds is a log parse operation binding the contract event 0xa577f4223f91f74e2dad65bbb8c30807587ae95d0d34288057bb3ec0d398a437.
//
// Solidity: event UpdateFinalizationPeriodSeconds(uint256 oldPeriod, uint256 newPeriod)
func (_Rollup *RollupFilterer) ParseUpdateFinalizationPeriodSeconds(log types.Log) (*RollupUpdateFinalizationPeriodSeconds, error) {
	event := new(RollupUpdateFinalizationPeriodSeconds)
	if err := _Rollup.contract.UnpackLog(event, "UpdateFinalizationPeriodSeconds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupUpdateMaxNumTxInChunkIterator is returned from FilterUpdateMaxNumTxInChunk and is used to iterate over the raw logs and unpacked data for UpdateMaxNumTxInChunk events raised by the Rollup contract.
type RollupUpdateMaxNumTxInChunkIterator struct {
	Event *RollupUpdateMaxNumTxInChunk // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupUpdateMaxNumTxInChunkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupUpdateMaxNumTxInChunk)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupUpdateMaxNumTxInChunk)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupUpdateMaxNumTxInChunkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupUpdateMaxNumTxInChunkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupUpdateMaxNumTxInChunk represents a UpdateMaxNumTxInChunk event raised by the Rollup contract.
type RollupUpdateMaxNumTxInChunk struct {
	OldMaxNumTxInChunk *big.Int
	NewMaxNumTxInChunk *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdateMaxNumTxInChunk is a free log retrieval operation binding the contract event 0x6d0f49971e462a2f78a25906f145cb29cd5e7bd01ebf681ac8f58cb814e5877a.
//
// Solidity: event UpdateMaxNumTxInChunk(uint256 oldMaxNumTxInChunk, uint256 newMaxNumTxInChunk)
func (_Rollup *RollupFilterer) FilterUpdateMaxNumTxInChunk(opts *bind.FilterOpts) (*RollupUpdateMaxNumTxInChunkIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "UpdateMaxNumTxInChunk")
	if err != nil {
		return nil, err
	}
	return &RollupUpdateMaxNumTxInChunkIterator{contract: _Rollup.contract, event: "UpdateMaxNumTxInChunk", logs: logs, sub: sub}, nil
}

// WatchUpdateMaxNumTxInChunk is a free log subscription operation binding the contract event 0x6d0f49971e462a2f78a25906f145cb29cd5e7bd01ebf681ac8f58cb814e5877a.
//
// Solidity: event UpdateMaxNumTxInChunk(uint256 oldMaxNumTxInChunk, uint256 newMaxNumTxInChunk)
func (_Rollup *RollupFilterer) WatchUpdateMaxNumTxInChunk(opts *bind.WatchOpts, sink chan<- *RollupUpdateMaxNumTxInChunk) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "UpdateMaxNumTxInChunk")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupUpdateMaxNumTxInChunk)
				if err := _Rollup.contract.UnpackLog(event, "UpdateMaxNumTxInChunk", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateMaxNumTxInChunk is a log parse operation binding the contract event 0x6d0f49971e462a2f78a25906f145cb29cd5e7bd01ebf681ac8f58cb814e5877a.
//
// Solidity: event UpdateMaxNumTxInChunk(uint256 oldMaxNumTxInChunk, uint256 newMaxNumTxInChunk)
func (_Rollup *RollupFilterer) ParseUpdateMaxNumTxInChunk(log types.Log) (*RollupUpdateMaxNumTxInChunk, error) {
	event := new(RollupUpdateMaxNumTxInChunk)
	if err := _Rollup.contract.UnpackLog(event, "UpdateMaxNumTxInChunk", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupUpdateProofWindowIterator is returned from FilterUpdateProofWindow and is used to iterate over the raw logs and unpacked data for UpdateProofWindow events raised by the Rollup contract.
type RollupUpdateProofWindowIterator struct {
	Event *RollupUpdateProofWindow // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupUpdateProofWindowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupUpdateProofWindow)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupUpdateProofWindow)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupUpdateProofWindowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupUpdateProofWindowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupUpdateProofWindow represents a UpdateProofWindow event raised by the Rollup contract.
type RollupUpdateProofWindow struct {
	OldWindow *big.Int
	NewWindow *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateProofWindow is a free log retrieval operation binding the contract event 0x1e3a2094feb4b696dd3d7caea38ad2f41dbdcac3fa3943c7a693aff8a64b0a61.
//
// Solidity: event UpdateProofWindow(uint256 oldWindow, uint256 newWindow)
func (_Rollup *RollupFilterer) FilterUpdateProofWindow(opts *bind.FilterOpts) (*RollupUpdateProofWindowIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "UpdateProofWindow")
	if err != nil {
		return nil, err
	}
	return &RollupUpdateProofWindowIterator{contract: _Rollup.contract, event: "UpdateProofWindow", logs: logs, sub: sub}, nil
}

// WatchUpdateProofWindow is a free log subscription operation binding the contract event 0x1e3a2094feb4b696dd3d7caea38ad2f41dbdcac3fa3943c7a693aff8a64b0a61.
//
// Solidity: event UpdateProofWindow(uint256 oldWindow, uint256 newWindow)
func (_Rollup *RollupFilterer) WatchUpdateProofWindow(opts *bind.WatchOpts, sink chan<- *RollupUpdateProofWindow) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "UpdateProofWindow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupUpdateProofWindow)
				if err := _Rollup.contract.UnpackLog(event, "UpdateProofWindow", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateProofWindow is a log parse operation binding the contract event 0x1e3a2094feb4b696dd3d7caea38ad2f41dbdcac3fa3943c7a693aff8a64b0a61.
//
// Solidity: event UpdateProofWindow(uint256 oldWindow, uint256 newWindow)
func (_Rollup *RollupFilterer) ParseUpdateProofWindow(log types.Log) (*RollupUpdateProofWindow, error) {
	event := new(RollupUpdateProofWindow)
	if err := _Rollup.contract.UnpackLog(event, "UpdateProofWindow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupUpdateProverIterator is returned from FilterUpdateProver and is used to iterate over the raw logs and unpacked data for UpdateProver events raised by the Rollup contract.
type RollupUpdateProverIterator struct {
	Event *RollupUpdateProver // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupUpdateProverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupUpdateProver)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupUpdateProver)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupUpdateProverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupUpdateProverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupUpdateProver represents a UpdateProver event raised by the Rollup contract.
type RollupUpdateProver struct {
	Account common.Address
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateProver is a free log retrieval operation binding the contract event 0x967f99d5d403870e4356ff46556df3a6b6ba1f50146639aaedfb9f248eb8661e.
//
// Solidity: event UpdateProver(address indexed account, bool status)
func (_Rollup *RollupFilterer) FilterUpdateProver(opts *bind.FilterOpts, account []common.Address) (*RollupUpdateProverIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "UpdateProver", accountRule)
	if err != nil {
		return nil, err
	}
	return &RollupUpdateProverIterator{contract: _Rollup.contract, event: "UpdateProver", logs: logs, sub: sub}, nil
}

// WatchUpdateProver is a free log subscription operation binding the contract event 0x967f99d5d403870e4356ff46556df3a6b6ba1f50146639aaedfb9f248eb8661e.
//
// Solidity: event UpdateProver(address indexed account, bool status)
func (_Rollup *RollupFilterer) WatchUpdateProver(opts *bind.WatchOpts, sink chan<- *RollupUpdateProver, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "UpdateProver", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupUpdateProver)
				if err := _Rollup.contract.UnpackLog(event, "UpdateProver", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateProver is a log parse operation binding the contract event 0x967f99d5d403870e4356ff46556df3a6b6ba1f50146639aaedfb9f248eb8661e.
//
// Solidity: event UpdateProver(address indexed account, bool status)
func (_Rollup *RollupFilterer) ParseUpdateProver(log types.Log) (*RollupUpdateProver, error) {
	event := new(RollupUpdateProver)
	if err := _Rollup.contract.UnpackLog(event, "UpdateProver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupUpdateVerifierIterator is returned from FilterUpdateVerifier and is used to iterate over the raw logs and unpacked data for UpdateVerifier events raised by the Rollup contract.
type RollupUpdateVerifierIterator struct {
	Event *RollupUpdateVerifier // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupUpdateVerifierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupUpdateVerifier)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupUpdateVerifier)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupUpdateVerifierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupUpdateVerifierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupUpdateVerifier represents a UpdateVerifier event raised by the Rollup contract.
type RollupUpdateVerifier struct {
	OldVerifier common.Address
	NewVerifier common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateVerifier is a free log retrieval operation binding the contract event 0x728af3d16a5760405e27a082c98ab272e9f0a1d02f0085d41532a26093aedd96.
//
// Solidity: event UpdateVerifier(address indexed oldVerifier, address indexed newVerifier)
func (_Rollup *RollupFilterer) FilterUpdateVerifier(opts *bind.FilterOpts, oldVerifier []common.Address, newVerifier []common.Address) (*RollupUpdateVerifierIterator, error) {

	var oldVerifierRule []interface{}
	for _, oldVerifierItem := range oldVerifier {
		oldVerifierRule = append(oldVerifierRule, oldVerifierItem)
	}
	var newVerifierRule []interface{}
	for _, newVerifierItem := range newVerifier {
		newVerifierRule = append(newVerifierRule, newVerifierItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "UpdateVerifier", oldVerifierRule, newVerifierRule)
	if err != nil {
		return nil, err
	}
	return &RollupUpdateVerifierIterator{contract: _Rollup.contract, event: "UpdateVerifier", logs: logs, sub: sub}, nil
}

// WatchUpdateVerifier is a free log subscription operation binding the contract event 0x728af3d16a5760405e27a082c98ab272e9f0a1d02f0085d41532a26093aedd96.
//
// Solidity: event UpdateVerifier(address indexed oldVerifier, address indexed newVerifier)
func (_Rollup *RollupFilterer) WatchUpdateVerifier(opts *bind.WatchOpts, sink chan<- *RollupUpdateVerifier, oldVerifier []common.Address, newVerifier []common.Address) (event.Subscription, error) {

	var oldVerifierRule []interface{}
	for _, oldVerifierItem := range oldVerifier {
		oldVerifierRule = append(oldVerifierRule, oldVerifierItem)
	}
	var newVerifierRule []interface{}
	for _, newVerifierItem := range newVerifier {
		newVerifierRule = append(newVerifierRule, newVerifierItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "UpdateVerifier", oldVerifierRule, newVerifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupUpdateVerifier)
				if err := _Rollup.contract.UnpackLog(event, "UpdateVerifier", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateVerifier is a log parse operation binding the contract event 0x728af3d16a5760405e27a082c98ab272e9f0a1d02f0085d41532a26093aedd96.
//
// Solidity: event UpdateVerifier(address indexed oldVerifier, address indexed newVerifier)
func (_Rollup *RollupFilterer) ParseUpdateVerifier(log types.Log) (*RollupUpdateVerifier, error) {
	event := new(RollupUpdateVerifier)
	if err := _Rollup.contract.UnpackLog(event, "UpdateVerifier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
