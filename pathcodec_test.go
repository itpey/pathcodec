package pathcodec

import (
	"testing"
)

func TestComprkessDecompress(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{"Valid path", "M257,455c-56,0-109-25-146-65-143-156,31-397,224-318,201,83,136,386-78,383z", "M257,455c-56,0-109-25-146-65-143-156,31-397,224-318,201,83,136,386-78,383z", false},
		{"Valid path", "M255,491h-182l-1-2-2-2,0-231,0-231,2-1,2-2h182,182l1,2,2,2v230,231l-2,2-2,2h-182z", "M255,491h-182l-1-2-2-2,0-231,0-231,2-1,2-2h182,182l1,2,2,2v230,231l-2,2-2,2h-182z", false},
		{"Valid path", "M256,330h-28l-2-2-2-2,1-2,0-2-13-119-14-119,0-4,0-4,3-3,2-2h56,55l2,2,2,2v2l1,2-14,123-14,123-2,3-3,2h-27z", "M256,330h-28l-2-2-2-2,1-2,0-2-13-119-14-119,0-4,0-4,3-3,2-2h56,55l2,2,2,2v2l1,2-14,123-14,123-2,3-3,2h-27z", false},
		{"Valid path", "M258,458h-211l-1-1-2-2,1-9,1-9,0-62,0-63,1-119,1-118,1-2,2-2,0-6,0-6,2-1,2-1h206,206l2,1,2,1v6,7h2l2,1,1,3v3l1,126,1,125v52,52,10l1,10-2,2-2,2h-210z", "M258,458h-211l-1-1-2-2,1-9,1-9,0-62,0-63,1-119,1-118,1-2,2-2,0-6,0-6,2-1,2-1h206,206l2,1,2,1v6,7h2l2,1,1,3v3l1,126,1,125v52,52,10l1,10-2,2-2,2h-210z", false},
		{"Valid path", "M134,465h-3l1-1,0-2,1-2,0-2,59-90,58-91v1h-69l-69,1,1-2,1-1,0-2,0-2,107-105,108-105,8-8,8-8,2,1h2l1,2v2l-39,81-40,82h71l71,1v2,2l-2,2-1,2-134,123-133,122h-3z", "M134,465h-3l1-1,0-2,1-2,0-2,59-90,58-91v1h-69l-69,1,1-2,1-1,0-2,0-2,107-105,108-105,8-8,8-8,2,1h2l1,2v2l-39,81-40,82h71l71,1v2,2l-2,2-1,2-134,123-133,122h-3z", false},
		{"Valid path", "M256,478c-108,0-111,14-112-98,0-26-6-173,0-184,2-3-10-32-2-40,5-5-3-84,9-102,6-10,18-17,30-18,30-3,136-7,161,2,8,3,15,8,20,15,16,20,7,89,7,115,0,79,0,163,0,245,0,15,4,34-6,47-21,30-75,18-106,18z", "M256,478c-108,0-111,14-112-98,0-26-6-173,0-184,2-3-10-32-2-40,5-5-3-84,9-102,6-10,18-17,30-18,30-3,136-7,161,2,8,3,15,8,20,15,16,20,7,89,7,115,0,79,0,163,0,245,0,15,4,34-6,47-21,30-75,18-106,18z", false},
		{"Valid path with letters", "MASTAVAz", "MASTAVAz", false},
		{"Invalid start character", "X1-3,5-7,9z", "", true},
		{"Invalid end character", "M1-3,5-7,9y", "", true},
		{"Invalid character in path", "M1-3,5-7,9$z", "", true},
		{"Empty path", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded, err := Compress(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("compress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}

			decoded, err := Decompress(encoded)
			if (err != nil) != tt.wantErr {
				t.Errorf("decompress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if decoded != tt.expected {
				t.Errorf("decompress() = %v, expected %v", decoded, tt.expected)
			}
		})
	}
}

func BenchmarkCompress(b *testing.B) {
	path := "M257,455c-56,0-109-25-146-65-143-156,31-397,224-318,201,83,136,386-78,383z"
	for i := 0; i < b.N; i++ {
		_, _ = Compress(path)
	}
}

func BenchmarkDecompress(b *testing.B) {
	encoded := []byte{2, 5, 7, 128, 4, 5, 5, 220, 64, 5, 6, 128, 0, 64, 1, 0, 9, 64, 2, 5, 64, 1, 4, 6, 64, 6, 5, 64, 1, 4, 3, 64, 1, 5, 6, 128, 3, 1, 64, 3, 9, 7, 128, 2, 2, 4, 64, 3, 1, 8, 128, 2, 0, 1, 128, 8, 3, 128, 1, 3, 6, 128, 3, 8, 6, 64, 7, 8, 128, 3, 8, 3}
	for i := 0; i < b.N; i++ {
		_, _ = Decompress(encoded)
	}
}
