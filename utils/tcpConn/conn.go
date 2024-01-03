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

	err := c.writer.U16(uint16(len(data))).Bytes(data).Error()
	if err != nil {
		return err
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
	return c.Send([]byte(SayHello), time.Now().Add(time.Second*5))
}

func (c *Conn) ClientSayHello() error {
	err := c.Send([]byte(SayHello), time.Now().Add(time.Second*5))
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
	t, err := c.reader.ReadU16()
	if err != nil {
		return nil, err
	}
	return c.reader.ReadBytes(int(t))
}
