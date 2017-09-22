// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package thriftgen

import (
	"bytes"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

// Attributes:
//  - Temperature
//  - EToday
//  - VPv1
//  - VPv2
//  - VPv3
//  - IPv1
//  - IPv2
//  - IPv3
//  - Iac1
//  - Iac2
//  - Iac3
//  - Vac1
//  - Vac2
//  - Vac3
//  - Fac
//  - Pac
//  - ETotal
//  - HTotal
//  - Error
//  - HistoryTime
//  - IP
//  - Model
type InvertorThriftBean struct {
	Temperature string `thrift:"temperature,1" json:"temperature"`
	EToday      string `thrift:"eToday,2" json:"eToday"`
	VPv1        string `thrift:"vPv1,3" json:"vPv1"`
	VPv2        string `thrift:"vPv2,4" json:"vPv2"`
	VPv3        string `thrift:"vPv3,5" json:"vPv3"`
	IPv1        string `thrift:"iPv1,6" json:"iPv1"`
	IPv2        string `thrift:"iPv2,7" json:"iPv2"`
	IPv3        string `thrift:"iPv3,8" json:"iPv3"`
	Iac1        string `thrift:"iac1,9" json:"iac1"`
	Iac2        string `thrift:"iac2,10" json:"iac2"`
	Iac3        string `thrift:"iac3,11" json:"iac3"`
	Vac1        string `thrift:"vac1,12" json:"vac1"`
	Vac2        string `thrift:"vac2,13" json:"vac2"`
	Vac3        string `thrift:"vac3,14" json:"vac3"`
	Fac         string `thrift:"fac,15" json:"fac"`
	Pac         string `thrift:"pac,16" json:"pac"`
	ETotal      string `thrift:"eTotal,17" json:"eTotal"`
	HTotal      string `thrift:"hTotal,18" json:"hTotal"`
	Error       string `thrift:"error,19" json:"error"`
	HistoryTime string `thrift:"historyTime,20" json:"historyTime"`
	IP          string `thrift:"ip,21" json:"ip"`
	Model       string `thrift:"model,22" json:"model"`
}

func NewInvertorThriftBean() *InvertorThriftBean {
	return &InvertorThriftBean{}
}

func (p *InvertorThriftBean) GetTemperature() string {
	return p.Temperature
}

func (p *InvertorThriftBean) GetEToday() string {
	return p.EToday
}

func (p *InvertorThriftBean) GetVPv1() string {
	return p.VPv1
}

func (p *InvertorThriftBean) GetVPv2() string {
	return p.VPv2
}

func (p *InvertorThriftBean) GetVPv3() string {
	return p.VPv3
}

func (p *InvertorThriftBean) GetIPv1() string {
	return p.IPv1
}

func (p *InvertorThriftBean) GetIPv2() string {
	return p.IPv2
}

func (p *InvertorThriftBean) GetIPv3() string {
	return p.IPv3
}

func (p *InvertorThriftBean) GetIac1() string {
	return p.Iac1
}

func (p *InvertorThriftBean) GetIac2() string {
	return p.Iac2
}

func (p *InvertorThriftBean) GetIac3() string {
	return p.Iac3
}

func (p *InvertorThriftBean) GetVac1() string {
	return p.Vac1
}

func (p *InvertorThriftBean) GetVac2() string {
	return p.Vac2
}

func (p *InvertorThriftBean) GetVac3() string {
	return p.Vac3
}

func (p *InvertorThriftBean) GetFac() string {
	return p.Fac
}

func (p *InvertorThriftBean) GetPac() string {
	return p.Pac
}

func (p *InvertorThriftBean) GetETotal() string {
	return p.ETotal
}

func (p *InvertorThriftBean) GetHTotal() string {
	return p.HTotal
}

func (p *InvertorThriftBean) GetError() string {
	return p.Error
}

func (p *InvertorThriftBean) GetHistoryTime() string {
	return p.HistoryTime
}

func (p *InvertorThriftBean) GetIP() string {
	return p.IP
}

func (p *InvertorThriftBean) GetModel() string {
	return p.Model
}
func (p *InvertorThriftBean) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.readField6(iprot); err != nil {
				return err
			}
		case 7:
			if err := p.readField7(iprot); err != nil {
				return err
			}
		case 8:
			if err := p.readField8(iprot); err != nil {
				return err
			}
		case 9:
			if err := p.readField9(iprot); err != nil {
				return err
			}
		case 10:
			if err := p.readField10(iprot); err != nil {
				return err
			}
		case 11:
			if err := p.readField11(iprot); err != nil {
				return err
			}
		case 12:
			if err := p.readField12(iprot); err != nil {
				return err
			}
		case 13:
			if err := p.readField13(iprot); err != nil {
				return err
			}
		case 14:
			if err := p.readField14(iprot); err != nil {
				return err
			}
		case 15:
			if err := p.readField15(iprot); err != nil {
				return err
			}
		case 16:
			if err := p.readField16(iprot); err != nil {
				return err
			}
		case 17:
			if err := p.readField17(iprot); err != nil {
				return err
			}
		case 18:
			if err := p.readField18(iprot); err != nil {
				return err
			}
		case 19:
			if err := p.readField19(iprot); err != nil {
				return err
			}
		case 20:
			if err := p.readField20(iprot); err != nil {
				return err
			}
		case 21:
			if err := p.readField21(iprot); err != nil {
				return err
			}
		case 22:
			if err := p.readField22(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *InvertorThriftBean) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Temperature = v
	}
	return nil
}

func (p *InvertorThriftBean) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.EToday = v
	}
	return nil
}

func (p *InvertorThriftBean) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.VPv1 = v
	}
	return nil
}

func (p *InvertorThriftBean) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.VPv2 = v
	}
	return nil
}

func (p *InvertorThriftBean) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.VPv3 = v
	}
	return nil
}

func (p *InvertorThriftBean) readField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.IPv1 = v
	}
	return nil
}

func (p *InvertorThriftBean) readField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		p.IPv2 = v
	}
	return nil
}

func (p *InvertorThriftBean) readField8(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 8: ", err)
	} else {
		p.IPv3 = v
	}
	return nil
}

func (p *InvertorThriftBean) readField9(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 9: ", err)
	} else {
		p.Iac1 = v
	}
	return nil
}

func (p *InvertorThriftBean) readField10(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 10: ", err)
	} else {
		p.Iac2 = v
	}
	return nil
}

func (p *InvertorThriftBean) readField11(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 11: ", err)
	} else {
		p.Iac3 = v
	}
	return nil
}

func (p *InvertorThriftBean) readField12(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 12: ", err)
	} else {
		p.Vac1 = v
	}
	return nil
}

func (p *InvertorThriftBean) readField13(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 13: ", err)
	} else {
		p.Vac2 = v
	}
	return nil
}

func (p *InvertorThriftBean) readField14(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 14: ", err)
	} else {
		p.Vac3 = v
	}
	return nil
}

func (p *InvertorThriftBean) readField15(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 15: ", err)
	} else {
		p.Fac = v
	}
	return nil
}

func (p *InvertorThriftBean) readField16(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 16: ", err)
	} else {
		p.Pac = v
	}
	return nil
}

func (p *InvertorThriftBean) readField17(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 17: ", err)
	} else {
		p.ETotal = v
	}
	return nil
}

func (p *InvertorThriftBean) readField18(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 18: ", err)
	} else {
		p.HTotal = v
	}
	return nil
}

func (p *InvertorThriftBean) readField19(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 19: ", err)
	} else {
		p.Error = v
	}
	return nil
}

func (p *InvertorThriftBean) readField20(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 20: ", err)
	} else {
		p.HistoryTime = v
	}
	return nil
}

func (p *InvertorThriftBean) readField21(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 21: ", err)
	} else {
		p.IP = v
	}
	return nil
}

func (p *InvertorThriftBean) readField22(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 22: ", err)
	} else {
		p.Model = v
	}
	return nil
}

func (p *InvertorThriftBean) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("InvertorThriftBean"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := p.writeField6(oprot); err != nil {
		return err
	}
	if err := p.writeField7(oprot); err != nil {
		return err
	}
	if err := p.writeField8(oprot); err != nil {
		return err
	}
	if err := p.writeField9(oprot); err != nil {
		return err
	}
	if err := p.writeField10(oprot); err != nil {
		return err
	}
	if err := p.writeField11(oprot); err != nil {
		return err
	}
	if err := p.writeField12(oprot); err != nil {
		return err
	}
	if err := p.writeField13(oprot); err != nil {
		return err
	}
	if err := p.writeField14(oprot); err != nil {
		return err
	}
	if err := p.writeField15(oprot); err != nil {
		return err
	}
	if err := p.writeField16(oprot); err != nil {
		return err
	}
	if err := p.writeField17(oprot); err != nil {
		return err
	}
	if err := p.writeField18(oprot); err != nil {
		return err
	}
	if err := p.writeField19(oprot); err != nil {
		return err
	}
	if err := p.writeField20(oprot); err != nil {
		return err
	}
	if err := p.writeField21(oprot); err != nil {
		return err
	}
	if err := p.writeField22(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *InvertorThriftBean) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("temperature", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:temperature: ", p), err)
	}
	if err := oprot.WriteString(string(p.Temperature)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.temperature (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:temperature: ", p), err)
	}
	return err
}

func (p *InvertorThriftBean) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("eToday", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:eToday: ", p), err)
	}
	if err := oprot.WriteString(string(p.EToday)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.eToday (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:eToday: ", p), err)
	}
	return err
}

func (p *InvertorThriftBean) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("vPv1", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:vPv1: ", p), err)
	}
	if err := oprot.WriteString(string(p.VPv1)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.vPv1 (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:vPv1: ", p), err)
	}
	return err
}

func (p *InvertorThriftBean) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("vPv2", thrift.STRING, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:vPv2: ", p), err)
	}
	if err := oprot.WriteString(string(p.VPv2)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.vPv2 (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:vPv2: ", p), err)
	}
	return err
}

func (p *InvertorThriftBean) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("vPv3", thrift.STRING, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:vPv3: ", p), err)
	}
	if err := oprot.WriteString(string(p.VPv3)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.vPv3 (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:vPv3: ", p), err)
	}
	return err
}

func (p *InvertorThriftBean) writeField6(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("iPv1", thrift.STRING, 6); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:iPv1: ", p), err)
	}
	if err := oprot.WriteString(string(p.IPv1)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.iPv1 (6) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 6:iPv1: ", p), err)
	}
	return err
}

func (p *InvertorThriftBean) writeField7(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("iPv2", thrift.STRING, 7); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:iPv2: ", p), err)
	}
	if err := oprot.WriteString(string(p.IPv2)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.iPv2 (7) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 7:iPv2: ", p), err)
	}
	return err
}

func (p *InvertorThriftBean) writeField8(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("iPv3", thrift.STRING, 8); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 8:iPv3: ", p), err)
	}
	if err := oprot.WriteString(string(p.IPv3)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.iPv3 (8) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 8:iPv3: ", p), err)
	}
	return err
}

func (p *InvertorThriftBean) writeField9(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("iac1", thrift.STRING, 9); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 9:iac1: ", p), err)
	}
	if err := oprot.WriteString(string(p.Iac1)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.iac1 (9) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 9:iac1: ", p), err)
	}
	return err
}

func (p *InvertorThriftBean) writeField10(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("iac2", thrift.STRING, 10); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 10:iac2: ", p), err)
	}
	if err := oprot.WriteString(string(p.Iac2)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.iac2 (10) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 10:iac2: ", p), err)
	}
	return err
}

func (p *InvertorThriftBean) writeField11(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("iac3", thrift.STRING, 11); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 11:iac3: ", p), err)
	}
	if err := oprot.WriteString(string(p.Iac3)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.iac3 (11) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 11:iac3: ", p), err)
	}
	return err
}

func (p *InvertorThriftBean) writeField12(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("vac1", thrift.STRING, 12); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 12:vac1: ", p), err)
	}
	if err := oprot.WriteString(string(p.Vac1)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.vac1 (12) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 12:vac1: ", p), err)
	}
	return err
}

func (p *InvertorThriftBean) writeField13(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("vac2", thrift.STRING, 13); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 13:vac2: ", p), err)
	}
	if err := oprot.WriteString(string(p.Vac2)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.vac2 (13) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 13:vac2: ", p), err)
	}
	return err
}

func (p *InvertorThriftBean) writeField14(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("vac3", thrift.STRING, 14); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 14:vac3: ", p), err)
	}
	if err := oprot.WriteString(string(p.Vac3)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.vac3 (14) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 14:vac3: ", p), err)
	}
	return err
}

func (p *InvertorThriftBean) writeField15(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("fac", thrift.STRING, 15); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 15:fac: ", p), err)
	}
	if err := oprot.WriteString(string(p.Fac)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.fac (15) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 15:fac: ", p), err)
	}
	return err
}

func (p *InvertorThriftBean) writeField16(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("pac", thrift.STRING, 16); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 16:pac: ", p), err)
	}
	if err := oprot.WriteString(string(p.Pac)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.pac (16) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 16:pac: ", p), err)
	}
	return err
}

func (p *InvertorThriftBean) writeField17(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("eTotal", thrift.STRING, 17); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 17:eTotal: ", p), err)
	}
	if err := oprot.WriteString(string(p.ETotal)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.eTotal (17) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 17:eTotal: ", p), err)
	}
	return err
}

func (p *InvertorThriftBean) writeField18(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("hTotal", thrift.STRING, 18); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 18:hTotal: ", p), err)
	}
	if err := oprot.WriteString(string(p.HTotal)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.hTotal (18) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 18:hTotal: ", p), err)
	}
	return err
}

func (p *InvertorThriftBean) writeField19(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("error", thrift.STRING, 19); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 19:error: ", p), err)
	}
	if err := oprot.WriteString(string(p.Error)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.error (19) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 19:error: ", p), err)
	}
	return err
}

func (p *InvertorThriftBean) writeField20(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("historyTime", thrift.STRING, 20); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 20:historyTime: ", p), err)
	}
	if err := oprot.WriteString(string(p.HistoryTime)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.historyTime (20) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 20:historyTime: ", p), err)
	}
	return err
}

func (p *InvertorThriftBean) writeField21(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("ip", thrift.STRING, 21); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 21:ip: ", p), err)
	}
	if err := oprot.WriteString(string(p.IP)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.ip (21) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 21:ip: ", p), err)
	}
	return err
}

func (p *InvertorThriftBean) writeField22(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("model", thrift.STRING, 22); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 22:model: ", p), err)
	}
	if err := oprot.WriteString(string(p.Model)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.model (22) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 22:model: ", p), err)
	}
	return err
}

func (p *InvertorThriftBean) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("InvertorThriftBean(%+v)", *p)
}