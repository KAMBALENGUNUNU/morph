// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"morph-l2/bindings/solc"
)

const WhitelistStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/libraries/common/Whitelist.sol:Whitelist\",\"label\":\"owner\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"contracts/libraries/common/Whitelist.sol:Whitelist\",\"label\":\"isWhitelisted\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_address,t_bool)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_mapping(t_address,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_bool\"}}}"

var WhitelistStorageLayout = new(solc.StorageLayout)

var WhitelistDeployedBin = "0x608060405234801561000f575f80fd5b5060043610610064575f3560e01c80638da5cb5b1161004d5780638da5cb5b14610085578063efc78401146100ce578063f2fde38b14610116575f80fd5b8063715018a61461006857806379586dd714610072575b5f80fd5b610070610129565b005b61007061008036600461050a565b6101b9565b5f546100a49073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6101066100dc3660046105fa565b73ffffffffffffffffffffffffffffffffffffffff165f9081526001602052604090205460ff1690565b60405190151581526020016100c5565b6100706101243660046105fa565b610329565b5f5473ffffffffffffffffffffffffffffffffffffffff1633146101ae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e657200000000000000000060448201526064015b60405180910390fd5b6101b75f610432565b565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610239576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e657200000000000000000060448201526064016101a5565b5f5b8251811015610324578160015f85848151811061025a5761025a61061a565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508281815181106102c3576102c361061a565b602002602001015173ffffffffffffffffffffffffffffffffffffffff167f8daaf060c3306c38e068a75c054bf96ecd85a3db1252712c4d93632744c42e0d83604051610314911515815260200190565b60405180910390a260010161023b565b505050565b5f5473ffffffffffffffffffffffffffffffffffffffff1633146103a9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e657200000000000000000060448201526064016101a5565b73ffffffffffffffffffffffffffffffffffffffff8116610426576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f6e6577206f776e657220697320746865207a65726f206164647265737300000060448201526064016101a5565b61042f81610432565b50565b5f805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b803573ffffffffffffffffffffffffffffffffffffffff811681146104f6575f80fd5b919050565b803580151581146104f6575f80fd5b5f806040838503121561051b575f80fd5b823567ffffffffffffffff80821115610532575f80fd5b818501915085601f830112610545575f80fd5b8135602082821115610559576105596104a6565b8160051b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f8301168101818110868211171561059c5761059c6104a6565b6040529283528183019350848101820192898411156105b9575f80fd5b948201945b838610156105de576105cf866104d3565b855294820194938201936105be565b96506105ed90508782016104fb565b9450505050509250929050565b5f6020828403121561060a575f80fd5b610613826104d3565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea164736f6c6343000818000a"

func init() {
	if err := json.Unmarshal([]byte(WhitelistStorageLayoutJSON), WhitelistStorageLayout); err != nil {
		panic(err)
	}

	layouts["Whitelist"] = WhitelistStorageLayout
	deployedBytecodes["Whitelist"] = WhitelistDeployedBin
}