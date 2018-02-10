package transaction

import (
	"sort"

	"golang.org/x/crypto/sha3"
	"github.com/mitchellh/mapstructure"
	"strings"
	crypto "github.com/tendermint/go-crypto"
	"github.com/pkg/errors"
)

/*
type Transaction struct {
	Type      TransactionType  `json:"type"`
	Timestamp string           `json:"timestamp"`
	Signature string           `json:"signature"`
	Nonce     uint32           `json:"nonce"`
	Data      []byte           `json:"data"`
}
 */

type Transaction struct {
	SignPubKey string        `json:"signature,omitempty"`
	Signature  string        `json:"signature,omitempty"`
	Type       string        `json:"type,omitempty"`
	Data       interface{}   `json:"data"`
}

func (t *Transaction) FromBytes(bs []byte) error {
	if err := json.Unmarshal(bs, t); err != nil {
		return err
	}
	return nil
}

func (t *Transaction) ToBytes() ([]byte, error) {
	return json.Marshal(t)
}

func (t *Transaction) ToDb() (*Db, error) {
	db := &Db{}
	if err := mapstructure.Decode(t.Data, db); err != nil {
		return nil, err
	}

	if strings.Compare(t.Signature, "") != 0 {
		pk := crypto.PubKeyEd25519{t.SignPubKey}
		if !pk.VerifyBytes(db.ToSortString(), crypto.SignatureEd25519FromBytes(t.SignPubKey)) {
			return nil, errors.New("验证签名失败")
		}
	}

	return db, nil
}

func (t *Transaction) ToAccount() (*Account, error) {
	account := &Account{}
	if err := mapstructure.Decode(t.Data, account); err != nil {
		return nil, err
	}

	if strings.Compare(t.Signature, "") == 0 {
		return nil, errors.New("验证签名为空")
	}

	pk := crypto.PubKeyEd25519{t.SignPubKey}
	if !pk.VerifyBytes(account.ToSortString(), crypto.SignatureEd25519FromBytes(t.SignPubKey)) {
		return nil, errors.New("验证签名失败")
	}

	return account, nil
}

func (t *Transaction) ToValidator() (*Validator, error) {
	val := &Validator{}
	if err := mapstructure.Decode(t.Data, val); err != nil {
		return nil, err
	}

	if strings.Compare(t.Signature, "") == 0 {
		return nil, errors.New("验证签名为空")
	}

	pk := crypto.PubKeyEd25519{t.SignPubKey}
	if !pk.VerifyBytes(val.ToSortString(), crypto.SignatureEd25519FromBytes(t.SignPubKey)) {
		return nil, errors.New("验证签名失败")
	}

	return val, nil
}

//func (t *Transaction) Hash() []byte {
//	hash := sha3.New512()
//	encoder := json.NewEncoder(hash)
//	encoder.Encode(t.Type)
//	encoder.Encode(t.Timestamp)
//	if hashable, ok := t.Data.(Hashable); ok {
//		hash.Write(hashable.Hash())
//	} else {
//		encoder.Encode(t.Data)
//	}
//	return hash.Sum(nil)
//}

//func (t *Transaction) ProofOfWork(cost byte) error {
//	for round := 0; round < (1 << 32); round++ {
//		t.Nonce = uint32(round)
//		if err := t.VerifyProofOfWork(cost); err == nil {
//			return nil
//		}
//	}
//	return errors.New("can not find pow")
//}

//func (t *Transaction) VerifyProofOfWork(cost byte) error {
//	hasher := sha3.New512()
//	hasher.Write(t.Hash())
//	binary.Write(hasher, binary.LittleEndian, t.Nonce)
//	tip := uint64(0)
//	buf := bytes.NewBuffer(hasher.Sum(nil))
//	binary.Read(buf, binary.LittleEndian, &tip)
//	if tip << (64 - cost) == 0 {
//		return nil
//	}
//	return errors.New("failed to validate proof of work")
//}

//func New(t TransactionType, data interface{}) *Transaction {
//	return &Transaction{t, time.Now(), "", 0, data}
//}

func hashStringMap(m map[string]interface{}) []byte {
	hash := sha3.New512()
	encoder := json.NewEncoder(hash)
	keys := make([]string, len(m))
	i := 0
	for id := range m {
		keys[i] = id
		i++
	}
	sort.Strings(keys)
	for _, key := range keys {
		encoder.Encode(key)
		encoder.Encode(m[key])
	}
	return hash.Sum(nil)
}

type RequestQuery struct {
	Data   []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Path   string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Height int64  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Prove  bool   `protobuf:"varint,4,opt,name=prove,proto3" json:"prove,omitempty"`
}