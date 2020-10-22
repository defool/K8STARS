//Package Tars comment
// This file war generated by tars2go 1.1
// Generated from tarsregistry.tars
package Tars

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

//KeepAliveReq strcut implement
type KeepAliveReq struct {
	NodeName string `json:"nodeName"`
	State    string `json:"state"`
	Application string `json:"application"`
	Server      string `json:"server"`
	SetID       string `json:"setID"`
}

func (st *KeepAliveReq) resetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *KeepAliveReq) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err = _is.Read_string(&st.NodeName, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.State, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Application, 2, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Server, 3, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.SetID, 4, false)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *KeepAliveReq) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require KeepAliveReq, but not exist. tag %d", tag)
		}
		return nil

	}

	st.ReadFrom(_is)

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *KeepAliveReq) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.NodeName, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.State, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Application, 2)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Server, 3)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.SetID, 4)
	if err != nil {
		return err
	}

	return nil
}

//WriteBlock encode struct
func (st *KeepAliveReq) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	st.WriteTo(_os)

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}
