package main

import (
	"os"
	"testing"
)

func TestTemplateRender_Render(t *testing.T) {
	tests := []struct {
		name    string
		tpl     string
		want    string
		wantErr bool
	}{
		{"no template tags", "hey", "hey", false},
		{"invalid template", "hey {{", "", true},
		{"with existing env", "hey {{env \"TEST_ENV\"}}", "hey test-env", false},
		{"with existing env with default", "hey {{envd \"TEST_ENV\" \"default-env\"}}", "hey test-env", false},
		{"with existing empty env with default", "hey {{envd \"TEST_ENV_EMPTY\" \"default-env\"}}", "hey ", false},
		{"with non-existing env", "hey {{env \"TEST_ENV2\"}}", "hey ", true},
		{"with non-existing env with default", "hey {{envd \"TEST_ENV2\" \"default-env\"}}", "hey default-env", false},
		{"with pipe func", "hey {{env \"TEST_ENV\" | upper}}", "hey TEST-ENV", false},
	}
	os.Setenv("TEST_ENV", "test-env")
	os.Setenv("TEST_ENV_EMPTY", "")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTemplateRender().Render(tt.tpl)
			if (err != nil) != tt.wantErr {
				t.Errorf("TemplateRender.render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TemplateRender.render() = %v, want %v", got, tt.want)
			}
		})
	}
	os.Unsetenv("TEST_ENV")
	os.Unsetenv("TEST_ENV_EMPTY")
}
