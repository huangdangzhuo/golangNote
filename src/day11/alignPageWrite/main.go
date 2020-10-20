package main

import "io"

// 以页为单元写入文件,或者通过flush写入
type PageWriter struct {
	w io.Writer
	// pageOffset为当前已写入页的偏移量(相对于buff的起始地址),因有可能非完整的页已写入文件(flush方式),所以新的写入应该先补齐这非完整的部分
	pageOffset int

	// 每页的字节数
	pageBytes int

	// buff中等待写入的字节数,bufferedBytes = len(buf)
	bufferedBytes int

	// buffer
	buf []byte

	// 当需要写入的字节
}