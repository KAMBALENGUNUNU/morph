// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const ProxyAdminStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/universal/ProxyAdmin.sol:ProxyAdmin\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"contracts/universal/ProxyAdmin.sol:ProxyAdmin\",\"label\":\"proxyType\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_address,t_enum(ProxyType)1006)\"},{\"astId\":1002,\"contract\":\"contracts/universal/ProxyAdmin.sol:ProxyAdmin\",\"label\":\"implementationName\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_address,t_string_storage)\"},{\"astId\":1003,\"contract\":\"contracts/universal/ProxyAdmin.sol:ProxyAdmin\",\"label\":\"addressManager\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_contract(AddressManager)1005\"},{\"astId\":1004,\"contract\":\"contracts/universal/ProxyAdmin.sol:ProxyAdmin\",\"label\":\"upgrading\",\"offset\":20,\"slot\":\"3\",\"type\":\"t_bool\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_contract(AddressManager)1005\":{\"encoding\":\"inplace\",\"label\":\"contract AddressManager\",\"numberOfBytes\":\"20\"},\"t_enum(ProxyType)1006\":{\"encoding\":\"inplace\",\"label\":\"enum ProxyAdmin.ProxyType\",\"numberOfBytes\":\"1\"},\"t_mapping(t_address,t_enum(ProxyType)1006)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e enum ProxyAdmin.ProxyType)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_enum(ProxyType)1006\"},\"t_mapping(t_address,t_string_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e string)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_string_storage\"},\"t_string_storage\":{\"encoding\":\"bytes\",\"label\":\"string\",\"numberOfBytes\":\"32\"}}}"

var ProxyAdminStorageLayout = new(solc.StorageLayout)

var ProxyAdminDeployedBin = "0x60806040526004361061010e5760003560e01c8063860f7cda116100a557806399a88ec411610074578063b794726211610059578063b794726214610329578063f2fde38b14610364578063f3b7dead1461038457600080fd5b806399a88ec4146102e95780639b2ea4bd1461030957600080fd5b8063860f7cda1461026b5780638d52d4a01461028b5780638da5cb5b146102ab5780639623609d146102d657600080fd5b80633ab76e9f116100e15780633ab76e9f146101cc5780636bd9f516146101f9578063715018a6146102365780637eff275e1461024b57600080fd5b80630652b57a1461011357806307c8f7b014610135578063204e1c7a14610155578063238181ae1461019f575b600080fd5b34801561011f57600080fd5b5061013361012e3660046111f9565b6103a4565b005b34801561014157600080fd5b50610133610150366004611216565b6103f3565b34801561016157600080fd5b506101756101703660046111f9565b610445565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156101ab57600080fd5b506101bf6101ba3660046111f9565b61066b565b60405161019691906112a6565b3480156101d857600080fd5b506003546101759073ffffffffffffffffffffffffffffffffffffffff1681565b34801561020557600080fd5b506102296102143660046111f9565b60016020526000908152604090205460ff1681565b60405161019691906112e8565b34801561024257600080fd5b50610133610705565b34801561025757600080fd5b50610133610266366004611329565b610719565b34801561027757600080fd5b50610133610286366004611484565b6108cc565b34801561029757600080fd5b506101336102a63660046114d4565b610903565b3480156102b757600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff16610175565b6101336102e4366004611506565b610977565b3480156102f557600080fd5b50610133610304366004611329565b610b8e565b34801561031557600080fd5b5061013361032436600461157c565b610e1e565b34801561033557600080fd5b5060035474010000000000000000000000000000000000000000900460ff166040519015158152602001610196565b34801561037057600080fd5b5061013361037f3660046111f9565b610eb4565b34801561039057600080fd5b5061017561039f3660046111f9565b610f6b565b6103ac6110e1565b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6103fb6110e1565b6003805491151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b73ffffffffffffffffffffffffffffffffffffffff811660009081526001602052604081205460ff1681816002811115610481576104816112b9565b036104fc578273ffffffffffffffffffffffffffffffffffffffff16635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104f591906115c3565b9392505050565b6001816002811115610510576105106112b9565b03610560578273ffffffffffffffffffffffffffffffffffffffff1663aaf10f426040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104d1573d6000803e3d6000fd5b6002816002811115610574576105746112b9565b036105fe5760035473ffffffffffffffffffffffffffffffffffffffff8481166000908152600260205260409081902090517fbf40fac1000000000000000000000000000000000000000000000000000000008152919092169163bf40fac1916105e1919060040161162d565b602060405180830381865afa1580156104d1573d6000803e3d6000fd5b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f50726f787941646d696e3a20756e6b6e6f776e2070726f78792074797065000060448201526064015b60405180910390fd5b50919050565b60026020526000908152604090208054610684906115e0565b80601f01602080910402602001604051908101604052809291908181526020018280546106b0906115e0565b80156106fd5780601f106106d2576101008083540402835291602001916106fd565b820191906000526020600020905b8154815290600101906020018083116106e057829003601f168201915b505050505081565b61070d6110e1565b6107176000611162565b565b6107216110e1565b73ffffffffffffffffffffffffffffffffffffffff821660009081526001602052604081205460ff169081600281111561075d5761075d6112b9565b036107e9576040517f8f28397000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8381166004830152841690638f283970906024015b600060405180830381600087803b1580156107cc57600080fd5b505af11580156107e0573d6000803e3d6000fd5b50505050505050565b60018160028111156107fd576107fd6112b9565b03610856576040517f13af403500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301528416906313af4035906024016107b2565b600281600281111561086a5761086a6112b9565b036105fe576003546040517ff2fde38b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301529091169063f2fde38b906024016107b2565b505050565b6108d46110e1565b73ffffffffffffffffffffffffffffffffffffffff821660009081526002602052604090206108c7828261171c565b61090b6110e1565b73ffffffffffffffffffffffffffffffffffffffff82166000908152600160208190526040909120805483927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff009091169083600281111561096e5761096e6112b9565b02179055505050565b61097f6110e1565b73ffffffffffffffffffffffffffffffffffffffff831660009081526001602052604081205460ff16908160028111156109bb576109bb6112b9565b03610a81576040517f4f1ef28600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851690634f1ef286903490610a169087908790600401611836565b60006040518083038185885af1158015610a34573d6000803e3d6000fd5b50505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610a7b919081019061186d565b50610b88565b610a8b8484610b8e565b60008473ffffffffffffffffffffffffffffffffffffffff163484604051610ab391906118e4565b60006040518083038185875af1925050503d8060008114610af0576040519150601f19603f3d011682016040523d82523d6000602084013e610af5565b606091505b5050905080610b86576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f50726f787941646d696e3a2063616c6c20746f2070726f78792061667465722060448201527f75706772616465206661696c6564000000000000000000000000000000000000606482015260840161065c565b505b50505050565b610b966110e1565b73ffffffffffffffffffffffffffffffffffffffff821660009081526001602052604081205460ff1690816002811115610bd257610bd26112b9565b03610c2b576040517f3659cfe600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8381166004830152841690633659cfe6906024016107b2565b6001816002811115610c3f57610c3f6112b9565b03610cbe576040517f9b0b0fda0000000000000000000000000000000000000000000000000000000081527f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc600482015273ffffffffffffffffffffffffffffffffffffffff8381166024830152841690639b0b0fda906044016107b2565b6002816002811115610cd257610cd26112b9565b03610e165773ffffffffffffffffffffffffffffffffffffffff831660009081526002602052604081208054610d07906115e0565b80601f0160208091040260200160405190810160405280929190818152602001828054610d33906115e0565b8015610d805780601f10610d5557610100808354040283529160200191610d80565b820191906000526020600020905b815481529060010190602001808311610d6357829003601f168201915b50506003546040517f9b2ea4bd00000000000000000000000000000000000000000000000000000000815294955073ffffffffffffffffffffffffffffffffffffffff1693639b2ea4bd9350610dde92508591508790600401611900565b600060405180830381600087803b158015610df857600080fd5b505af1158015610e0c573d6000803e3d6000fd5b5050505050505050565b6108c7611938565b610e266110e1565b6003546040517f9b2ea4bd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690639b2ea4bd90610e7e9085908590600401611900565b600060405180830381600087803b158015610e9857600080fd5b505af1158015610eac573d6000803e3d6000fd5b505050505050565b610ebc6110e1565b73ffffffffffffffffffffffffffffffffffffffff8116610f5f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161065c565b610f6881611162565b50565b73ffffffffffffffffffffffffffffffffffffffff811660009081526001602052604081205460ff1681816002811115610fa757610fa76112b9565b03610ff7578273ffffffffffffffffffffffffffffffffffffffff1663f851a4406040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104d1573d6000803e3d6000fd5b600181600281111561100b5761100b6112b9565b0361105b578273ffffffffffffffffffffffffffffffffffffffff1663893d20e86040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104d1573d6000803e3d6000fd5b600281600281111561106f5761106f6112b9565b036105fe57600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104d1573d6000803e3d6000fd5b60005473ffffffffffffffffffffffffffffffffffffffff163314610717576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161065c565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b73ffffffffffffffffffffffffffffffffffffffff81168114610f6857600080fd5b60006020828403121561120b57600080fd5b81356104f5816111d7565b60006020828403121561122857600080fd5b813580151581146104f557600080fd5b60005b8381101561125357818101518382015260200161123b565b50506000910152565b60008151808452611274816020860160208601611238565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006104f5602083018461125c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6020810160038310611323577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b6000806040838503121561133c57600080fd5b8235611347816111d7565b91506020830135611357816111d7565b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156113d8576113d8611362565b604052919050565b600067ffffffffffffffff8211156113fa576113fa611362565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b6000611439611434846113e0565b611391565b905082815283838301111561144d57600080fd5b828260208301376000602084830101529392505050565b600082601f83011261147557600080fd5b6104f583833560208501611426565b6000806040838503121561149757600080fd5b82356114a2816111d7565b9150602083013567ffffffffffffffff8111156114be57600080fd5b6114ca85828601611464565b9150509250929050565b600080604083850312156114e757600080fd5b82356114f2816111d7565b915060208301356003811061135757600080fd5b60008060006060848603121561151b57600080fd5b8335611526816111d7565b92506020840135611536816111d7565b9150604084013567ffffffffffffffff81111561155257600080fd5b8401601f8101861361156357600080fd5b61157286823560208401611426565b9150509250925092565b6000806040838503121561158f57600080fd5b823567ffffffffffffffff8111156115a657600080fd5b6115b285828601611464565b9250506020830135611357816111d7565b6000602082840312156115d557600080fd5b81516104f5816111d7565b600181811c908216806115f457607f821691505b602082108103610665577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000602080835260008454611641816115e0565b80848701526040600180841660008114611662576001811461169a576116c8565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008516838a01528284151560051b8a010195506116c8565b896000528660002060005b858110156116c05781548b82018601529083019088016116a5565b8a0184019650505b509398975050505050505050565b601f8211156108c757600081815260208120601f850160051c810160208610156116fd5750805b601f850160051c820191505b81811015610eac57828155600101611709565b815167ffffffffffffffff81111561173657611736611362565b61174a8161174484546115e0565b846116d6565b602080601f83116001811461179d57600084156117675750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555610eac565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b828110156117ea578886015182559484019460019091019084016117cb565b508582101561182657878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201526000611865604083018461125c565b949350505050565b60006020828403121561187f57600080fd5b815167ffffffffffffffff81111561189657600080fd5b8201601f810184136118a757600080fd5b80516118b5611434826113e0565b8181528560208385010111156118ca57600080fd5b6118db826020830160208601611238565b95945050505050565b600082516118f6818460208701611238565b9190910192915050565b604081526000611913604083018561125c565b905073ffffffffffffffffffffffffffffffffffffffff831660208301529392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fdfea164736f6c6343000810000a"

func init() {
	if err := json.Unmarshal([]byte(ProxyAdminStorageLayoutJSON), ProxyAdminStorageLayout); err != nil {
		panic(err)
	}

	layouts["ProxyAdmin"] = ProxyAdminStorageLayout
	deployedBytecodes["ProxyAdmin"] = ProxyAdminDeployedBin
}