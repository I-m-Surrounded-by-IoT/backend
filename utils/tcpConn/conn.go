package tcpconn

import (
	"bufio"
	"errors"
	"net"
	"sync"
	"time"

	"github.com/zijiren233/stream"
)

const (
	SayHello = "hello"
)

type Conn struct {
	netConn net.Conn
	rl      sync.Mutex
	wl      sync.Mutex
	buf     *bufio.ReadWriter
	reader  *stream.Reader
	writer  *stream.Writer
}

func NewConn(netConn net.Conn) *Conn {
	c := &Conn{
		netConn: netConn,
		buf:     bufio.NewReadWriter(bufio.NewReader(netConn), bufio.NewWriter(netConn)),
	}
	c.reader = stream.NewReader(c.buf, stream.BigEndian)
	c.writer = stream.NewWriter(c.buf, stream.BigEndian)
	return c
}

func (c *Conn) Send(data []byte, deadline ...time.Time) error {
	c.wl.Lock()
	defer c.wl.Unlock()

	if len(deadline) > 0 {
		err := c.netConn.SetWriteDeadline(deadline[0])
		if err != nil {
			return err
		}
		defer func() {
			_ = c.netConn.SetWriteDeadline(time.Time{})
		}()
	} else {
		_ = c.netConn.SetWriteDeadline(time.Time{})
	}

	l := len(data)
	if l > 0x7fff {
		return errors.New("data too long")
	} else if l > 0x7f {
		err := c.writer.Byte(0x80 | uint8(l>>8)).Byte(uint8(l & 0xff)).Bytes(data).Error()
		if err != nil {
			return err
		}
	} else {
		err := c.writer.Byte(uint8(l)).Bytes(data).Error()
		if err != nil {
			return err
		}
	}

	return c.buf.Flush()
}

func (c *Conn) SayHello() error {
	s, err := c.NextMessage(time.Now().Add(time.Second * 5))
	if err != nil {
		return err
	}
	if string(s) != SayHello {
		return errors.New("invalid hello")
	}
	return c.Send(stream.StringToBytes(SayHello), time.Now().Add(time.Second*5))
}

func (c *Conn) ClientSayHello() error {
	err := c.Send(stream.StringToBytes(SayHello), time.Now().Add(time.Second*5))
	if err != nil {
		return err
	}
	s, err := c.NextMessage(time.Now().Add(time.Second * 5))
	if err != nil {
		return err
	}
	if string(s) != SayHello {
		return errors.New("invalid hello")
	}
	return nil
}

func (c *Conn) Close() error {
	defer c.netConn.Close()
	return c.buf.Flush()
}

func (c *Conn) NextMessage(deadline ...time.Time) ([]byte, error) {
	c.rl.Lock()
	defer c.rl.Unlock()

	if len(deadline) > 0 {
		err := c.netConn.SetReadDeadline(deadline[0])
		if err != nil {
			return nil, err
		}
		defer func() {
			_ = c.netConn.SetReadDeadline(time.Time{})
		}()
	} else {
		_ = c.netConn.SetReadDeadline(time.Time{})
	}
	var t int
	// 如果第一个字节的最高位为1，则表示后面还有一个字节，两个字节拼接为int16
	i, err := c.reader.ReadByte()
	if err != nil {
		return nil, err
	}
	if i>>7 == 1 {
		i2, err := c.reader.ReadByte()
		if err != nil {
			return nil, err
		}
		t = int(int16(i&0x7f)<<8) + int(i2)
	} else {
		t = int(i)
	}
	return c.reader.ReadBytes(t)
}
