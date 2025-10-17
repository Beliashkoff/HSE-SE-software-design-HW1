package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Beliashkoff/HSE-SE-software-design-HW1/zoo-erp/internal/repo"

	"github.com/Beliashkoff/HSE-SE-software-design-HW1/zoo-erp/internal/services"

	"github.com/Beliashkoff/HSE-SE-software-design-HW1/zoo-erp/internal/domain"

	"github.com/Beliashkoff/HSE-SE-software-design-HW1/zoo-erp/internal/app"
)

func main() {
	c := app.Build()

	if err := c.Invoke(func(z services.ZooService, gen services.NumberGen) {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Println("\n--- Московский зоопарк ---")
			fmt.Println("1) Добавить животное")
			fmt.Println("2) Добавить вещь")
			fmt.Println("3) Показать суммарный корм (кг/сутки)")
			fmt.Println("4) Список для контактного зоопарка")
			fmt.Println("5) Инвентаризация (имя и номер)")
			fmt.Println("0) Выход")
			fmt.Print("Выбор: ")

			choice := readInt(reader)
			switch choice {
			case 1:
				addAnimalFlow(reader, z, gen)
			case 2:
				addThingFlow(reader, z, gen)
			case 3:
				fmt.Printf("Суммарный корм: %d кг/сутки\n", z.TotalFoodKG())
			case 4:
				fmt.Println("Контактный зоопарк (>5 доброты):")
				for _, it := range z.ContactZoo() {
					fmt.Printf("- %s #%d\n", it.Name(), it.Number())
				}
			case 5:
				fmt.Println("Инвентаризация:")
				for _, it := range z.Inventory() {
					fmt.Printf("- %s #%d\n", it.Name(), it.Number())
				}
			case 0:
				fmt.Println("Пока!")
				return
			default:
				fmt.Println("Неизвестная команда")
			}
		}
	}); err != nil {
		panic(err)
	}
}

func addAnimalFlow(r *bufio.Reader, z services.ZooService, gen services.NumberGen) {
	fmt.Println("\nВид животного: 1) Rabbit  2) Monkey  3) Tiger  4) Wolf")
	fmt.Print("Ваш выбор: ")
	kind := readInt(r)

	fmt.Print("Кг/сутки (Food): ")
	food := readInt(r)

	var kindness int
	if kind == 1 || kind == 2 {
		fmt.Print("Доброта (0..10): ")
		kindness = readInt(r)
	}

	number := gen.Next() // единый источник номеров

	var item domain.InventoryItem
	var err error

	switch kind {
	case 1:
		var a *domain.Rabbit
		a, err = domain.NewRabbit(number, food, kindness)
		item = a
	case 2:
		var a *domain.Monkey
		a, err = domain.NewMonkey(number, food, kindness)
		item = a
	case 3:
		var a *domain.Tiger
		a, err = domain.NewTiger(number, food)
		item = a
	case 4:
		var a *domain.Wolf
		a, err = domain.NewWolf(number, food)
		item = a
	default:
		fmt.Println("Неизвестный вид")
		return
	}

	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	if err := z.AcceptAnimal(item); err != nil {
		switch err {
		case services.ErrRejectedByVet:
			fmt.Println("Ветклиника отказала в приёме животного")
		case repo.ErrDuplicateNumber:
			fmt.Println("Ошибка: дублируется инвентарный номер")
		case repo.ErrNotAnimal:
			fmt.Println("Ошибка: это не животное")
		default:
			fmt.Println("Ошибка:", err)
		}
		return
	}

	fmt.Printf("Принято: %s #%d, корм: %d кг/сутки\n", item.Name(), item.Number(), food)
}

func addThingFlow(r *bufio.Reader, z services.ZooService, gen services.NumberGen) {
	fmt.Println("\nВещь: 1) Table  2) Computer")
	fmt.Print("Ваш выбор: ")
	kind := readInt(r)

	number := gen.Next()

	var item domain.InventoryItem
	var err error

	switch kind {
	case 1:
		var t *domain.Table
		t, err = domain.NewTable(number)
		item = t
	case 2:
		var t *domain.Computer
		t, err = domain.NewComputer(number)
		item = t
	default:
		fmt.Println("Неизвестная вещь")
		return
	}

	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	if err := z.AddThing(item); err != nil {
		switch err {
		case repo.ErrDuplicateNumber:
			fmt.Println("Ошибка: дублируется инвентарный номер")
		default:
			fmt.Println("Ошибка:", err)
		}
		return
	}

	fmt.Printf("Добавлено: %s #%d\n", item.Name(), item.Number())
}

func readInt(r *bufio.Reader) int {
	for {
		line, _ := r.ReadString('\n')
		line = strings.TrimSpace(line)
		if n, err := strconv.Atoi(line); err == nil {
			return n
		}
		fmt.Print("Введите целое число: ")
	}
}
