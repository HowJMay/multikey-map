package multikey_test

import (
	"multikey-map"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	mkey := multikey.New()

	assert.NotNil(t, mkey)
}

func TestSetGet(t *testing.T) {
	type args struct {
		keys  []int
		value string
	}
	type wants struct {
		value string
	}

	type test struct {
		args  args
		wants wants
	}

	tests := map[string]func(t *testing.T) test{
		"successfully": func(t *testing.T) test {

			return test{
				args: args{
					keys:  []int{1, 2, 3},
					value: "espionage",
				},
				wants: wants{
					value: "espionage",
				},
			}
		},
	}

	for name, fn := range tests {
		t.Run(name, func(t *testing.T) {
			tt := fn(t)

			mkey := multikey.New()
			mkey.Set(tt.args.keys, tt.args.value)

			vals := mkey.Get(tt.args.keys).(string)
			assert.Equal(t, tt.wants.value, vals)
		})
	}
}
