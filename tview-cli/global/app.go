package global

import (
	m "pc-uptime/tview/models"

	"github.com/nats-io/nats.go"
	"github.com/rivo/tview"
)

type appGlobalHandler struct{}

func (a *appGlobalHandler) HandleFooterMsg(txt string) tview.Primitive {
	return drawAppFooterText(txt)
}

func NewGlobalElems() *appGlobalHandler {
	return &appGlobalHandler{}
}

func NewNatsConn(url string) *nats.Conn {
	nc, err := nats.Connect(url)
	if err != nil {
		panic(err)
	}
	return nc
}

func NewApp(natsURL string) *m.AppHolder {
	ah := &m.AppHolder{
		TviewApp:    tview.NewApplication(),
		AppFlex:     tview.NewFlex(),
		GlobalElems: NewGlobalElems(),
		NC:          NewNatsConn(natsURL),
	}
	setCurrentKeyHandler(ah)
	return ah
}
