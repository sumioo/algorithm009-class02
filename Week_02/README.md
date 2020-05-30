# week2 学习笔记

## 哈希表

哈希表是一种快速查找的数据结构，理想情况下查找和插入都为O(1)。哈希即哈希函数，可以将任意长度的键转化为固定长度的值，在哈希表中一般转化为整数。表即数组，可以随机访问。

### 查找过程
键通过哈希函数转化为数组索引，得到索引后便可随机访问数组内容。但不同的键可能得到相同的数组索引（数组空间是有限的），这时便需要进行碰撞处理

### 碰撞处理

#### 拉链法
数组元素的值指向一个链表，链表节点保存key、value。

数据结构：数组 + 链表

#### 开放地址（线性探测法）

开放地址的具体思想就是与其将内存用作链表，不如将它们作为在散列表的空元素，空元素可以作为查找结束的标志。

数据结构：key数组 + value数组


## HashMap
HashMap主要属性有
* buckets 数组，即表
* loadFactor 转载因子 threshold = (int) (initialCapacity * loadFactor);
* threshold rehash的临界值
* size 表大小

HashMap解决碰撞的方法是拉链法，数组元素保存的是HashEntry<K, V>对象。在计算key的哈希值时会调用key的hashCode方法，然后将返回值%len(bukets)的出数组索引。
水人理论上采用拉链法解决碰撞的哈希表是不需要rehash的，但是为了避免性能下降HashMap会在插入一个新键的时候进行size++,当size大于threshold将触发rehash。

### get方法
* 调用hash方法得出数组索引
* 数组随机访问取得HashEntry<K, V>
* 沿着链表检索key

### put方法
* 调用hash方法得出数组索引
* 数组随机访问取得HashEntry<K, V>
* 如果HashEntry不为空，则在链表尾部插入新的key
* 如果HashEntry为空，则先检查是否需要rehash，然后再addEntry
