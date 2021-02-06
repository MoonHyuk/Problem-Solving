# 주식을 사고팔기 가장 좋은 시점

## 풀이1)

```python
3, 4, 6, 1, 2, 3
```

먼저 시작할 때는 사고 본다.

1일째에 3원에 산다. 2일째와 3일째에 가격이 올라서 이익은 3원까지 됐다.

4일째에는 가격이 샀을 때보다 더 떨어졌다. 이 이후 수익률은 3원을 그대로 가지고 있는 것 보단 1원일 때 사는게 무조건 좋다.

4일째에 1원에 사고, 5일째 6일째에 가격이 올라 이익이 2원이 됐다.

이익들 중 최대값은 3원이다.

최저가가 나오면 항상 사고, 이익을 구해서 이익들의 최댓값을 구하면 된다.

```python
    def maxProfit(self, prices: List[int]) -> int:
        max_profit = 0
        min_price = 100000
        
        for price in prices:
            min_price = min(min_price, price)
            max_profit = max(max_profit, price - min_price)
            
        return max_profit
```

