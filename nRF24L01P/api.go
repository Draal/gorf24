package nRF24L01P

import (
	"github.com/galaktor/gorf24/reg/addrw"
	"github.com/galaktor/gorf24/reg/autoack"
	"github.com/galaktor/gorf24/reg/config"
	"github.com/galaktor/gorf24/reg/dynpd"
	"github.com/galaktor/gorf24/reg/enrxaddr"
	"github.com/galaktor/gorf24/reg/feature"
	"github.com/galaktor/gorf24/reg/fifo"
	"github.com/galaktor/gorf24/reg/retrans"
	"github.com/galaktor/gorf24/reg/rfchan"
	"github.com/galaktor/gorf24/reg/rfsetup"
	"github.com/galaktor/gorf24/reg/rpd"
	"github.com/galaktor/gorf24/reg/rxpw"
	"github.com/galaktor/gorf24/reg/status"
	"github.com/galaktor/gorf24/reg/txobs"
	"github.com/galaktor/gorf24/reg/xaddr"
)

type ModeSwitcher interface {
	Standby() StandbyMode
	Tx() TxMode
	Rx() RxMode
}

/* STANDBY

   Only mode in which one can write registers. */
type StandbyMode interface {
	ModeSwitcher
	StandbyRegGetter
	StandbyRegSetter
}

/* for immutability, provides values, not pointers */
type StandbyRegGetter interface {
	GetConfig() (config.C, error)
	GetAutoAck() (autoack.AA, error)
	GetEnRxAddr() (enrxaddr.E, error)
	GetAddrWid() (addrw.AW, error)
	GetRetrans() (retrans.R, error)
	GetRfchan() (rfchan.R, error)
	GetRfsetup() (rfsetup.R, error)
	GetStatus() (status.S, error)
	GetTrans() (txobs.O, error)
	// reading RPD only makes sense when in Rx mode
	GetRxAddrP0() (xaddr.Full, error)
	GetRxAddrP1() (xaddr.Full, error)
	GetRxAddrP2() (xaddr.Partial, error)
	GetRxAddrP3() (xaddr.Partial, error)
	GetRxAddrP4() (xaddr.Partial, error)
	GetRxAddrP5() (xaddr.Partial, error)
	GetTxAddr() (xaddr.Full, error)
	GetRxPwP0() (rxpw.W, error)
	GetRxPwP1() (rxpw.W, error)
	GetRxPwP2() (rxpw.W, error)
	GetRxPwP3() (rxpw.W, error)
	GetRxPwP4() (rxpw.W, error)
	GetRxPwP5() (rxpw.W, error)
	GetFifo() (fifo.F, error)
	GetDynpd() (dynpd.DP, error)
	GetFeat() (feature.F, error)
}

type StandbyRegSetter interface {
	SetConfig(f func(c *config.C)) error
	SetAutoAck(f func(a *autoack.AA)) error
	SetEnRxAddr(f func(e *enrxaddr.E)) error
	SetAddrWid(f func(a *addrw.AW)) error
	SetRetrans(f func(r *retrans.R)) error
	SetRfchan(f func(r *rfchan.R)) error
	SetRfsetup(f func(r *rfsetup.R)) error
	SetStatus(f func(s *status.S)) error
	SetTrans(f func(t *txobs.O)) error
	SetRpd(f func(r *rpd.R)) error
	SetRxAddrP0(f func(a *xaddr.Full)) error
	SetRxAddrP1(f func(a *xaddr.Full)) error
	SetRxAddrP2(f func(a *xaddr.Partial)) error
	SetRxAddrP3(f func(a *xaddr.Partial)) error
	SetRxAddrP4(f func(a *xaddr.Partial)) error
	SetRxAddrP5(f func(a *xaddr.Partial)) error
	SetTxAddr(f func(a *xaddr.Full)) error
	SetRxPwP0(f func(w *rxpw.W)) error
	SetRxPwP1(f func(w *rxpw.W)) error
	SetRxPwP2(f func(w *rxpw.W)) error
	SetRxPwP3(f func(w *rxpw.W)) error
	SetRxPwP4(f func(w *rxpw.W)) error
	SetRxPwP5(f func(w *rxpw.W)) error
	SetFifo(f func(*fifo.F)) error
	SetDynpd(f func(d *dynpd.DP)) error
	SetFeat(f func(f *feature.F)) error
}

/* TX

   Cannot write registers. */
type TxMode interface {
	ModeSwitcher
	TxRegGetter
}

/* for immutability, provides values, not pointers */
type TxRegGetter interface {
	StandbyRegGetter
	// not sure if needs special interface...
}

/* RX

   Cannot write registers. */
type RxMode interface {
	ModeSwitcher
	RxRegGetter
}

/* for immutability, provides values, not pointers */
type RxRegGetter interface {
	StandbyRegGetter

	GetRpd() (rpd.R, error) // was 'CD' before nRF24L01+
}
