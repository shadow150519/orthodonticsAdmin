package model

import (
	"database/sql/driver"
	"strconv"
)

type ReservationState string
const (
	UNFINISHED ReservationState = "0"
	FINISHED   ReservationState = "1"
)

// Value 写入数据库之前，转为string
func (rs ReservationState) Value()( driver.Value, error) {
	state := string(rs)
	return state, nil
}

// Scan 读出数据库时，转为ReservationState
func (rs *ReservationState) Scan(v interface{})(error) {
	vtSlice := v.([]uint8)
	vt := int(vtSlice[0])
	vt = vt - 48
	*rs = ReservationState(strconv.Itoa(vt))
	return nil
}