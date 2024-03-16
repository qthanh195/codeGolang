/*
Game đơn giản:

Tạo struct NhanVat với các thuộc tính: tên, máu, sát thương.
Tạo method TanCong để tấn công nhân vật khác.
Tạo method PhongThu để phòng thủ khi bị tấn công.
Tạo struct QuaiVat với các thuộc tính: tên, máu, sát thương.
Tạo struct TroChoi với các thuộc tính: nhân vật, quái vật.
Viết code để thực hiện trận chiến giữa nhân vật và quái vật.
*/

package main

import (
	"fmt"
	"strconv"
	"time"
)

type Character struct {
	name   string
	blood  int8
	damage uint8
}

func (s *Character) addCharacter() {
	fmt.Println("Add Character of you.")
	for {
		fmt.Print("Name: ")
		fmt.Scan(&s.name)
		//kiem tra name phai la string
		if _, err := strconv.Atoi(s.name); err != nil {
			fmt.Println("Name must be string.")
			continue
		}
		fmt.Printf("Blood of %s: ", s.name)
		fmt.Scan(&s.blood)
		//kiem tra blood phair laf int
		if _, err := strconv.Atoi(fmt.Sprintf("%d", s.blood)); err != nil {
			fmt.Println("Blood must be int.")
			continue
		}
		//kiem tra damage phai la int
		fmt.Printf("Damage of %s: ", s.name)
		fmt.Scan(&s.damage)
		if _, err := strconv.Atoi(fmt.Sprintf("%d", s.damage)); err != nil {
			fmt.Println("Damage must be int.")
			continue
		}
		break

	}
}

func (s *Character) printInforCharacter() {
	fmt.Println("___________________________________")
	fmt.Println("_____Your character information____")
	fmt.Println("Name:   ", s.name)
	fmt.Println("Blood:  ", s.blood)
	fmt.Println("Damage: ", s.damage)
	fmt.Println("___________________________________")

}
func (s *Character) attackCharacter() {
	fmt.Printf("%s attack %d damage.\n", s.name, s.damage)
}

// func (s Character) defenseCharacter() {
// 	fmt.Printf("%s defense %d damage.\n", s.name, s.damage)

// }

type Monster struct {
	name    string
	blood   int8
	defense uint8
}

func (s *Monster) addMonster() {
	fmt.Println("Add monster you want combat.")
	for {
		fmt.Print("Name: ")
		fmt.Scan(&s.name)
		//kiem tra name phai la string
		if _, err := strconv.Atoi(s.name); err != nil {
			fmt.Println("Name must be string.")
			continue
		}
		fmt.Printf("Blood of %s: ", s.name)
		fmt.Scan(&s.blood)
		//kiem tra blood phair laf int
		if _, err := strconv.Atoi(fmt.Sprintf("%d", s.blood)); err != nil {
			fmt.Println("Blood must be int.")
			continue
		}
		//kiem tra damage phai la int
		fmt.Printf("Defense of %s: ", s.name)
		fmt.Scan(&s.defense)
		if _, err := strconv.Atoi(fmt.Sprintf("%d", s.defense)); err != nil {
			fmt.Println("Defense must be int.")
			continue
		}
		break
	}
}

// func (s Monster) attackMonster() {
// 	fmt.Printf("%s attack %d damage.\n", s.name, s.damage)
// }

func (s *Monster) defenseMonster() {
	fmt.Printf("%s defense %d damage.\n", s.name, s.defense)

}
func (s *Monster) printInforMonster() {
	fmt.Println("___________________________________")
	fmt.Println("________Monster information________")
	fmt.Println("Name:   ", s.name)
	fmt.Println("Blood:  ", s.blood)
	fmt.Println("Defense: ", s.defense)
	fmt.Println("___________________________________")

}

func main() {
	var character Character
	var monster Monster
	character.addCharacter()
	character.printInforCharacter()
	time.Sleep(time.Second)

	monster.addMonster()
	monster.printInforMonster()
	fmt.Println("The battle will start in 5 seconds.")
	time.Sleep(time.Second * 5)

	fmt.Println("Fight monsters!")

	for {

		character.attackCharacter()
		monster.defenseMonster()
		monster.blood = monster.blood - (int8(character.damage) - int8(monster.defense))
		fmt.Printf("The monster has %d health remaining", monster.blood)
		if monster.blood <= 0 {
			fmt.Println("The monster is dead!")
			fmt.Println("The battle is over.")
			break
		}

		time.Sleep(time.Second * 2)
	}
}
