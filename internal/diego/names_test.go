package diego

import (
	"testing"
)

func TestValidateEnvPrefix(t *testing.T) {
	tests := []struct {
		name          string
		schemaPrefix  string
		want          Prefix
		wantErr       bool
		expectedError string
	}{
		{
			name:         "valid prefix",
			schemaPrefix: "MY_APP",
			want:         "MY_APP",
			wantErr:      false,
		},
		{
			name:         "valid prefix lowercase",
			schemaPrefix: "my_app",
			want:         "MY_APP",
			wantErr:      false,
		},
		{
			name:          "invalid prefix with hyphen",
			schemaPrefix:  "my-app",
			wantErr:       true,
			expectedError: "invalid environment prefix 'my-app'; should match ^[\\dA-z_]*$",
		},
		{
			name:          "invalid prefix with special characters",
			schemaPrefix:  "my$app",
			wantErr:       true,
			expectedError: "invalid environment prefix 'my$app'; should match ^[\\dA-z_]*$",
		},
		{
			name:         "empty prefix",
			schemaPrefix: "",
			want:         "",
			wantErr:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidatePrefix(tt.schemaPrefix)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateEnvPrefix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.expectedError {
				t.Errorf("ValidateEnvPrefix() error = %v, expectedError %v", err.Error(), tt.expectedError)
			}
			if got != tt.want {
				t.Errorf("ValidateEnvPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildEnvVar(t *testing.T) {
	tests := []struct {
		name   string
		prefix Prefix
		arg    string
		want   string
	}{
		{
			name:   "simple prefix and name",
			prefix: "MY_APP",
			arg:    "foo",
			want:   "MY_APP_FOO",
		},
		{
			name:   "name with hyphens",
			prefix: "MY_APP",
			arg:    "foo-bar-baz",
			want:   "MY_APP_FOO_BAR_BAZ",
		},
		{
			name:   "empty name",
			prefix: "MY_APP",
			arg:    "",
			want:   "MY_APP_",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildEnvVar(tt.prefix, tt.arg); got != tt.want {
				t.Errorf("BuildEnvVar() = %v, want %v", got, tt.want)
			}
		})
	}
}
