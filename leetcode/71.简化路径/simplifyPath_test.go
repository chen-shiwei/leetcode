package _1_简化路径

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: `输入：path = "/a/./b/../../c/"输出："/c"`, args: args{path: "/a/./b/../../c/"}, want: "/c"},
		{name: `"/home//foo/" "/home/foo"`, args: args{path: "/home//foo/"}, want: "/home/foo"},
		{name: `"/../" "/"`, args: args{path: "/home//foo/"}, want: "/"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
