// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.7.0 <0.9.0;

import "./tradeBot.sol";

contract TradeBotFactory {

    address public immutable owner;

    mapping(uint256 => address) private tradeBots;
    uint256 private tradeBotCount;

    event NewTradeBot(address tradeBot);

    constructor() {
        owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Unauthorized access.");
        _;
    }

    function createTradeBot(address wbnb, address busd, address usdt, address usdc) external returns (address){
        address newTradeBot = address(new TradeBot(msg.sender,wbnb,busd,usdt,usdc));
        emit NewTradeBot(newTradeBot);
        tradeBotCount++;
        tradeBots[tradeBotCount] = address(newTradeBot);
        return address(newTradeBot);
    }

    function getTradeBotCount() external view returns (uint256) {
        return tradeBotCount;
    }

    function getTradeBots(uint256 pageIndex, uint256 pageSize) external view returns (address[] memory) {
        uint256 startIndex = (pageIndex - 1) * pageSize + 1;
        uint256 endIndex = pageIndex * pageSize;
        if (endIndex > tradeBotCount) {
            endIndex = tradeBotCount;
        }
        uint256 length = endIndex - startIndex + 1;
        address[] memory result = new address[](length);
        for (uint256 i = startIndex; i <= endIndex; i++) {
            result[i - startIndex] = tradeBots[i];
        }
        return result;
    }
}