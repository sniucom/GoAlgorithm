package hashmap

import "hash/crc32"

type Bitmap struct {
	bits  []byte
	limit int
}

func NewBitmap(limit int) *Bitmap {
	bits := make([]byte, (limit>>3)+1)
	return &Bitmap{
		bits,
		limit,
	}
}

// 存储
func (bitmap *Bitmap) Add(data string) {
	code := bitmap.HashCode(data)
	num := code & (uint32(bitmap.limit) - 1) // 对hashcode取余
	index := num >> 3
	position := num & 0x07
	bitmap.bits[index] |= 1 << position
}

// 是否存在
func (bitmap *Bitmap) Exists(data string) bool {
	code := bitmap.HashCode(data)
	num := code & (uint32(bitmap.limit) - 1) // 对hashcode取余
	index := num >> 3
	position := num & 0x07
	return bitmap.bits[index]&(1<<position) != 0
}

// 删除
func (bitmap *Bitmap) Remove(data string) {
	code := bitmap.HashCode(data)
	num := code & (uint32(bitmap.limit) - 1) // 对hashcode取余
	index := num >> 3
	position := num & 0x07
	bitmap.bits[index] = bitmap.bits[index] & ^(1 << position)
}

func (bitmap *Bitmap) HashCode(data string) uint32 {
	return uint32(crc32.ChecksumIEEE([]byte(data)))
}
