package main

import (
	"log"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	b, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	err = b.AddFromFile("main.glade")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	obj, err := b.GetObject("window_main")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	win := obj.(*gtk.Window)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	obj, _ = b.GetObject("entry_1")
	entry1 := obj.(*gtk.Entry)

	// Получаем кнопку
	obj, _ = b.GetObject("button_1")
	button1 := obj.(*gtk.Button)

	// Получаем метку
	obj, _ = b.GetObject("label_1")
	label1 := obj.(*gtk.Label)

	// Сигнал по нажатию на кнопку
	button1.Connect("clicked", func() {
		text, err := entry1.GetText()
		if err == nil {
			// Устанавливаем текст из поля ввода метке
			label1.SetText(text)
		}
	})

	// Отображаем все виджеты в окне
	win.ShowAll()

	// Выполняем главный цикл GTK (для отрисовки). Он остановится когда
	// выполнится gtk.MainQuit()
	gtk.Main()
}
