// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.7.0 <0.9.0;

library SafeMath {
    function add(uint256 a, uint256 b) internal pure returns (uint256) {
        uint256 c = a + b;
        require(c >= a, "SafeMath: addition overflow");

        return c;
    }

    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b <= a, "SafeMath: subtraction overflow");
        uint256 c = a - b;

        return c;
    }

    function mul(uint256 a, uint256 b) internal pure returns (uint256) {
        if (a == 0 || b == 0) {
            return 0;
        }

        uint256 c = a * b;
        require(c / a == b, "SafeMath: multiplication overflow");

        return c;
    }

    function div(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b > 0, "SafeMath: division by zero");
        uint256 c = a / b;

        return c;
    }
}

interface IUniswapV2Factory {
    function getPair(address tokenA, address tokenB) external view returns (address pair);
}

interface IUniswapV2Pair {
    function getReserves() external view returns (uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast);
    function token0() external view returns (address);
    function swap(uint amount0Out, uint amount1Out, address to, bytes calldata data) external;
}

interface IUniswapV2Router02 {
    function factory() external pure returns (address);
    function getAmountOut(uint amountIn, uint reserveIn, uint reserveOut) external pure returns (uint amountOut);
    function getAmountIn(uint amountOut, uint reserveIn, uint reserveOut) external pure returns (uint amountIn);
    function getAmountsOut(uint amountIn, address[] calldata path) external view returns (uint[] memory amounts);
    function getAmountsIn(uint amountOut, address[] calldata path) external view returns (uint[] memory amounts);
    function swapExactTokensForTokens(
        uint amountIn,
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external returns (uint[] memory amounts);
}

interface IERC20 {
    function decimals() external view returns (uint8);
    function balanceOf(address owner) external view returns (uint);
    function transfer(address to, uint value) external returns (bool);
    function approve(address _spender, uint _value) external returns (bool);
}

interface IChiToken {
    function freeUpTo(uint256 value) external returns (uint256 freed);
    function mint(uint256 value) external;
}

contract TradeBot {


    address public immutable owner;

    // address public immutable WBNB;

    // address public immutable BUSD;

    // address public immutable USDC;

    // address public immutable USDT;

    address payable private administrator;

    mapping(address => bool) private whiteList;

    IChiToken constant chiToken = IChiToken(0x0000000000004946c0e9F43F4Dee607b0eF1fA1c);


    receive() external payable {}

    modifier onlyOwner() {
        require(msg.sender == owner, "Unauthorized access.");
        _;
    }

    modifier onlyWhite() {
        require(whiteList[msg.sender], "Unauthorized action.");
        _;
    }

    modifier gasTokenRefund {
        uint256 gasStart = gasleft();
        _;
        chiToken.freeUpTo((21000 + gasStart - gasleft() + 16 * msg.data.length + 14154) / 41947);
    }


    constructor(address sender) {
        owner = sender;
        administrator = payable(sender);
        whiteList[sender] = true;
    }

    function mintChiToken(uint256 amount) external virtual onlyOwner {
        chiToken.mint(amount);
    }

    function sendTokenBack(address token, uint256 amount) external virtual onlyOwner {
        IERC20(token).transfer(owner, amount);
    }

    function sendTokenBackAll(address token) external virtual onlyOwner {
        IERC20(token).transfer(owner, IERC20(token).balanceOf(address(this)));
    }

    function sendBnbBack() external virtual onlyOwner {
        administrator.transfer(address(this).balance);
    }

    function setWhite(address account) external virtual onlyOwner {
        whiteList[account] = true;
    }

    function getPair(address router, address token0, address token1) public view returns (address pair){
        // 获取工厂合约实例
        IUniswapV2Factory _uniswapV2Factory = IUniswapV2Factory(IUniswapV2Router02(router).factory());
        // 调用工厂合约的 `getPair` 函数获取交易对地址
        return _uniswapV2Factory.getPair(token0, token1);
    }


    function swap(address router,address token0, address token1, uint256 amount) external virtual onlyWhite gasTokenRefund {
        // 确认代币已经被购买
        require(IERC20(token0).balanceOf(address(this)) > 0, "Swap: Token not enought.");

        // 将 token0 授权给 pancakeswapRouter 进行交易
        IERC20(token0).approve(router, amount);

        // 获取交易路径，从 token0 到 token1
        address[] memory path = new address[](2);
        path[0] = token0;
        path[1] = token1;
        IUniswapV2Router02 pancakeswapRouter = IUniswapV2Router02(router);
        // 计算 amountOutMin 参数，保证至少有 1% 的交易价格保护
        uint256[] memory amounts = pancakeswapRouter.getAmountsOut(amount, path);
        uint256 amountOutMin = amounts[1] - (amounts[1] / 100);

        // 执行交易
        pancakeswapRouter.swapExactTokensForTokens(
            amount,
            amountOutMin,
            path,
            address(this),
            block.timestamp + 60 // 交易截止时间，一般设置为当前区块时间加上 60 秒
        );
    }

}