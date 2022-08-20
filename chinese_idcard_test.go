package gofunc

import (
	"testing"
	"time"
)

func TestChineseIdCardVerify(t *testing.T) {
	type args struct {
		idCard string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "testCase-01", args: args{idCard: "110101190001011009"}, want: true},
		{name: "testCase-02", args: args{idCard: "110101190001011002"}, want: false},
		{name: "testCase-03", args: args{idCard: "110225196403026127"}, want: true},
		{name: "testCase-04", args: args{idCard: "350201198701146613"}, want: true},
		{name: "testCase-05", args: args{idCard: "350201199308194277"}, want: true},
		{name: "testCase-06", args: args{idCard: "150201199308194277"}, want: false},
		{name: "testCase-07", args: args{idCard: "350205198711076036"}, want: true},
		{name: "testCase-08", args: args{idCard: "513029199006070453"}, want: true},
		{name: "testCase-08", args: args{idCard: "513029201006070453"}, want: true},
		{name: "testCase-09", args: args{idCard: "32062419760212671X"}, want: true},
		{name: "testCase-10", args: args{idCard: "32068319760212671X"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChineseIdCardVerify(tt.args.idCard); got != tt.want {
				t.Errorf("ChineseIdCardVerify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChineseNameVerify(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "testCase-01", args: args{name: "窦民佑"}, want: true},
		{name: "testCase-02", args: args{name: "窦民佑窦民佑"}, want: true},
		{name: "testCase-03", args: args{name: "窦民佑窦民佑佑"}, want: false},
		{name: "testCase-04", args: args{name: "窦民"}, want: true},
		{name: "testCase-05", args: args{name: "窦"}, want: false},
		{name: "testCase-06", args: args{name: "臧誉胜"}, want: true},
		{name: "testCase-07", args: args{name: "呵呵呵"}, want: true},
		{name: "testCase-08", args: args{name: "卡拉潘克·晓"}, want: true},
		{name: "testCase-09", args: args{name: "库大库路马拉可林不拉夫斯基"}, want: true},
		{name: "testCase-10", args: args{name: "胡柏·布里恩·沃夫斯里积士丁可辛贝格多夫"}, want: true},
		{name: "testCase-11", args: args{name: "叶利钦·叶尔波力"}, want: true},
		{name: "testCase-12", args: args{name: "这是个什么鬼名字·我也不知道"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChineseNameVerify(tt.args.name); got != tt.want {
				t.Errorf("ChineseNameVerify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseChineIdCardAge(t *testing.T) {
	type args struct {
		idCard string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "testCase-01", args: args{idCard: "110101190001011009"}, want: 121, wantErr: false},
		{name: "testCase-02", args: args{idCard: "110101190001011002"}, want: 0, wantErr: true},
		{name: "testCase-03", args: args{idCard: "110225196403026127"}, want: 57, wantErr: false},
		{name: "testCase-04", args: args{idCard: "350201198701146613"}, want: 34, wantErr: false},
		{name: "testCase-05", args: args{idCard: "350201199308194277"}, want: 27, wantErr: false},
		{name: "testCase-06", args: args{idCard: "150201199308194277"}, want: 0, wantErr: true},
		{name: "testCase-07", args: args{idCard: "350205198711076036"}, want: 33, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseChineIdCardAge(tt.args.idCard)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseChineIdCardAge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseChineIdCardAge() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenBirthDayByAge(t *testing.T) {
	type args struct {
		age int
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{name: "testCase-01", args: args{age: 1}},
		{name: "testCase-02", args: args{age: 99}},
		{name: "testCase-03", args: args{age: 60}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenBirthDayByAge(tt.args.age); Age(got) != tt.args.age {
				t.Errorf("GenBirthDayByAge() = %v", got.Format(time.RFC3339))
			}

		})
	}
}
