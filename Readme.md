# Todo CLI

Todo CLI is a simple yet powerful Command Line Interface (CLI) tool to manage your daily tasks. Built with Go, it provides an easy and efficient way to add, list, complete, and delete todos directly from the terminal. With Todo CLI, you can keep track of your tasks and stay organized, all without leaving your command line.

## Features

- **Add Todos**: Quickly add tasks to your todo list with a single command.
- **List Todos**: Display all your tasks in a clean, tabular format with color-coded status indicators.
- **Mark as Complete**: Easily mark tasks as completed and track their completion time.
- **Delete Todos**: Remove tasks you no longer need from your list.
- **Persistent Storage**: All tasks are stored in a local JSON file, so your list is available even after restarting the CLI.
- **Clear Todos**: Reset your todo list with a single command.

## Future Improvements

- **Task Prioritization**: Add the ability to assign priorities to tasks for better organization.
- **Due Dates**: Include due dates for tasks and display overdue tasks.
- **Recurring Tasks**: Implement support for recurring tasks with configurable intervals.
- **Export and Import**: Provide options to export tasks to a file or import them from other sources.

## Installation

1. **Download**: Get the `.zip` file for Windows from the [releases page](https://github.com/AayushBhusari/Todo/releases).
2. **Extract**: Right-click the downloaded `.zip` file and select "Extract All..." to unzip the file to a location of your choice.
3. **Setup**: Move the extracted `.exe` file to a directory of your choice. To make it accessible from anywhere in the terminal, add this directory to your system's PATH environment variable:
   - Open the Start Menu and search for "Environment Variables."
   - Select "Edit the system environment variables."
   - In the System Properties window, click on "Environment Variables..."
   - In the Environment Variables window, locate the `Path` variable under "System variables," select it, and click "Edit..."
   - Click "New" and add the path to the directory where you placed the `.exe` file.
   - Click "OK" to close all dialog boxes.
4. **Verify Installation**: Open Command Prompt and type `todo` to ensure the application runs correctly.

Here’s the revised usage section with the correct commands from your image:

---

## Usage

To manage your tasks with Todo CLI, use the following commands:

- **Add a Task**:

  ```bash
  gotodo -add
  ```

  Adds a new todo by prompting for task details.

- **Clear All Tasks**:

  ```bash
  gotodo -clear
  ```

  Clears all existing todos.

- **Mark a Task as Complete**:

  ```bash
  gotodo -complete <task_index>
  ```

  Marks the specified task as complete.

- **Delete a Task**:

  ```bash
  gotodo -del <task_index>
  ```

  Deletes the specified task.

- **Exit the Application**:

  ```bash
  gotodo -exit
  ```

  Exits the todo app.

- **List All Tasks**:
  ```bash
  gotodo -list
  ```
  Displays all the todos in the list.

---

Would you like me to add more examples, or expand on any section?

## Examples

- **Adding a Task**:

  ```bash
  gotodo -add "Complete the project documentation"
  ```

- **Listing Tasks**:  
  Displays a table with task details like index, description, status, creation time, and completion time (if applicable).

- **Completing a Task**:

  ```bash
  gotodo -com 1
  ```

  Marks the first task in the list as completed.

- **Deleting a Task**:

  ```bash
  gotodo -del 2
  ```

  Deletes the second task from the list.

- **Clearing All Tasks**:
  ```bash
  gotodo -cl
  ```

## Contributing

Contributions to Todo CLI are welcome! If you have suggestions for new features or improvements, please submit an issue or pull request on [GitHub](https://github.com/AayushBhusari/Todo).
