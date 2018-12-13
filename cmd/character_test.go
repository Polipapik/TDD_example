package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var characters = map[string]Character{
	"hero":    Character{Name: "Hero", Hp: 200, Dmg: 10, IsAlive: true},
	"monster": Character{Name: "Monster", Hp: 25, Dmg: 5, IsAlive: true},
	"boss":    Character{Name: "Boss", Hp: 50, Dmg: 20, IsAlive: true},
}

func TestJSONEncoder(t *testing.T) {
	// 1 - инициализируем структуру, с помощью которой будем проверять
	expected := characters["hero"]

	// 2 - преобразовываем его в JSON формат
	encode, err := JSONEncoder(expected)
	if err != nil {
		t.Errorf("func JSONEncoder returned error: %s", err.Error())
	}

	// 3 - проверяем преобразование через стандартные функции
	actual := Character{}
	if err = json.Unmarshal(encode, &actual); err != nil {

	}

	// 4 - сравниваем
	if !assert.True(t, actual == expected) {
		t.Errorf("func JSONEncoder returned unexpected result: got %v want %v", actual, expected)
	}

}

//TestJSONDecoder - тестируем обратное преобразование с JSON
func TestJSONDecoders(t *testing.T) {
	// Выполняем действия обратные Encoding
	expected := characters["hero"]

	encode, _ := json.Marshal(expected)

	actual := Character{}
	if err := JSONDecoder(encode, &actual); err != nil {
		t.Errorf("func JSONDecoder returned error: %s", err.Error())
	}

	if !assert.True(t, actual == expected) {
		t.Errorf("func JSONDecoder returned unexpected result: got %v want %v", actual, expected)
	}
}

//TestHitFunction - тестируем математику функции удара
func TestHitFunction_CheckDamage(t *testing.T) {
	hero := characters["hero"]
	monster := characters["monster"]

	expected := (monster.Hp - hero.Dmg)

	hero.Hit(&monster)
	actual := monster.Hp

	if !assert.True(t, actual == expected) {
		t.Errorf("func Hit returned unexpected monster hp value: got %vhp, want %vhp", actual, expected)
	}
	// Добавляем математику
}

//TestHitFunction - тестируем, умирает ли цель после удара
func TestHitFunction_CheckIsAlive(t *testing.T) {
	hero := characters["hero"]
	monster := characters["monster"]

	// Ожидаем, что монстр умрет
	expected := false

	hero.Hit(&monster)
	hero.Hit(&monster)
	hero.Hit(&monster)
	actual := monster.IsAlive

	if !assert.True(t, actual == expected) {
		t.Errorf("func Hit returned unexpected monster isAlive value: got %v, want %v", actual, expected)
	}
	//В результате этого теста была добавлена проверка, если hp < 0, то персонаж умирает
}

//TestHitFunction - тестируем, удары по мертвым (лежачих не бьют)
func TestHitFunction_CheckIsAliveHitting(t *testing.T) {
	hero := characters["hero"]
	monster := characters["monster"]

	hero.Hit(&monster)
	hero.Hit(&monster)
	hero.Hit(&monster)
	expected := monster.Hp // Записываем хп во время смерти

	hero.Hit(&monster)
	actual := monster.Hp

	if !assert.True(t, actual == expected) {
		t.Errorf("func Hit returned unexpected monster hp value: got %vhp, want %vhp", actual, expected)
	}
	// В результате этого теста в функции Hit появилась обертка в виде первого if target.IsAlive == true {..}
}

// TestHitFunction_Result - в нашей рпг должен быть хоть какой-то вывод. Эта функция проверяет выводы
func TestHitFunction_Result(t *testing.T) {
	hero := characters["hero"]
	monster := characters["monster"]

	expected := fmt.Sprintf("%v hit %v, now %vs hp is %v", hero.Name, monster.Name, monster.Name, (monster.Hp - hero.Dmg))
	actual := hero.Hit(&monster)

	if !assert.True(t, actual == expected) {
		t.Errorf("func Hit returned unexpected result: got { %v }, want { %v }", actual, expected)
	}

	hero.Hit(&monster)

	expected = fmt.Sprintf("%v hit %v, now %vs hp is %v and %v died", hero.Name, monster.Name, monster.Name, (monster.Hp - hero.Dmg), monster.Name)
	actual = hero.Hit(&monster)

	if !assert.True(t, actual == expected) {
		t.Errorf("func Hit returned unexpected result: got { %v }, want { %v }", actual, expected)
	}

	expected = fmt.Sprintf("%v is already dead", monster.Name)
	actual = hero.Hit(&monster)

	if !assert.True(t, actual == expected) {
		t.Errorf("func Hit returned unexpected result: got { %v }, want { %v }", actual, expected)
	}
	// В результате этого теста добавилось возвращаемое значение функции Hit, и логика связанная с ним
}

func TestFightFunction(t *testing.T) {
	hero := characters["hero"]
	monster := characters["monster"]

	Fight(&hero, &monster)
	expected := false
	actual := monster.IsAlive

	if !assert.True(t, actual == expected) {
		t.Errorf("func Fight returned unexpected monster isAlive value: got { %v }, want { %v }", actual, expected)
	}

	expectedHp := hero.Hp
	Fight(&hero, &monster)
	actualHp := hero.Hp

	if !assert.True(t, actualHp == expectedHp) {
		t.Errorf("func Fight returned unexpected monster isAlive value: got { %v }, want { %v }", actual, expected)
	}
}
