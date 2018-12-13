package main

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONEncoder(t *testing.T) {
	type args struct {
		c Character
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "test1",
			args: args{c: Character{Name: "hero", Hp: 200, Dmg: 10}},
			want: func() []byte {
				ret, _ := json.Marshal(Character{Name: "hero", Hp: 200, Dmg: 10})
				return ret
			}(),
			wantErr: false,
		},
		{
			name: "test2",
			args: args{c: Character{Name: "monster", Hp: 2, Dmg: 5}},
			want: func() []byte {
				ret, _ := json.Marshal(Character{Name: "monster", Hp: 2, Dmg: 5})
				return ret
			}(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JSONEncoder(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONEncoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONDecoder(t *testing.T) {
	type args struct {
		s []byte
		c *Character
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := JSONDecoder(tt.args.s, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("JSONDecoder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJSONEncoders(t *testing.T) {
	//1 - инициализируем структуру, с помощью которой будем проверять
	expected := Character{Name: "hero", Hp: 200, Dmg: 10}

	//2 - преобразовываем его в JSON формат
	encode, err := JSONEncoder(expected)
	if err != nil {
		t.Errorf("func JSONEncoder returned error: %s", err.Error())
	}

	//3 - проверяем преобразование через стандартные функции
	actual := Character{}
	json.Unmarshal(encode, &actual)

	//4 - сравниваем
	if !assert.True(t, actual == expected) {
		t.Errorf("func JSONEncoder returned unexpected result: got %v want %v", actual, expected)
	}

}

//TestJSONDecoder - тестируем обратное преобразование с JSON
func TestJSONDecoders(t *testing.T) {
	expected := Character{Name: "hero", Hp: 200, Dmg: 10}

	encode, _ := json.Marshal(expected)

	actual := Character{}
	if err := JSONDecoder(encode, &actual); err != nil {
		t.Errorf("func JSONDecoder returned error: %s", err.Error())
	}

	if !assert.True(t, actual == expected) {
		t.Errorf("func JSONDecoder returned unexpected result: got %v want %v", actual, expected)
	}
}
