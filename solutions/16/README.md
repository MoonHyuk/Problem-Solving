# 큐를 이용한 스택 구현

## 풀이1) push할 때마다 재정렬하기

만약 `[1, 2, 3]` 순으로 push를 한다면 큐에는 `[3, 2, 1]` 순서로 들어가 있어야 pop을 할 때 스택과 같은 순서로 나온다.

현재 큐가 어떤 방식으로든 `[3, 2, 1]` 로 잘 정렬되어 있다고 하자. 이 때 큐에 새로운 수 4가 push 된다면 `[3, 2, 1, 4]` 가 될텐데, 이를 `[4, 3, 2, 1]` 로 만들어줘야 한다.

큐에 4가 push된 이후에 4앞에 있던 수 3개를 큐에서 차례대로 꺼내서 다시 큐에 넣어주면 된다.

| 처음           | step1          | step2          | step3          |
| -------------- | -------------- | -------------- | -------------- |
| `[3, 2, 1, 4]` | `[2, 1, 4, 3]` | `[1, 4, 3, 2]` | `[4, 3, 2, 1]` |

```python

class MyStack:

    def __init__(self):
        self.q = deque()
        

    def push(self, x: int) -> None:
        self.q.append(x)
        
        for _ in range(len(self.q) - 1):
            self.q.append(self.q.popleft())

    def pop(self) -> int:
        return self.q.popleft()
        

    def top(self) -> int:
        return self.q[0]
        

    def empty(self) -> bool:
        return not self.q

```

