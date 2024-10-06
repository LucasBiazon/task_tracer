package main

import (
	"flag"
	"fmt"
	"os"

	entity "github.com/lucasBiazon/task_tracker/internal/entities"
	"github.com/lucasBiazon/task_tracker/internal/repository"
	usecases "github.com/lucasBiazon/task_tracker/internal/use-cases"
)

func main() {
	repo := repository.NewTaskRepository("tasks.json")
	taskUseCase := usecases.NewTaskCases(*repo)

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	markInProgressCmd := flag.NewFlagSet("mark-in-progress", flag.ExitOnError)
	markDoneCmd := flag.NewFlagSet("mark-done", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	addDescription := addCmd.String("description", "", "Descrição da tarefa")
	updateID := updateCmd.String("id", "", "ID da tarefa para atualizar")
	updateDescription := updateCmd.String("description", "", "Nova descrição da tarefa")
	deleteID := deleteCmd.String("id", "", "ID da tarefa para deletar")
	markInProgressID := markInProgressCmd.String("id", "", "ID da tarefa para marcar como 'in-progress'")
	markDoneID := markDoneCmd.String("id", "", "ID da tarefa para marcar como 'done'")
	getID := getCmd.String("id", "", "ID da tarefa para obter")

	if len(os.Args) < 2 {
		fmt.Println("Esperado um comando (add, update, delete, mark-in-progress, mark-done, list, get)")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		err := taskUseCase.CreateTask(*addDescription)
		if err != nil {
			fmt.Println("Erro ao adicionar tarefa:", err)
		} else {
			fmt.Println("Tarefa adicionada com sucesso!")
		}

	case "update":
		updateCmd.Parse(os.Args[2:])
		err := taskUseCase.UpdateTask(*updateID, *updateDescription)
		if err != nil {
			fmt.Println("Erro ao atualizar tarefa:", err)
		} else {
			fmt.Println("Tarefa atualizada com sucesso!")
		}

	case "delete":
		deleteCmd.Parse(os.Args[2:])
		err := taskUseCase.DeleteTask(*deleteID)
		if err != nil {
			fmt.Println("Erro ao deletar tarefa:", err)
		} else {
			fmt.Println("Tarefa deletada com sucesso!")
		}

	case "mark-in-progress":
		markInProgressCmd.Parse(os.Args[2:])
		err := taskUseCase.InProgressTask(*markInProgressID)
		if err != nil {
			fmt.Println("Erro ao marcar tarefa como 'in-progress':", err)
		} else {
			fmt.Println("Tarefa marcada como 'in-progress'!")
		}

	case "mark-done":
		markDoneCmd.Parse(os.Args[2:])
		err := taskUseCase.CompleteTask(*markDoneID)
		if err != nil {
			fmt.Println("Erro ao marcar tarefa como 'done':", err)
		} else {
			fmt.Println("Tarefa marcada como 'done'!")
		}

	case "list":
		listCmd.Parse(os.Args[2:])
		if listCmd.NArg() > 0 {
			status := listCmd.Arg(0)
			switch status {
			case "done":
				tasks, err := taskUseCase.GetDoneTasks()
				if err != nil {
					fmt.Println("Erro ao listar tarefas 'done':", err)
				} else {
					printTasks(tasks)
				}
			case "in-progress":
				tasks, err := taskUseCase.GetInProgressTasks()
				if err != nil {
					fmt.Println("Erro ao listar tarefas 'in-progress':", err)
				} else {
					printTasks(tasks)
				}
			case "todo":
				tasks, err := taskUseCase.GetTodoTasks()
				if err != nil {
					fmt.Println("Erro ao listar tarefas 'todo':", err)
				} else {
					printTasks(tasks)
				}
			default:
				fmt.Println("Status inválido. Use 'done', 'in-progress' ou 'todo'.")
			}
		} else {
			tasks, err := taskUseCase.GetTasks()
			if err != nil {
				fmt.Println("Erro ao listar todas as tarefas:", err)
			} else {
				printTasks(tasks)
			}
		}

	case "get":
		getCmd.Parse(os.Args[2:])
		if *getID == "" {
			fmt.Println("ID é obrigatório para obter uma tarefa.")
			os.Exit(1)
		}
		task, err := taskUseCase.GetTask(*getID)
		if err != nil {
			fmt.Println("Erro ao obter tarefa:", err)
		} else {
			printTask(task)
		}

	default:
		fmt.Println("Comando não reconhecido:", os.Args[1])
		os.Exit(1)
	}
}

func printTasks(tasks []*entity.Task) {
	for _, task := range tasks {
		fmt.Printf("ID: %s, Descrição: %s, Status: %s, Criada em: %s, Atualizada em: %s\n", task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
	}
}

func printTask(task *entity.Task) {
	if task != nil {
		fmt.Printf("ID: %s, Descrição: %s, Status: %s, Criada em: %s, Atualizada em: %s\n", task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
	} else {
		fmt.Println("Tarefa não encontrada.")
	}
}
