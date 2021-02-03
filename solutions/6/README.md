# 주식을 사고팔기 가장 좋은 시점

## 풀이1)

```python
    def maxProfit(self, prices: List[int]) -> int:
        max_profit = 0
        min_price = 100000
        
        for price in prices:
            min_price = min(min_price, price)
            max_profit = max(max_profit, price - min_price)
            
        return max_profit
```

