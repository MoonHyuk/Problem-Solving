# 두 수의 덧셈



## 풀이1) 전가산기 구현

L1은

```
2 -> 4 -> 3 -> 1
```

L2는

```
5 -> 9 -> 4
```

일 때 L1과 L2의 각 자릿수마다 L1, L2의 값과 carry를 더하여 합 리스트를 만들어 나가는데,

만약 합이 10이 넘으면 다음 자릿수 합에 carry=1을 넘겨주고 자신은 일의 자리만 남긴다.

(예: L1이 4, L2가 9일 때 carry=1, sum=3)

만약 합이 10이 넘지 않으면 carry=0이고 합은 그대로이다.

| L1    | 2    | 4    | 3    | 1    |
| ----- | ---- | ---- | ---- | ---- |
| L2    | 5    | 9    | 4    | 0    |
| carry | 0    | 0    | 1    | 0    |
| sum   | 7    | 3    | 8    | 1    |

```python
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        head = head_copy = ListNode()
      
        carry = 0
        while l1 or l2 or carry:
            s = carry
            if l1:
                s += l1.val
                l1 = l1.next
            if l2:
                s += l2.val
                l2 = l2.next
            
            carry, s = s // 10, s % 10
            head.next = ListNode(s)
            head = head.next
        
        return head_copy.next
```



`carry, s = s // 10, s % 10` 구문은 `divmod` 함수를 사용해서 더 간단히 만들 수 있다.

```python
carry, s = divmod(s, 10)
```

