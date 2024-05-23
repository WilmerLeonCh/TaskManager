# CLI Task Manager
* This is a CLI task manager that allows you to add, update, remove, list, show details and mark tasks as completed. It is written in golang and uses cobra to interact with the Postgres database.
* It was stylized using bubbletea with **bubbles** and **lipgloss** extensions.

# Demo
![Demo](demo.gif)

# Run locally
* Clone the project
    ```
    git clone https://github.com/WilmerLeonCh/TaskManager.git
    ```
* Set your environment variables (You can guide from the _**.env.example**_ file)
    ```
    TASK_MANAGER_POSTGRES_HOST="localhost"
    TASK_MANAGER_POSTGRES_PORT="5410"
    TASK_MANAGER_POSTGRES_USER="r00t"
    TASK_MANAGER_POSTGRES_PASSWORD="passw0rd"
    TASK_MANAGER_POSTGRES_DB="task_manager_db"
    ```
* Run docker-compose
    ```
    docker-compose up
    ```
* Let start to administrate your tasks
   ```
   go run main.go <command> <args>
   ```
    Available commands:

  | Command   | args |
  |-----------|:----:|
  | add       |  -   |
  | completed |  id  |
  | delete    |  id  |
  | details   |  id  |
  | list      |  -   |
  | update    |  id  |
* Enjoy it!
