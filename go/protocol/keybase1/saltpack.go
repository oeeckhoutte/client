// Auto-generated by avdl-compiler v1.3.24 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/saltpack.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type AuthenticityType int

const (
	AuthenticityType_SIGNED     AuthenticityType = 0
	AuthenticityType_REPUDIABLE AuthenticityType = 1
	AuthenticityType_ANONYMOUS  AuthenticityType = 2
)

func (o AuthenticityType) DeepCopy() AuthenticityType { return o }

var AuthenticityTypeMap = map[string]AuthenticityType{
	"SIGNED":     0,
	"REPUDIABLE": 1,
	"ANONYMOUS":  2,
}

var AuthenticityTypeRevMap = map[AuthenticityType]string{
	0: "SIGNED",
	1: "REPUDIABLE",
	2: "ANONYMOUS",
}

func (e AuthenticityType) String() string {
	if v, ok := AuthenticityTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type SaltpackEncryptOptions struct {
	Recipients                []string         `codec:"recipients" json:"recipients"`
	TeamRecipients            []string         `codec:"teamRecipients" json:"teamRecipients"`
	AuthenticityType          AuthenticityType `codec:"authenticityType" json:"authenticityType"`
	UseEntityKeys             bool             `codec:"useEntityKeys" json:"useEntityKeys"`
	UseDeviceKeys             bool             `codec:"useDeviceKeys" json:"useDeviceKeys"`
	UsePaperKeys              bool             `codec:"usePaperKeys" json:"usePaperKeys"`
	NoSelfEncrypt             bool             `codec:"noSelfEncrypt" json:"noSelfEncrypt"`
	Binary                    bool             `codec:"binary" json:"binary"`
	SaltpackVersion           int              `codec:"saltpackVersion" json:"saltpackVersion"`
	UseKBFSKeysOnlyForTesting bool             `codec:"useKBFSKeysOnlyForTesting" json:"useKBFSKeysOnlyForTesting"`
}

func (o SaltpackEncryptOptions) DeepCopy() SaltpackEncryptOptions {
	return SaltpackEncryptOptions{
		Recipients: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.Recipients),
		TeamRecipients: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.TeamRecipients),
		AuthenticityType:          o.AuthenticityType.DeepCopy(),
		UseEntityKeys:             o.UseEntityKeys,
		UseDeviceKeys:             o.UseDeviceKeys,
		UsePaperKeys:              o.UsePaperKeys,
		NoSelfEncrypt:             o.NoSelfEncrypt,
		Binary:                    o.Binary,
		SaltpackVersion:           o.SaltpackVersion,
		UseKBFSKeysOnlyForTesting: o.UseKBFSKeysOnlyForTesting,
	}
}

type SaltpackDecryptOptions struct {
	Interactive      bool `codec:"interactive" json:"interactive"`
	ForceRemoteCheck bool `codec:"forceRemoteCheck" json:"forceRemoteCheck"`
	UsePaperKey      bool `codec:"usePaperKey" json:"usePaperKey"`
}

func (o SaltpackDecryptOptions) DeepCopy() SaltpackDecryptOptions {
	return SaltpackDecryptOptions{
		Interactive:      o.Interactive,
		ForceRemoteCheck: o.ForceRemoteCheck,
		UsePaperKey:      o.UsePaperKey,
	}
}

type SaltpackSignOptions struct {
	Detached        bool `codec:"detached" json:"detached"`
	Binary          bool `codec:"binary" json:"binary"`
	SaltpackVersion int  `codec:"saltpackVersion" json:"saltpackVersion"`
}

func (o SaltpackSignOptions) DeepCopy() SaltpackSignOptions {
	return SaltpackSignOptions{
		Detached:        o.Detached,
		Binary:          o.Binary,
		SaltpackVersion: o.SaltpackVersion,
	}
}

type SaltpackVerifyOptions struct {
	SignedBy  string `codec:"signedBy" json:"signedBy"`
	Signature []byte `codec:"signature" json:"signature"`
}

func (o SaltpackVerifyOptions) DeepCopy() SaltpackVerifyOptions {
	return SaltpackVerifyOptions{
		SignedBy: o.SignedBy,
		Signature: (func(x []byte) []byte {
			if x == nil {
				return nil
			}
			return append([]byte{}, x...)
		})(o.Signature),
	}
}

type SaltpackEncryptedMessageInfo struct {
	Devices          []Device       `codec:"devices" json:"devices"`
	NumAnonReceivers int            `codec:"numAnonReceivers" json:"numAnonReceivers"`
	ReceiverIsAnon   bool           `codec:"receiverIsAnon" json:"receiverIsAnon"`
	Sender           SaltpackSender `codec:"sender" json:"sender"`
}

func (o SaltpackEncryptedMessageInfo) DeepCopy() SaltpackEncryptedMessageInfo {
	return SaltpackEncryptedMessageInfo{
		Devices: (func(x []Device) []Device {
			if x == nil {
				return nil
			}
			ret := make([]Device, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Devices),
		NumAnonReceivers: o.NumAnonReceivers,
		ReceiverIsAnon:   o.ReceiverIsAnon,
		Sender:           o.Sender.DeepCopy(),
	}
}

type SaltpackOperationType int

const (
	SaltpackOperationType_DECRYPT SaltpackOperationType = 0
	SaltpackOperationType_VERIFY  SaltpackOperationType = 1
)

func (o SaltpackOperationType) DeepCopy() SaltpackOperationType { return o }

var SaltpackOperationTypeMap = map[string]SaltpackOperationType{
	"DECRYPT": 0,
	"VERIFY":  1,
}

var SaltpackOperationTypeRevMap = map[SaltpackOperationType]string{
	0: "DECRYPT",
	1: "VERIFY",
}

func (e SaltpackOperationType) String() string {
	if v, ok := SaltpackOperationTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type SaltpackEncryptArg struct {
	SessionID int                    `codec:"sessionID" json:"sessionID"`
	Source    Stream                 `codec:"source" json:"source"`
	Sink      Stream                 `codec:"sink" json:"sink"`
	Opts      SaltpackEncryptOptions `codec:"opts" json:"opts"`
}

type SaltpackDecryptArg struct {
	SessionID int                    `codec:"sessionID" json:"sessionID"`
	Source    Stream                 `codec:"source" json:"source"`
	Sink      Stream                 `codec:"sink" json:"sink"`
	Opts      SaltpackDecryptOptions `codec:"opts" json:"opts"`
}

type SaltpackSignArg struct {
	SessionID int                 `codec:"sessionID" json:"sessionID"`
	Source    Stream              `codec:"source" json:"source"`
	Sink      Stream              `codec:"sink" json:"sink"`
	Opts      SaltpackSignOptions `codec:"opts" json:"opts"`
}

type SaltpackVerifyArg struct {
	SessionID int                   `codec:"sessionID" json:"sessionID"`
	Source    Stream                `codec:"source" json:"source"`
	Sink      Stream                `codec:"sink" json:"sink"`
	Opts      SaltpackVerifyOptions `codec:"opts" json:"opts"`
}

type NotifySaltpackSuccessArg struct {
	SessionID int                   `codec:"sessionID" json:"sessionID"`
	Typ       SaltpackOperationType `codec:"typ" json:"typ"`
	Message   string                `codec:"message" json:"message"`
}

type SaltpackInterface interface {
	SaltpackEncrypt(context.Context, SaltpackEncryptArg) error
	SaltpackDecrypt(context.Context, SaltpackDecryptArg) (SaltpackEncryptedMessageInfo, error)
	SaltpackSign(context.Context, SaltpackSignArg) error
	SaltpackVerify(context.Context, SaltpackVerifyArg) (SaltpackSender, error)
	NotifySaltpackSuccess(context.Context, NotifySaltpackSuccessArg) error
}

func SaltpackProtocol(i SaltpackInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.saltpack",
		Methods: map[string]rpc.ServeHandlerDescription{
			"saltpackEncrypt": {
				MakeArg: func() interface{} {
					ret := make([]SaltpackEncryptArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SaltpackEncryptArg)
					if !ok {
						err = rpc.NewTypeError((*[]SaltpackEncryptArg)(nil), args)
						return
					}
					err = i.SaltpackEncrypt(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"saltpackDecrypt": {
				MakeArg: func() interface{} {
					ret := make([]SaltpackDecryptArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SaltpackDecryptArg)
					if !ok {
						err = rpc.NewTypeError((*[]SaltpackDecryptArg)(nil), args)
						return
					}
					ret, err = i.SaltpackDecrypt(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"saltpackSign": {
				MakeArg: func() interface{} {
					ret := make([]SaltpackSignArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SaltpackSignArg)
					if !ok {
						err = rpc.NewTypeError((*[]SaltpackSignArg)(nil), args)
						return
					}
					err = i.SaltpackSign(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"saltpackVerify": {
				MakeArg: func() interface{} {
					ret := make([]SaltpackVerifyArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SaltpackVerifyArg)
					if !ok {
						err = rpc.NewTypeError((*[]SaltpackVerifyArg)(nil), args)
						return
					}
					ret, err = i.SaltpackVerify(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"notifySaltpackSuccess": {
				MakeArg: func() interface{} {
					ret := make([]NotifySaltpackSuccessArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]NotifySaltpackSuccessArg)
					if !ok {
						err = rpc.NewTypeError((*[]NotifySaltpackSuccessArg)(nil), args)
						return
					}
					err = i.NotifySaltpackSuccess(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type SaltpackClient struct {
	Cli rpc.GenericClient
}

func (c SaltpackClient) SaltpackEncrypt(ctx context.Context, __arg SaltpackEncryptArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.saltpack.saltpackEncrypt", []interface{}{__arg}, nil)
	return
}

func (c SaltpackClient) SaltpackDecrypt(ctx context.Context, __arg SaltpackDecryptArg) (res SaltpackEncryptedMessageInfo, err error) {
	err = c.Cli.Call(ctx, "keybase.1.saltpack.saltpackDecrypt", []interface{}{__arg}, &res)
	return
}

func (c SaltpackClient) SaltpackSign(ctx context.Context, __arg SaltpackSignArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.saltpack.saltpackSign", []interface{}{__arg}, nil)
	return
}

func (c SaltpackClient) SaltpackVerify(ctx context.Context, __arg SaltpackVerifyArg) (res SaltpackSender, err error) {
	err = c.Cli.Call(ctx, "keybase.1.saltpack.saltpackVerify", []interface{}{__arg}, &res)
	return
}

func (c SaltpackClient) NotifySaltpackSuccess(ctx context.Context, __arg NotifySaltpackSuccessArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.saltpack.notifySaltpackSuccess", []interface{}{__arg}, nil)
	return
}
