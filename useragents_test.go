package useragents

import (
	"testing"
)

func mergeSlice[T comparable](slices ...[]T) (result []T) {
	for _, s := range slices {
		result = append(result, s...)
	}

	return
}

func contains[T comparable](list []T, v T) bool {
	for _, l := range list {
		if l == v {
			return true
		}
	}

	return false
}

func Test_random(t *testing.T) {
	type args struct {
		list []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "chrome",
			args: args{
				list: List.Chrome,
			},
		},
		{
			name: "edge",
			args: args{
				list: List.Edge,
			},
		},
		{
			name: "firefox",
			args: args{
				list: List.Firefox,
			},
		},
		{
			name: "safari",
			args: args{
				list: List.Safari,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := random(tt.args.list)

			if !contains(tt.args.list, got) {
				t.Errorf("random() = slice not contain %v", got)
			}
		})
	}
}

func TestRandom(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Success",
		},
	}

	ua := mergeSlice(
		List.Chrome,
		List.Edge,
		List.Firefox,
		List.Safari,
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Random()

			if !contains(ua, got) {
				t.Errorf("Random() = slice not contain %v", got)
			}
		})
	}
}

func TestRandomLatest(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Success",
		},
	}

	ua := mergeSlice(
		List.Chrome,
		List.Edge,
		List.Firefox,
		List.Safari,
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomLatest()

			if !contains(ua, got) {
				t.Errorf("RandomLatest() = slice not contain %v", got)
			}
		})
	}
}

func TestChrome(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Success",
		},
	}

	ua := mergeSlice(
		List.Chrome,
		List.Edge,
		List.Firefox,
		List.Safari,
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Chrome()

			if !contains(ua, got) {
				t.Errorf("Chrome() = slice not contain %v", got)
			}
		})
	}
}

func TestChromeLatest(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Success",
		},
	}

	ua := mergeSlice(
		List.Chrome,
		List.Edge,
		List.Firefox,
		List.Safari,
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ChromeLatest()

			if !contains(ua, got) {
				t.Errorf("ChromeLatest() = slice not contain %v", got)
			}
		})
	}
}

func TestEdge(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Success",
		},
	}

	ua := mergeSlice(
		List.Chrome,
		List.Edge,
		List.Firefox,
		List.Safari,
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Edge()

			if !contains(ua, got) {
				t.Errorf("Edge() = slice not contain %v", got)
			}
		})
	}
}

func TestEdgeLatest(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Success",
		},
	}

	ua := mergeSlice(
		List.Chrome,
		List.Edge,
		List.Firefox,
		List.Safari,
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EdgeLatest()

			if !contains(ua, got) {
				t.Errorf("EdgeLatest() = slice not contain %v", got)
			}
		})
	}
}

func TestFirefox(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Success",
		},
	}

	ua := mergeSlice(
		List.Chrome,
		List.Edge,
		List.Firefox,
		List.Safari,
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Firefox()

			if !contains(ua, got) {
				t.Errorf("Firefox() = slice not contain %v", got)
			}
		})
	}
}

func TestFirefoxLatest(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Success",
		},
	}

	ua := mergeSlice(
		List.Chrome,
		List.Edge,
		List.Firefox,
		List.Safari,
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FirefoxLatest()

			if !contains(ua, got) {
				t.Errorf("FirefoxLatest() = slice not contain %v", got)
			}
		})
	}
}

func TestSafari(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Success",
		},
	}

	ua := mergeSlice(
		List.Chrome,
		List.Edge,
		List.Firefox,
		List.Safari,
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Safari()

			if !contains(ua, got) {
				t.Errorf("Safari() = slice not contain %v", got)
			}
		})
	}
}

func TestSafariLatest(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Success",
		},
	}

	ua := mergeSlice(
		List.Chrome,
		List.Edge,
		List.Firefox,
		List.Safari,
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SafariLatest()

			if !contains(ua, got) {
				t.Errorf("SafariLatest() = slice not contain %v", got)
			}
		})
	}
}
