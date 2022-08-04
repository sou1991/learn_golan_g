package main

import (
	"fmt"
)

type Animal struct{
	Name, Sound string
}

type IAnimal interface{
	MakeSound() string
	//・・・メソッドシグニチャ郡を作る
}

func (animal Animal) MakeSound() string{
	return animal.Sound
}

func main(){
	var cat IAnimal = Animal{
		Name : "ねこ",
		Sound : "にゃー",
	}
	var rion IAnimal = Animal{
		Name : "ライオン",
		Sound : "がおー",
	}
	var dog IAnimal = Animal{
		Name : "いぬ",
		Sound : "ワン",
	}

	fmt.Println(DoSomeThing(cat))
	fmt.Println(DoSomeThing(rion))
	fmt.Println(DoSomeThing(dog))
}

func DoSomeThing(animal IAnimal) string{
	return animal.MakeSound()
}
