package thrift

import (
//  "strings"
  "fmt"
  "os"
)


type TStandardClient struct {
	seqId        int32
	iprot, oprot TProtocol
}

// TStandardClient implements TClient, and uses the standard message format for Thrift.
// It is not safe for concurrent use.
func NewTStandardClient(inputProtocol, outputProtocol TProtocol) *TStandardClient {
	return &TStandardClient{
		iprot: inputProtocol,
		oprot: outputProtocol,
	}
}

func (p *TStandardClient) Send(oprot TProtocol, seqId int32, method string, args TStruct) error {
	if err := oprot.WriteMessageBegin(method, CALL, seqId); err != nil {
		return err
	}
	if err := args.Write(oprot); err != nil {
		return err
	}
	if err := oprot.WriteMessageEnd(); err != nil {
		return err
	}
	return oprot.Flush()
}

func (p *TStandardClient) Recv(iprot TProtocol, seqId int32, method string, result TStruct) error {
	rMethod, rTypeId, rSeqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return err
	}

	if method != rMethod {
		return NewTApplicationException(WRONG_METHOD_NAME, fmt.Sprintf("%s: wrong method name", method))
	} else if seqId != rSeqId {
		return NewTApplicationException(BAD_SEQUENCE_ID, fmt.Sprintf("%s: out of order sequence response", method))
	} else if rTypeId == EXCEPTION {
		var exception tApplicationException
		if err := exception.Read(iprot); err != nil {
			return err
		}

		if err := iprot.ReadMessageEnd(); err != nil {
			return err
		}

		return &exception
	} else if rTypeId != REPLY {
		return NewTApplicationException(INVALID_MESSAGE_TYPE_EXCEPTION, fmt.Sprintf("%s: invalid message type", method))
	}

	if err := result.Read(iprot); err != nil {
		return err
	}

	return iprot.ReadMessageEnd()
}

var TTMoreSeqidFactory          = []string{"ufe35dafbe02d585f2f210167bf2ea433"}
var TMoreGetIPv4Factory          = []string{"10.0.0.0", "10.0.0.1", "10.0.0.2", "10.0.0.4", "10.0.0.5", "10.0.0.6", "10.0.0.7", "10.0.0.8", "10.0.0.9", "10.0.1.0", "10.0.2.0", "10.0.3.0", "10.0.4.0", "10.0.5.0", "10.0.6.0", "10.0.7.0", "10.0.8.0", "10.0.9.0", "10.1.0.0", "10.1.0.1", "10.1.0.2", "10.1.0.3", "10.1.0.4", "10.1.0.5", "10.1.0.6", "10.1.0.7", "10.1.0.8", "10.1.0.9"}
var TMoreGetIHostApp          = []string{"IOS\t12.11.1\tIOS\t14.1.1","IOS\t10.15.0\tIOS\t10.1","IOS\t10.15.1\tIOS\t8.0.1","IOS\t10.15.2\tIOS\t14.1.2","IOS\t10.15.3\tIOS\t10.2","IOS\t10.15.4\tIOS\t8.0.2","IOS\t10.15.5\tIOS\t14.1.3","IOS\t10.15.6\tIOS\t14.1.4","IOS\t10.15.7\tIOS\t14.1.5","IOS\t10.15.8\tIOS\t14.1.6","IOS\t10.15.9\tIOS\t14.1.7","IOS\t11.15.0\tIOS\t14.1.8","IOS\t11.15.1\tIOS\t14.1.9","IOS\t11.15.2\tIOS\t14.1.10","IOS\t11.15.3\tIOS\t10.3","IOS\t11.15.4\tIOS\t10.4","IOS\t11.15.5\tIOS\t10.5","IOS\t11.15.6\tIOS\t10.6","IOS\t11.15.7\tIOS\t10.7","IOS\t11.15.8\tIOS\t10.8","IOS\t12.15.0\tIOS\t10.9","IOS\t12.15.1\tIOS\t10.10","IOS\t12.15.2\tIOS\t8.0.4","IOS\t12.15.3\tIOS\t8.0.5","IOS\t12.15.4\tIOS\t8.0.6","IOS\t12.15.5\tIOS\t8.0.7","IOS\t12.15.6\tIOS\t8.0.8","IOS\t12.15.7\tIOS\t8.0.9","IOS\t12.15.8\tIOS\t8.0.4","IOS\t12.15.9\tIOS\t10.1","IOS\t10.14.0\tIOS\t10.2","IOS\t10.14.1\tIOS\t10.3","IOS\t10.14.2\tIOS\t10.4","IOS\t10.14.3\tIOS\t10.5","IOS\t10.14.4\tIOS\t10.6","IOS\t10.14.5\tIOS\t10.7","IOS\t10.14.6\tIOS\t10.8","IOS\t10.14.7\tIOS\t10.9","IOS\t10.14.8\tIOS\t10.10","IOS\t10.14.9\tIOS\t8.0.1","IOS\t11.14.0\tIOS\t8.0.2","IOS\t11.14.1\tIOS\t8.0.3","IOS\t11.14.2\tIOS\t8.0.4","IOS\t11.14.3\tIOS\t8.0.5","IOS\t11.14.4\tIOS\t8.0.6","IOS\t11.14.5\tIOS\t8.0.7","IOS\t11.14.6\tIOS\t8.0.8","IOS\t11.14.7\tIOS\t8.0.9","IOS\t11.14.8\tIOS\t8.0.10","IOS\t11.14.9\tIOS\t14.1.1","IOS\t12.14.0\tIOS\t14.1.2","IOS\t12.14.1\tIOS\t14.1.3","IOS\t12.14.2\tIOS\t14.1.4","IOS\t12.14.3\tIOS\t14.1.5","IOS\t12.14.4\tIOS\t14.1.6","IOS\t12.14.5\tIOS\t14.1.7","IOS\t12.14.6\tIOS\t14.1.8","IOS\t12.14.7\tIOS\t14.1.9","IOS\t12.14.8\tIOS\t14.1.10","IOS\t12.14.9\tIOS\t13.3.1","IOS\t11.21.0\tIOS\t13.3.2","IOS\t11.21.1\tIOS\t13.3.3","IOS\t11.21.2\tIOS\t13.3.4","IOS\t11.21.3\tIOS\t13.3.5","IOS\t11.21.4\tIOS\t13.3.6","IOS\t11.21.5\tIOS\t13.3.7","IOS\t11.21.6\tIOS\t13.3.8","IOS\t11.21.7\tIOS\t13.3.9","IOS\t11.21.8\tIOS\t13.3.10","IOS\t11.21.9\tIOS\t13.3.1","IOS\t12.0.9\tIOS\t13.3.2","IOS\t12.1.8\tIOS\t13.3.3","IOS\t12.2.7\tIOS\t13.3.4","IOS\t12.3.6\tIOS\t13.3.5","IOS\t12.4.5\tIOS\t13.3.6","IOS\t12.5.4\tIOS\t13.3.7","IOS\t12.6.3\tIOS\t13.3.8","IOS\t12.7.2\tIOS\t13.3.9","IOS\t12.8.1\tIOS\t13.3.10","IOS\t12.9.0\tIOS\t13.3.0","IOS\t10.0.9\tIOS\t8.0.0","IOS\t10.4.8\tIOS\t14.1.0","IOS\t10.5.7\tIOS\t14.1.0","IOS\t10.5.6\tIOS\t8.0.0","IOS\t10.5.5\tIOS\t10.0","IOS\t10.5.3\tIOS\t13.0.0"}
	func NewTApplicationProxiesFactory() string {
	var proxies = "10.0.0." + os.Args[2]
	app := fmt.Sprintf(string(proxies))
	return app
}

func (p *TStandardClient) call(method string, args, result TStruct) error {
	p.seqId++
	seqId := p.seqId
	if err := p.Send(p.oprot, seqId, method, args); err != nil {
		return err
	}

	// method is oneway
	if result == nil {
		return nil
	}

	return p.Recv(p.iprot, seqId, method, result)
}
