'use strict';
// 各種DOMを取得
const btn = document.querySelector('#btn');
const textBox = document.querySelector('#textbox');
const tbody = document.querySelector('#tbody');
const all = document.querySelector('#all');
const inprogress = document.querySelector('#inprogress');
const done = document.querySelector('#done');

// 「すべて」ラジオボタンを押した時の動作
all.addEventListener('click', () => {
  tbody.innerHTML = '';
  todos.forEach((todo) => {
    displayTodos(todo);
  });
})

// 「作業中」ラジオボタンを押した時の動作
inprogress.addEventListener('click', () => {
  tbody.innerHTML = '';
  const inprogressTodos = todos.filter((currentStatus) => {
    return currentStatus.status === '作業中';
  });
  inprogressTodos.forEach((todo) => {
    displayTodos(todo);
  });
});

// 「完了」ラジオボタンを押した時の動作
done.addEventListener('click', () => {
  tbody.innerHTML = '';
  const doneTodos = todos.filter((currentStatus) => {
    return currentStatus.status === '完了';
  });
  doneTodos.forEach((todo) => {
    displayTodos(todo);
  });
});

// タスクを格納する配列を定義
const todos = [];

// カウンターを定義
let counter = 0;

// todoを追加する関数
function createTask() {
  const todo = {
    id: counter,
    task: textBox.value,
    status: '作業中'
  };
  todos.push(todo);
  counter++;
  textBox.value = '';
};

// 状態ボタンを生成する関数
function createStatusBtn(currentTodo) {
  const statusBtn = document.createElement('td');
  const btn = document.createElement('button');
  btn.innerText = currentTodo.status;
  btn.addEventListener('click', (event) => {
    const arrayId = event.path[2].childNodes[0].innerText;
    if(todos[arrayId].status === '作業中') {
      todos[arrayId].status = '完了';
    }else{
      todos[arrayId].status = '作業中';
    };
    tbody.innerHTML = '';
    const inprogressTodos = todos.filter((currentStatus) => {
      return currentStatus.status === '作業中';
    });
    const doneTodos = todos.filter((currentStatus) => {
      return currentStatus.status === '完了';
    });
    if(all.checked) {
      todos.forEach((todo) => {
        displayTodos(todo);
      });
    }else if(inprogress.checked) {
      inprogressTodos.forEach((todo) => {
        displayTodos(todo);
      });
    }else if(done.checked) {
      doneTodos.forEach((todo) => {
        displayTodos(todo);
      });
    };
  })
  statusBtn.appendChild(btn);
  return statusBtn;
};

// タスクを削除する関数
function deleteTask(event) {
  const arrayId = event.path[2].childNodes[0].innerText;
  todos.splice(arrayId, 1);
  tbody.innerHTML = '';
  for(let i = 0; i < todos.length; i++){
   todos[i].id = i;
  };
  const inprogressTodos = todos.filter((currentStatus) => {
    return currentStatus.status === '作業中';
  });
  const doneTodos = todos.filter((currentStatus) => {
    return currentStatus.status === '完了';
  });
  if(all.checked) {
    todos.forEach((todo) => {
      displayTodos(todo);
    });
  }else if(inprogress.checked) {
    inprogressTodos.forEach((todo) => {
      displayTodos(todo);
    });
  }else if(done.checked) {
    doneTodos.forEach((todo) => {
      displayTodos(todo);
    });
  };
  counter += -1;
};

// 削除ボタンを生成する関数
function createDelBtn() {
  const delBtn = document.createElement('td');
  const btn = document.createElement('button');
  btn.innerText = '削除';
  delBtn.appendChild(btn);
  btn.addEventListener('click', (event) => {
    deleteTask(event);
  });
  return delBtn;
};

// todoを表示する関数
const displayTodos = (todo) => {
  const tr = document.createElement('tr');
  // インデックス
  const indexTd = document.createElement('td');
  indexTd.innerText = todo.id;
  // 入力内容
  const taskTd = document.createElement('td');
  taskTd.innerText = todo.task;
  // 配列をHTMLへ追加
  tr.appendChild(indexTd);
  tr.appendChild(taskTd);
  tr.appendChild(createStatusBtn(todo));
  tr.appendChild(createDelBtn());
  tbody.appendChild(tr);
};

// 追加ボタンを押下した時の動作
btn.addEventListener('click', () => {
  tbody.innerHTML = '';
  createTask();
  const inprogressTodos = todos.filter((currentStatus) => {
    return currentStatus.status === '作業中';
  });
  const doneTodos = todos.filter((currentStatus) => {
    return currentStatus.status === '完了';
  });
  if(all.checked) {
    todos.forEach((todo) => {
      displayTodos(todo);
    });
  }else if(inprogress.checked) {
    inprogressTodos.forEach((todo) => {
      displayTodos(todo);
    });
  }else if(done.checked) {
    doneTodos.forEach((todo) => {
      displayTodos(todo);
    });
  };
});