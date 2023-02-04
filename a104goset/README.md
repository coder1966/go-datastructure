# go-set
Set implementation for go
```azure
    
一些集合需要使用的方法

add(element):向集合添加一个元素
delete(element)：删除集合里的某个元素
has(element):判断集合里是否有某个元素
clear(): 清空集合（移除集合里的所有元素）
size(): 返回集合所包含元素的数量
values():返回一个包含集合所有元素的数组

A∪B并集union ：对于给定的俩个集合，返回一个包含俩个集合中所有元素的新集合
A∩B 交集intersection ：对于给定的俩个集合，返回一个包含俩个集合中所有共有元素的新集合
A - B差集difference ：对于给定的俩个集合，返回一个包含所有存在第一个集合且不存在第二个集合的元素的新集合
A > B子集ifSub ：验证一个给定的集合是否是另一个集合的子集

```