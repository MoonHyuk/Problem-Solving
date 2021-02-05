# 페어의 노드 스왑

## 풀이1) 새 연결 리스트를 만들기 (시간: O(n), 공간: O(n))

새 ListNode를 만들어 기존 ListNode의 페어 중 두번째 노드를 첫번째 노드보다 먼저 넣는 것을 반복한다.

직관적이고 어려울 게 없지만 별도 공간을 더 쓴다는 점이 단점이다.

```python
   def swapPairs(self, head: ListNode) -> ListNode:
        ans_copy = ans = ListNode()
        
        while head and head.next:
            ans.next = ListNode(head.next.val)
            ans.next.next = ListNode(head.val)
            
            ans = ans.next.next
            head = head.next.next
            
        if head:
            ans.next = head
            
        return ans_copy.next
```

## 풀이2) 값을 바꾸기 (시간: O(n), 공간: O(1))

리스트 노드들이 아래처럼 주어졌다고 해보자.

```
1 -> 2 -> 3 -> 4
```

만약 노드의 next를 바꾸는 방법으로 한다면

```
2 -> 1   4 -> 3
```

위 그림처럼 1 -> 2를 2 -> 1로, 3 -> 4를 4 -> 3으로 바꾸는 것 뿐만 아니라 1의 next를 4로 해줘야 정답이 완성된다.

하지만 노드의 next를 바꾸는 것이 아니라 value를 바꾼다면 연결 관계는 그대로 유지해도 정답이 된다.

```
2 -> 1 -> 4 -> 3
```



```python
    def swapPairs(self, head: ListNode) -> ListNode:
        head_copy = head
        
        while head and head.next:
            head.val, head.next.val = head.next.val, head.val
            head = head.next.next
            
        return head_copy
```



## 풀이3) 직접 뒤집기 (시간: O(n), 공간: O(1))

1번 풀이를 조금만 수정하면 공간복잡도 O(1)로 풀이가 가능하다.

```python
    def swapPairs(self, head: ListNode) -> ListNode:
        ans_copy = ans = ListNode()
        
        while head and head.next:
            nxt = head.next
            nxtnxt = head.next.next
            
            head.next = None
            ans.next = nxt
            ans.next.next = head
            
            ans = ans.next.next
            head = nxtnxt
            
        if head:
            ans.next = head
            
        return ans_copy.next
```

