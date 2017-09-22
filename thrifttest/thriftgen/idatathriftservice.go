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

type IDataThriftService interface {
	// Parameters:
	//  - GatewaySn
	//  - InvSn
	//  - StartDate
	//  - EndDate
	//  - TimeZone
	GetInvertorBeanByDate(gatewaySn string, invSn string, startDate string, endDate string, timeZone int32) (r []*InvertorThriftBean, err error)
	Echo() (r string, err error)
}

type IDataThriftServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewIDataThriftServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *IDataThriftServiceClient {
	return &IDataThriftServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewIDataThriftServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *IDataThriftServiceClient {
	return &IDataThriftServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - GatewaySn
//  - InvSn
//  - StartDate
//  - EndDate
//  - TimeZone
func (p *IDataThriftServiceClient) GetInvertorBeanByDate(gatewaySn string, invSn string, startDate string, endDate string, timeZone int32) (r []*InvertorThriftBean, err error) {
	if err = p.sendGetInvertorBeanByDate(gatewaySn, invSn, startDate, endDate, timeZone); err != nil {
		return
	}
	return p.recvGetInvertorBeanByDate()
}

func (p *IDataThriftServiceClient) sendGetInvertorBeanByDate(gatewaySn string, invSn string, startDate string, endDate string, timeZone int32) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("getInvertorBeanByDate", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := IDataThriftServiceGetInvertorBeanByDateArgs{
		GatewaySn: gatewaySn,
		InvSn:     invSn,
		StartDate: startDate,
		EndDate:   endDate,
		TimeZone:  timeZone,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *IDataThriftServiceClient) recvGetInvertorBeanByDate() (value []*InvertorThriftBean, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "getInvertorBeanByDate" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "getInvertorBeanByDate failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getInvertorBeanByDate failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error1 error
		error1, err = error0.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error1
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "getInvertorBeanByDate failed: invalid message type")
		return
	}
	result := IDataThriftServiceGetInvertorBeanByDateResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

func (p *IDataThriftServiceClient) Echo() (r string, err error) {
	if err = p.sendEcho(); err != nil {
		return
	}
	return p.recvEcho()
}

func (p *IDataThriftServiceClient) sendEcho() (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("echo", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := IDataThriftServiceEchoArgs{}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *IDataThriftServiceClient) recvEcho() (value string, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "echo" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "echo failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "echo failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error3 error
		error3, err = error2.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error3
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "echo failed: invalid message type")
		return
	}
	result := IDataThriftServiceEchoResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type IDataThriftServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      IDataThriftService
}

func (p *IDataThriftServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *IDataThriftServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *IDataThriftServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewIDataThriftServiceProcessor(handler IDataThriftService) *IDataThriftServiceProcessor {

	self4 := &IDataThriftServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self4.processorMap["getInvertorBeanByDate"] = &iDataThriftServiceProcessorGetInvertorBeanByDate{handler: handler}
	self4.processorMap["echo"] = &iDataThriftServiceProcessorEcho{handler: handler}
	return self4
}

func (p *IDataThriftServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x5 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x5.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x5

}

type iDataThriftServiceProcessorGetInvertorBeanByDate struct {
	handler IDataThriftService
}

func (p *iDataThriftServiceProcessorGetInvertorBeanByDate) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := IDataThriftServiceGetInvertorBeanByDateArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getInvertorBeanByDate", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := IDataThriftServiceGetInvertorBeanByDateResult{}
	var retval []*InvertorThriftBean
	var err2 error
	if retval, err2 = p.handler.GetInvertorBeanByDate(args.GatewaySn, args.InvSn, args.StartDate, args.EndDate, args.TimeZone); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getInvertorBeanByDate: "+err2.Error())
		oprot.WriteMessageBegin("getInvertorBeanByDate", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getInvertorBeanByDate", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type iDataThriftServiceProcessorEcho struct {
	handler IDataThriftService
}

func (p *iDataThriftServiceProcessorEcho) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := IDataThriftServiceEchoArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("echo", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := IDataThriftServiceEchoResult{}
	var retval string
	var err2 error
	if retval, err2 = p.handler.Echo(); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing echo: "+err2.Error())
		oprot.WriteMessageBegin("echo", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("echo", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - GatewaySn
//  - InvSn
//  - StartDate
//  - EndDate
//  - TimeZone
type IDataThriftServiceGetInvertorBeanByDateArgs struct {
	GatewaySn string `thrift:"gatewaySn,1" json:"gatewaySn"`
	InvSn     string `thrift:"invSn,2" json:"invSn"`
	StartDate string `thrift:"startDate,3" json:"startDate"`
	EndDate   string `thrift:"endDate,4" json:"endDate"`
	TimeZone  int32  `thrift:"timeZone,5" json:"timeZone"`
}

func NewIDataThriftServiceGetInvertorBeanByDateArgs() *IDataThriftServiceGetInvertorBeanByDateArgs {
	return &IDataThriftServiceGetInvertorBeanByDateArgs{}
}

func (p *IDataThriftServiceGetInvertorBeanByDateArgs) GetGatewaySn() string {
	return p.GatewaySn
}

func (p *IDataThriftServiceGetInvertorBeanByDateArgs) GetInvSn() string {
	return p.InvSn
}

func (p *IDataThriftServiceGetInvertorBeanByDateArgs) GetStartDate() string {
	return p.StartDate
}

func (p *IDataThriftServiceGetInvertorBeanByDateArgs) GetEndDate() string {
	return p.EndDate
}

func (p *IDataThriftServiceGetInvertorBeanByDateArgs) GetTimeZone() int32 {
	return p.TimeZone
}
func (p *IDataThriftServiceGetInvertorBeanByDateArgs) Read(iprot thrift.TProtocol) error {
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

func (p *IDataThriftServiceGetInvertorBeanByDateArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.GatewaySn = v
	}
	return nil
}

func (p *IDataThriftServiceGetInvertorBeanByDateArgs) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.InvSn = v
	}
	return nil
}

func (p *IDataThriftServiceGetInvertorBeanByDateArgs) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.StartDate = v
	}
	return nil
}

func (p *IDataThriftServiceGetInvertorBeanByDateArgs) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.EndDate = v
	}
	return nil
}

func (p *IDataThriftServiceGetInvertorBeanByDateArgs) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.TimeZone = v
	}
	return nil
}

func (p *IDataThriftServiceGetInvertorBeanByDateArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getInvertorBeanByDate_args"); err != nil {
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
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *IDataThriftServiceGetInvertorBeanByDateArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("gatewaySn", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:gatewaySn: ", p), err)
	}
	if err := oprot.WriteString(string(p.GatewaySn)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.gatewaySn (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:gatewaySn: ", p), err)
	}
	return err
}

func (p *IDataThriftServiceGetInvertorBeanByDateArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("invSn", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:invSn: ", p), err)
	}
	if err := oprot.WriteString(string(p.InvSn)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.invSn (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:invSn: ", p), err)
	}
	return err
}

func (p *IDataThriftServiceGetInvertorBeanByDateArgs) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("startDate", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:startDate: ", p), err)
	}
	if err := oprot.WriteString(string(p.StartDate)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.startDate (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:startDate: ", p), err)
	}
	return err
}

func (p *IDataThriftServiceGetInvertorBeanByDateArgs) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("endDate", thrift.STRING, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:endDate: ", p), err)
	}
	if err := oprot.WriteString(string(p.EndDate)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.endDate (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:endDate: ", p), err)
	}
	return err
}

func (p *IDataThriftServiceGetInvertorBeanByDateArgs) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("timeZone", thrift.I32, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:timeZone: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.TimeZone)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.timeZone (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:timeZone: ", p), err)
	}
	return err
}

func (p *IDataThriftServiceGetInvertorBeanByDateArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("IDataThriftServiceGetInvertorBeanByDateArgs(%+v)", *p)
}

// Attributes:
//  - Success
type IDataThriftServiceGetInvertorBeanByDateResult struct {
	Success []*InvertorThriftBean `thrift:"success,0" json:"success,omitempty"`
}

func NewIDataThriftServiceGetInvertorBeanByDateResult() *IDataThriftServiceGetInvertorBeanByDateResult {
	return &IDataThriftServiceGetInvertorBeanByDateResult{}
}

var IDataThriftServiceGetInvertorBeanByDateResult_Success_DEFAULT []*InvertorThriftBean

func (p *IDataThriftServiceGetInvertorBeanByDateResult) GetSuccess() []*InvertorThriftBean {
	return p.Success
}
func (p *IDataThriftServiceGetInvertorBeanByDateResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *IDataThriftServiceGetInvertorBeanByDateResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.readField0(iprot); err != nil {
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

func (p *IDataThriftServiceGetInvertorBeanByDateResult) readField0(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*InvertorThriftBean, 0, size)
	p.Success = tSlice
	for i := 0; i < size; i++ {
		_elem6 := &InvertorThriftBean{}
		if err := _elem6.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem6), err)
		}
		p.Success = append(p.Success, _elem6)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *IDataThriftServiceGetInvertorBeanByDateResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getInvertorBeanByDate_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
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

func (p *IDataThriftServiceGetInvertorBeanByDateResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.LIST, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Success)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Success {
			if err := v.Write(oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *IDataThriftServiceGetInvertorBeanByDateResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("IDataThriftServiceGetInvertorBeanByDateResult(%+v)", *p)
}

type IDataThriftServiceEchoArgs struct {
}

func NewIDataThriftServiceEchoArgs() *IDataThriftServiceEchoArgs {
	return &IDataThriftServiceEchoArgs{}
}

func (p *IDataThriftServiceEchoArgs) Read(iprot thrift.TProtocol) error {
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
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
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

func (p *IDataThriftServiceEchoArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("echo_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *IDataThriftServiceEchoArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("IDataThriftServiceEchoArgs(%+v)", *p)
}

// Attributes:
//  - Success
type IDataThriftServiceEchoResult struct {
	Success *string `thrift:"success,0" json:"success,omitempty"`
}

func NewIDataThriftServiceEchoResult() *IDataThriftServiceEchoResult {
	return &IDataThriftServiceEchoResult{}
}

var IDataThriftServiceEchoResult_Success_DEFAULT string

func (p *IDataThriftServiceEchoResult) GetSuccess() string {
	if !p.IsSetSuccess() {
		return IDataThriftServiceEchoResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *IDataThriftServiceEchoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *IDataThriftServiceEchoResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.readField0(iprot); err != nil {
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

func (p *IDataThriftServiceEchoResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *IDataThriftServiceEchoResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("echo_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
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

func (p *IDataThriftServiceEchoResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *IDataThriftServiceEchoResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("IDataThriftServiceEchoResult(%+v)", *p)
}
