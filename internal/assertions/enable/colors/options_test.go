// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package colors

import "testing"

func TestOptions(t *testing.T) {
	t.Run("option should enable colors", func(t *testing.T) {
		o := optionsWithDefaults([]Option{
			WithEnable(true),
		})

		if !o.enabled {
			t.Errorf("expected option to enable colors")
		}
	})

	t.Run("option should set Dark theme", func(t *testing.T) {
		o := optionsWithDefaults([]Option{
			WithDark(),
		})

		if o.theme != ThemeDark {
			t.Errorf("expected ThemeDark got: %v", o.theme)
		}
	})

	t.Run("option should set Light theme", func(t *testing.T) {
		o := optionsWithDefaults([]Option{
			WithLight(),
		})

		if o.theme != ThemeLight {
			t.Errorf("expected ThemeLight got: %v", o.theme)
		}
	})

	t.Run("option should sanitize string theme", func(t *testing.T) {
		t.Run("with light", func(t *testing.T) {
			o := optionsWithDefaults([]Option{
				WithSanitizedTheme("light"),
			})

			if o.theme != ThemeLight {
				t.Errorf("expected ThemeLight got: %v", o.theme)
			}
		})

		t.Run("with dark", func(t *testing.T) {
			o := optionsWithDefaults([]Option{
				WithSanitizedTheme("dark"),
			})

			if o.theme != ThemeDark {
				t.Errorf("expected ThemeDark got: %v", o.theme)
			}
		})

		t.Run("with invalid value", func(t *testing.T) {
			o := optionsWithDefaults([]Option{
				WithSanitizedTheme("invalid"),
			})

			defaultOptions := optionsWithDefaults(nil)

			if o.theme != defaultOptions.theme {
				t.Errorf("expected %v (the default) got: %v", defaultOptions.theme, o.theme)
			}
		})
	})
}

func TestOptionsTheme(t *testing.T) {
	t.Run("Theme should be a stringer", func(t *testing.T) {
		th := ThemeDark
		if str := th.String(); str != "dark" {
			t.Errorf(`expected ThemeDark to stringify as "dark", but got: %q`, str)
		}
	})
}
