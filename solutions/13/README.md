# 역순 연결 리스트2

## 풀이1)

```python
    def reverseBetween(self, head: ListNode, m: int, n: int) -> ListNode:        
        start = head
        
        i = 1
        prev = None
        while head:
            nxt = head.next
            if i == m:
                x = prev
                y = head
                
            if i > m and i <= n:
                head.next = prev
                
            if i == n:
                y.next = nxt
                if m > 1:
                    x.next = head
                else:
                    start = head

            i += 1
            prev = head
            head = nxt
        
        return start
```

