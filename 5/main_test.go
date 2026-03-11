package main

import (
	"testing"
)

func TestCheckPassword(t *testing.T) {
	tests := []struct {
		name    string
		pwd     string
		wantErr bool
		msg     string
	}{
		{
			name:    "Empty password",
			pwd:     "",
			wantErr: true,
			msg:     "password is empty",
		},
		{
			name:    "Too short",
			pwd:     "123",
			wantErr: true,
			msg:     "password length is less than 5",
		},
		{
			name:    "Too long",
			pwd:     "this_is_a_very_long_password_more_than_20",
			wantErr: true,
			msg:     "password length is more than 20",
		},
		{
			name:    "Has forbidden chars",
			pwd:     "pass;word",
			wantErr: true,
			msg:     "password contains forbidden chars",
		},
		{
			name:    "Valid password",
			pwd:     "secure123",
			wantErr: false,
			msg:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CheckPassword(tt.pwd)

			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && err.Error() != tt.msg {
				t.Errorf("CheckPassword() error msg = %v, want %v", err.Error(), tt.msg)
			}
		})
	}
}
