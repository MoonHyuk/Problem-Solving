# 유효한 괄호

## 풀이1) 스택 이용

우선 빈 스택을 만들어준다.

문자열 s를 순회하면서 현재 문자가 여는 괄호이면 현재 문자를 스택에 넣는다.

닫는 괄호이면

- 현재 스택이 비어있거나
- 스택을 pop했을 때 나오는 문자가 현재 문자와 쌍이 아니면

False를 반환한다.

문자열 순회가 끝난 후 스택이 비어있지 않으면 False를 반환하고, 비어있으면 True를 반환한다.

```python
def isValid(self, s: str) -> bool:
    p = {
        '(': ')',
        '[': ']',
        '{': '}'
    }
    
    stack = []
    for c in s:
        if c in p:
            stack.append(c)
        elif len(stack) == 0 or p[stack.pop()] != c:
            return False
    
    if len(stack) != 0:
        return False
    
    return True
```