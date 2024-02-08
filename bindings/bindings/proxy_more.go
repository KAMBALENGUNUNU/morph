// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const ProxyStorageLayoutJSON = "{\"storage\":null,\"types\":{}}"

var ProxyStorageLayout = new(solc.StorageLayout)

var ProxyDeployedBin = "0x60806040526004361061005d575f3560e01c80635c60da1b116100425780635c60da1b146100bc5780638f283970146100f5578063f851a440146101145761006c565b80633659cfe6146100745780634f1ef286146100935761006c565b3661006c5761006a610128565b005b61006a610128565b34801561007f575f80fd5b5061006a61008e3660046106c4565b610218565b6100a66100a13660046106dd565b61028a565b6040516100b39190610759565b60405180910390f35b3480156100c7575f80fd5b506100d0610409565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100b3565b348015610100575f80fd5b5061006a61010f3660046106c4565b61049f565b34801561011f575f80fd5b506100d0610506565b5f6101517f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b905073ffffffffffffffffffffffffffffffffffffffff81166101fb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f50726f78793a20696d706c656d656e746174696f6e206e6f7420696e6974696160448201527f6c697a656400000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b365f80375f80365f845af43d5f803e80610213573d5ffd5b503d5ff35b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480610271575033155b156102825761027f81610591565b50565b61027f610128565b60606102b47fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035490565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806102eb575033155b156103fa576102f984610591565b5f808573ffffffffffffffffffffffffffffffffffffffff1685856040516103229291906107c3565b5f60405180830381855af49150503d805f811461035a576040519150601f19603f3d011682016040523d82523d5f602084013e61035f565b606091505b5091509150816103f1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603960248201527f50726f78793a2064656c656761746563616c6c20746f206e657720696d706c6560448201527f6d656e746174696f6e20636f6e7472616374206661696c65640000000000000060648201526084016101f2565b91506104029050565b610402610128565b9392505050565b5f6104327fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035490565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480610469575033155b1561049457507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b61049c610128565b90565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806104f8575033155b156102825761027f816105f8565b5f61052f7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035490565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480610566575033155b1561049457507fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035490565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc81905560405173ffffffffffffffffffffffffffffffffffffffff8216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b5f6106217fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035490565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61038390556040805173ffffffffffffffffffffffffffffffffffffffff8084168252851660208201529192507f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f910160405180910390a15050565b803573ffffffffffffffffffffffffffffffffffffffff811681146106bf575f80fd5b919050565b5f602082840312156106d4575f80fd5b6104028261069c565b5f805f604084860312156106ef575f80fd5b6106f88461069c565b9250602084013567ffffffffffffffff80821115610714575f80fd5b818601915086601f830112610727575f80fd5b813581811115610735575f80fd5b876020828501011115610746575f80fd5b6020830194508093505050509250925092565b5f602080835283518060208501525f5b8181101561078557858101830151858201604001528201610769565b505f6040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b818382375f910190815291905056fea164736f6c6343000818000a"

func init() {
	if err := json.Unmarshal([]byte(ProxyStorageLayoutJSON), ProxyStorageLayout); err != nil {
		panic(err)
	}

	layouts["Proxy"] = ProxyStorageLayout
	deployedBytecodes["Proxy"] = ProxyDeployedBin
}
