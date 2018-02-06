package multiplexer

import (
	"errors"
	"sync"
	"sync/atomic"
	"fmt"
)

var (
	ERR_EXIT = errors.New("exit")
)

type Socket interface {
	Read() ([]byte, error)
	Write([]byte) error
	Close()
}

type DataHandler interface {
	Process([]byte)
}

type ErrorHandler interface {
	OnError(error)
}

type IdentityHandler interface {
	GetIdentity([]byte) uint32
	SetIdentity([]byte, uint32)
}

type Connection struct {
	conn       Socket
	wg         sync.WaitGroup
	mutex      sync.Mutex
	applicants map[uint32]chan []byte
	chexit     chan bool
	chsend     chan []byte
	chch       chan chan []byte
	dh         DataHandler
	ih         IdentityHandler
	eh         ErrorHandler
	identity   uint32
}

func NewConnection(conn Socket, maxcount int, dh DataHandler, ih IdentityHandler, eh ErrorHandler) *Connection {
	count := maxcount
	if count < 1024 {
		count = 1024
	}
	chch := make(chan chan []byte, count)
	for i := 0; i < count; i++ {
		chch <- make(chan []byte, 1)
	}
	return &Connection{
		conn:       conn,
		applicants: make(map[uint32]chan []byte, count),
		chsend:     make(chan []byte, count),
		chexit:     make(chan bool),
		chch:       chch,
		dh:         dh,
		ih:         ih,
		eh:         eh,
	}
}


func (p *Connection) Start() {
	p.wg.Add(2)
	go func() {
		defer p.wg.Done()
		p.recv()
	}()
	go func() {
		defer p.wg.Done()
		p.send()
	}()
}

func (p *Connection) Close() {
	close(p.chexit)
	p.conn.Close()
	p.wg.Wait()
}

func (p *Connection) Query(data []byte) (res []byte, err error) {
	var ch chan []byte
	select {
	case <-p.chexit:
		return nil, ERR_EXIT
	case ch = <-p.chch:
		defer func() {
			p.chch <- ch
		}()
	}
	id := p.newIdentity()
	p.ih.SetIdentity(data, id)
	p.addApplicant(id, ch)
	defer func() {
		if err != nil {
			p.popApplicant(id)
		}
	}()
	if err := p.Write(data); err != nil {
		return nil, err
	}
	select {
	case <-p.chexit:
		return nil, ERR_EXIT
	case res = <-ch:
		break
	}
	return res, nil
}

func (p *Connection) Reply(query, answer []byte) error {
	// put back the identity attached to the query
	id := p.ih.GetIdentity(query)
	p.ih.SetIdentity(answer, id)
	return p.Write(answer)
}

func (p *Connection) Write(data []byte) error {
	select {
	case <-p.chexit:
		return ERR_EXIT
	case p.chsend <- data:
		break
	}
	return nil
}

func (p *Connection) send() {
	for {
		select {
		case <-p.chexit:
			return
		case data := <-p.chsend:
			if p.conn.Write(data) != nil {
				return
			}
		}
	}
}

func (p *Connection) recv() (err error) {
	defer func() {
		if err != nil {
			select {
			case <-p.chexit:
				err = nil
			default:
				p.eh.OnError(err)
			}
		}
	}()
	for {
		select {
		case <-p.chexit:
			return nil
		default:
			break
		}
		data, err := p.conn.Read()
		if err != nil {
			return err
		}
		if id := p.ih.GetIdentity(data); id > 0 {
			ch, ok := p.popApplicant(id)
			if ok {
				ch <- data
				continue
			}
		}
		p.dh.Process(data)
	}
	return nil
}

func (p *Connection) newIdentity() uint32 {
	return atomic.AddUint32(&p.identity, 1)
}

func (p *Connection) addApplicant(identity uint32, ch chan []byte) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.applicants[identity] = ch
}

func (p *Connection) popApplicant(identity uint32) (chan []byte, bool) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	ch, ok := p.applicants[identity]
	if !ok {
		return nil, false
	}
	delete(p.applicants, identity)
	return ch, true
}



func FmtErr(err error){
	fmt.Println(err)
}

func SetId(data []byte,id int)  {
      return
}

func GetId([]byte) int32  {
     return  1
}

