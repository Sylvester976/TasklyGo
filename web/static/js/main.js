// Elements
const todoInput = document.getElementById("todoInput");
const addBtn = document.getElementById("addBtn");
const todoList = document.getElementById("todoList");
const emptyState = document.getElementById("emptyState");
const clearBtn = document.getElementById("clearBtn");
const clearSection = document.getElementById("clearSection");
const filterButtons = document.querySelectorAll(".segmented-control .btn");

// State
let todos = [];
let filter = "all";

// Fetch tasks from backend
async function loadTodos() {
    try {
        const res = await fetch("/todos");
        todos = await res.json();
        renderTodos();
    } catch (err) {
        console.error("Error loading todos:", err);
    }
}

// Render tasks
function renderTodos() {
    todoList.innerHTML = "";
    const filteredTodos = todos.filter(todo => {
        if (filter === "active") return !todo.status;
        if (filter === "completed") return todo.status;
        return true;
    });

    if (filteredTodos.length === 0) {
        emptyState.style.display = "block";
    } else {
        emptyState.style.display = "none";
    }

    filteredTodos.forEach(todo => {
        const item = document.createElement("div");
        item.className = "todo-item d-flex justify-content-between align-items-center mb-2 p-2 border rounded";

        // Task title (editable)
        const title = document.createElement("span");
        title.textContent = todo.title;
        title.style.textDecoration = todo.status ? "line-through" : "none";
        title.style.cursor = "pointer";

        title.addEventListener("click", () => {
            const input = document.createElement("input");
            input.type = "text";
            input.value = todo.title;
            input.className = "form-control";
            item.replaceChild(input, title);
            input.focus();

            input.addEventListener("blur", () => saveTitle(todo, input.value, item, title));
            input.addEventListener("keydown", (e) => {
                if (e.key === "Enter") input.blur();
            });
        });

        // Actions
        const actions = document.createElement("div");
        actions.className = "d-flex gap-2";

        // Toggle done/undo button
        const toggleBtn = document.createElement("button");
        toggleBtn.className = "btn btn-sm btn-success";
        toggleBtn.innerHTML = todo.status ? '<i class="fas fa-undo"></i>' : '<i class="fas fa-check"></i>';
        toggleBtn.onclick = () => toggleTodo(todo.id);

        // Delete button
        const deleteBtn = document.createElement("button");
        deleteBtn.className = "btn btn-sm btn-danger";
        deleteBtn.innerHTML = '<i class="fas fa-trash"></i>';
        deleteBtn.onclick = () => deleteTodo(todo.id);

        actions.appendChild(toggleBtn);
        actions.appendChild(deleteBtn);

        item.appendChild(title);
        item.appendChild(actions);

        todoList.appendChild(item);
    });

    // Show clear completed if any completed tasks exist
    const completedCount = todos.filter(t => t.status).length;
    clearSection.classList.toggle("d-none", completedCount === 0);

    // Update stats
    document.getElementById("totalCount").textContent = todos.length;
    document.getElementById("activeCount").textContent = todos.filter(t => !t.status).length;
    document.getElementById("completedCount").textContent = completedCount;
}

// Save edited title to backend
async function saveTitle(todo, newTitle, item, titleEl) {
    if (!newTitle.trim()) {
        renderTodos();
        return;
    }

    try {
        await fetch(`/todos/update/${todo.id}`, {
            method: "PUT",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
                title: newTitle,
                description: todo.description,
                status: todo.status
            })
        });
        loadTodos();
    } catch (err) {
        console.error("Error updating title:", err);
        renderTodos();
    }
}

// Add new todo
addBtn.addEventListener("click", async () => {
    const title = todoInput.value.trim();
    if (!title) return;
    try {
        await fetch("/todos/create", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ title, description: "" })
        });
        todoInput.value = "";
        loadTodos();
    } catch (err) {
        console.error("Error adding todo:", err);
    }
});

// Toggle done status
async function toggleTodo(id) {
    const todo = todos.find(t => t.id === id);
    if (!todo) return;
    try {
        await fetch(`/todos/update/${id}`, {
            method: "PUT",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ title: todo.title, description: todo.description, status: !todo.status })
        });
        loadTodos();
    } catch (err) {
        console.error("Error updating todo:", err);
    }
}

// Delete todo
async function deleteTodo(id) {
    try {
        await fetch(`/todos/delete/${id}`, { method: "DELETE" });
        loadTodos();
    } catch (err) {
        console.error("Error deleting todo:", err);
    }
}

// Clear all completed
clearBtn.addEventListener("click", async () => {
    const completedTodos = todos.filter(t => t.status);
    for (let t of completedTodos) {
        await fetch(`/todos/delete/${t.id}`, { method: "DELETE" });
    }
    loadTodos();
});

// Filter buttons
filterButtons.forEach(btn => {
    btn.addEventListener("click", () => {
        filterButtons.forEach(b => b.classList.remove("active"));
        btn.classList.add("active");
        filter = btn.dataset.filter;
        renderTodos();
    });
});

// Initial load
loadTodos();
