# go-smart-contract
这是一个go与以太坊合约交互的demo，参考了[以太坊开发文档](https://ethereum.org/zh/developers/docs/programming-languages/golang/)

## 1.前置环境

- 安装solc编译工具: https://docs.soliditylang.org/en/v0.8.2/installing-solidity.html#linux-packages
- 编译abigen命令: https://github.com/ethereum/go-ethereum
- 安装本地测试链: https://trufflesuite.com/ganache

## 2. abi与bin生成

- bin生成: solcjs --optimize --bin ./contracts/Bank.sol  -o build
- abi生成: solcjs --optimize --abi  ./contracts/Bank.sol  -o build


## 3. 生成合约go文件

- go文件生成: abigen --abi=./build/contracts_Bank_sol_Bank.abi   --bin=./build/contracts_Bank_sol_Bank.bin  --pkg=api --out=/api/Bank.go

# 架构图
![My Local Image](./架构图.png)