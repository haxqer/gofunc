package gofunc

import "testing"

func TestReplaceMacro(t *testing.T) {
	type args struct {
		str        string
		macro      map[string][]string
		macroValue map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testCase01", args: args{
			str: "a.gif?event=start&appid=1&capid=16099&adid=18057&assetid=18216&token=49B228EEA940E02927F8222BEFF35B7B16016A29B47E29F17533B5C7948CE86134A7CBA5C63A227CB13A169649CD1ED1&mac=222C856A32BDE1A2D7044E800FA957E6&adtoken=EF0B6CD417C1D13E18970A0FD6BCDAEC041C51A513DB351F4070426B631084F69B50EC3411A677A271D8D6B6AA1F609B&ip=__IP__&ts=__TS__&mac_raw=__MACR__&app_version=__APP__&version=__VERSION__&chan=__CHAN__&dra=__DRA__&iesid=__IESID__&hiesid=__HIESID__&branch=__BRANCH__&sys=__SYSTEM__&mn=__MN__&device_type=__DEVICETYPE__&os=__OS__&ipdx=__IPDX__&pos=__POS__&ctref=__CTREF__&adspace=${ADSPACE_ID}&camp=${CAMPAIGN_ID}&creative=${CREATIVE_ID}&uuid=__UUID__&cookie=${COOKIE}&gppid=__GPPID__&ua=__UA__",
			macro: map[string][]string{
				"mac_raw": {"__MACR__", "__M6O__"},
				"ua":      {"__UA__"},
			},
			macroValue: map[string]string{
				"mac_raw": "28-35-45-19-CC-D9",
				"ua":      "Mozilla%2F5.0+%28Linux%3B+Android+6.1%3B+MagicBox_M13+Build%2FLMY47V%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Version%2F4.0+Chrome%2F39.0.0.0+Safari%2F537.36",
			},
		}, want: `a.gif?event=start&appid=1&capid=16099&adid=18057&assetid=18216&token=49B228EEA940E02927F8222BEFF35B7B16016A29B47E29F17533B5C7948CE86134A7CBA5C63A227CB13A169649CD1ED1&mac=222C856A32BDE1A2D7044E800FA957E6&adtoken=EF0B6CD417C1D13E18970A0FD6BCDAEC041C51A513DB351F4070426B631084F69B50EC3411A677A271D8D6B6AA1F609B&ip=__IP__&ts=__TS__&mac_raw=28-35-45-19-CC-D9&app_version=__APP__&version=__VERSION__&chan=__CHAN__&dra=__DRA__&iesid=__IESID__&hiesid=__HIESID__&branch=__BRANCH__&sys=__SYSTEM__&mn=__MN__&device_type=__DEVICETYPE__&os=__OS__&ipdx=__IPDX__&pos=__POS__&ctref=__CTREF__&adspace=${ADSPACE_ID}&camp=${CAMPAIGN_ID}&creative=${CREATIVE_ID}&uuid=__UUID__&cookie=${COOKIE}&gppid=__GPPID__&ua=Mozilla%2F5.0+%28Linux%3B+Android+6.1%3B+MagicBox_M13+Build%2FLMY47V%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Version%2F4.0+Chrome%2F39.0.0.0+Safari%2F537.36`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceMacro(tt.args.str, tt.args.macro, tt.args.macroValue); got != tt.want {
				t.Errorf("ReplaceMacro() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkReplaceMacro(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ReplaceMacro("a.gif?event=start&appid=1&capid=16099&adid=18057&assetid=18216&token=49B228EEA940E02927F8222BEFF35B7B16016A29B47E29F17533B5C7948CE86134A7CBA5C63A227CB13A169649CD1ED1&mac=222C856A32BDE1A2D7044E800FA957E6&adtoken=EF0B6CD417C1D13E18970A0FD6BCDAEC041C51A513DB351F4070426B631084F69B50EC3411A677A271D8D6B6AA1F609B&ip=__IP__&ts=__TS__&mac_raw=__MACR__&app_version=__APP__&version=__VERSION__&chan=__CHAN__&dra=__DRA__&iesid=__IESID__&hiesid=__HIESID__&branch=__BRANCH__&sys=__SYSTEM__&mn=__MN__&device_type=__DEVICETYPE__&os=__OS__&ipdx=__IPDX__&pos=__POS__&ctref=__CTREF__&adspace=${ADSPACE_ID}&camp=${CAMPAIGN_ID}&creative=${CREATIVE_ID}&uuid=__UUID__&cookie=${COOKIE}&gppid=__GPPID__&ua=__UA__",
			map[string][]string{
				"mac_raw": {"__MACR__", "__M6O__"},
				"ua":      {"__UA__"},
				"branch":  {"__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__"},
				"branch1": {"__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__"},
				"branch2": {"__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__"},
				"branch3": {"__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__"},
				"branch4": {"__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__", "__BRANCH__"},
			},
			map[string]string{
				"mac_raw": "28-35-45-19-CC-D9",
				"branch":  "Xiaomi",
				"branch2": "Xiaomi",
				"branch1": "Xiaomi",
				"branch3": "Xiaomi",
				"branch4": "Xiaomi",
				"ua":      "Mozilla%2F5.0+%28Linux%3B+Android+6.1%3B+MagicBox_M13+Build%2FLMY47V%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Version%2F4.0+Chrome%2F39.0.0.0+Safari%2F537.36",
			})

	}
}
