package repo

import "testing"

func TestValidateAddPositionStruct(t *testing.T) {
	validSalary := 100
	validName := "Software Engieneer"

	type args struct {
		p Position
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Null name -> Invalid",
			args:    args{p: Position{Name: nil, Salary: &validSalary}},
			wantErr: true,
		},
		{
			name:    "Null salary -> Invalid",
			args:    args{p: Position{Name: &validName, Salary: nil}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateAddPositionStruct(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("ValidateAddPositionStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
