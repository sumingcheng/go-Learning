import os
import re

md_structure = '''
- cmd/
  - main.go
- internal/
  - domain/
    - model/
      - order.go
    - repository/
      - order_repository.go
  - application/
    - service/
      - order_service.go
  - infrastructure/
    - repository/
      - order_repository_impl.go
    - db/
      - database.go
  - interfaces/
    - handler/
      - order_handler.go
  - di/
    - wire.go
'''

def parse_md_structure(md_text):
    lines = md_text.strip().split('\n')
    stack = []
    paths = []
    for line in lines:
        # 计算缩进级别（每两个空格为一级）
        indent = len(line) - len(line.lstrip(' '))
        level = indent // 2

        # 提取名称，去除前面的符号和空格
        name = line.strip().lstrip('-').strip()

        # 更新堆栈
        while len(stack) > level:
            stack.pop()
        stack.append(name)

        # 构建路径
        path = os.path.join(*stack)
        paths.append(path)

    return paths

def create_dirs_and_files(paths):
    for path in paths:
        if path.endswith('/'):
            # 是目录
            dir_path = path.rstrip('/')
            os.makedirs(dir_path, exist_ok=True)
            print(f"Created directory: {dir_path}")
        elif '.' in os.path.basename(path):
            # 是文件
            dir_name = os.path.dirname(path)
            if dir_name and not os.path.exists(dir_name):
                os.makedirs(dir_name, exist_ok=True)
                print(f"Created directory: {dir_name}")
            with open(path, 'w') as f:
                pass
            print(f"Created file: {path}")
        else:
            # 是目录（没有斜杠，但也不是文件）
            os.makedirs(path, exist_ok=True)
            print(f"Created directory: {path}")

if __name__ == "__main__":
    paths = parse_md_structure(md_structure)
    create_dirs_and_files(paths)
