package minimumwindowsubstring

import "testing"

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{
			name: "classic ascii example",
			s:    "ADOBECODEBANC",
			t:    "ABC",
			want: "BANC",
		},
		{
			name: "unicode simple",
			s:    "żółw and żaba",
			t:    "żb",
			want: "żab",
		},
		{
			name: "unicode duplicates required",
			s:    "żżaba",
			t:    "żż",
			want: "żż",
		},
		{
			name: "emoji runes",
			s:    "a🙂b🙂c",
			t:    "🙂c",
			want: "🙂c",
		},
		{
			name: "cjk runes",
			s:    "我喜欢编程也喜欢咖啡",
			t:    "喜欢咖",
			want: "喜欢咖",
		},
		{
			name: "no possible window",
			s:    "abc🙂def",
			t:    "😎",
			want: "",
		},
		{
			name: "t longer than s by runes",
			s:    "🙂a",
			t:    "🙂a🙂",
			want: "",
		},
		{
			name: "multiple windows choose smallest (unicode)",
			s:    "áβγáδβ",
			t:    "βá",
			want: "áβ",
		},
		{
			name: "empty t",
			s:    "anything",
			t:    "",
			want: "",
		},
		{
			name: "empty s",
			s:    "",
			t:    "x",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinWindow(tt.s, tt.t)
			if got != tt.want {
				t.Fatalf("MinWindow(%q, %q) = %q, want %q", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
