# 스택을 이용한 큐 구현

## 풀이1) 스택 2개를 이용

이전 문제인 [큐를 이용한 스택 구현](../16/README.md)과 유사하지만 이번 문제에서는 스택 하나만으로는 풀 수 없다. 앞에서 빼고 뒤에서 넣는 큐와는 달리 스택은 뒤에서 빼고 뒤에서 넣기 때문에 스택 하나만으로는 순서를 재정렬 할 수 없다.

스택을 두 개(각각 s1, s2라 한다) 만들고 push할 때는 s1에만 넣어주다가 pop또는 peek할 때는 s2에서 꺼내는데, 만약 s2가 비어있다면 그 전에 s1에 있던 원소들을 차례로 꺼내어 s2에 넣어준다. 그럼 순서가 바뀌어 큐처럼 나오게 된다!

아래 표는 s1에 있던 원소들을 차례로 꺼내어 s2에 넣었을 때 순서가 바뀌는 모습을 나타낸 것이다.

|      | 1           | 2        | 3        | 4           |
| ---- | ----------- | -------- | -------- | ----------- |
| s1   | `[1, 2, 3]` | `[1, 2]` | `[1]`    | `[]`        |
| s2   | `[]`        | `[3]`    | `[3, 2]` | `[3, 2, 1]` |



```python
    def __init__(self):
        self.s1 = []
        self.s2 = []

    def push(self, x: int) -> None:
        self.s1.append(x)
        

    def pop(self) -> int:
				self.peek()
        return self.s2.pop()
        

    def peek(self) -> int:
        if not self.s2:
            while self.s1:
                self.s2.append(self.s1.pop())
            
        return self.s2[-1]



    def empty(self) -> bool:
        return not (self.s1 or self.s2)


```

