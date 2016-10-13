package snrpost

import (
	"errors"
	"reflect"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

const (
	postIDLength     = 69
	aviserDataLength = 331
)

var (
	// ErrNotAStructPtr is returned if you pass something that is not a pointer to a
	// Struct to Parse
	ErrNotAStructPtr = errors.New("Expected a pointer to a Struct")
	// ErrUnsupportedType if the struct field type is not supported
	ErrUnsupportedType = errors.New("Type is not supported")
	// ErrCannotBeNil if any of the Unmarshal parameters are nil
	ErrCannotBeNil = errors.New("Neither data or v can be nil")
	// ErrTagUnsupportedType if the tag snr cannot be converted to int
	ErrTagUnsupportedType = errors.New("snr tag can't be converted to int")
	decoder               = charmap.ISO8859_15.NewDecoder()
)

// Unmarshal will unmarshal a snr row into the given struct
func unmarshal(data []byte, v interface{}) (int, error) {
	if data == nil || v == nil {
		return 0, ErrCannotBeNil
	}
	ptrRef := reflect.ValueOf(v)
	if ptrRef.Kind() != reflect.Ptr {
		return 0, ErrNotAStructPtr
	}
	ref := ptrRef.Elem()
	if ref.Kind() != reflect.Struct {
		return 0, ErrNotAStructPtr
	}
	return doUnmarshal(data, ref)
}

// UnmarshalPostID will unmarshal the first section of the snr row known as the PostID
func UnmarshalPostID(data []byte, postID *PostID) {
	unmarshal(data, postID)
}

// UnmarshalData will unmarshal the second section of the snr row known as aviserdata
func UnmarshalData(data []byte, v interface{}) {
	unmarshal(data[postIDLength:], v)
}

func doUnmarshal(data []byte, ref reflect.Value) (int, error) {
	refType := ref.Type()
	readPos := 0
	for i := 0; i < refType.NumField(); i++ {
		field := refType.Field(i)
		length, err := getLen(field)
		if err != nil {
			return readPos, err
		}
		switch field.Type.Kind() {
		case reflect.String:
			d := data[readPos : readPos+length]
			b, err := decoder.Bytes(d)
			if err != nil {
				return readPos, err
			}
			fn := ref.FieldByName(field.Name)
			if err != nil {
				return readPos, err
			}
			decoded, err := decoder.Bytes(b)
			if err != nil {
				return readPos, err
			}
			fn.SetString(strings.TrimSpace(string(decoded)))
			readPos += length
		case reflect.Array:
			fn := ref.FieldByName(field.Name)
			for j := 0; j < fn.Len(); j++ {
				str := fn.Index(j)
				addRPos, err := doUnmarshal(data[readPos:], str)
				if err != nil {
					return readPos, err
				}
				readPos += addRPos
			}
		default:
			readPos += length
		}
	}
	return readPos, nil
}

func getLen(field reflect.StructField) (int, error) {
	key := field.Tag.Get("snr")
	i, err := strconv.Atoi(key)
	if err != nil {
		return 0, ErrTagUnsupportedType
	}
	return i, nil
}
