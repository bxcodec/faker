package faker

import "testing"

type Underlying string
type TestInterface interface {
	Test()
}
type TStruct1 struct {
	A Underlying
	B Underlying
	C *Underlying
	D *Underlying
}

type TStruct2 struct {
	A TStruct1
	B TStruct1
	C *TStruct1
	D *TStruct1
}

type TStruct3 struct {
	A []TStruct1
	B []TStruct1
	C *[]TStruct1
	D *[]TStruct1
}

func TestFakeSkipFieldsStrings(t *testing.T) {
	t.Run("test skip fields", func(t *testing.T) {
		var a TStruct1
		err := FakeDataSkipFields(&a, []string{"A", "C", "E"})
		if err != nil {
			t.Errorf("failed to fake strings skip fields type: %v", err)
		}
		if a.A != "" || a.C != nil {
			t.Errorf("failed to fake strings skip fields type: %v", err)
		}
	})
}

func TestFakeSkipFieldsStructs(t *testing.T) {
	t.Run("test skip fields", func(t *testing.T) {
		var a TStruct2
		err := FakeDataSkipFields(&a, []string{"A", "C"})
		if err != nil {
			t.Errorf("failed to fake strings skip fields type: %v", err)
		}
		empty := TStruct1{}
		if a.A != empty || a.C != nil {
			t.Errorf("failed to fake structs skip fields type: %v", err)
		}
	})
}

func TestFakeSkipFieldsStructsPath(t *testing.T) {
	t.Run("test skip fields", func(t *testing.T) {
		var a TStruct2
		err := FakeDataSkipFields(&a, []string{"A.A", "C.A"})
		if err != nil {
			t.Errorf("failed to fake strings skip fields type: %v", err)
		}
		if a.A.A != "" || a.C.A != "" {
			t.Errorf("failed to fake structs skip fields : %v", err)
		}
	})
}

func TestFakeSkipFieldsStructsSlice(t *testing.T) {
	t.Run("test skip fields", func(t *testing.T) {
		var a TStruct3
		err := FakeDataSkipFields(&a, []string{"A.A", "C.C"})
		if err != nil {
			t.Errorf("failed to fake strings skip fields type: %v", err)
		}
		if a.A[0].A != "" || (*a.C)[0].C != nil {
			t.Errorf("failed to fake slices skip fields : %v", err)
		}
	})
}
