package hashtablearray

import (
	"crypto/sha256"
	"fmt"
)

const (
	Deleted      = iota // 数据已经被删除
	MinTableSize = 100  // 哈希表大小
	legimate     = iota // 已经穿在的合法数据
	Empty        = iota // 数据为空
)

// 函数指针|函数签名
type HashFunc func(data interface{}, tableSize int) int

// 哈希表元素
type HashEntry struct {
	data interface{} // 数据
	kind int         // 数据类型： 数据已经被删除 已经穿在的合法数据 数据为空
}

// 哈希表
type HashTable struct {
	tableSize int          //
	cells     []*HashEntry // 数组，每一个元素是指针，指向一个哈希结构
	hashFunc  HashFunc     // 哈希函数
}

// 接口
type HashTableGO interface {
	Find(data interface{}) int      // 查找
	Insert(data interface{})        //
	Clear()                         // 清空
	GetValue(index int) interface{} //
}

func New(size int, hashFunc HashFunc) (*HashTable, error) {
	if size < MinTableSize {
		return nil, fmt.Errorf("哈希表太小")
	}
	if hashFunc == nil {
		return nil, fmt.Errorf("缺少哈希函数")
	}
	hashTable := new(HashTable)
	hashTable.tableSize = size
	hashTable.cells = make([]*HashEntry, size) // 数组分配内存
	hashTable.hashFunc = hashFunc
	// 开辟单元格
	for i := 0; i < hashTable.tableSize; i++ {
		hashTable.cells[i] = new(HashEntry)
		hashTable.cells[i].data = nil
		hashTable.cells[i].kind = Empty
	}

	return hashTable, nil
}

// 实现接口的方法
func (ht *HashTable) Find(data interface{}) int {
	var collId int = 0
	curPos := ht.hashFunc(data, ht.tableSize) // 计算哈希位置
	if ht.cells[curPos].kind != Empty && ht.cells[curPos].data != data {
		collId++
		curPos = 2*curPos - 1 // 平方探测
		if curPos > ht.tableSize {
			curPos -= ht.tableSize // 越界处理
		}
	}
	return curPos
} // 查找
func (ht *HashTable) Insert(data interface{}) {
	pos := ht.Find(data)   // 查找位置，找不到咋办？平方探测一定能找到
	entry := ht.cells[pos] // 插入独居记录状态
	if entry.kind != legimate {
		entry.kind = legimate
		entry.data = data // 插入了数据
	}
} //
func (ht *HashTable) Clear() {
	// 循环清空数据
	for i := 0; i < ht.tableSize; i++ {
		if ht.cells[i] == nil {
			continue
		}
		// ht.cells[i].data = nil
		ht.cells[i].kind = Deleted
	}
} // 清空
func (ht *HashTable) GetValue(index int) interface{} {
	if index > ht.tableSize {
		return nil
	}
	entry := ht.cells[index] // 取出数据
	if entry.kind == legimate {
		return entry.data
	}
	return nil
} //

// MySHA 哈希函数
func MySHA(str interface{}, tableSize int) int {
	var hash int = 0
	var chars []byte

	if strings, ok := str.(string); ok {
		chars = []byte(strings) // 字符串转化字节数组
	}
	for _, v := range chars {
		hash = (hash << 17) + int(v) // 哈希算法
		// hash = (hash<<17 | 123) + int(v)             // 哈希算法 与
		// hash = (hash<<17 | 123&12345) + int(v)       // 哈希算法 或
		// hash = (hash<<17 | 123&12345 ^ 567) + int(v) // 哈希算法 异或
	}
	return hash % MinTableSize
}

// MySHA256 真正的哈希函数
func MySHA256(str string, tableSize int) int {
	shaObj := sha256.New()
	shaObj.Write([]byte(str))

	chars := shaObj.Sum(nil)

	var hash int = 0

	for _, v := range chars {
		hash = (hash << 17) + int(v) // 哈希算法
		// hash = (hash<<17 | 123) + int(v)             // 哈希算法 与
		// hash = (hash<<17 | 123&12345) + int(v)       // 哈希算法 或
		// hash = (hash<<17 | 123&12345 ^ 567) + int(v) // 哈希算法 异或
	}
	return hash % MinTableSize
}
