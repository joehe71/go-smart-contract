// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

contract Bank {
    // 存储每个账户的余额
    mapping(address => uint256) private balances;

    // 所有账户
    address[] private accounts;

    // 存款函数，允许用户存入一定数量的以太币
    function deposit() public payable {
        require(msg.value > 0, "Deposit amount must be greater than zero");

        // 如果用户是第一次存款，则将其添加到账户列表中
        if (balances[msg.sender] == 0) {
            accounts.push(msg.sender);
        }

        balances[msg.sender] += msg.value;
    }

    // 取款函数，允许用户从其账户中提取一定数量的以太币
    function withdraw(uint256 amount) public {
        require(amount > 0, "Withdrawal amount must be greater than zero");
        require(balances[msg.sender] >= amount, "Insufficient balance");

        balances[msg.sender] -= amount;
        payable(msg.sender).transfer(amount);
    }

    // 查询余额函数，允许用户查看其账户的余额
    function getBalance() public view returns (uint256) {
        return balances[msg.sender];
    }

    // 查询银行总余额函数，允许查看合约中存储的所有以太币的总和
    function getTotalBalance() public view returns (uint256) {
        uint256 totalBalance = 0;
        for (uint256 i = 0; i < accounts.length; i++) {
            totalBalance += balances[accounts[i]];
        }
        return totalBalance;
    }
}
