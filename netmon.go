package main

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"time"
)

func value(iface, file string) uint64 {
	path := filepath.Join("/sys", "class", "net", iface, "statistics", file)
	b, err := ioutil.ReadFile(path)

	if err != nil {
		return 0
	}

	b = bytes.TrimSuffix(b, []byte("\n"))
	i, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return 0
	}

	return i
}

func rx(iface string) uint64 {
	return value(iface, "rx_bytes")
}

func tx(iface string) uint64 {
	return value(iface, "tx_bytes")
}

func speedTx(iface string) uint64 {
	txBefore := tx(iface)
	time.Sleep(time.Second)
	txAfter := tx(iface)

	if (txAfter == 0) || (txBefore == 0) {
		return 0
	}
	return txAfter - txBefore
}

func speedRx(iface string) uint64 {
	rxBefore := rx(iface)
	time.Sleep(time.Second)
	rxAfter := rx(iface)

	if (rxAfter == 0) || (rxBefore == 0) {
		return 0
	}
	return rxAfter - rxBefore
}
