package business_hours

import "testing"

func TestIsValidBusinessHours(t *testing.T) {
	type args struct {
		hours []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "【時刻表記のチェック】4桁の数字で時刻を表記",
			args: args{
				hours: []string{"07:00", "22:00"},
			},
			want: true,
		},
		{
			name: "【時刻表記のチェック】0が省略された形",
			args: args{
				hours: []string{"7:00", "22:00"},
			},
			want: true,
		},
		{
			name: "【時刻表記のチェック】時間の表記が不適当な値",
			args: args{
				hours: []string{"30:00", "40:00"},
			},
			want: false,
		},
		{
			name: "【時刻表記のチェック】分の表記が不適当な値",
			args: args{
				hours: []string{"07:30", "10:69"},
			},
			want: false,
		},
		{
			name: "【時刻表記のチェック】分の0が省略されている",
			args: args{
				hours: []string{"7:3", "10:00"},
			},
			want: false,
		},
		{
			name: "【時刻表記のチェック】桁数がおかしい",
			args: args{
				hours: []string{"07:300", "10:30"},
			},
			want: false,
		},
		{
			name: "【時刻表記のチェック】コロンがない",
			args: args{
				hours: []string{"07:30", "1030"},
			},
			want: false,
		},
		{
			name: "【時刻表記のチェック】コロン以外の文字が使われている",
			args: args{
				hours: []string{"07.30", "10.30"},
			},
			want: false,
		},
		{
			name: "【時刻表記のチェック】AM/PMなど文字が入っている",
			args: args{
				hours: []string{"AM07:30", "PM10:30"},
			},
			want: false,
		},
		{
			name: "【組み合わせ】開始時間が抜けている",
			args: args{
				hours: []string{"", "10:30"},
			},
			want: true,
		},
		{
			name: "【組み合わせ】終了時間が抜けている",
			args: args{
				hours: []string{"7:30", ""},
			},
			want: true,
		},
		{
			name: "【組み合わせ】終了時間と開始時間が逆転している",
			args: args{
				hours: []string{"10:00", "9:00"},
			},
			want: false,
		},
		{
			name: "【組み合わせ】前の終了時間と開始時間が重複している",
			args: args{
				hours: []string{"10:30", "15:00", "14:00", "19:00"},
			},
			want: false,
		},
		{
			name: "【その他】空文字列が後半に入っている",
			args: args{
				hours: []string{"07:00", "22:00", "", "", "", ""},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidBusinessHours(tt.args.hours); got != tt.want {
				t.Errorf("IsValidBusinessHours() = %v, want %v", got, tt.want)
			}
		})
	}
}
