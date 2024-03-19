package snrpost

import (
	"bytes"
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

// UnmarshalStartPost will unmarshal the first line of the snr file
func UnmarshalStartPost(data []byte, post *StartPost) (int, error) {
	return unmarshal(data, post)
}

// UnmarshalFirPost will unmarshal the first section of the snr row known as the FirPost
func UnmarshalFirPost(data []byte, post *FirPost) (int, error) {
	return unmarshal(data, post)
}

// UnmarshalAviserPost will unmarshal the first section of the snr row known as the AviserPost
func UnmarshalAviserPost(data []byte, post *AviserPost) (int, error) {
	return unmarshal(data, post)
}

// UnmarshalData will unmarshal the second section of the snr row known as aviserdata
func UnmarshalData(data []byte, v interface{}) (int, error) {
	return unmarshal(data[postIDLength:], v)
}

// SnrSize is the sum of snr tag values.
func SnrSize(v interface{}) (int, error) {
	snr := reflect.ValueOf(v)
	if snr.Kind() == reflect.Ptr {
		snr = snr.Elem()
	}
	return structSize(snr.Type())
}

//
// Internal functions used only in this file
//

func structSize(refType reflect.Type) (int, error) {
	result := 0
	for i := 0; i < refType.NumField(); i++ {
		field := refType.Field(i)
		switch field.Type.Kind() {
		case reflect.String, reflect.Interface:
			length, _, err := getLen(field)
			if err != nil {
				return result, err
			}
			result += length
		case reflect.Array:
			s, err := structSize(field.Type.Elem())
			if err != nil {
				return s, err
			}
			result += s * field.Type.Len()
		default:
			return int(field.Type.Kind()), ErrUnsupportedType
		}
	}
	return result, nil
}

func doUnmarshal(data []byte, ref reflect.Value) (int, error) {
	refType := ref.Type()
	readPos := 0
	for i := 0; i < refType.NumField(); i++ {
		field := refType.Field(i)
		length, splitAt, err := getLen(field)
		if err != nil && field.Type.Kind() != reflect.Array {
			return readPos, err
		}
		switch field.Type.Kind() {
		case reflect.String:
			d := data[readPos : readPos+length]
			b, err := decoder.Bytes(bytes.Trim(d, "\x00"))
			if err != nil {
				return readPos, err
			}
			fn := ref.FieldByName(field.Name)
			var value string
			if splitAt < length {
				v := string(b)
				value = v[:splitAt] + "." + v[splitAt:]
			} else {
				value = string(b)
			}
			fn.SetString(strings.TrimSpace(value))
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

func getLen(field reflect.StructField) (int, int, error) {
	v := field.Tag.Get("snr")
	vs := strings.Split(v, ",")
	var number int
	var decimal int
	var err error
	number, err = strconv.Atoi(vs[0])
	if err != nil {
		return 0, 0, ErrTagUnsupportedType
	}
	if len(vs) > 1 {
		var err error
		decimal, err = strconv.Atoi(vs[1])
		if err != nil {
			return 0, 0, ErrTagUnsupportedType
		}
	}
	length := number + decimal
	splitAt := length - decimal
	return length, splitAt, nil
}
