package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarLista(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) eliminarDeLista(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)

}
func (t *taskList) mostrarTareas() {
	for _, v := range t.tasks {
		fmt.Println("Nombre", v.nombre)
		fmt.Println("Descripcion", v.descricion)
	}

}
func (t *taskList) tareasCompletadas() {
	for _, v := range t.tasks {
		if v.completado {
			fmt.Println("Nombre: ", v.nombre)
			fmt.Println("Descripcion: ", v.descricion)
		}

	}

}

type task struct {
	nombre     string
	descricion string
	completado bool
}

func (t *task) marcaCompleta() {
	t.completado = true
}
func (t *task) actualizarDescripcion(descripcion string) {
	t.descricion = descripcion
}
func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	t1 := &task{
		nombre:     "completar mi cusrso de go",
		descricion: "completar mi cuarso de Go de platzu en esta semana",
	}
	t2 := &task{
		nombre:     "completar mi cusrso de python",
		descricion: "completar mi cuarso de python de platzu en esta semana",
	}
	t3 := &task{
		nombre:     "completar mi cusrso de javaScript",
		descricion: "completar mi cuarso de javaScript de platzu en esta semana",
	}

	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	lista.agregarLista(t3)

	lista.eliminarDeLista(1)

	lista.tasks[1].marcaCompleta()
	lista.tareasCompletadas()
	mapaTareas := make(map[string]*taskList)

	mapaTareas["Nestor"] = lista

	t4 := &task{
		nombre:     "completar mi cusrso de java",
		descricion: "completar mi cuarso de java de platzu en esta semana",
	}
	t5 := &task{
		nombre:     "completar mi cusrso de c#",
		descricion: "completar mi cuarso de c# de platzu en esta semana",
	}

	lista2 := &taskList{
		tasks: []*task{
			t4, t5,
		},
	}
	mapaTareas["Ricardo"] = lista2

	mapaTareas["Ricardo"].mostrarTareas()
}
