// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"morph-l2/bindings/solc"
)

const L2TxFeeVaultStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/l2/system/L2TxFeeVault.sol:L2TxFeeVault\",\"label\":\"owner\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"contracts/l2/system/L2TxFeeVault.sol:L2TxFeeVault\",\"label\":\"minWithdrawAmount\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"},{\"astId\":1002,\"contract\":\"contracts/l2/system/L2TxFeeVault.sol:L2TxFeeVault\",\"label\":\"messenger\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_address\"},{\"astId\":1003,\"contract\":\"contracts/l2/system/L2TxFeeVault.sol:L2TxFeeVault\",\"label\":\"recipient\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"contracts/l2/system/L2TxFeeVault.sol:L2TxFeeVault\",\"label\":\"totalProcessed\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_uint256\"},{\"astId\":1005,\"contract\":\"contracts/l2/system/L2TxFeeVault.sol:L2TxFeeVault\",\"label\":\"transferAllowed\",\"offset\":0,\"slot\":\"5\",\"type\":\"t_mapping(t_address,t_bool)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_mapping(t_address,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_bool\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var L2TxFeeVaultStorageLayout = new(solc.StorageLayout)

var L2TxFeeVaultDeployedBin = "0x6080604052600436106100f2575f3560e01c806384411d6511610087578063da13f9a211610057578063da13f9a2146102c6578063f2fde38b146102e5578063feec756c14610304578063ff4f354614610323575f80fd5b806384411d65146102485780638da5cb5b1461025d5780639e7adc7914610288578063a03fa7e3146102a7575f80fd5b80633ccfd60b116100c25780633ccfd60b146101d1578063457e1a49146101e557806366d003ac14610208578063715018a614610234575f80fd5b8063151eeb55146100fd5780632ccb1b30146101405780632e1a7d4d146101615780633cb747bf14610180575f80fd5b366100f957005b5f80fd5b348015610108575f80fd5b5061012b610117366004611139565b60056020525f908152604090205460ff1681565b60405190151581526020015b60405180910390f35b34801561014b575f80fd5b5061015f61015a366004611159565b610342565b005b34801561016c575f80fd5b5061015f61017b366004611181565b61063f565b34801561018b575f80fd5b506002546101ac9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610137565b3480156101dc575f80fd5b5061015f6109ac565b3480156101f0575f80fd5b506101fa60015481565b604051908152602001610137565b348015610213575f80fd5b506003546101ac9073ffffffffffffffffffffffffffffffffffffffff1681565b34801561023f575f80fd5b5061015f610a39565b348015610253575f80fd5b506101fa60045481565b348015610268575f80fd5b505f546101ac9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610293575f80fd5b5061015f6102a2366004611139565b610ac4565b3480156102b2575f80fd5b5061015f6102c1366004611139565b610bba565b3480156102d1575f80fd5b5061015f6102e03660046111d4565b610c62565b3480156102f0575f80fd5b5061015f6102ff366004611139565b610ddc565b34801561030f575f80fd5b5061015f61031e366004611139565b610ee2565b34801561032e575f80fd5b5061015f61033d366004611181565b610fd8565b335f9081526005602052604090205460ff168061037557505f5473ffffffffffffffffffffffffffffffffffffffff1633145b6103e0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f63616c6c6572206973206e6f7420616c6c6f776564000000000000000000000060448201526064015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216610483576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f4665655661756c743a20726563697069656e7420616464726573732063616e6e60448201527f6f7420626520616464726573732830290000000000000000000000000000000060648201526084016103d7565b4780821115610514576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f4665655661756c743a20696e73756666696369656e742062616c616e6365207460448201527f6f207472616e736665720000000000000000000000000000000000000000000060648201526084016103d7565b60048054830190556003546040805184815273ffffffffffffffffffffffffffffffffffffffff90921660208301523382820152517f0a429aba3d89849a2db0153e4534d95c46a1d83c8109d73893f55ebc44010ff49181900360600190a15f8373ffffffffffffffffffffffffffffffffffffffff16836040515f6040518083038185875af1925050503d805f81146105c9576040519150601f19603f3d011682016040523d82523d5f602084013e6105ce565b606091505b5050905080610639576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f4665655661756c743a20455448207472616e73666572206661696c656400000060448201526064016103d7565b50505050565b5f5473ffffffffffffffffffffffffffffffffffffffff1633146106bf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e657200000000000000000060448201526064016103d7565b60035473ffffffffffffffffffffffffffffffffffffffff16610764576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f4665655661756c743a20726563697069656e7420616464726573732063616e6e60448201527f6f7420626520616464726573732830290000000000000000000000000000000060648201526084016103d7565b60015481101561081c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f4665655661756c743a207769746864726177616c20616d6f756e74206d75737460448201527f2062652067726561746572207468616e206d696e696d756d207769746864726160648201527f77616c20616d6f756e7400000000000000000000000000000000000000000000608482015260a4016103d7565b47808211156108ad576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f4665655661756c743a20696e73756666696369656e742062616c616e6365207460448201527f6f2077697468647261770000000000000000000000000000000000000000000060648201526084016103d7565b60048054830190556003546040805184815273ffffffffffffffffffffffffffffffffffffffff90921660208301523382820152517fc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba9181900360600190a1600254600354604080516020810182525f80825291517fb2267a7b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9485169463b2267a7b94889461097a9491909216928592906004016112c4565b5f604051808303818588803b158015610991575f80fd5b505af11580156109a3573d5f803e3d5ffd5b50505050505050565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610a2c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e657200000000000000000060448201526064016103d7565b47610a368161063f565b50565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610ab9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e657200000000000000000060448201526064016103d7565b610ac25f61109d565b565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610b44576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e657200000000000000000060448201526064016103d7565b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f1c928c417a10a21c3cddad148c5dba5d710e4b1442d6d8a36de345935ad84612905f90a35050565b335f9081526005602052604090205460ff1680610bed57505f5473ffffffffffffffffffffffffffffffffffffffff1633145b610c53576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f63616c6c6572206973206e6f7420616c6c6f776564000000000000000000000060448201526064016103d7565b47610c5e8282610342565b5050565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610ce2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e657200000000000000000060448201526064016103d7565b5f5b8251811015610dd7578160055f858481518110610d0357610d0361135a565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055507fbb5d3e442e800faa1089a4f57bae4f36808d3cf15d051033d78a72147782f24c838281518110610d8d57610d8d61135a565b602002602001015183604051610dc792919073ffffffffffffffffffffffffffffffffffffffff9290921682521515602082015260400190565b60405180910390a1600101610ce4565b505050565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610e5c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e657200000000000000000060448201526064016103d7565b73ffffffffffffffffffffffffffffffffffffffff8116610ed9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f6e6577206f776e657220697320746865207a65726f206164647265737300000060448201526064016103d7565b610a368161109d565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610f62576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e657200000000000000000060448201526064016103d7565b6003805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f7e1e96961a397c8aa26162fe259cc837afc95e33aad4945ddc61c18dabb7a6ad905f90a35050565b5f5473ffffffffffffffffffffffffffffffffffffffff163314611058576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e657200000000000000000060448201526064016103d7565b600180549082905560408051828152602081018490527f0d3c80219fe57713b9f9c83d1e51426792d0c14d8e330e65b102571816140965910160405180910390a15050565b5f805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611134575f80fd5b919050565b5f60208284031215611149575f80fd5b61115282611111565b9392505050565b5f806040838503121561116a575f80fd5b61117383611111565b946020939093013593505050565b5f60208284031215611191575f80fd5b5035919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b80358015158114611134575f80fd5b5f80604083850312156111e5575f80fd5b823567ffffffffffffffff808211156111fc575f80fd5b818501915085601f83011261120f575f80fd5b813560208282111561122357611223611198565b8160051b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f8301168101818110868211171561126657611266611198565b604052928352818301935084810182019289841115611283575f80fd5b948201945b838610156112a85761129986611111565b85529482019493820193611288565b96506112b790508782016111c5565b9450505050509250929050565b73ffffffffffffffffffffffffffffffffffffffff851681525f60208560208401526080604084015284518060808501525f5b818110156113135786810183015185820160a0015282016112f7565b505f60a0828601015260a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505082606083015295945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea164736f6c6343000818000a"

func init() {
	if err := json.Unmarshal([]byte(L2TxFeeVaultStorageLayoutJSON), L2TxFeeVaultStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2TxFeeVault"] = L2TxFeeVaultStorageLayout
	deployedBytecodes["L2TxFeeVault"] = L2TxFeeVaultDeployedBin
}
