# 홀짝 연결 리스트

## 풀이1) 

ListNode가

```
1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7
```

로 주어졌다고 하자.

우선은 `odd = head`, `even = head.next` 로 각각 1과 2에서 시작한다.

|      | 1      | 2      | 3         |
| ---- | ------ | ------ | --------- |
| odd  | 1 -> 3 | 3 -> 5 | 5 -> 7    |
| even | 2 -> 4 | 4 -> 6 | 6 -> None |

그 다음엔 위 표 처럼 odd와 even의 next를 바꾸는 작업을 반복한다.

마지막으론 odd의 next를 even의 처음 원소로 가게 한다.

단, 반복문을 돌면서 even의 현재 노드는 가장 끝에 있으므로 반복문을 돌기 전에 even의 시작 노드를 복사해둔다.

```python
def oddEvenList(self, head: ListNode) -> ListNode:
    if not head:
        return head
    
    odd = head
    even = even_copy = head.next
    
    while even and even.next:
        odd.next = odd.next.next
        odd = odd.next
        
        even.next = even.next.next
        even = even.next
        
    odd.next = even_copy
    
    return head 
```
