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
)

// L2WETHGatewayMetaData contains all meta data concerning the L2WETHGateway contract.
var L2WETHGatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1WETH\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinalizeDepositERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"WithdrawERC20\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeDepositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getL1ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getL2ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20AndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b5060405162001dd038038062001dd0833981016040819052620000349162000134565b6200003e62000056565b6001600160a01b0391821660a052166080526200016c565b600054610100900460ff1615620000c35760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161462000115576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b80516001600160a01b03811681146200012f57600080fd5b919050565b600080604083850312156200014857600080fd5b620001538362000117565b9150620001636020840162000117565b90509250929050565b60805160a051611c0e620001c26000396000818160f4015281816102a201528181610336015281816105dc0152610a920152600081816101cf015281816102d60152818161055b0152610bca0152611c0e6000f3fe6080604052600436106100ec5760003560e01c80638da5cb5b1161008a578063c0c53b8b11610059578063c0c53b8b146102f8578063c676ad2914610318578063f2fde38b14610358578063f887ea401461037857600080fd5b80638da5cb5b1461025f578063a93a4af91461027d578063ad5c464814610290578063b32d8c65146102c457600080fd5b80636c07ea43116100c65780636c07ea4314610204578063715018a614610217578063797594b01461022c5780638431f5c11461024c57600080fd5b80633cb747bf1461017557806354bbd59c146101b1578063575361b6146101f157600080fd5b3661017057337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03161461016e5760405162461bcd60e51b815260206004820152600960248201527f6f6e6c792057455448000000000000000000000000000000000000000000000060448201526064015b60405180910390fd5b005b600080fd5b34801561018157600080fd5b50609954610195906001600160a01b031681565b6040516001600160a01b03909116815260200160405180910390f35b3480156101bd57600080fd5b506101956101cc36600461164a565b507f000000000000000000000000000000000000000000000000000000000000000090565b61016e6101ff3660046116b7565b610398565b61016e610212366004611732565b6103e4565b34801561022357600080fd5b5061016e610423565b34801561023857600080fd5b50609754610195906001600160a01b031681565b61016e61025a366004611767565b610437565b34801561026b57600080fd5b506065546001600160a01b0316610195565b61016e61028b3660046117ff565b6107c2565b34801561029c57600080fd5b506101957f000000000000000000000000000000000000000000000000000000000000000081565b3480156102d057600080fd5b506101957f000000000000000000000000000000000000000000000000000000000000000081565b34801561030457600080fd5b5061016e610313366004611845565b6107d5565b34801561032457600080fd5b5061019561033336600461164a565b507f000000000000000000000000000000000000000000000000000000000000000090565b34801561036457600080fd5b5061016e61037336600461164a565b6109a8565b34801561038457600080fd5b50609854610195906001600160a01b031681565b6103dc86868686868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250889250610a38915050565b505050505050565b61041e83338460005b6040519080825280601f01601f191660200182016040528015610417576020820181803683370190505b5085610a38565b505050565b61042b610d84565b6104356000610dde565b565b6099546001600160a01b03163381146104925760405162461bcd60e51b815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c0000000000000000006044820152606401610165565b806001600160a01b0316636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104f491906118bf565b6097546001600160a01b039081169116146105515760405162461bcd60e51b815260206004820152601860248201527f6f6e6c792063616c6c20627920636f756e7465727061727400000000000000006044820152606401610165565b610559610e48565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316886001600160a01b0316146105da5760405162461bcd60e51b815260206004820152601160248201527f6c3120746f6b656e206e6f7420574554480000000000000000000000000000006044820152606401610165565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316876001600160a01b03161461065b5760405162461bcd60e51b815260206004820152601160248201527f6c3220746f6b656e206e6f7420574554480000000000000000000000000000006044820152606401610165565b3484146106aa5760405162461bcd60e51b815260206004820152601260248201527f6d73672e76616c7565206d69736d6174636800000000000000000000000000006044820152606401610165565b866001600160a01b031663d0e30db0856040518263ffffffff1660e01b81526004016000604051808303818588803b1580156106e557600080fd5b505af11580156106f9573d6000803e3d6000fd5b50610714935050506001600160a01b03891690508686610ea1565b6107548584848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610f6892505050565b856001600160a01b0316876001600160a01b0316896001600160a01b03167f165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34888888886040516107a794939291906118dc565b60405180910390a46107b860018055565b5050505050505050565b6107cf84848460006103ed565b50505050565b600054610100900460ff16158080156107f55750600054600160ff909116105b8061080f5750303b15801561080f575060005460ff166001145b6108815760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610165565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156108df57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6001600160a01b0383166109355760405162461bcd60e51b815260206004820152601360248201527f7a65726f20726f757465722061646472657373000000000000000000000000006044820152606401610165565b610940848484611005565b80156107cf57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050565b6109b0610d84565b6001600160a01b038116610a2c5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610165565b610a3581610dde565b50565b610a40610e48565b60008311610a905760405162461bcd60e51b815260206004820152601460248201527f7769746864726177207a65726f20616d6f756e740000000000000000000000006044820152606401610165565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316856001600160a01b031614610b115760405162461bcd60e51b815260206004820152601460248201527f6f6e6c79205745544820697320616c6c6f7765640000000000000000000000006044820152606401610165565b60985433906001600160a01b0316819003610b3f5782806020019051810190610b3a9190611964565b935090505b610b546001600160a01b038716823087611148565b6040517f2e1a7d4d000000000000000000000000000000000000000000000000000000008152600481018590526001600160a01b03871690632e1a7d4d90602401600060405180830381600087803b158015610baf57600080fd5b505af1158015610bc3573d6000803e3d6000fd5b50506040517f0000000000000000000000000000000000000000000000000000000000000000925060009150610c079083908a9086908b908b908b90602401611a8f565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f84bd13b0000000000000000000000000000000000000000000000000000000001790526099549091506001600160a01b031663b2267a7b610c9b3489611add565b6097546040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168152610ce6916001600160a01b0316908b9087908b90600401611b1d565b6000604051808303818588803b158015610cff57600080fd5b505af1158015610d13573d6000803e3d6000fd5b5050505050826001600160a01b0316886001600160a01b0316836001600160a01b03167fd8d3a3f4ab95694bef40475997598bcf8acd3ed9617a4c1013795429414c27e88a8a8a604051610d6993929190611b56565b60405180910390a4505050610d7d60018055565b5050505050565b6065546001600160a01b031633146104355760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610165565b606580546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600260015403610e9a5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610165565b6002600155565b6040516001600160a01b03831660248201526044810182905261041e9084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611199565b60008151118015610f8357506000826001600160a01b03163b115b15610ffb576040517f444b281f0000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063444b281f90610fcd908490600401611b87565b600060405180830381600087803b158015610fe757600080fd5b505af11580156103dc573d6000803e3d6000fd5b5050565b60018055565b6001600160a01b03831661105b5760405162461bcd60e51b815260206004820152601860248201527f7a65726f20636f756e74657270617274206164647265737300000000000000006044820152606401610165565b6001600160a01b0381166110b15760405162461bcd60e51b815260206004820152601660248201527f7a65726f206d657373656e6765722061646472657373000000000000000000006044820152606401610165565b6110b9611281565b6110c1611306565b609780546001600160a01b038086167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560998054848416921691909117905582161561041e57609880546001600160a01b0384167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055505050565b6040516001600160a01b03808516602483015283166044820152606481018290526107cf9085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401610ee6565b60006111ee826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661138b9092919063ffffffff16565b905080516000148061120f57508080602001905181019061120f9190611b9a565b61041e5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610165565b600054610100900460ff166112fe5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610165565b6104356113a2565b600054610100900460ff166113835760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610165565b61043561141f565b606061139a84846000856114a5565b949350505050565b600054610100900460ff16610fff5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610165565b600054610100900460ff1661149c5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610165565b61043533610dde565b60608247101561151d5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610165565b600080866001600160a01b031685876040516115399190611bbc565b60006040518083038185875af1925050503d8060008114611576576040519150601f19603f3d011682016040523d82523d6000602084013e61157b565b606091505b509150915061158c87838387611597565b979650505050505050565b606083156116065782516000036115ff576001600160a01b0385163b6115ff5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610165565b508161139a565b61139a838381511561161b5781518083602001fd5b8060405162461bcd60e51b81526004016101659190611b87565b6001600160a01b0381168114610a3557600080fd5b60006020828403121561165c57600080fd5b813561166781611635565b9392505050565b60008083601f84011261168057600080fd5b50813567ffffffffffffffff81111561169857600080fd5b6020830191508360208285010111156116b057600080fd5b9250929050565b60008060008060008060a087890312156116d057600080fd5b86356116db81611635565b955060208701356116eb81611635565b945060408701359350606087013567ffffffffffffffff81111561170e57600080fd5b61171a89828a0161166e565b979a9699509497949695608090950135949350505050565b60008060006060848603121561174757600080fd5b833561175281611635565b95602085013595506040909401359392505050565b600080600080600080600060c0888a03121561178257600080fd5b873561178d81611635565b9650602088013561179d81611635565b955060408801356117ad81611635565b945060608801356117bd81611635565b93506080880135925060a088013567ffffffffffffffff8111156117e057600080fd5b6117ec8a828b0161166e565b989b979a50959850939692959293505050565b6000806000806080858703121561181557600080fd5b843561182081611635565b9350602085013561183081611635565b93969395505050506040820135916060013590565b60008060006060848603121561185a57600080fd5b833561186581611635565b9250602084013561187581611635565b9150604084013561188581611635565b809150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000602082840312156118d157600080fd5b815161166781611635565b6001600160a01b038516815283602082015260606040820152816060820152818360808301376000818301608090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01601019392505050565b60005b8381101561195b578181015183820152602001611943565b50506000910152565b6000806040838503121561197757600080fd5b825161198281611635565b602084015190925067ffffffffffffffff808211156119a057600080fd5b818501915085601f8301126119b457600080fd5b8151818111156119c6576119c6611890565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715611a0c57611a0c611890565b81604052828152886020848701011115611a2557600080fd5b611a36836020830160208801611940565b80955050505050509250929050565b60008151808452611a5d816020860160208601611940565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60006001600160a01b0380891683528088166020840152808716604084015280861660608401525083608083015260c060a0830152611ad160c0830184611a45565b98975050505050505050565b80820180821115611b17577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b92915050565b6001600160a01b0385168152836020820152608060408201526000611b456080830185611a45565b905082606083015295945050505050565b6001600160a01b0384168152826020820152606060408201526000611b7e6060830184611a45565b95945050505050565b6020815260006116676020830184611a45565b600060208284031215611bac57600080fd5b8151801515811461166757600080fd5b60008251611bce818460208701611940565b919091019291505056fea2646970667358221220b7f2809b964cfd96e9ef4d928a065c218629b1ed257bb9b1fe8c13e37abf893364736f6c63430008100033",
}

// L2WETHGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use L2WETHGatewayMetaData.ABI instead.
var L2WETHGatewayABI = L2WETHGatewayMetaData.ABI

// L2WETHGatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2WETHGatewayMetaData.Bin instead.
var L2WETHGatewayBin = L2WETHGatewayMetaData.Bin

// DeployL2WETHGateway deploys a new Ethereum contract, binding an instance of L2WETHGateway to it.
func DeployL2WETHGateway(auth *bind.TransactOpts, backend bind.ContractBackend, _WETH common.Address, _l1WETH common.Address) (common.Address, *types.Transaction, *L2WETHGateway, error) {
	parsed, err := L2WETHGatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2WETHGatewayBin), backend, _WETH, _l1WETH)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2WETHGateway{L2WETHGatewayCaller: L2WETHGatewayCaller{contract: contract}, L2WETHGatewayTransactor: L2WETHGatewayTransactor{contract: contract}, L2WETHGatewayFilterer: L2WETHGatewayFilterer{contract: contract}}, nil
}

// L2WETHGateway is an auto generated Go binding around an Ethereum contract.
type L2WETHGateway struct {
	L2WETHGatewayCaller     // Read-only binding to the contract
	L2WETHGatewayTransactor // Write-only binding to the contract
	L2WETHGatewayFilterer   // Log filterer for contract events
}

// L2WETHGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2WETHGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2WETHGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2WETHGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2WETHGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2WETHGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2WETHGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2WETHGatewaySession struct {
	Contract     *L2WETHGateway    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2WETHGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2WETHGatewayCallerSession struct {
	Contract *L2WETHGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// L2WETHGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2WETHGatewayTransactorSession struct {
	Contract     *L2WETHGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// L2WETHGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2WETHGatewayRaw struct {
	Contract *L2WETHGateway // Generic contract binding to access the raw methods on
}

// L2WETHGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2WETHGatewayCallerRaw struct {
	Contract *L2WETHGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// L2WETHGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2WETHGatewayTransactorRaw struct {
	Contract *L2WETHGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2WETHGateway creates a new instance of L2WETHGateway, bound to a specific deployed contract.
func NewL2WETHGateway(address common.Address, backend bind.ContractBackend) (*L2WETHGateway, error) {
	contract, err := bindL2WETHGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2WETHGateway{L2WETHGatewayCaller: L2WETHGatewayCaller{contract: contract}, L2WETHGatewayTransactor: L2WETHGatewayTransactor{contract: contract}, L2WETHGatewayFilterer: L2WETHGatewayFilterer{contract: contract}}, nil
}

// NewL2WETHGatewayCaller creates a new read-only instance of L2WETHGateway, bound to a specific deployed contract.
func NewL2WETHGatewayCaller(address common.Address, caller bind.ContractCaller) (*L2WETHGatewayCaller, error) {
	contract, err := bindL2WETHGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2WETHGatewayCaller{contract: contract}, nil
}

// NewL2WETHGatewayTransactor creates a new write-only instance of L2WETHGateway, bound to a specific deployed contract.
func NewL2WETHGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L2WETHGatewayTransactor, error) {
	contract, err := bindL2WETHGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2WETHGatewayTransactor{contract: contract}, nil
}

// NewL2WETHGatewayFilterer creates a new log filterer instance of L2WETHGateway, bound to a specific deployed contract.
func NewL2WETHGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*L2WETHGatewayFilterer, error) {
	contract, err := bindL2WETHGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2WETHGatewayFilterer{contract: contract}, nil
}

// bindL2WETHGateway binds a generic wrapper to an already deployed contract.
func bindL2WETHGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2WETHGatewayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2WETHGateway *L2WETHGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2WETHGateway.Contract.L2WETHGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2WETHGateway *L2WETHGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.L2WETHGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2WETHGateway *L2WETHGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.L2WETHGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2WETHGateway *L2WETHGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2WETHGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2WETHGateway *L2WETHGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2WETHGateway *L2WETHGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_L2WETHGateway *L2WETHGatewayCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2WETHGateway.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_L2WETHGateway *L2WETHGatewaySession) WETH() (common.Address, error) {
	return _L2WETHGateway.Contract.WETH(&_L2WETHGateway.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_L2WETHGateway *L2WETHGatewayCallerSession) WETH() (common.Address, error) {
	return _L2WETHGateway.Contract.WETH(&_L2WETHGateway.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2WETHGateway *L2WETHGatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2WETHGateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2WETHGateway *L2WETHGatewaySession) Counterpart() (common.Address, error) {
	return _L2WETHGateway.Contract.Counterpart(&_L2WETHGateway.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2WETHGateway *L2WETHGatewayCallerSession) Counterpart() (common.Address, error) {
	return _L2WETHGateway.Contract.Counterpart(&_L2WETHGateway.CallOpts)
}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address ) view returns(address)
func (_L2WETHGateway *L2WETHGatewayCaller) GetL1ERC20Address(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2WETHGateway.contract.Call(opts, &out, "getL1ERC20Address", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address ) view returns(address)
func (_L2WETHGateway *L2WETHGatewaySession) GetL1ERC20Address(arg0 common.Address) (common.Address, error) {
	return _L2WETHGateway.Contract.GetL1ERC20Address(&_L2WETHGateway.CallOpts, arg0)
}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address ) view returns(address)
func (_L2WETHGateway *L2WETHGatewayCallerSession) GetL1ERC20Address(arg0 common.Address) (common.Address, error) {
	return _L2WETHGateway.Contract.GetL1ERC20Address(&_L2WETHGateway.CallOpts, arg0)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) view returns(address)
func (_L2WETHGateway *L2WETHGatewayCaller) GetL2ERC20Address(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2WETHGateway.contract.Call(opts, &out, "getL2ERC20Address", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) view returns(address)
func (_L2WETHGateway *L2WETHGatewaySession) GetL2ERC20Address(arg0 common.Address) (common.Address, error) {
	return _L2WETHGateway.Contract.GetL2ERC20Address(&_L2WETHGateway.CallOpts, arg0)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) view returns(address)
func (_L2WETHGateway *L2WETHGatewayCallerSession) GetL2ERC20Address(arg0 common.Address) (common.Address, error) {
	return _L2WETHGateway.Contract.GetL2ERC20Address(&_L2WETHGateway.CallOpts, arg0)
}

// L1WETH is a free data retrieval call binding the contract method 0xb32d8c65.
//
// Solidity: function l1WETH() view returns(address)
func (_L2WETHGateway *L2WETHGatewayCaller) L1WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2WETHGateway.contract.Call(opts, &out, "l1WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1WETH is a free data retrieval call binding the contract method 0xb32d8c65.
//
// Solidity: function l1WETH() view returns(address)
func (_L2WETHGateway *L2WETHGatewaySession) L1WETH() (common.Address, error) {
	return _L2WETHGateway.Contract.L1WETH(&_L2WETHGateway.CallOpts)
}

// L1WETH is a free data retrieval call binding the contract method 0xb32d8c65.
//
// Solidity: function l1WETH() view returns(address)
func (_L2WETHGateway *L2WETHGatewayCallerSession) L1WETH() (common.Address, error) {
	return _L2WETHGateway.Contract.L1WETH(&_L2WETHGateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2WETHGateway *L2WETHGatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2WETHGateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2WETHGateway *L2WETHGatewaySession) Messenger() (common.Address, error) {
	return _L2WETHGateway.Contract.Messenger(&_L2WETHGateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2WETHGateway *L2WETHGatewayCallerSession) Messenger() (common.Address, error) {
	return _L2WETHGateway.Contract.Messenger(&_L2WETHGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2WETHGateway *L2WETHGatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2WETHGateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2WETHGateway *L2WETHGatewaySession) Owner() (common.Address, error) {
	return _L2WETHGateway.Contract.Owner(&_L2WETHGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2WETHGateway *L2WETHGatewayCallerSession) Owner() (common.Address, error) {
	return _L2WETHGateway.Contract.Owner(&_L2WETHGateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2WETHGateway *L2WETHGatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2WETHGateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2WETHGateway *L2WETHGatewaySession) Router() (common.Address, error) {
	return _L2WETHGateway.Contract.Router(&_L2WETHGateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2WETHGateway *L2WETHGatewayCallerSession) Router() (common.Address, error) {
	return _L2WETHGateway.Contract.Router(&_L2WETHGateway.CallOpts)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2WETHGateway *L2WETHGatewayTransactor) FinalizeDepositERC20(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2WETHGateway.contract.Transact(opts, "finalizeDepositERC20", _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2WETHGateway *L2WETHGatewaySession) FinalizeDepositERC20(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.FinalizeDepositERC20(&_L2WETHGateway.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2WETHGateway *L2WETHGatewayTransactorSession) FinalizeDepositERC20(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.FinalizeDepositERC20(&_L2WETHGateway.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L2WETHGateway *L2WETHGatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2WETHGateway.contract.Transact(opts, "initialize", _counterpart, _router, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L2WETHGateway *L2WETHGatewaySession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.Initialize(&_L2WETHGateway.TransactOpts, _counterpart, _router, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L2WETHGateway *L2WETHGatewayTransactorSession) Initialize(_counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.Initialize(&_L2WETHGateway.TransactOpts, _counterpart, _router, _messenger)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2WETHGateway *L2WETHGatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2WETHGateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2WETHGateway *L2WETHGatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _L2WETHGateway.Contract.RenounceOwnership(&_L2WETHGateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2WETHGateway *L2WETHGatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2WETHGateway.Contract.RenounceOwnership(&_L2WETHGateway.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2WETHGateway *L2WETHGatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2WETHGateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2WETHGateway *L2WETHGatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.TransferOwnership(&_L2WETHGateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2WETHGateway *L2WETHGatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.TransferOwnership(&_L2WETHGateway.TransactOpts, newOwner)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2WETHGateway *L2WETHGatewayTransactor) WithdrawERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2WETHGateway.contract.Transact(opts, "withdrawERC20", _token, _amount, _gasLimit)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2WETHGateway *L2WETHGatewaySession) WithdrawERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.WithdrawERC20(&_L2WETHGateway.TransactOpts, _token, _amount, _gasLimit)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2WETHGateway *L2WETHGatewayTransactorSession) WithdrawERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.WithdrawERC20(&_L2WETHGateway.TransactOpts, _token, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2WETHGateway *L2WETHGatewayTransactor) WithdrawERC200(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2WETHGateway.contract.Transact(opts, "withdrawERC200", _token, _to, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2WETHGateway *L2WETHGatewaySession) WithdrawERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.WithdrawERC200(&_L2WETHGateway.TransactOpts, _token, _to, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2WETHGateway *L2WETHGatewayTransactorSession) WithdrawERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.WithdrawERC200(&_L2WETHGateway.TransactOpts, _token, _to, _amount, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2WETHGateway *L2WETHGatewayTransactor) WithdrawERC20AndCall(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2WETHGateway.contract.Transact(opts, "withdrawERC20AndCall", _token, _to, _amount, _data, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2WETHGateway *L2WETHGatewaySession) WithdrawERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.WithdrawERC20AndCall(&_L2WETHGateway.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2WETHGateway *L2WETHGatewayTransactorSession) WithdrawERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2WETHGateway.Contract.WithdrawERC20AndCall(&_L2WETHGateway.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2WETHGateway *L2WETHGatewayTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2WETHGateway.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2WETHGateway *L2WETHGatewaySession) Receive() (*types.Transaction, error) {
	return _L2WETHGateway.Contract.Receive(&_L2WETHGateway.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2WETHGateway *L2WETHGatewayTransactorSession) Receive() (*types.Transaction, error) {
	return _L2WETHGateway.Contract.Receive(&_L2WETHGateway.TransactOpts)
}

// L2WETHGatewayFinalizeDepositERC20Iterator is returned from FilterFinalizeDepositERC20 and is used to iterate over the raw logs and unpacked data for FinalizeDepositERC20 events raised by the L2WETHGateway contract.
type L2WETHGatewayFinalizeDepositERC20Iterator struct {
	Event *L2WETHGatewayFinalizeDepositERC20 // Event containing the contract specifics and raw log

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
func (it *L2WETHGatewayFinalizeDepositERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2WETHGatewayFinalizeDepositERC20)
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
		it.Event = new(L2WETHGatewayFinalizeDepositERC20)
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
func (it *L2WETHGatewayFinalizeDepositERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2WETHGatewayFinalizeDepositERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2WETHGatewayFinalizeDepositERC20 represents a FinalizeDepositERC20 event raised by the L2WETHGateway contract.
type L2WETHGatewayFinalizeDepositERC20 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFinalizeDepositERC20 is a free log retrieval operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2WETHGateway *L2WETHGatewayFilterer) FilterFinalizeDepositERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2WETHGatewayFinalizeDepositERC20Iterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2WETHGateway.contract.FilterLogs(opts, "FinalizeDepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2WETHGatewayFinalizeDepositERC20Iterator{contract: _L2WETHGateway.contract, event: "FinalizeDepositERC20", logs: logs, sub: sub}, nil
}

// WatchFinalizeDepositERC20 is a free log subscription operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2WETHGateway *L2WETHGatewayFilterer) WatchFinalizeDepositERC20(opts *bind.WatchOpts, sink chan<- *L2WETHGatewayFinalizeDepositERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2WETHGateway.contract.WatchLogs(opts, "FinalizeDepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2WETHGatewayFinalizeDepositERC20)
				if err := _L2WETHGateway.contract.UnpackLog(event, "FinalizeDepositERC20", log); err != nil {
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

// ParseFinalizeDepositERC20 is a log parse operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2WETHGateway *L2WETHGatewayFilterer) ParseFinalizeDepositERC20(log types.Log) (*L2WETHGatewayFinalizeDepositERC20, error) {
	event := new(L2WETHGatewayFinalizeDepositERC20)
	if err := _L2WETHGateway.contract.UnpackLog(event, "FinalizeDepositERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2WETHGatewayInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2WETHGateway contract.
type L2WETHGatewayInitializedIterator struct {
	Event *L2WETHGatewayInitialized // Event containing the contract specifics and raw log

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
func (it *L2WETHGatewayInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2WETHGatewayInitialized)
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
		it.Event = new(L2WETHGatewayInitialized)
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
func (it *L2WETHGatewayInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2WETHGatewayInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2WETHGatewayInitialized represents a Initialized event raised by the L2WETHGateway contract.
type L2WETHGatewayInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2WETHGateway *L2WETHGatewayFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2WETHGatewayInitializedIterator, error) {

	logs, sub, err := _L2WETHGateway.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2WETHGatewayInitializedIterator{contract: _L2WETHGateway.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2WETHGateway *L2WETHGatewayFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2WETHGatewayInitialized) (event.Subscription, error) {

	logs, sub, err := _L2WETHGateway.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2WETHGatewayInitialized)
				if err := _L2WETHGateway.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L2WETHGateway *L2WETHGatewayFilterer) ParseInitialized(log types.Log) (*L2WETHGatewayInitialized, error) {
	event := new(L2WETHGatewayInitialized)
	if err := _L2WETHGateway.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2WETHGatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L2WETHGateway contract.
type L2WETHGatewayOwnershipTransferredIterator struct {
	Event *L2WETHGatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L2WETHGatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2WETHGatewayOwnershipTransferred)
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
		it.Event = new(L2WETHGatewayOwnershipTransferred)
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
func (it *L2WETHGatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2WETHGatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2WETHGatewayOwnershipTransferred represents a OwnershipTransferred event raised by the L2WETHGateway contract.
type L2WETHGatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2WETHGateway *L2WETHGatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L2WETHGatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2WETHGateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L2WETHGatewayOwnershipTransferredIterator{contract: _L2WETHGateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2WETHGateway *L2WETHGatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L2WETHGatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2WETHGateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2WETHGatewayOwnershipTransferred)
				if err := _L2WETHGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L2WETHGateway *L2WETHGatewayFilterer) ParseOwnershipTransferred(log types.Log) (*L2WETHGatewayOwnershipTransferred, error) {
	event := new(L2WETHGatewayOwnershipTransferred)
	if err := _L2WETHGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2WETHGatewayWithdrawERC20Iterator is returned from FilterWithdrawERC20 and is used to iterate over the raw logs and unpacked data for WithdrawERC20 events raised by the L2WETHGateway contract.
type L2WETHGatewayWithdrawERC20Iterator struct {
	Event *L2WETHGatewayWithdrawERC20 // Event containing the contract specifics and raw log

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
func (it *L2WETHGatewayWithdrawERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2WETHGatewayWithdrawERC20)
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
		it.Event = new(L2WETHGatewayWithdrawERC20)
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
func (it *L2WETHGatewayWithdrawERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2WETHGatewayWithdrawERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2WETHGatewayWithdrawERC20 represents a WithdrawERC20 event raised by the L2WETHGateway contract.
type L2WETHGatewayWithdrawERC20 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawERC20 is a free log retrieval operation binding the contract event 0xd8d3a3f4ab95694bef40475997598bcf8acd3ed9617a4c1013795429414c27e8.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2WETHGateway *L2WETHGatewayFilterer) FilterWithdrawERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2WETHGatewayWithdrawERC20Iterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2WETHGateway.contract.FilterLogs(opts, "WithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2WETHGatewayWithdrawERC20Iterator{contract: _L2WETHGateway.contract, event: "WithdrawERC20", logs: logs, sub: sub}, nil
}

// WatchWithdrawERC20 is a free log subscription operation binding the contract event 0xd8d3a3f4ab95694bef40475997598bcf8acd3ed9617a4c1013795429414c27e8.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2WETHGateway *L2WETHGatewayFilterer) WatchWithdrawERC20(opts *bind.WatchOpts, sink chan<- *L2WETHGatewayWithdrawERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2WETHGateway.contract.WatchLogs(opts, "WithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2WETHGatewayWithdrawERC20)
				if err := _L2WETHGateway.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
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

// ParseWithdrawERC20 is a log parse operation binding the contract event 0xd8d3a3f4ab95694bef40475997598bcf8acd3ed9617a4c1013795429414c27e8.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2WETHGateway *L2WETHGatewayFilterer) ParseWithdrawERC20(log types.Log) (*L2WETHGatewayWithdrawERC20, error) {
	event := new(L2WETHGatewayWithdrawERC20)
	if err := _L2WETHGateway.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}