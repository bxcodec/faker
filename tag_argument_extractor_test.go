package faker

import (
	"fmt"
	"testing"

	fakerErrors "github.com/bxcodec/faker/v4/pkg/errors"
)

func TestExtractFloat64(t *testing.T) {

	t.Run("happy path", func(t *testing.T) {
		const f1 = 658897324.4626827
		const f2 = 77592747.726643
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
		}
		res, err := extractFloat64FromTagArgs(args)
		if err != nil {
			t.Errorf("expected no error but got %v", err)
		}
		res64, ok := res.(float64)
		if !ok {
			t.Errorf("expected a float64, but got something else")
		}
		one := res64 == f1
		two := res64 == f2
		if !one && !two {
			t.Errorf(
				"expected %v or %v but got %v",
				f1,
				f2,
				res64,
			)
		}
	})

	t.Run("ErrUnsupportedTagArguments", func(t *testing.T) {
		const f1 = 658897324.4626827
		const f2 = 77592747.726643
		const f3 = "oops"
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
			fmt.Sprintf("%v", f3),
		}
		_, err := extractFloat64FromTagArgs(args)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
		actual := err.Error()
		expected := fakerErrors.ErrUnsupportedTagArguments
		if actual != expected {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})

}

func TestExtractFloat32(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		const f1 = 6500347324.4627466827
		const f2 = 7757290047.772026643
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
		}
		res, err := extractFloat32FromTagArgs(args)
		if err != nil {
			t.Errorf("expected no error but got %v", err)
		}
		res32, ok := res.(float32)
		if !ok {
			t.Errorf("expected a float32, but got something else")
		}
		one := res32 == f1
		two := res32 == f2
		if !one && !two {
			t.Errorf(
				"expected %v or %v but got %v",
				f1,
				f2,
				res32,
			)
		}
	})

	t.Run("ErrUnsupportedTagArguments", func(t *testing.T) {
		const f1 = 6500034324.4646626827
		const f2 = 775743047.7757926643
		const f3 = "oops"
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
			fmt.Sprintf("%v", f3),
		}
		_, err := extractFloat64FromTagArgs(args)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
		actual := err.Error()
		expected := fakerErrors.ErrUnsupportedTagArguments
		if actual != expected {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})

}

func TestExtractInt64(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		const f1 = 6500347324
		const f2 = 7757290047
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
		}
		res, err := extractInt64FromTagArgs(args)
		if err != nil {
			t.Errorf("expected no error but got %v", err)
		}
		res64, ok := res.(int64)
		if !ok {
			t.Errorf("expected a float64, but got something else")
		}
		one := res64 == f1
		two := res64 == f2
		if !one && !two {
			t.Errorf(
				"expected %v or %v but got %v",
				f1,
				f2,
				res64,
			)
		}
	})

	t.Run("ErrUnsupportedTagArguments", func(t *testing.T) {
		const f1 = 6500347324
		const f2 = 7757290047
		const f3 = "oops"
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
			fmt.Sprintf("%v", f3),
		}
		_, err := extractInt64FromTagArgs(args)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
		actual := err.Error()
		expected := fakerErrors.ErrUnsupportedTagArguments
		if actual != expected {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})

}

func TestExtractInt32(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		const f1 = -160347324
		const f2 = -75290047
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
		}
		res, err := extractInt32FromTagArgs(args)
		if err != nil {
			t.Errorf("expected no error but got %v", err)
		}
		res32, ok := res.(int32)
		if !ok {
			t.Errorf("expected a float32, but got something else")
		}
		one := res32 == f1
		two := res32 == f2
		if !one && !two {
			t.Errorf(
				"expected %v or %v but got %v",
				f1,
				f2,
				res32,
			)
		}
	})

	t.Run("ErrUnsupportedTagArguments", func(t *testing.T) {
		const f1 = -6500347324
		const f2 = 7757290047
		const f3 = "oops"
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
			fmt.Sprintf("%v", f3),
		}
		_, err := extractInt32FromTagArgs(args)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
		actual := err.Error()
		expected := fakerErrors.ErrUnsupportedTagArguments
		if actual != expected {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})

}

func TestExtractInt16(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		const f1 = -19474
		const f2 = 5047
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
		}
		res, err := extractInt16FromTagArgs(args)
		if err != nil {
			t.Errorf("expected no error but got %v", err)
		}
		res16, ok := res.(int16)
		if !ok {
			t.Errorf("expected a float16, but got something else")
		}
		one := res16 == f1
		two := res16 == f2
		if !one && !two {
			t.Errorf(
				"expected %v or %v but got %v",
				f1,
				f2,
				res16,
			)
		}
	})

	t.Run("ErrUnsupportedTagArguments", func(t *testing.T) {
		const f1 = -65003
		const f2 = 77572
		const f3 = "oops"
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
			fmt.Sprintf("%v", f3),
		}
		_, err := extractInt16FromTagArgs(args)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
		actual := err.Error()
		expected := fakerErrors.ErrUnsupportedTagArguments
		if actual != expected {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})

	t.Run("int overflow", func(t *testing.T) {
		const f1 = -650875703 // these are too big for int16
		const f2 = 775784842
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
		}
		_, err := extractInt16FromTagArgs(args)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
		actual := err.Error()
		expected := fakerErrors.ErrUnsupportedTagArguments
		if actual != expected {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})

}

func TestExtractInt8(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		const f1 = -81
		const f2 = 104
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
		}
		res, err := extractInt8FromTagArgs(args)
		if err != nil {
			t.Errorf("expected no error but got %v", err)
		}
		res8, ok := res.(int8)
		if !ok {
			t.Errorf("expected a float8, but got something else")
		}
		one := res8 == f1
		two := res8 == f2
		if !one && !two {
			t.Errorf(
				"expected %v or %v but got %v",
				f1,
				f2,
				res8,
			)
		}
	})

	t.Run("ErrUnsupportedTagArguments", func(t *testing.T) {
		const f1 = -103
		const f2 = 72
		const f3 = "oops"
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
			fmt.Sprintf("%v", f3),
		}
		_, err := extractInt8FromTagArgs(args)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("int overflow", func(t *testing.T) {
		const f1 = -650875703 // these are too big for int8
		const f2 = 775784842
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
		}
		_, err := extractInt8FromTagArgs(args)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
		actual := err.Error()
		expected := fakerErrors.ErrUnsupportedTagArguments
		if actual != expected {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})

}

func TestExtractInt(t *testing.T) {

	t.Run("happy path", func(t *testing.T) {
		const f1 = -17575
		const f2 = 10467463
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
		}
		res, err := extractIntFromTagArgs(args)
		if err != nil {
			t.Errorf("expected no error but got %v", err)
		}
		res, ok := res.(int)
		if !ok {
			t.Errorf("expected a float, but got something else")
		}
		one := res == f1
		two := res == f2
		if !one && !two {
			t.Errorf(
				"expected %v or %v but got %v",
				f1,
				f2,
				res,
			)
		}
	})

	t.Run("ErrUnsupportedTagArguments", func(t *testing.T) {
		const f1 = -1037474
		const f2 = 72747346
		const f3 = "oops"
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
			fmt.Sprintf("%v", f3),
		}
		_, err := extractIntFromTagArgs(args)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

}

func TestExtractUint64(t *testing.T) {

	t.Run("happy path", func(t *testing.T) {
		const f1 = 1658584
		const f2 = 10467463
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
		}
		res, err := extractUint64FromTagArgs(args)
		if err != nil {
			t.Errorf("expected no error but got %v", err)
		}
		res, ok := res.(uint64)
		if !ok {
			t.Errorf("expected a uint64, but got something else")
		}
		one := res == uint64(f1)
		two := res == uint64(f2)
		if !one && !two {
			t.Errorf(
				"expected %v or %v but got %v",
				f1,
				f2,
				res,
			)
		}
	})

	t.Run("ErrUnsupportedTagArguments", func(t *testing.T) {
		const f1 = 175757
		const f2 = 727473474
		const f3 = "oops"
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
			fmt.Sprintf("%v", f3),
		}
		_, err := extractUint64FromTagArgs(args)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("negatives", func(t *testing.T) {
		const f1 = -1757665
		const f2 = 727473474
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
		}
		_, err := extractUint64FromTagArgs(args)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

}

func TestExtractUint32(t *testing.T) {

	t.Run("happy path", func(t *testing.T) {
		const f1 = 1658584
		const f2 = 10467463
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
		}
		res, err := extractUint32FromTagArgs(args)
		if err != nil {
			t.Errorf("expected no error but got %v", err)
		}
		res, ok := res.(uint32)
		if !ok {
			t.Errorf("expected a uint32, but got something else")
		}
		one := res == uint32(f1)
		two := res == uint32(f2)
		if !one && !two {
			t.Errorf(
				"expected %v or %v but got %v",
				f1,
				f2,
				res,
			)
		}
	})

	t.Run("ErrUnsupportedTagArguments", func(t *testing.T) {
		const f1 = 175757
		const f2 = 727473474
		const f3 = "oops"
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
			fmt.Sprintf("%v", f3),
		}
		_, err := extractUint32FromTagArgs(args)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("negatives", func(t *testing.T) {
		const f1 = -175757
		const f2 = 727473474
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
		}
		_, err := extractUint32FromTagArgs(args)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

}

func TestExtractUint16(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		const f1 = 16580
		const f2 = 10467
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
		}
		res, err := extractUint16FromTagArgs(args)
		if err != nil {
			t.Errorf("expected no error but got %v", err)
		}
		res, ok := res.(uint16)
		if !ok {
			t.Errorf("expected a uint16, but got something else")
		}
		one := res == uint16(f1)
		two := res == uint16(f2)
		if !one && !two {
			t.Errorf(
				"expected %v or %v but got %v",
				f1,
				f2,
				res,
			)
		}
	})

	t.Run("ErrUnsupportedTagArguments", func(t *testing.T) {
		const f1 = 175757
		const f2 = 727473474
		const f3 = "oops"
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
			fmt.Sprintf("%v", f3),
		}
		_, err := extractUint16FromTagArgs(args)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("negatives", func(t *testing.T) {
		const f1 = -1757
		const f2 = 2747
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
		}
		_, err := extractUint16FromTagArgs(args)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})
}

func TestExtractUint8(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		const f1 = 254
		const f2 = 107
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
		}
		res, err := extractUint8FromTagArgs(args)
		if err != nil {
			t.Errorf("expected no error but got %v", err)
		}
		res, ok := res.(uint8)
		if !ok {
			t.Errorf("expected a uint8, but got something else")
		}
		one := res == uint8(f1)
		two := res == uint8(f2)
		if !one && !two {
			t.Errorf(
				"expected %v or %v but got %v",
				f1,
				f2,
				res,
			)
		}
	})

	t.Run("ErrUnsupportedTagArguments", func(t *testing.T) {
		const f1 = 175
		const f2 = 227
		const f3 = "oops"
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
			fmt.Sprintf("%v", f3),
		}
		_, err := extractUint8FromTagArgs(args)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("negatives", func(t *testing.T) {
		const f1 = -175
		const f2 = 204
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
		}
		_, err := extractUint8FromTagArgs(args)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})
}

func TestExtractUint(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		const f1 = 25467565
		const f2 = 1072562460
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
		}
		res, err := extractUintFromTagArgs(args)
		if err != nil {
			t.Errorf("expected no error but got %v", err)
		}
		res, ok := res.(uint)
		if !ok {
			t.Errorf("expected a uint, but got something else")
		}
		one := res == uint(f1)
		two := res == uint(f2)
		if !one && !two {
			t.Errorf(
				"expected %v or %v but got %v",
				f1,
				f2,
				res,
			)
		}
	})

	t.Run("ErrUnsupportedTagArguments", func(t *testing.T) {
		const f1 = 175774290
		const f2 = 2270064663
		const f3 = "oops"
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
			fmt.Sprintf("%v", f3),
		}
		_, err := extractUintFromTagArgs(args)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("negatives", func(t *testing.T) {
		const f1 = -1756746346
		const f2 = 204466272
		args := []string{
			fmt.Sprintf("%v", f1),
			fmt.Sprintf("%v", f2),
		}
		_, err := extractUintFromTagArgs(args)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})
}
