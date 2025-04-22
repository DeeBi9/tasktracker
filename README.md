# Task Tracker CLI
This is task tracker cli build using golang, through this we can add, update and delete task, which are stored in JSON file.

## Additional feature
1. Mark a task as in progress or done
2. List all tasks
3. List all tasks that are done
4. List all tasks that are not done
5. List all tasks that are in progress

## Tech Stack
For this project I have used golang and cobra framework/library. Cobra is very convenient and very poweful library to create modern CLI applications.

## Commands 
1. add
2. update
3. delete
4. changestatus (`done`,`in-progress`,`to-do`)
5. listall 
6. list done, list to-do, list in-progress (we can list status wise also)

## How to use 
1. Clone the repo
    ```
    git clone https://github.com/DeeBi9/tasktracker
    ```
2. Install dependecy 
    ```
    go mod tidy
    ```
3. Build the Project
    ```
    go build -o tasktracker
    ```
4. Run the commands 
    * Add
        ```
        ./tasktracker add add "<Task-Description>"
        ```
    * Update
        ```
        ./tasktracker update update <task-id> "<New-Task-Description>"
        ```
    * Delete
        ```
        ./tasktracker delete delete <task-id>
        ```
    * Change status
        ```
        ./tasktracker changestatus "in-progress/to-do/done"
        ```
    * List All tasks
        ```
        ./tasktracker listall 
        ```
    * List specific Task status wise
        ```
        ./tasktracker list "To-DO/complete/in-progress"
        ```
## Conclusion

I have made this project to get deeper understanding in golang and cobra framework, to make more powerful CLI applictions.