// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const L1MessageQueueWithGasPriceOracleStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L1/rollup/L1MessageQueueWithGasPriceOracle.sol:L1MessageQueueWithGasPriceOracle\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L1/rollup/L1MessageQueueWithGasPriceOracle.sol:L1MessageQueueWithGasPriceOracle\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L1/rollup/L1MessageQueueWithGasPriceOracle.sol:L1MessageQueueWithGasPriceOracle\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)1013_storage\"},{\"astId\":1003,\"contract\":\"contracts/L1/rollup/L1MessageQueueWithGasPriceOracle.sol:L1MessageQueueWithGasPriceOracle\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"contracts/L1/rollup/L1MessageQueueWithGasPriceOracle.sol:L1MessageQueueWithGasPriceOracle\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)1012_storage\"},{\"astId\":1005,\"contract\":\"contracts/L1/rollup/L1MessageQueueWithGasPriceOracle.sol:L1MessageQueueWithGasPriceOracle\",\"label\":\"l2BaseFee\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_uint256\"},{\"astId\":1006,\"contract\":\"contracts/L1/rollup/L1MessageQueueWithGasPriceOracle.sol:L1MessageQueueWithGasPriceOracle\",\"label\":\"messageQueue\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_bytes32)dyn_storage\"},{\"astId\":1007,\"contract\":\"contracts/L1/rollup/L1MessageQueueWithGasPriceOracle.sol:L1MessageQueueWithGasPriceOracle\",\"label\":\"pendingQueueIndex\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_uint256\"},{\"astId\":1008,\"contract\":\"contracts/L1/rollup/L1MessageQueueWithGasPriceOracle.sol:L1MessageQueueWithGasPriceOracle\",\"label\":\"maxGasLimit\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_uint256\"},{\"astId\":1009,\"contract\":\"contracts/L1/rollup/L1MessageQueueWithGasPriceOracle.sol:L1MessageQueueWithGasPriceOracle\",\"label\":\"droppedMessageBitmap\",\"offset\":0,\"slot\":\"105\",\"type\":\"t_struct(BitMap)1014_storage\"},{\"astId\":1010,\"contract\":\"contracts/L1/rollup/L1MessageQueueWithGasPriceOracle.sol:L1MessageQueueWithGasPriceOracle\",\"label\":\"skippedMessageBitmap\",\"offset\":0,\"slot\":\"106\",\"type\":\"t_mapping(t_uint256,t_uint256)\"},{\"astId\":1011,\"contract\":\"contracts/L1/rollup/L1MessageQueueWithGasPriceOracle.sol:L1MessageQueueWithGasPriceOracle\",\"label\":\"whitelistChecker\",\"offset\":0,\"slot\":\"107\",\"type\":\"t_address\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_bytes32)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"bytes32[]\",\"numberOfBytes\":\"32\"},\"t_array(t_uint256)1012_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1013_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_uint256,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_uint256\"},\"t_struct(BitMap)1014_storage\":{\"encoding\":\"inplace\",\"label\":\"struct BitMapsUpgradeable.BitMap\",\"numberOfBytes\":\"32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1MessageQueueWithGasPriceOracleStorageLayout = new(solc.StorageLayout)

var L1MessageQueueWithGasPriceOracleDeployedBin = "0x608060405234801561000f575f80fd5b50600436106101a5575f3560e01c80639b159782116100e8578063d5ad4a9711610093578063e172d3a11161006e578063e172d3a1146103ca578063e3176bd5146103e4578063f2fde38b146103ed578063fd0ad31e14610400575f80fd5b8063d5ad4a9714610391578063d99bc80e146103a4578063da35a26f146103b7575f80fd5b8063bb7862ca116100c3578063bb7862ca14610337578063bdc6f0a014610357578063cb23bcb51461036a575f80fd5b80639b15978214610308578063a85006ca1461031b578063ae453cd514610324575f80fd5b80635ad9945a11610153578063715018a61161012e578063715018a6146102bc5780637d82191a146102c45780638da5cb5b146102d757806391652461146102f5575f80fd5b80635ad9945a1461028d5780635e45da23146102a05780635f9cd92e146102a9575f80fd5b80633e6dada1116101835780633e6dada11461022e5780633e83496c1461025157806355f613ce14610278575f80fd5b806329aa604b146101a95780633cb747bf146101cf5780633e4cbbe61461021b575b5f80fd5b6101bc6101b73660046116d0565b610408565b6040519081526020015b60405180910390f35b6101f67f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101c6565b6101bc61022936600461170a565b610427565b61024161023c3660046116d0565b6104db565b60405190151581526020016101c6565b6101f67f000000000000000000000000000000000000000000000000000000000000000081565b61028b610286366004611732565b610521565b005b6101bc61029b3660046117a0565b610763565b6101bc60685481565b61028b6102b736600461181c565b610953565b61028b6109e8565b6102416102d23660046116d0565b6109fb565b60335473ffffffffffffffffffffffffffffffffffffffff166101f6565b61028b6103033660046116d0565b610a2e565b61028b610316366004611835565b610cb5565b6101bc60675481565b6101bc6103323660046116d0565b610db1565b606b546101f69073ffffffffffffffffffffffffffffffffffffffff1681565b61028b61036536600461188b565b610e41565b6101f67f000000000000000000000000000000000000000000000000000000000000000081565b61028b61039f3660046116d0565b610fa6565b61028b6103b23660046116d0565b610ff4565b61028b6103c53660046118fe565b61103a565b6101bc6103d8366004611928565b60100261520801919050565b6101bc60655481565b61028b6103fb36600461181c565b611209565b6066546101bc565b60668181548110610417575f80fd5b5f91825260209091200154905081565b606b546040517fefc7840100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301525f92169063efc7840190602401602060405180830381865afa158015610495573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104b99190611967565b156104c557505f6104d5565b6065546104d2908361198d565b90505b92915050565b600881901c5f908152606a6020526040812054600160ff84161b16151580156104d55750600882901c5f90815260696020526040902054600160ff84161b1615156104d5565b337f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16146105c5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f4f6e6c792063616c6c61626c652062792074686520726f6c6c7570000000000060448201526064015b60405180910390fd5b610100821115610631576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f706f7020746f6f206d616e79206d65737361676573000000000000000000000060448201526064016105bc565b826067541461069c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f737461727420696e646578206d69736d6174636800000000000000000000000060448201526064016105bc565b600883901c5f818152606a6020526040902080546001851b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0193841660ff871681811b90921790925590929190610100818601111561071357600182015f908152606a6020526040902061010082900385901c90555b50505081830160675560408051848152602081018490529081018290527fc77f792f838ae38399ac31acc3348389aeb110ce7bedf3cfdbdd5e6679267970906060015b60405180910390a1505050565b5f607e8161080d565b5f8161077a57506001919050565b5b81156107905760089190911c9060010161077b565b919050565b8060808310600181146107cd576107ab8461076c565b60808101835360018301925084816020036008021b83528083019250506107ee565b84841516600181146107e1578483536107e6565b608083535b506001820191505b509392505050565b806094815360609290921b60018301525060150190565b6005604051018061082060018c83610795565b905061082e60018983610795565b905061083a89826107f6565b905061084860018b83610795565b905060018614600181146108b05760388710600181146108955761086b8861076c565b8060b701845360018401935088816020036008021b845280840193505087898437918701916108aa565b87608001835360018301925087898437918701915b506108c1565b6108be5f89355f1a84610795565b91505b506108cc8c826107f6565b90508181035f8060388310600181146108ff576108e88461076c565b60f78101600882021b85179350600101915061090a565b8360c0019250600191505b5086816008021b821791506001810190508060080292508451831c8284610100031b17915080850394505080845250508181038220925050508092505050979650505050505050565b61095b6112c0565b606b5460405173ffffffffffffffffffffffffffffffffffffffff8084169216907ff91b2a410a89d46f14ee984a57e6d7892c217f116905371180998e20cef237e5905f90a3606b80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6109f06112c0565b6109f95f611341565b565b5f6067548210610a0c57505f919050565b600882901c5f908152606a6020526040902054600160ff84161b1615156104d5565b337f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1614610af3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f4f6e6c792063616c6c61626c6520627920746865204c3143726f7373446f6d6160448201527f696e4d657373656e67657200000000000000000000000000000000000000000060648201526084016105bc565b6067548110610b5e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f63616e6e6f742064726f702070656e64696e67206d657373616765000000000060448201526064016105bc565b600881901c5f908152606a6020526040902054600160ff83161b16610bdf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f64726f70206e6f6e2d736b6970706564206d657373616765000000000000000060448201526064016105bc565b600881901c5f90815260696020526040902054600160ff83161b1615610c61576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6d65737361676520616c72656164792064726f7070656400000000000000000060448201526064016105bc565b600881901c5f9081526069602052604090208054600160ff84161b1790556040518181527f43a375005206d20a83abc71722cba68c24434a8dc1f583775be7c3fde0396cbf9060200160405180910390a150565b337f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1614610d7a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f4f6e6c792063616c6c61626c6520627920746865204c3143726f7373446f6d6160448201527f696e4d657373656e67657200000000000000000000000000000000000000000060648201526084016105bc565b610d858383836113b7565b3373111100000000000000000000000000000000111101610daa81865f8787876114e7565b5050505050565b6066545f908210610e1e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f6d65737361676520696e646578206f7574206f662072616e676500000000000060448201526064016105bc565b60668281548110610e3157610e316119c9565b905f5260205f2001549050919050565b337f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1614610f06576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f6e6c792063616c6c61626c652062792074686520456e666f7263656454784760448201527f617465776179000000000000000000000000000000000000000000000000000060648201526084016105bc565b73ffffffffffffffffffffffffffffffffffffffff86163b15610f85576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f6f6e6c7920454f4100000000000000000000000000000000000000000000000060448201526064016105bc565b610f908383836113b7565b610f9e8686868686866114e7565b505050505050565b610fae6112c0565b606880549082905560408051828152602081018490527fa030881e03ff723954dd0d35500564afab9603555d09d4456a32436f2b2373c591015b60405180910390a15050565b610ffc6112c0565b606580549082905560408051828152602081018490527fc5271ba80b67178cc31f04a3755325121400925878dc608432b6fcaead3663299101610fe8565b5f54610100900460ff161580801561105857505f54600160ff909116105b806110715750303b15801561107157505f5460ff166001145b6110fd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016105bc565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015611159575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b611161611598565b6068839055606b80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790558015611204575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602001610756565b505050565b6112116112c0565b73ffffffffffffffffffffffffffffffffffffffff81166112b4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016105bc565b6112bd81611341565b50565b60335473ffffffffffffffffffffffffffffffffffffffff1633146109f9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016105bc565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b606854831115611449576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f476173206c696d6974206d757374206e6f7420657863656564206d617847617360448201527f4c696d697400000000000000000000000000000000000000000000000000000060648201526084016105bc565b6010810261520801808410156114e1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603360248201527f496e73756666696369656e7420676173206c696d69742c206d7573742062652060448201527f61626f766520696e7472696e736963206761730000000000000000000000000060648201526084016105bc565b50505050565b6066545f6114fa8883888a898989610763565b606680546001810182555f919091527f46501879b8ca8525e8c2fd519e2fbfcfa2ebea26501294aa02cbfcfb12e943540181905560405190915073ffffffffffffffffffffffffffffffffffffffff80891691908a16907f69cfcb8e6d4192b8aba9902243912587f37e550d75c1fa801491fce26717f37e90611586908a9087908b908b908b906119f6565b60405180910390a35050505050505050565b5f54610100900460ff1661162e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105bc565b6109f95f54610100900460ff166116c7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105bc565b6109f933611341565b5f602082840312156116e0575f80fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610790575f80fd5b5f806040838503121561171b575f80fd5b611724836116e7565b946020939093013593505050565b5f805f60608486031215611744575f80fd5b505081359360208301359350604090920135919050565b5f8083601f84011261176b575f80fd5b50813567ffffffffffffffff811115611782575f80fd5b602083019150836020828501011115611799575f80fd5b9250929050565b5f805f805f805f60c0888a0312156117b6575f80fd5b6117bf886116e7565b965060208801359550604088013594506117db606089016116e7565b93506080880135925060a088013567ffffffffffffffff8111156117fd575f80fd5b6118098a828b0161175b565b989b979a50959850939692959293505050565b5f6020828403121561182c575f80fd5b6104d2826116e7565b5f805f8060608587031215611848575f80fd5b611851856116e7565b935060208501359250604085013567ffffffffffffffff811115611873575f80fd5b61187f8782880161175b565b95989497509550505050565b5f805f805f8060a087890312156118a0575f80fd5b6118a9876116e7565b95506118b7602088016116e7565b94506040870135935060608701359250608087013567ffffffffffffffff8111156118e0575f80fd5b6118ec89828a0161175b565b979a9699509497509295939492505050565b5f806040838503121561190f575f80fd5b8235915061191f602084016116e7565b90509250929050565b5f8060208385031215611939575f80fd5b823567ffffffffffffffff81111561194f575f80fd5b61195b8582860161175b565b90969095509350505050565b5f60208284031215611977575f80fd5b81518015158114611986575f80fd5b9392505050565b80820281158282048414176104d5577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b85815267ffffffffffffffff8516602082015283604082015260806060820152816080820152818360a08301375f81830160a090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016010194935050505056fea164736f6c6343000818000a"

func init() {
	if err := json.Unmarshal([]byte(L1MessageQueueWithGasPriceOracleStorageLayoutJSON), L1MessageQueueWithGasPriceOracleStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1MessageQueueWithGasPriceOracle"] = L1MessageQueueWithGasPriceOracleStorageLayout
	deployedBytecodes["L1MessageQueueWithGasPriceOracle"] = L1MessageQueueWithGasPriceOracleDeployedBin
}
