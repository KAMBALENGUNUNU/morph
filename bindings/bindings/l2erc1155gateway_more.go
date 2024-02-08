// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const L2ERC1155GatewayStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/gateways/L2ERC1155Gateway.sol:L2ERC1155Gateway\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L2/gateways/L2ERC1155Gateway.sol:L2ERC1155Gateway\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L2/gateways/L2ERC1155Gateway.sol:L2ERC1155Gateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)1018_storage\"},{\"astId\":1003,\"contract\":\"contracts/L2/gateways/L2ERC1155Gateway.sol:L2ERC1155Gateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_array(t_uint256)1018_storage\"},{\"astId\":1004,\"contract\":\"contracts/L2/gateways/L2ERC1155Gateway.sol:L2ERC1155Gateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_array(t_uint256)1018_storage\"},{\"astId\":1005,\"contract\":\"contracts/L2/gateways/L2ERC1155Gateway.sol:L2ERC1155Gateway\",\"label\":\"_status\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_uint256\"},{\"astId\":1006,\"contract\":\"contracts/L2/gateways/L2ERC1155Gateway.sol:L2ERC1155Gateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_array(t_uint256)1017_storage\"},{\"astId\":1007,\"contract\":\"contracts/L2/gateways/L2ERC1155Gateway.sol:L2ERC1155Gateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"201\",\"type\":\"t_array(t_uint256)1018_storage\"},{\"astId\":1008,\"contract\":\"contracts/L2/gateways/L2ERC1155Gateway.sol:L2ERC1155Gateway\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"251\",\"type\":\"t_address\"},{\"astId\":1009,\"contract\":\"contracts/L2/gateways/L2ERC1155Gateway.sol:L2ERC1155Gateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"252\",\"type\":\"t_array(t_uint256)1017_storage\"},{\"astId\":1010,\"contract\":\"contracts/L2/gateways/L2ERC1155Gateway.sol:L2ERC1155Gateway\",\"label\":\"counterpart\",\"offset\":0,\"slot\":\"301\",\"type\":\"t_address\"},{\"astId\":1011,\"contract\":\"contracts/L2/gateways/L2ERC1155Gateway.sol:L2ERC1155Gateway\",\"label\":\"router\",\"offset\":0,\"slot\":\"302\",\"type\":\"t_address\"},{\"astId\":1012,\"contract\":\"contracts/L2/gateways/L2ERC1155Gateway.sol:L2ERC1155Gateway\",\"label\":\"messenger\",\"offset\":0,\"slot\":\"303\",\"type\":\"t_address\"},{\"astId\":1013,\"contract\":\"contracts/L2/gateways/L2ERC1155Gateway.sol:L2ERC1155Gateway\",\"label\":\"__rateLimiter\",\"offset\":0,\"slot\":\"304\",\"type\":\"t_address\"},{\"astId\":1014,\"contract\":\"contracts/L2/gateways/L2ERC1155Gateway.sol:L2ERC1155Gateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"305\",\"type\":\"t_array(t_uint256)1016_storage\"},{\"astId\":1015,\"contract\":\"contracts/L2/gateways/L2ERC1155Gateway.sol:L2ERC1155Gateway\",\"label\":\"tokenMapping\",\"offset\":0,\"slot\":\"351\",\"type\":\"t_mapping(t_address,t_address)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1016_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[46]\",\"numberOfBytes\":\"1472\"},\"t_array(t_uint256)1017_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1018_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_mapping(t_address,t_address)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e address)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_address\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L2ERC1155GatewayStorageLayout = new(solc.StorageLayout)

var L2ERC1155GatewayDeployedBin = "0x608060405260043610610123575f3560e01c80638c23d5b2116100a1578063eaa72ad911610071578063f2fde38b11610057578063f2fde38b146103be578063f887ea40146103dd578063fac752eb1461040a575f80fd5b8063eaa72ad91461035b578063f23a6e611461037a575f80fd5b80638c23d5b2146102675780638da5cb5b1461027a578063ba27f50b146102a4578063bc197c81146102e6575f80fd5b80634764cc62116100f657806348de03de116100dc57806348de03de14610213578063715018a614610226578063797594b01461023a575f80fd5b80634764cc62146101d5578063485cc955146101f4575f80fd5b806301ffc9a7146101275780630f2da0801461015b57806321fedfc9146101705780633cb747bf14610183575b5f80fd5b348015610132575f80fd5b50610146610141366004611d8f565b610429565b60405190151581526020015b60405180910390f35b61016e610169366004611df6565b6104c1565b005b61016e61017e366004611e2e565b6104d4565b34801561018e575f80fd5b5061012f546101b09073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610152565b3480156101e0575f80fd5b5061016e6101ef366004611e7b565b6104e8565b3480156101ff575f80fd5b5061016e61020e366004611ee5565b610888565b61016e610221366004611f64565b610a29565b348015610231575f80fd5b5061016e610a40565b348015610245575f80fd5b5061012d546101b09073ffffffffffffffffffffffffffffffffffffffff1681565b61016e610275366004611fe9565b610a53565b348015610285575f80fd5b5060fb5473ffffffffffffffffffffffffffffffffffffffff166101b0565b3480156102af575f80fd5b506101b06102be366004612080565b61015f6020525f908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b3480156102f1575f80fd5b5061032a61030036600461221e565b7fbc197c810000000000000000000000000000000000000000000000000000000095945050505050565b6040517fffffffff000000000000000000000000000000000000000000000000000000009091168152602001610152565b348015610366575f80fd5b5061016e6103753660046122c5565b610a62565b348015610385575f80fd5b5061032a610394366004612378565b7ff23a6e610000000000000000000000000000000000000000000000000000000095945050505050565b3480156103c9575f80fd5b5061016e6103d8366004612080565b610e0f565b3480156103e8575f80fd5b5061012e546101b09073ffffffffffffffffffffffffffffffffffffffff1681565b348015610415575f80fd5b5061016e610424366004611ee5565b610ec6565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f4e2312e00000000000000000000000000000000000000000000000000000000014806104bb57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6104ce8433858585610fd4565b50505050565b6104e18585858585610fd4565b5050505050565b61012f5473ffffffffffffffffffffffffffffffffffffffff16338114610570576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c00000000000000000060448201526064015b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105b9573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105dd91906123dc565b61012d5473ffffffffffffffffffffffffffffffffffffffff908116911614610662576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6f6e6c792063616c6c20627920636f756e7465727061727400000000000000006044820152606401610567565b61066a61132a565b73ffffffffffffffffffffffffffffffffffffffff87166106e7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f746f6b656e20616464726573732063616e6e6f742062652030000000000000006044820152606401610567565b73ffffffffffffffffffffffffffffffffffffffff8087165f90815261015f602052604090205488821691161461077a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f6c3220746f6b656e206d69736d617463680000000000000000000000000000006044820152606401610567565b6040517f731133e900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301526024820185905260448201849052608060648301525f608483015287169063731133e99060a4015f604051808303815f87803b1580156107fb575f80fd5b505af115801561080d573d5f803e3d5ffd5b50506040805173ffffffffffffffffffffffffffffffffffffffff88811682526020820188905291810186905281891693508982169250908a16907f5399dc7b86d085e50a28946dbc213966bb7a7ac78d312aedd6018c791ad6cef99060600160405180910390a461087f6001609755565b50505050505050565b5f54610100900460ff16158080156108a657505f54600160ff909116105b806108bf5750303b1580156108bf57505f5460ff166001145b61094b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610567565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156109a7575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6109af6113a4565b6109b76113a4565b6109c2835f8461143a565b8015610a24575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b610a38863387878787876115e8565b505050505050565b610a48611a27565b610a515f611aa8565b565b61087f878787878787876115e8565b61012f5473ffffffffffffffffffffffffffffffffffffffff16338114610ae5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c0000000000000000006044820152606401610567565b8073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b2e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b5291906123dc565b61012d5473ffffffffffffffffffffffffffffffffffffffff908116911614610bd7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6f6e6c792063616c6c20627920636f756e7465727061727400000000000000006044820152606401610567565b610bdf61132a565b73ffffffffffffffffffffffffffffffffffffffff8916610c5c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f746f6b656e20616464726573732063616e6e6f742062652030000000000000006044820152606401610567565b73ffffffffffffffffffffffffffffffffffffffff8089165f90815261015f60205260409020548a8216911614610cef576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f6c3220746f6b656e206d69736d617463680000000000000000000000000000006044820152606401610567565b6040517fb48ab8b600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff89169063b48ab8b690610d499089908990899089908990600401612440565b5f604051808303815f87803b158015610d60575f80fd5b505af1158015610d72573d5f803e3d5ffd5b505050508673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff168a73ffffffffffffffffffffffffffffffffffffffff167ff07745bfeb45fb1184165136e9148689adf57ba578a5b90dde949f26066b77568989898989604051610df295949392919061249f565b60405180910390a4610e046001609755565b505050505050505050565b610e17611a27565b73ffffffffffffffffffffffffffffffffffffffff8116610eba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610567565b610ec381611aa8565b50565b610ece611a27565b73ffffffffffffffffffffffffffffffffffffffff8116610f4b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f746f6b656e20616464726573732063616e6e6f742062652030000000000000006044820152606401610567565b73ffffffffffffffffffffffffffffffffffffffff8083165f81815261015f602052604080822080548686167fffffffffffffffffffffffff0000000000000000000000000000000000000000821681179092559151919094169392849290917f2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d9190a4505050565b610fdc61132a565b5f8211611045576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7769746864726177207a65726f20616d6f756e740000000000000000000000006044820152606401610567565b73ffffffffffffffffffffffffffffffffffffffff8086165f90815261015f602052604090205416806110d4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f6e6f20636f72726573706f6e64696e67206c3120746f6b656e000000000000006044820152606401610567565b5f336040517ff5298aca00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff808316600483015260248201889052604482018790529192509088169063f5298aca906064015f604051808303815f87803b15801561114e575f80fd5b505af1158015611160573d5f803e3d5ffd5b505060405173ffffffffffffffffffffffffffffffffffffffff8086166024830152808b16604483015280851660648301528916608482015260a4810188905260c481018790525f925060e4019050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f730608b30000000000000000000000000000000000000000000000000000000017905261012f5461012d5491517fb2267a7b00000000000000000000000000000000000000000000000000000000815292935073ffffffffffffffffffffffffffffffffffffffff9081169263b2267a7b92349261128b929116905f9087908b906004016124ed565b5f604051808303818588803b1580156112a2575f80fd5b505af11580156112b4573d5f803e3d5ffd5b50506040805173ffffffffffffffffffffffffffffffffffffffff8c81168252602082018c90529181018a905281871694508c8216935090871691507f1f9dcda7fce6f73a13055f044ffecaed2032a7a844e0a37a3eb8bbb17488d01a9060600160405180910390a45050506104e16001609755565b600260975403611396576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610567565b6002609755565b6001609755565b5f54610100900460ff16610a51576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610567565b73ffffffffffffffffffffffffffffffffffffffff83166114b7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f7a65726f20636f756e74657270617274206164647265737300000000000000006044820152606401610567565b73ffffffffffffffffffffffffffffffffffffffff8116611534576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f7a65726f206d657373656e6765722061646472657373000000000000000000006044820152606401610567565b61153c611b1e565b611544611bbc565b61012d805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925561012f80548484169216919091179055821615610a245761012e805473ffffffffffffffffffffffffffffffffffffffff84167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055505050565b6115f061132a565b83611657576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6e6f20746f6b656e20746f2077697468647261770000000000000000000000006044820152606401610567565b8382146116c0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f6c656e677468206d69736d6174636800000000000000000000000000000000006044820152606401610567565b5f5b82811015611753575f8484838181106116dd576116dd612583565b905060200201351161174b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7769746864726177207a65726f20616d6f756e740000000000000000000000006044820152606401610567565b6001016116c2565b5073ffffffffffffffffffffffffffffffffffffffff8088165f90815261015f602052604090205416806117e3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f6e6f20636f72726573706f6e64696e67206c3120746f6b656e000000000000006044820152606401610567565b6040517ff6eb127a000000000000000000000000000000000000000000000000000000008152339073ffffffffffffffffffffffffffffffffffffffff8a169063f6eb127a9061183f9084908b908b908b908b9060040161249f565b5f604051808303815f87803b158015611856575f80fd5b505af1158015611868573d5f803e3d5ffd5b505050505f828a838b8b8b8b8b60405160240161188c9897969594939291906125b0565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff92748d30000000000000000000000000000000000000000000000000000000017905261012f5461012d5491517fb2267a7b00000000000000000000000000000000000000000000000000000000815292935073ffffffffffffffffffffffffffffffffffffffff9081169263b2267a7b923492611968929116905f9087908b906004016124ed565b5f604051808303818588803b15801561197f575f80fd5b505af1158015611991573d5f803e3d5ffd5b50505050508173ffffffffffffffffffffffffffffffffffffffff168a73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f5d2d5d4cdbf7b115e43f0b9986644dd8b9514b10be6a019ab6a4a87f122909708c8c8c8c8c604051611a1295949392919061249f565b60405180910390a450505061087f6001609755565b60fb5473ffffffffffffffffffffffffffffffffffffffff163314610a51576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610567565b60fb805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f54610100900460ff16611bb4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610567565b610a51611c5a565b5f54610100900460ff16611c52576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610567565b610a51611cf0565b5f54610100900460ff1661139d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610567565b5f54610100900460ff16611d86576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610567565b610a5133611aa8565b5f60208284031215611d9f575f80fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114611dce575f80fd5b9392505050565b73ffffffffffffffffffffffffffffffffffffffff81168114610ec3575f80fd5b5f805f8060808587031215611e09575f80fd5b8435611e1481611dd5565b966020860135965060408601359560600135945092505050565b5f805f805f60a08688031215611e42575f80fd5b8535611e4d81611dd5565b94506020860135611e5d81611dd5565b94979496505050506040830135926060810135926080909101359150565b5f805f805f8060c08789031215611e90575f80fd5b8635611e9b81611dd5565b95506020870135611eab81611dd5565b94506040870135611ebb81611dd5565b93506060870135611ecb81611dd5565b9598949750929560808101359460a0909101359350915050565b5f8060408385031215611ef6575f80fd5b8235611f0181611dd5565b91506020830135611f1181611dd5565b809150509250929050565b5f8083601f840112611f2c575f80fd5b50813567ffffffffffffffff811115611f43575f80fd5b6020830191508360208260051b8501011115611f5d575f80fd5b9250929050565b5f805f805f8060808789031215611f79575f80fd5b8635611f8481611dd5565b9550602087013567ffffffffffffffff80821115611fa0575f80fd5b611fac8a838b01611f1c565b90975095506040890135915080821115611fc4575f80fd5b50611fd189828a01611f1c565b979a9699509497949695606090950135949350505050565b5f805f805f805f60a0888a031215611fff575f80fd5b873561200a81611dd5565b9650602088013561201a81611dd5565b9550604088013567ffffffffffffffff80821115612036575f80fd5b6120428b838c01611f1c565b909750955060608a013591508082111561205a575f80fd5b506120678a828b01611f1c565b989b979a50959894979596608090950135949350505050565b5f60208284031215612090575f80fd5b8135611dce81611dd5565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561210f5761210f61209b565b604052919050565b5f82601f830112612126575f80fd5b8135602067ffffffffffffffff8211156121425761214261209b565b8160051b6121518282016120c8565b928352848101820192828101908785111561216a575f80fd5b83870192505b8483101561218957823582529183019190830190612170565b979650505050505050565b5f82601f8301126121a3575f80fd5b813567ffffffffffffffff8111156121bd576121bd61209b565b6121ee60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016120c8565b818152846020838601011115612202575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f805f60a08688031215612232575f80fd5b853561223d81611dd5565b9450602086013561224d81611dd5565b9350604086013567ffffffffffffffff80821115612269575f80fd5b61227589838a01612117565b9450606088013591508082111561228a575f80fd5b61229689838a01612117565b935060808801359150808211156122ab575f80fd5b506122b888828901612194565b9150509295509295909350565b5f805f805f805f8060c0898b0312156122dc575f80fd5b88356122e781611dd5565b975060208901356122f781611dd5565b9650604089013561230781611dd5565b9550606089013561231781611dd5565b9450608089013567ffffffffffffffff80821115612333575f80fd5b61233f8c838d01611f1c565b909650945060a08b0135915080821115612357575f80fd5b506123648b828c01611f1c565b999c989b5096995094979396929594505050565b5f805f805f60a0868803121561238c575f80fd5b853561239781611dd5565b945060208601356123a781611dd5565b93506040860135925060608601359150608086013567ffffffffffffffff8111156123d0575f80fd5b6122b888828901612194565b5f602082840312156123ec575f80fd5b8151611dce81611dd5565b8183525f7f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831115612427575f80fd5b8260051b80836020870137939093016020019392505050565b73ffffffffffffffffffffffffffffffffffffffff86168152608060208201525f61246f6080830186886123f7565b82810360408401526124828185876123f7565b83810360609094019390935250505f815260200195945050505050565b73ffffffffffffffffffffffffffffffffffffffff86168152606060208201525f6124ce6060830186886123f7565b82810360408401526124e18185876123f7565b98975050505050505050565b73ffffffffffffffffffffffffffffffffffffffff851681525f60208560208401526080604084015284518060808501525f5b8181101561253c5786810183015185820160a001528201612520565b505f60a0828601015260a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505082606083015295945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f73ffffffffffffffffffffffffffffffffffffffff808b168352808a166020840152808916604084015280881660608401525060c060808301526125f960c0830186886123f7565b82810360a084015261260c8185876123f7565b9b9a505050505050505050505056fea164736f6c6343000818000a"

func init() {
	if err := json.Unmarshal([]byte(L2ERC1155GatewayStorageLayoutJSON), L2ERC1155GatewayStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2ERC1155Gateway"] = L2ERC1155GatewayStorageLayout
	deployedBytecodes["L2ERC1155Gateway"] = L2ERC1155GatewayDeployedBin
}
