package main

import (
	"io"
)

////////////////////分页对齐写入文件——golang实现////////////////

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

	// 当需要写入的字节数达到buffWatermarkBytes,则触发写入,该值小于len(buf),保证写入文件中的数据页对齐
	buffWatermarkBytes int
}

func NewPageWriter(w io.Writer, pageBytes, pageOffset int) *PageWriter {
	return &PageWriter{
		w:                  w,
		pageOffset:         pageOffset,
		pageBytes:          pageBytes,
		buf:                make([]byte, defaultBufferBytes+pageBytes), // 因页偏移,需多申请一页的空间
		buffWatermarkBytes: defaultBufferBytes,
	}
}

func (pw *PageWriter) Flush() error {
	_, err := pw.flush()
	return err
}

func (pw *PageWriter) FlushN() (int, error) {
	return pw.flush()
}

func (pw *PageWriter) flush() (int, error) {
	if pw.bufferedBytes == 0 {
		return 0, nil
	}
	n, err := pw.w.Write(pw.buf[:pw.bufferedBytes]) // 调整pageOffset,考虑此次flush可能并不是整页写入,因而要记录已写入的不完整页的字节数
	pw.pageOffset = (pw.pageOffset + pw.bufferedBytes) % pw.pageBytes
	pw.bufferedBytes = 0
	return n, err
}

func (pw *PageWriter) Write(p []byte) (n int, err error) {
	if len(p)+pw.bufferedBytes <= pw.buffWatermarkBytes {
		// 未超过,则写入buff
		copy(pw.buf[pw.bufferedBytes:], p)
		pw.bufferedBytes += len(p)
		return len(p), nil
	}
	// 计算出末页对齐的字节数
	slack := pw.pageBytes - ((pw.pageOffset + pw.bufferedBytes) % pw.pageBytes)
	if slack != pw.pageBytes {
		partial := slack > len(p)
		if partial {
			// 没有足够数据来对齐
			slack = len(p)
		}
		// Append数据
		copy(pw.buf[pw.bufferedBytes:], p[:slack])
		pw.bufferedBytes += slack
		n = slack
		p = p[slack:]
		if partial {
			//若还未达到对齐条件,则直接返回
			return 0, nil
		}
	}
	// buffer中已经是页对齐,可以写入文件
	if err = pw.Flush(); err != nil {
		return n, err
	}
	// 若p中字节数大于页大小,则直接写入文件
	if len(p) > pw.pageBytes {
		pages := len(p) / pw.pageBytes
		c, werr := pw.w.Write(p[:pages*pw.pageBytes])
		n += c
		if werr != nil {
			return n, werr
		}
		p = p[pages*pw.pageBytes:] // 剩余的部分
	}
	// 将剩余的部分写入buffer
	c, werr := pw.Write(p)
	n += c
	return n, werr

}
