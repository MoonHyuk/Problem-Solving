# 원형 큐 디자인

## 풀이1) 메모리를 많이 쓰는 풀이

문제에서 큐를 사용하지 말고 구현해보라 했으니, 클래스 내부에 배열과 큐의 start와 end를 나타내는 투 포인터 변수를 만든다.

enQueue(x)를 하면 우선 큐가 꽉 차있는지 확인하고 차있지 않다면 end 포인터를 하나 증가시키고 배열에 x를 넣는다.

deQueue()를 하면 우선 큐가 비어있는지 확인하고 비어있지 않다면 start 포인터를 하나 증가시킨다.

큐가 비어있는지 여부는 start 포인터와 end 포인터가 같은 위치에 있는지로 확인할 수 있다.

큐가 꽉 차있는지 여부는 start 포인터와 end 포인터의 차이가 k와 같은지로 확인할 수 있다.

```python
class MyCircularQueue:

    def __init__(self, k: int):
        self.fullSize = k
        self.li = []
        self.start = 0
        self.end = 0

    def enQueue(self, value: int) -> bool:
        if self.isFull():
            return False
        
        self.end += 1
        self.li.append(value)
        return True
        
    def deQueue(self) -> bool:
        if self.isEmpty():
            return False
        
        self.start += 1
        return True

    def Front(self) -> int:
        if self.isEmpty():
            return -1
        
        return self.li[self.start]
        
    def Rear(self) -> int:
        if self.isEmpty():
            return -1
        
        return self.li[self.end - 1]

    def isEmpty(self) -> bool:
        return self.start == self.end    

    def isFull(self) -> bool:
        return self.end - self.start == self.fullSize
```

이 풀이는 deQueue를 할 때 리스트의 앞 부분을 메모리에서 지우는 것이 아니라서 deQueue, enQueue를 반복할 수록 메모리 사용량이 무한히 늘어난다. 이 풀이로도 leetcode에서는 충분히 통과가 가능하지만 메모리를 덜 쓰도록 개선된 방법이 필요하다.



## 풀이2) 메모리 사용량을 개선한 풀이

이 풀이는 원형 큐에서 포인터를 k(원형 큐의 사이즈)만큼 움직이면 다시 원래 자리로 돌아온다는 아이디어에서 착안한다.

배열의 크기를 k로 고정시키고 더이상 늘리지도 줄이지도 않으며 특정 자리가 비어있다는 것은 배열의 값을 `None`으로 하여 나타낸다.

```python
class MyCircularQueue:

    def __init__(self, k: int):
        self.fullSize = k
        self.li = [None] * k
        self.start = 0
        self.end = 0

    def enQueue(self, value: int) -> bool:
        if self.isFull():
            return False
        
        self.li[self.end] = value
        self.end = (self.end + 1) % self.fullSize
        
        return True
        
    def deQueue(self) -> bool:
        if self.isEmpty():
            return False
        
        self.li[self.start] = None
        self.start = (self.start + 1) % self.fullSize
        
        return True

    def Front(self) -> int:
        if self.isEmpty():
            return -1
        
        return self.li[self.start]
        
    def Rear(self) -> int:
        if self.isEmpty():
            return -1
        
        return self.li[self.end - 1]

    def isEmpty(self) -> bool:
        return self.start == self.end and self.li[self.start] is None    

    def isFull(self) -> bool:
        return self.start == self.end and self.li[self.start] is not None
```

첫번째 풀이와 유사하나 몇가지 다른 점은

- 포인터를 이동시킬 때 mod k를 한번 해줘야 한다는 점
- enQueue(x)시 배열에 append가 아니라 end 포인터기 가르키는 인덱스에 값을 넣는다는 점
- 큐가 비어있을 때와 큐가 꽉 찼을 때 모두 start와 end 포인터의 위치가 같다는 점. 따라서 추가적으로 `self.li[self.start]`가 None 인지 아닌지로 구분한다는 점

이다.