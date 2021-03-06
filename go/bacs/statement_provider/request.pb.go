// Code generated by protoc-gen-go.
// source: bacs/statement_provider/request.proto
// DO NOT EDIT!

/*
Package statement_provider is a generated protocol buffer package.

It is generated from these files:
	bacs/statement_provider/request.proto

It has these top-level messages:
	Request
*/
package statement_provider

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import bacs_problem "github.com/bacsorg/problem/go/bacs/problem"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Request struct {
	Package  string                 `protobuf:"bytes,1,opt,name=package" json:"package,omitempty"`
	Revision *bacs_problem.Revision `protobuf:"bytes,2,opt,name=revision" json:"revision,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetRevision() *bacs_problem.Revision {
	if m != nil {
		return m.Revision
	}
	return nil
}
