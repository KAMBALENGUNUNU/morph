// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const MintManagerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/governance/MintManager.sol:MintManager\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"contracts/governance/MintManager.sol:MintManager\",\"label\":\"mintPermittedAfter\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var MintManagerStorageLayout = new(solc.StorageLayout)

var MintManagerDeployedBin = "0x608060405234801561000f575f80fd5b50600436106100b8575f3560e01c80638da5cb5b1161007257806398f1312e1161005857806398f1312e1461015a578063f2fde38b14610162578063f96dae0a14610175575f80fd5b80638da5cb5b14610113578063918f867414610151575f80fd5b806340c10f19116100a257806340c10f19146100ed578063715018a61461010057806383ea6e9714610108575f80fd5b8062f8900c146100bc5780630900f010146100d8575b5f80fd5b6100c560015481565b6040519081526020015b60405180910390f35b6100eb6100e636600461075f565b61019c565b005b6100eb6100fb36600461077f565b6102eb565b6100eb610579565b6100c56301e1338081565b5f5473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100cf565b6100c56103e881565b6100c5601481565b6100eb61017036600461075f565b61058c565b61012c7f000000000000000000000000000000000000000000000000000000000000000081565b6101a4610643565b73ffffffffffffffffffffffffffffffffffffffff811661024c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603460248201527f4d696e744d616e616765723a206d696e74206d616e616765722063616e6e6f7460448201527f20626520746865207a65726f206164647265737300000000000000000000000060648201526084015b60405180910390fd5b6040517ff2fde38b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063f2fde38b906024015f604051808303815f87803b1580156102d2575f80fd5b505af11580156102e4573d5f803e3d5ffd5b5050505050565b6102f3610643565b600154156104c15742600154111561038d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4d696e744d616e616765723a206d696e74696e67206e6f74207065726d69747460448201527f65642079657400000000000000000000000000000000000000000000000000006064820152608401610243565b6103e860147f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103fb573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061041f91906107a7565b61042991906107eb565b6104339190610808565b8111156104c1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f4d696e744d616e616765723a206d696e7420616d6f756e74206578636565647360448201527f20636170000000000000000000000000000000000000000000000000000000006064820152608401610243565b6104cf6301e1338042610840565b6001556040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8381166004830152602482018390527f000000000000000000000000000000000000000000000000000000000000000016906340c10f19906044015f604051808303815f87803b15801561055f575f80fd5b505af1158015610571573d5f803e3d5ffd5b505050505050565b610581610643565b61058a5f6106c3565b565b610594610643565b73ffffffffffffffffffffffffffffffffffffffff8116610637576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610243565b610640816106c3565b50565b5f5473ffffffffffffffffffffffffffffffffffffffff16331461058a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610243565b5f805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461075a575f80fd5b919050565b5f6020828403121561076f575f80fd5b61077882610737565b9392505050565b5f8060408385031215610790575f80fd5b61079983610737565b946020939093013593505050565b5f602082840312156107b7575f80fd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8082028115828204841417610802576108026107be565b92915050565b5f8261083b577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b80820180821115610802576108026107be56fea164736f6c6343000818000a"

func init() {
	if err := json.Unmarshal([]byte(MintManagerStorageLayoutJSON), MintManagerStorageLayout); err != nil {
		panic(err)
	}

	layouts["MintManager"] = MintManagerStorageLayout
	deployedBytecodes["MintManager"] = MintManagerDeployedBin
}
